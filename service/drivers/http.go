package drivers

import (
	"io/ioutil"
	"net/http"
	"bytes"
)

type HTTPService struct {
	URL string
}

func (h *HTTPService) FetchByte(payload *[]byte) ([]byte, error) {
	//resp, err := http.Post ("https://api.hollowverse.com/graphql","application/json", strings.NewReader(string(*payload)))
	resp, err := http.Post (h.URL,"application/json", bytes.NewBuffer(*payload))
	if err != nil {
		// handle error
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println (string(body), h)
	return body, err
}

/*func (h *HTTPService) Fetch(query string, variables interface{}) ([]byte, error) {
	return h.FetchByte([]byte(query))
}*/