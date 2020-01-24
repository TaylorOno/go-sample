package endpoints

import (
	"context"

	"github.com/TaylorOno/go-sample/internal/pkg/plex"
	"github.com/go-kit/kit/endpoint"
)

type HookResponse struct {
	Message string `json:"message"`
}

type PlexHook struct {
	Event   string `json:"event"`
	User    bool   `json:"user"`
	Owner   bool   `json:"owner"`
	Account struct {
		ID    int    `json:"id"`
		Thumb string `json:"thumb"`
		Title string `json:"title"`
	} `json:"Account"`
	Server struct {
		Title string `json:"title"`
		UUID  string `json:"uuid"`
	} `json:"Server"`
	Player struct {
		Local         bool   `json:"local"`
		PublicAddress string `json:"publicAddress"`
		Title         string `json:"title"`
		UUID          string `json:"uuid"`
	} `json:"Player"`
	Metadata struct {
		LibrarySectionType    string `json:"librarySectionType"`
		RatingKey             string `json:"ratingKey"`
		Key                   string `json:"key"`
		ParentRatingKey       string `json:"parentRatingKey"`
		GrandparentRatingKey  string `json:"grandparentRatingKey"`
		GUID                  string `json:"guid"`
		ParentGUID            string `json:"parentGuid"`
		GrandparentGUID       string `json:"grandparentGuid"`
		Type                  string `json:"type"`
		Title                 string `json:"title"`
		GrandparentKey        string `json:"grandparentKey"`
		ParentKey             string `json:"parentKey"`
		GrandparentTitle      string `json:"grandparentTitle"`
		ParentTitle           string `json:"parentTitle"`
		Summary               string `json:"summary"`
		Index                 int    `json:"index"`
		ParentIndex           int    `json:"parentIndex"`
		LastViewedAt          int    `json:"lastViewedAt"`
		Year                  int    `json:"year"`
		Thumb                 string `json:"thumb"`
		Art                   string `json:"art"`
		GrandparentThumb      string `json:"grandparentThumb"`
		GrandparentArt        string `json:"grandparentArt"`
		OriginallyAvailableAt string `json:"originallyAvailableAt"`
		AddedAt               int    `json:"addedAt"`
		UpdatedAt             int    `json:"updatedAt"`
	} `json:"Metadata"`
}

func MakeListenEndpoint(s plex.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		switch {
		case request == nil:
			s.GetListen()
			return HookResponse{Message: "GET OK"}, nil
		case request != nil:
			s.PostListen(request)
			return HookResponse{Message: "POST OK"}, nil
		}
		return HookResponse{Message: "ERROR"}, nil
	}
}
