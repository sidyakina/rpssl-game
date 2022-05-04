package generator

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

type Generator struct {
	url string
}

func New(url string) *Generator {
	return &Generator{
		url: url,
	}
}

type Response struct {
	RandomNumber int32 `json:"random_number"`
}

func (g *Generator) GetRandomNumber() (int32, error) {
	response, err := http.Get(g.url)
	if err != nil {
		return 0, errors.Wrap(err, "get request")
	}

	if response.Body != nil {
		defer func() {
			err := response.Body.Close()
			if err != nil {
				log.Printf("failed to close response body: %v", err)
			}
		}()
	}

	if response.StatusCode != http.StatusOK {
		return 0, errors.Errorf("status code %v is not OK", response.StatusCode)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, errors.Wrap(err, "failed to read body")
	}

	rawResponse := Response{}

	err = json.Unmarshal(data, &rawResponse)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to unmarshal %s", err)
	}

	return rawResponse.RandomNumber, nil
}
