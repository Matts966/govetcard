package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetZIP(module, version string) string {
	url := strings.Join([]string{
		"https://proxy.golang.org",
		module,
		"@v",
		strings.Join([]string{
			version,
			".zip",
		}, ""),
	}, "/")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(byteArray)
}
