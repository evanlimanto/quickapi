package crawlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/evanlimanto/quickapi/src/database"
)

type Crawler interface {
	LoginAndGetAccounts() ([]database.Account, error)
}

type RequestOptions struct {
	uri  string
	body map[string]interface{}
}

func requestJson(options RequestOptions) ([]byte, error) {
	jsonBody, err := json.Marshal(options.body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", options.uri, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
