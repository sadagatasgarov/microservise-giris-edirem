package main

import (
	"fmt"
	"net/http"
)

func main() {
	store := NewMemoryStore()
	var (
		svc = NewInvoiceAggregator(store)
	)
	// burada dekorator patern dedi arasdiracagam
	http.HandleFunc("/aggregate", handleAggregate(svc))
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaa")
}

func handleAggregate(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
