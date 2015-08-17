package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "0.0.1"

var bind = flag.String("bind", ":80", "bind address")
var showVersion = flag.Bool("version", false, "print version and exit")

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://" + r.Host + r.URL.Path, http.StatusMovedPermanently)
}

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Println(version)
		os.Exit(2)
	}
	http.HandleFunc("/", redirectHandler)
	log.Printf("starting on %v", *bind)
	log.Fatal(http.ListenAndServe(*bind, nil))
}
