package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/TaylorOno/go-sample/internal/pkg/plex"
	"github.com/TaylorOno/go-sample/internal/pkg/plex/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeListenHandler(svc plex.Service) *httptransport.Server {
	return httptransport.NewServer(
		endpoints.MakeListenEndpoint(svc),
		decodeListenRequest,
		encodeResponse,
	)
}

func decodeListenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoints.PlexHook
	switch method := r.Method; method {
	case http.MethodPost:
		r.ParseMultipartForm(1 << 20) //1Mb
		payload := r.MultipartForm.Value["payload"][0]
		form := r.MultipartForm
		value := form.Value
		fmt.Println(value)
		file := form.File["thumb"]
		fmt.Println(file[0])

		err := json.NewDecoder(strings.NewReader(payload)).Decode(&request)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	case http.MethodGet:
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
