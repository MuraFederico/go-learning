package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	res, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	lw := logWriter{}

	io.Copy(lw, res.Body)
}

func (logWriter) Write(b []byte) (int, error) {
	fmt.Println(string(b))
	return len(b), nil
}
