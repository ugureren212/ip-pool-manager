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

// AllAvailbleIPs Return all available IP's (IP's that start with "a-" not "na-")
func AllAvailbleIPs(rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allAvailbleIPs := getAllIPs(rdb)
		// Marshal data to return result to user
		strAllAvailbleIPs, err := json.Marshal(allAvailbleIPs)
		if err != nil {
			log.Println(err)
		}

		// Setting response headers and content
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(strAllAvailbleIPs) //nolint:errcheck
	}
}

// Returns all IP's that start with "a-" and "available = true"
func getAllIPs(rdb *redis.Client) []ip.IPpost {
	ctx := context.Background()
	allIPs := []ip.IPpost{}

	// Loop used to iterate other each key that stars with "a-" in DB
	iter := rdb.Scan(ctx, 0, "a-*", 0).Iterator()
	for iter.Next(ctx) {
		// Storing each IP in DB
		foundIP, err := rdb.Get(ctx, iter.Val()).Result()
		if err != nil {
			log.Println("IP not found. ERR: ", err)
			continue
		}

		// Gob to Struct
		bufDe := &bytes.Buffer{}
		bufDe.WriteString(foundIP)

		// Decode returned Gob format into IP struct
		var dataDecode ip.IPpost
		if err := gob.NewDecoder(bufDe).Decode(&dataDecode); err != nil {
			log.Println(err)
			continue
		}

		allIPs = append(allIPs, dataDecode)
	}

	allAvailbleIPs := findAvailbleIP(allIPs)
	return allAvailbleIPs
}

// Returns IP's  with "available = true"
func findAvailbleIP(allIPs []ip.IPpost) []ip.IPpost {
	allAvailbleIPs := []ip.IPpost{}

	// Check to see if IP is available or not
	for _, IP := range allIPs {
		if IP.Detail.Available {
			allAvailbleIPs = append(allAvailbleIPs, IP)
		} else {
			log.Println(IP.IPaddress, " is not availble right now")
		}
	}

	return allAvailbleIPs
}
