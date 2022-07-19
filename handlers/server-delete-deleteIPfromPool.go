package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

// DeleteIPfromPool Deletes the specified IP from the IP pool
func DeleteIPfromPool(rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		// Storing "key" url param that contains IP key/id
		param := r.URL.Query().Get("key")

		// Delete specified IP key from db
		if param == "" {
			log.Println("Empty URL parameter was passed. Need valid URL key param.")
			w.Write([]byte("Empty URL parameter")) //nolint:errcheck
			w.WriteHeader(http.StatusBadRequest)
		}
		// If IP doesn't exist throw an err
		if err := rdb.Del(ctx, param).Err(); err != nil {
			http.Error(w, http.StatusText(404), 404)
			log.Println(param, " Error: ", err)
		}
		// Send back ok status response and "User Deleted" message
		responseMsg := param + " IP deleted \n"
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(responseMsg)) //nolint:errcheck
	}
}
