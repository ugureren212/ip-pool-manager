# ip-pool-manager

This ip-pool-manger Micro-Service manages IP addresses that are typically used to reserve resources for specific users or groups, or to ensure that certain ranges are used for specific purposes.

![overview](https://github.com/UErenReply/ip-pool-manager/blob/main/documentation/ipPool.jpg)

## Installation

```bash
git clone https://github.com/UErenReply/ip-pool-manager
go build
```

## Usage

First run the server:

```bash
./server
```

or

```bash
go run server.go --port 8080 --address 0.0.0.0 --redisPort 6378 --redisAddress 0.0.0.0
```

Use client:

```bash
go run client.go
```

## Testing

Run Unitests

```bash
go test
```

### List of available curl requests

```bash
curl "localhost:3000/allAvailbleIPs"

curl "localhost:3000/getIP?key=a-185.9.249.220"

curl -X DELETE "localhost:3000/deleteIPfromPool?key=na-102.131.46.22"

curl -X POST -H 'content-type: application/json' --data '{"ip":"a-222.2.222.222","detail":{"MACaddress":"89-43-5F-60-DC-76","leaseTime":"2021-12-13T11:11:52.106975Z","available":true}}' http://localhost:3000/addIPtoPool

curl -X PUT -H "Content-Type: application/json" -d '{"targetIp":"a-185.9.249.220","ip":"na-185.9.249.220","detail":{"MACaddress":"11-11-11-11-11-","leaseTime":"2021-12-13T11:11:52.106975Z","available":true}}' http://localhost:3000/createNewIPpool

curl "localhost:3000/healthz"

curl "localhost:3000/readyz"
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
[![Docker](https://github.com/net-reply-future-networks/ip-pool-manager/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/net-reply-future-networks/ip-pool-manager/actions/workflows/docker-publish.yml)
