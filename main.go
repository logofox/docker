package main

import (
	"io"
	"net/http"
	"time"

	"fmt"
)

func main() {
	http.HandleFunc("/docker", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello Docker! "+time.Now().String())
	})

	http.HandleFunc("hello", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, Welcome! "+time.Now().Format("20060102 15:04:05"))
	})

	http.HandleFunc("info", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
			io.WriteString(w, "Get Info! "+time.Now().Format("20060102 15:04:05"))
		} else if req.Method == http.MethodPost {
			io.WriteString(w, "Post Info! "+time.Now().Format("20060102 15:04:05"))
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Err:%v", err)
	}
}
