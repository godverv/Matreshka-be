// Code generated by RedSock CLI. DO NOT EDIT

package grpc

import (
	"context"
	"net"
	"net/http"
	"sync"

	"github.com/godverv/matreshka/api"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/godverv/matreshka-be/internal/config"
	"github.com/godverv/matreshka-be/internal/data"
	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

type Server struct {
	serverMux cmux.CMux

	grpcServer *grpc.Server

	serverAddress string
	m             sync.Mutex
}

func NewServer(cfg config.Config, server *api.GRPC, storage data.Data) (*Server, error) {
	srv := grpc.NewServer()

	// Register your servers here
	matreshka_api.RegisterMatreshkaBeAPIServer(srv, &App{
		version: cfg.AppInfo().Version,
		storage: storage,
	})

	return &Server{
		grpcServer: srv,

		serverAddress: ":" + server.GetPortStr(),
	}, nil
}

func (s *Server) Start(_ context.Context) error {
	s.m.Lock()
	defer s.m.Unlock()

	listener, err := net.Listen("tcp", s.serverAddress)
	if err != nil {
		return errors.Wrap(err, "error opening listener")
	}
	s.serverMux = cmux.New(listener)

	go s.startGRPC()

	go s.startGateway()

	go func() {
		err := s.serverMux.Serve()
		if err != nil {
			logrus.Errorf("error service server %s", err)
		}
	}()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.grpcServer.GracefulStop()
	logrus.Infof("Server at %s is stopped", s.serverAddress)

	return nil
}

func (s *Server) startGRPC() {
	grpcListener := s.serverMux.
		MatchWithWriters(
			cmux.HTTP2MatchHeaderFieldSendSettings(
				"content-type",
				"application/grpc",
			))

	err := s.grpcServer.Serve(grpcListener)
	if err != nil {
		logrus.Errorf("error starting grpc server: %s", err)
	}
}

func (s *Server) startGateway() {
	httpListener := s.serverMux.Match(cmux.Any())

	httpMux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard, &runtime.JSONPb{}))

	err := matreshka_api.RegisterMatreshkaBeAPIHandlerFromEndpoint(
		context.Background(),
		httpMux,
		s.serverAddress,
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		})
	if err != nil {
		logrus.Errorf("error registering grpc2http handler: %s", err)
	}
	server := &http.Server{
		Addr:    s.serverAddress,
		Handler: httpMux,
	}

	err = server.Serve(httpListener)
	if err != nil {
		logrus.Errorf("error starting gateway: %s", err)
	}
}
