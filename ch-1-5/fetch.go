package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		if strings.HasPrefix(url, "http://") == false && strings.HasPrefix(url, "https://") == false {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}
		// b, err := ioutil.ReadAll(resp.Body)

		// statusCode := resp.StatusCode(resp.StatusCode)
		if resp.StatusCode != 200 {
			fmt.Println("Page cannot be loaded")
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: read %s: %v\n", url, err)
		}
		// fmt.Printf("%s", b)
	}
}
