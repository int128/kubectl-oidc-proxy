// Package portforwarder provides port forwarding between local and Kubernetes.
package portforwarder

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"net/url"
	"os"

	"github.com/google/wire"
	"github.com/int128/kauthproxy/pkg/logger"
	"golang.org/x/xerrors"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
)

var Set = wire.NewSet(
	wire.Struct(new(PortForwarder), "*"),
	wire.Bind(new(Interface), new(*PortForwarder)),
)

//go:generate mockgen -destination mock_portforwarder/mock_portforwarder.go github.com/int128/kauthproxy/pkg/portforwarder Interface

// Options represents an option of PortForwarder.
type Options struct {
	Config *rest.Config
	Source Source
	Target Target
}

// Source represents a local source.
type Source struct {
	Port int
}

// Target represents a target pod.
type Target struct {
	Pod           *v1.Pod
	ContainerPort int
}

type Interface interface {
	Run(ctx context.Context, o Options) error
}

type PortForwarder struct {
	Logger logger.Interface
}

func (pf *PortForwarder) Run(ctx context.Context, o Options) error {
	pfURL, err := url.Parse(o.Config.Host + o.Target.Pod.GetSelfLink() + "/portforward")
	if err != nil {
		return xerrors.Errorf("could not build URL for portforward: %w", err)
	}
	rt, upgrader, err := spdy.RoundTripperFor(o.Config)
	if err != nil {
		return xerrors.Errorf("could not create a round tripper: %w", err)
	}
	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: rt}, http.MethodPost, pfURL)
	portPair := fmt.Sprintf("%d:%d", o.Source.Port, o.Target.ContainerPort)
	stopChan, readyChan := make(chan struct{}), make(chan struct{})
	forwarder, err := portforward.NewOnAddresses(dialer, []string{"127.0.0.1"}, []string{portPair}, stopChan, readyChan, os.Stdout, os.Stderr)
	if err != nil {
		return xerrors.Errorf("could not create a port forwarder: %w", err)
	}

	finalizeChan := make(chan struct{})
	var eg errgroup.Group
	eg.Go(func() error {
		defer close(finalizeChan)
		pf.Logger.V(1).Infof("starting a port forwarder at %s", portPair)
		if err := forwarder.ForwardPorts(); err != nil {
			return xerrors.Errorf("could not run the forwarder at %s: %w", portPair, err)
		}
		pf.Logger.V(1).Infof("stopped the port forwarder at %s", portPair)
		return nil
	})
	eg.Go(func() error {
		defer close(stopChan)
		select {
		case <-ctx.Done():
			pf.Logger.V(1).Infof("stopping the port forwarder at %s", portPair)
			return xerrors.Errorf("stopping the port forwarder: %w", ctx.Err())
		case <-finalizeChan:
			pf.Logger.V(1).Infof("finished goroutine of the port forwarder at %s", portPair)
			return nil
		}
	})
	if err := eg.Wait(); err != nil {
		return xerrors.Errorf("error while running a port forwarder: %w", err)
	}
	return nil
}
