package api

import (
	context "context"
	"flag"
	"fmt"
	"net"

	gp "github.com/andrei-dascalu/go-shortener/src/grpcserve"
	"github.com/andrei-dascalu/go-shortener/src/shortener"
	"github.com/rs/zerolog/log"
	grpc "google.golang.org/grpc"
)

type ShortenerServer struct {
	gp.UnimplementedShortenerServiceServer
	redirectService shortener.RedirectService
}

func StartServer(redirectService shortener.RedirectService) {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8282))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	gp.RegisterShortenerServiceServer(grpcServer, newServer(redirectService))

	log.Warn().Msg("Started GRPC")
	grpcServer.Serve(lis)
}

func newServer(redirectService shortener.RedirectService) gp.ShortenerServiceServer {
	return ShortenerServer{
		redirectService: redirectService,
	}
}

func (ss ShortenerServer) Short(c context.Context, req *gp.ShortRequest) (*gp.ShortResponse, error) {
	redirect := &shortener.Redirect{
		URL: req.GetURL(),
	}

	err := ss.redirectService.Store(redirect)

	if err != nil {
		log.Error().Err(err).Msg("Failed short")
		return nil, err
	}

	return &gp.ShortResponse{
		URL:       redirect.URL,
		CreatedAt: redirect.CreatedAt,
		Code:      redirect.Code,
	}, nil
}

func (ss ShortenerServer) ShortMany(serv gp.ShortenerService_ShortManyServer) error {
	return nil
}
