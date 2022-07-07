package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

type UploadHandler struct {
	UploadDir string
}

func main() {
	uploadHandler := &UploadHandler{
		UploadDir: "/build/bin/upload",
	}

	newSrv := &http.Server{
		Addr:         ":8000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("http server started on: ", newSrv.Addr)

	http.HandleFunc("/list", uploadHandler.listGetFiles)

	log.Fatal(newSrv.ListenAndServe())

}

func (h *UploadHandler) listGetFiles(w http.ResponseWriter, r *http.Request) {
	dirFiles, err := ioutil.ReadDir(h.UploadDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, files := range dirFiles {
		fmt.Fprintln(w,
			files.Name(),
			filepath.Ext(filepath.Ext(files.Name())),
			files.Size())
	}
}
