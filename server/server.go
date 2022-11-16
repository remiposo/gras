package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/jmoiron/sqlx"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	srv *http.Server
}

func NewServer(port string, db *sqlx.DB) *Server {
	srv := &http.Server{
		Addr:    port,
		Handler: NewRouter(db),
	}
	return &Server{
		srv: srv,
	}
}

func (s *Server) Run(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})

	<-ctx.Done()

	if err := s.srv.Shutdown(context.Background()); err != nil {
		return err
	}
	return eg.Wait()
}
