package handlers

import (
	"context"
	"encoding/json"
	"ip-pool-manager/ip"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

type putIPpost struct {
	TargetIP  string       `json:"targetIp"`
	IPaddress string       `json:"ip"`
	Detail    putIPdetails `json:"detail"`
}

type putIPdetails struct {
	MACaddress string    `json:"MACaddress"`
	LeaseTime  time.Time `json:"leaseTime"`
	Available  bool      `json:"available"`
}

func CreateNewIPinPool(rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		log.Println(r.Body)

		// Creating a empty user post called "uIP"
		var uPut putIPpost

		// Decodes response JSON into a userPostIP object and catches any errors
		if err := json.NewDecoder(r.Body).Decode(&uPut); err != nil {
			log.Println("ERR: ", err)
			w.WriteHeader(http.StatusBadGateway)
			w.Write([]byte("Cannot decode request")) //nolint:errcheck

			return
		}

		// Checking if user IP value is correct lengh
		if len(uPut.IPaddress) != 15 && len(uPut.IPaddress) != 16 {
			w.WriteHeader(http.StatusBadGateway)
			log.Println(len(uPut.IPaddress))
			resp := uPut.IPaddress + " IP is not correct length . May need to contain a- or na-"
			w.Write([]byte(resp)) //nolint:errcheck
		}

		ctx := context.Background()
		_, err := rdb.Get(ctx, uPut.TargetIP).Result()
		switch {
		case err == redis.Nil: // TODO: Will fail wrapped errors using == use errors.Is
			log.Println("key does not exist")
			log.Println(uPut.TargetIP)
			return
		case err != nil:
			log.Println("Get failed", err)
			return
		}

		tempIPpost := ip.IPpost{
			IPaddress: uPut.IPaddress,
			Detail: ip.IPdetails{
				MACaddress: uPut.Detail.MACaddress,
				LeaseTime:  uPut.Detail.LeaseTime,
				Available:  uPut.Detail.Available,
			},
		}

		newIPencoded := encodeIP(tempIPpost)

		rdb.Rename(ctx, uPut.TargetIP, uPut.IPaddress)
		// Storing user key & value into db
		rdb.Set(ctx, uPut.IPaddress, newIPencoded, 0)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("IP has changed \n")) //nolint:errcheck
	}
}
