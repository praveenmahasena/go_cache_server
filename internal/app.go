// Package internal ...
package internal

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/praveenmahasena/gocacheserver/internal/dbms"
	s "github.com/praveenmahasena/gocacheserver/internal/server"
	utils "github.com/praveenmahasena/gocacheserver/internal/util"
)

// Run to start app
func Run() error {
	port := ""
	flag.StringVar(&port, "port", ":42069", "port id to start server")
	flag.Parse()

	messageCh := make(chan utils.Node)
	responseCh := make(chan utils.Node)
	server := s.NewServer(port)
	db := dbms.New()

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	go func(cancelFunc context.CancelFunc) {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		cancelFunc()
	}(cancelFunc)

	if err := server.Listen(); err != nil {
		return err
	}

	wg := &sync.WaitGroup{}
	wg.Add(3)
	go server.Accept(ctx, messageCh, wg)
	go db.Job(messageCh, responseCh, wg)
	go s.Response(responseCh, wg)

	wg.Wait()
	return nil
}
