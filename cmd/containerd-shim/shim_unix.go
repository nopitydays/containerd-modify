// +build !linux,!windows,!darwin

package main

import (
	"os"
	"os/signal"
	"fmt"

	"github.com/containerd/containerd/reaper"
	runc "github.com/containerd/go-runc"
	"github.com/stevvooe/ttrpc"
)

// setupSignals creates a new signal handler for all signals and sets the shim as a
// sub-reaper so that the container processes are reparented
func setupSignals() (chan os.Signal, error) {
	fmt.Printf("%q\n", "this is unix setupSignals!")
	signals := make(chan os.Signal, 2048)
	signal.Notify(signals)
	// make sure runc is setup to use the monitor
	// for waiting on processes
	runc.Monitor = reaper.Default
	return signals, nil
}

func newServer() (*ttrpc.Server, error) {
	fmt.Printf("%q\n", "this is unix setupSignals!")
	return ttrpc.NewServer(ttrpc.WithServerHandshaker(ttrpc.UnixSocketRequireSameUser()))
}
