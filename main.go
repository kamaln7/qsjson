package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
)

var listenAddress = flag.String("listenAddress", "localhost:3000", "address to run the HTTP server on")

func flatten(inp map[string][]string) map[string]string {
	ret := make(map[string]string, len(inp))
	for k, v := range inp {
		ret[k] = v[0]
	}

	return ret
}

func marshalQueryString(w http.ResponseWriter, r *http.Request) {
	obj, err := json.Marshal(flatten(r.URL.Query()))
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("internal server error"))
		return
	}

	w.Write(obj)
}

func main() {
	flag.Parse()
	http.HandleFunc("/", marshalQueryString)
	log.Printf("listening on %s\n", *listenAddress)
	err := http.ListenAndServe(*listenAddress, nil)
	if err != nil {
		log.Panic(err)
	}
}
