package service

import (
	"flag"
	"net"
	"net/http"
	"os"

	"github.com/TaylorOno/go-sample/internal/pkg/plex"

	"github.com/go-kit/kit/log"
)

func Run() error {
	fs := flag.NewFlagSet("addsvc", flag.ExitOnError)
	var (
		httpAddr = fs.String("http-addr", ":8000", "HTTP listen address")
	)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	server, err := net.Listen("tcp", *httpAddr)
	if err != nil {
		logger.Log("transport", "HTTP", "during", "Listen", "err", err)
		return err
	}

	plexService := plex.NewPlexService()

	routes := routes(&plexService)
	err = http.Serve(server, routes)
	if err != nil {
		logger.Log("transport", "HTTP", "during", "Listen", "err", err)
		return err
	}

	return nil
}
