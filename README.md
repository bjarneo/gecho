# gecho

This service will echo back the the incoming request to you. Often while working on services that are hidden behind a proxy, or multiple proxies, a load balancer, an API gateway or however your infrastructure is built, you might want to know how the end request would look like. I.e. what headers has been removed, manipulated, or added.

## Example by using HTTPIE

Post JSON

```bash
http localhost:1337/echo name="Something" hobby="something else" X-My-Header:gecho --auth username:password
HTTP/1.1 200 OK
Content-Encoding: gzip
Content-Length: 250
Content-Type: text/plain; charset=utf-8
Date: Mon, 18 Jul 2022 07:59:33 GMT
Vary: Accept-Encoding

#### REQUEST HEADERS ####
Connection: keep-alive
Content-Type: application/json
X-My-Header: gecho
Content-Length: 51
User-Agent: HTTPie/2.4.0
Accept: application/json, */*;q=0.5
Authorization: Basic dXNlcm5hbWU6cGFzc3dvcmQ=


#### REQUEST BODY ####
{"name": "Something", "hobby": "something else"}
```

Post a FORM

```bash
http --form localhost:1337/echo file@"test.txt" X-My-Header:gecho --auth username:password
HTTP/1.1 200 OK
Content-Encoding: gzip
Content-Length: 322
Content-Type: text/plain; charset=utf-8
Date: Mon, 18 Jul 2022 08:00:57 GMT
Vary: Accept-Encoding

#### REQUEST HEADERS ####
X-My-Header: gecho
User-Agent: HTTPie/2.4.0
Connection: keep-alive
Content-Length: 196
Authorization: Basic dXNlcm5hbWU6cGFzc3dvcmQ=
Accept: */*
Content-Type: multipart/form-data; boundary=d2a6a71154154fde9698582d2d8b36a8


#### REQUEST BODY ####
--d2a6a71154154fde9698582d2d8b36a8
Content-Disposition: form-data; name="file"; filename="test.txt"
Content-Type: text/plain

This is the file content

--d2a6a71154154fde9698582d2d8b36a8--
```

## Usage

```bash
dotenvjson -h
Usage of dotenvjson:
  -o  string  Save the output to file. -o=file.json
  -p  bool    Pretty print the output. (default false)
  -pt bool    Print to terminal. (default true)
  -y  bool    Transform it to YAML
```

## Build

```bash
go build -o bin/gecho
```

## Releases

Have a look at the releases page for binaries.

## License

See [LICENSE](LICENSE)
