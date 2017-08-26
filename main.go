package main

import (
	"io"
	"net/http"
	"time"

	"fmt"
)

func main() {
	http.HandleFunc("/docker", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello Docker:"+time.Now().String())
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Err:%v", err)
	}
}
