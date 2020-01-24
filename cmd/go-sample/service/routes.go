package service

import (
	"net/http"

	"github.com/TaylorOno/go-sample/internal/pkg/plex"
	handlers "github.com/TaylorOno/go-sample/internal/pkg/plex/http"
)

func routes(svc plex.Service) http.Handler {
	mux := http.NewServeMux()
	//mux.Handle("/", http2.MakePostListenHandler(svc))
	mux.Handle("/", handlers.MakeListenHandler(svc))
	return mux
}
