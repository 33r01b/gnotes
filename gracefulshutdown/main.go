package main

// Graceful Shutdown realization by sigterm or some app errors

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	w := worker{errors: make(chan error, 1)}
	ctx, cancel := context.WithCancel(context.Background())

	// used WaitGroup to control goroutines from root
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		w.run(ctx) // run worker
		wg.Done()
	}()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)

	// catch errors by chanel
	select {
	case x := <-sigc:
		log.Info("received a signal.", x.String())
		break
	case err := <-w.notify():
		log.Error("received an error: ", err)
		break
	}

	cancel()  // signal cancel to context
	wg.Wait() // wait childs processes

	log.Info("worker stopped")
}

type worker struct {
	errors chan error
}

// wrap worker main handler, to catch errors
func (w *worker) run(ctx context.Context) {
	log.Info("run worker")
	w.errors <- w.handle(ctx)
	close(w.errors)
}

// run main handler with child processes
func (w *worker) handle(ctx context.Context) error {
	childs := []childProcess{
		{name: "A"},
		{name: "B"},
	}

	var wg sync.WaitGroup
	wg.Add(len(childs))

	// to simulate unexpected child processes error
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	// run child processes
	for _, child := range childs {
		c := child
		go func() {
			c.run(ctx)
			wg.Done()
		}()
	}

	wg.Wait()

	log.Info("worker handler stopped")

	return errors.New("worker handler error")
}

// notify about error handled
func (w *worker) notify() <-chan error {
	return w.errors
}

type childProcess struct {
	name string
}

// simulate processes, with context listener
func (c *childProcess) run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Infof("Shutdown processes: %s", c.name)
			return
		}
	}
}
