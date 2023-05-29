package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iriscompanyio/awslex-bot/internal/core"
	"github.com/iriscompanyio/awslex-bot/pkg/middleware"
)

type Server struct {
	engine          *gin.Engine
	httpAddr        string
	ShutdownTimeout time.Duration
}

func NewServer(ctx context.Context, host string, port uint, shutdownTimeout time.Duration) (context.Context, Server) {
	srv := Server{
		engine:          gin.New(),
		httpAddr:        fmt.Sprintf("%s:%d", host, port),
		ShutdownTimeout: shutdownTimeout,
	}
	srv.registerRoutes()
	srv.engine.Use(middleware.CORSMiddleware())
	return serverContext(ctx), srv
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}

func (s *Server) RunServer(ctx context.Context) error {
	log.Println("Server running on", s.httpAddr)

	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server shut down", err)
		}
	}()

	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.ShutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutDown)
}

func (s *Server) registerRoutes() {
	s.engine.POST("/bot", core.WebhookHandler())
}
