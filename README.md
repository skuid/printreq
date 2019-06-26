# printreq

Echoes out a request. Useful for dumping a data dog request locally, but it can be use to dump any request made to it.
 
It will always return a `200:OK` response unless it encounters an error.

For each call, it will dump:
- **METHOD**: GET, PUT, POST, DELETE, etc...
- **PATH**: What path was called
- **HEADERS**: All of the headers on the request
- **BODY**: The request body

Currently it does not try to parse the body based on content type. It just dumps whatever it was given.

There are two optional flags that can be passed to the command:
- **host**: The host or IP address to serve from. Defaults to `127.0.0.1`
- **port**: The port to listen on. Defaults to `8126` which is the default port for datadogs TCP agent

# Build

Build using `go`

```bash
go build
```

# Run

After building, just execute the binary.

```bash
./printreq
```

```bash
./printreq -host localhost -port 18222
```
