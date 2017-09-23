package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"errors"
	"github.com/marthjod/binquiry/convert"
	"github.com/marthjod/binquiry/getter"
)

var (
	converter = convert.JSONConverter{}
	urlPrefix = "http://dev.phpbin.ja.is/ajax_leit.php"
)

func getQuery(r *http.Request) (string, error) {
	q := r.URL.Query().Get("q")
	if q == "" {
		return q, errors.New("empty query")
	}
	return q, nil
}

func handleQuery(w http.ResponseWriter, r *http.Request) {
	query, err := getQuery(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(converter.ConvertBytes(&getter.Getter{URLPrefix: urlPrefix}, query))
}

func main() {
	var (
		port = flag.Int("port", 4242, "Local port to listen on")
	)

	flag.Parse()

	http.HandleFunc("/", handleQuery)

	fmt.Printf("Serving on port %d\n", *port)
	err := http.ListenAndServe(":"+strconv.Itoa(*port), nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
