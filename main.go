package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// curl --data-binary @path/to/file.bin http://192.168.155.228:9090
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Create("file.bin")
		if err != nil {
			http.Error(w, "failed to create file", http.StatusInternalServerError)
			return
		}
		h := sha256.New()
		_, err = io.Copy(io.MultiWriter(f, h), r.Body)
		if err != nil {
			http.Error(w, "failed to write file", http.StatusInternalServerError)
			return
		}
		sum := fmt.Sprintf("%x", h.Sum(nil))
		fmt.Println(sum)
		io.WriteString(w, sum)
	})
	http.ListenAndServe(":9090", h)
}
