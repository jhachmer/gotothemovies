package media

import (
	"fmt"
	"github.com/jhachmer/gotothemovies/pkg/config"
	"github.com/jhachmer/gotothemovies/pkg/util"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type OmdbRequest interface {
	SendRequest() (*Movie, error)
}

type OmdbIDRequest struct {
	imdbID string
}

type OmdbTitleRequest struct {
	title string
	year  int
}

func NewOmdbIDRequest(imdbID string) *OmdbIDRequest {
	return &OmdbIDRequest{
		imdbID: imdbID,
	}
}

func NewOmdbTitleRequest(title string, year int) *OmdbTitleRequest {
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
	reqURL, err := url.Parse(fmt.Sprintf("http://www.omdbapi.com/?apikey=%s", config.Envs.OmdbApiKey))
	if err != nil {
		return "", err
	}
	switch v := r.(type) {
	case OmdbTitleRequest:
		values := reqURL.Query()
		values.Add("t", v.title)
		values.Add("y", strconv.Itoa(v.year))
		if err != nil {
			return "", err
		}
		reqURL.RawQuery = values.Encode()
		return reqURL.String(), nil
	case OmdbIDRequest:
		values := reqURL.Query()
		values.Add("i", v.imdbID)
		if err != nil {
			return "", err
		}
		reqURL.RawQuery = values.Encode()
		return reqURL.String(), nil
	default:
		return "", fmt.Errorf("no valid request type")
	}
}

func decodeRequest(requestURL string) (Movie, error) {
	var mov Movie
	req, err := getRequest(requestURL)
	if err != nil {
		return mov, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return mov, err
	}
	mov, err = util.Decode[Movie](res)
	if err != nil {
		return mov, err
	}
	if !checkIfResponseTrue(mov) {
		return Movie{}, fmt.Errorf("response value is false")
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

func checkIfResponseTrue(mov Movie) bool {
	return strings.ToLower(mov.Response) == "true"
}
