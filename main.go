package Service

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func SendRecv(url string, data []byte) ([]byte, error) {
	r := bytes.NewReader(data)
	resp, err := http.Post(url, "application/json", r)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	data, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return data, nil
}
