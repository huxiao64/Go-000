package main

import(
	"golang.org/x/sync/errgroup"
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
	"time"
	"fmt"
)

type Server struct{
	srv *http.Server
}

func NewServer() *Server {
	mux := http.NewServerMux()
	mux.Handle("/", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request){
			fmt.Println("receive request")
			time.Sleep(100*time.Second)
		},
	))
	srv := &http.Server{
		Addr: ":8080"
		Handle: mux,
	}
	return &Server{srv: srv}
}

func (s *Server) Start() error {
	fmt.Println("[HTTP] Listening on: %s\n", s.srv.Addr)
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func main() {
	stop := make(chan struct{})
	g, ctx := errgroup.WithContext(context.Background())
	srv := NewServer()

	g.Go(func() error {
		fmt.Println("start http")
		go func() {
			<-ctx.Done()
			fmt.Println("http ctx done")
			ctx2, cnacel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := srv.Shutdown(ctx2); err!=nil {
				fmt.Println("Server forced to be shutdown:", err)
			}
			stop <- struct{}{}
			fmt.Println("server exiting")
		}()
		return srv.Start
	})

	g.Go(func() error {
		quit:=make(chan os.Signal)
		signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM,syscall.SIGQUIT,os.Interrupt)
		for {
			fmt.Println("waiting for quit signal")
			select {
			case <-ctx.Done():
				fmt.Println("signal ctx done")
				return ctx.Err()
			case <-quit:
				return errors.New("receive quit signal")
			}
		}
	})

	g.Go(func() error{
		for{
			select {
			case <-ctx.Done():
				fmt.Println("background ctx done")
				return ctx.Err()
			default:
				fmt.Println("do something")
				time.Sleep(1*time.Second)
			}
		}
	})
	err := g.Wait()
	fmt.Println(err)
	<-stop
	fmt.Println("server completely stopped!")
}
