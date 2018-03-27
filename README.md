# go-http-format

Formats a stream of http request / response payloads with JSON bodies.

## Install
```sh
go get github.com/travisperson/go-http-format
```

## Usage

```sh
$ go-http-format < payload.http_response
```

Designed to be used with [ipfs-proxy](https://github.com/travisperson/.dot/blob/master/bin/ipfs-proxy) utility.

```sh
tail -f /tmp/out | go-http-format
```

## License

MIT
