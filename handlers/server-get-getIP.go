package handlers

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"ip-pool-manager/ip"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

// GetIP Returns specified IP's
func GetIP(rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		ctx := context.Background()

		// Storing "key" url param that contains IP key/id
		param := r.URL.Query().Get("key")

		// Check if URL param is empty or if specified IP is not available
		if param == "" {
			log.Println("Empty URL parameter")
			w.Write([]byte("Empty URL parameter")) //nolint:errcheck
			w.WriteHeader(http.StatusBadRequest)
			return
		} else if strings.HasPrefix(param, "na-") {
			log.Println("Must select available IP")
			w.Write([]byte("Must select available IP")) //nolint:errcheck
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Retrieving IP stored in DB
		val, err := rdb.Get(ctx, param).Result()
		if err != nil {
			log.Println("Cannot find IP")
			w.Write([]byte("Cannot find IP")) //nolint:errcheck
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Decode returned Gob format to IP Struct
		bufDe := &bytes.Buffer{}
		bufDe.WriteString(val)
		var valDecode ip.IPpost
		if err := gob.NewDecoder(bufDe).Decode(&valDecode); err != nil {
			log.Println(err)
		}

		// returning copy of IP with IPaddress == na and availability = false
		returnIP := ip.IPpost{
			IPaddress: strings.Replace(valDecode.IPaddress, "a", "na", 1),
			Detail: ip.IPdetails{
				MACaddress: "89-43-5F-60-DC-76",
				LeaseTime:  time.Now(),
				Available:  false,
			},
		}

		// Convert IP struct into Gob format to store in DB
		bufEn := &bytes.Buffer{}
		if err := gob.NewEncoder(bufEn).Encode(returnIP); err != nil {
			log.Println(err)
		}
		returnIPdecode := bufEn.String()

		// Storing user key & value into db
		rdb.Set(ctx, returnIP.IPaddress, returnIPdecode, 0)

		// If IP doesn't exist throw an err
		if err := rdb.Del(ctx, valDecode.IPaddress).Err(); err != nil {
			log.Println(param, "Cannot delete original IP: ", err)
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Cannot delete original IP")) //nolint:errcheck
		}

		// Converting Struct into JSON byte to return to user
		responseIP, err := json.Marshal(returnIP)
		if err != nil {
			log.Println(err)
		}

		// Send back ok status response and specified IP
		w.WriteHeader(http.StatusOK)
		w.Write(responseIP) //nolint:errcheck
	}
}
