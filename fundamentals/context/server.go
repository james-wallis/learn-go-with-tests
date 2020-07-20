package context

import (
	"context"
	"fmt"
	"net/http"
)

// Store : describes a data store with a Fetch method
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// Server : returns a HTTP handler function that either prints returned data or returns
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return
		}

		fmt.Fprint(w, data)
	}
}
