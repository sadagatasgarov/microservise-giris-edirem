package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/sadagatasgarov/toll-calc/types"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "the listen address of the HTTP server")
	flag.Parse()
	store := NewMemoryStore()
	var (
		svc = NewInvoiceAggregator(store)
	)
	// burada dekorator patern dedi arasdiracagam
	http.HandleFunc("/aggregate", handleAggregate(svc))
	http.ListenAndServe(*listenAddr, nil)
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaa")
}

func handleAggregate(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return 
		}
	}
}
