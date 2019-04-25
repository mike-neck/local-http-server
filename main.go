package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 4000, "server http port, default: 4000")
	directory := flag.String("dir", "", "root directory, default: current directory")

	flag.Parse()

	dir, err := determineRootDir(directory)
	if err != nil {
		log.Fatalln("terminate server")
	}

	root := http.Dir(*dir)
	fileServer := http.FileServer(root)

	addr := fmt.Sprintf(":%d", *port)

	log.Println("starting file server on", addr, "on directory", *dir)

	err = http.ListenAndServe(addr, fileServer)
	if err != nil {
		log.Fatalln("failed to start http file server", err)
	}
}

func determineRootDir(directory *string) (*string, error) {
	if directory == nil {
		dir, err := os.Getwd()
		if err != nil {
			log.Println("cannot get current directory", err)
			return nil, err
		}
		return &dir, nil
	} else {
		return directory, nil
	}
}
