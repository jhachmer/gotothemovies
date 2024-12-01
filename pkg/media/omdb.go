package media

import (
	"fmt"
	"github.com/jhachmer/gotothemovies/internal/config"
	"github.com/jhachmer/gotothemovies/internal/server"
	"net/http"
	"net/url"
)

type OmdbRequest interface {
	SendRequest() (*Movie, error)
}

type OmdbIDRequest struct {
	imdbID string
}

type OmdbTitleRequest struct {
	title string
	year  string
}

func NewOmdbIDRequest(imdbID string) *OmdbIDRequest {
	return &OmdbIDRequest{
		imdbID: imdbID,
	}
}

func NewOmdbTitleRequest(title, year string) *OmdbTitleRequest {
	return &OmdbTitleRequest{
		title: title,
		year:  year,
	}
}

func (r OmdbTitleRequest) SendRequest() (*Movie, error) {
	requestURL, err := makeRequestURL(r)
	if err != nil {
		return nil, err
	}
	var mov Movie
	mov, err = decodeRequest(requestURL)
	if err != nil {
		return nil, err
	}
	return &mov, nil
}

func (r OmdbIDRequest) SendRequest() (*Movie, error) {
	requestURL, err := makeRequestURL(r)
	if err != nil {
		return nil, err
	}
	var mov Movie
	mov, err = decodeRequest(requestURL)
	if err != nil {
		return nil, err
	}
	return &mov, nil
}

func makeRequestURL(r OmdbRequest) (string, error) {
	baseURL := "http://www.omdbapi.com/?apikey="
	apiKey := config.Envs.OmdbApiKey
	switch v := r.(type) {
	case OmdbTitleRequest:
		joinedURL, err := url.JoinPath(baseURL, apiKey, "&t=", v.title, "&y=", v.year)
		if err != nil {
			return "", err
		}
		return joinedURL, nil
	case OmdbIDRequest:
		joinedURL, err := url.JoinPath(baseURL, apiKey, "&i=", v.imdbID)
		if err != nil {
			return "", err
		}
		return joinedURL, nil
	default:
		return "", fmt.Errorf("no valid request type")
	}
}

func decodeRequest(requestURL string) (Movie, error) {
	var mov Movie
	res, err := getRequest(requestURL)
	if err != nil {
		return mov, err
	}
	mov, err = server.Decode[Movie](res)
	if err != nil {
		return mov, err
	}
	return mov, nil
}

func getRequest(requestURL string) (*http.Request, error) {
	res, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}
