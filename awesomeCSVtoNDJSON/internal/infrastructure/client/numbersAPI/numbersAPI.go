package numbersAPI

import (
	"context"
	"fmt"
	"gitlab.neoway.com.br/awesomeCSVtoNDJSON/internal/infrastructure/client/restClient"
	"io"
	"io/ioutil"
)

type NumbersAPI struct {
	URL string
}

const (
	URLNumbersAPI = "http://numbersapi.com/random/math"
)

func NewNumbersAPI() *NumbersAPI {
	return &NumbersAPI{
		URL: URLNumbersAPI,
	}
}

func (n *NumbersAPI) Request() (string, error) {
	ctx := context.Background()

	resp, err := restClient.Get(ctx, n.URL)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		err := fmt.Errorf("url: [%s] status: [%d]", URLNumbersAPI, resp.StatusCode)
		return "", err
	}

	return string(bodyBytes), nil
}
