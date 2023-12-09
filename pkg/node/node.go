package node

import (
	"golang.org/x/net/html"
	"net/http"
)

func Get(url string) (n *html.Node, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	n, err = html.Parse(resp.Body)

	return
}
