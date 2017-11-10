package mypkg

import (
	"io/ioutil"
	"net/http"
)

func download(url string) (ret string, err error) {
	response, err := http.Get(url)
	if err != nil {
		return
	}

	defer response.Body.Close()
	ret, err = ioutil.ReadAll(response.Body)
}
