package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	var maxsize int
	var maxSizeName string

	for range os.Args[1:] {
		datafromchannel := <-ch
		parts := strings.Fields(datafromchannel)
		// fmt.Println(parts[1])
		size, err := strconv.Atoi(parts[1])
		if err == nil && size > maxsize {
			maxSizeName = parts[2]
		}
		// fmt.Println(parts[2])
		fmt.Println(datafromchannel)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	fmt.Fprintf(os.Stdout, "The biggest size of a page is on site: %s\n", maxSizeName)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
