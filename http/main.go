package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if resp.StatusCode == 200 {
		// data := make([]byte, 1e6)
		// resp.Body.Read(data)
		// fmt.Println(n, e)
		// fmt.Println(string(data))

		lw := logWriter{}
		io.Copy(lw, resp.Body)
	}

	resp.Body.Close()
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("just wrote this many bytes:", len(bs))
	return len(bs), nil
}
