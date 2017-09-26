package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/marthjod/binquiry/convert"
	"github.com/marthjod/binquiry/getter"
)

const (
	urlPrefix = "http://dev.phpbin.ja.is/ajax_leit.php"
	indexPage = "index.html"
)

var (
	converter = convert.JSONConverter{}
)

func handleQuery(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.ServeFile(w, r, indexPage)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(converter.ConvertBytes(&getter.Getter{URLPrefix: urlPrefix}, query))
}

func main() {
	var (
		port = flag.Int("port", 4242, "Local port to listen on")
	)

	flag.Parse()

	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.HandleFunc("/", handleQuery)

	fmt.Printf("Serving on port %d\n", *port)
	err := http.ListenAndServe(":"+strconv.Itoa(*port), nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
