package handlers

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"ip-pool-manager/ip"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

func AddToIPtoPool(rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")

		// Creating a empty user post called "u"
		var u ip.IPpost

		// Decodes response JSON into a userPostIP object and catches any errors
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			log.Println("ERR: ", err)
			w.WriteHeader(http.StatusBadGateway)
			w.Write([]byte("Cannot decode request")) //nolint:errcheck
			return
		}

		// Checking if user IP value is correct lengh
		if len(u.IPaddress) != 15 {
			w.WriteHeader(http.StatusBadGateway)
			resp := "IP is not correct length " + u.IPaddress
			w.Write([]byte(resp)) //nolint:errcheck
			return
		}

		ctx := context.Background()

		// These print are for debug purposes
		log.Printf("Values of new IP. IP address : %v. Value: %v", u, u.IPaddress)

		encodedU := encodeIP(u)
		// Storing user key & value into db
		rdb.Set(ctx, u.IPaddress, encodedU, 0)

		userResponse := u.IPaddress + "IP has been added to DB"
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(userResponse)) //nolint:errcheck
	}
}

// Encodes IP into glob format
func encodeIP(ip ip.IPpost) string {
	// struct to Gob
	bufEn := &bytes.Buffer{}
	if err := gob.NewEncoder(bufEn).Encode(ip); err != nil {
		log.Println(err)
	}
	BufEnString := bufEn.String()

	return BufEnString
}
