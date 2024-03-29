# gecho

This HTTP service will echo back the incoming request to you. Often while working on services that are hidden behind a proxy, or multiple proxies, a load balancer, an API gateway or however your infrastructure is built, you might want to know how the end request would look like. I.e. what headers has been removed, manipulated, or added.

## Use cases

Without adding too much debugging to your application, you can use this service as a drop in replacement for the app you are debugging.

* It has wildcard support for routes, so you can just change your application host with the host of this service
* What is the end request headers after going through gateways/proxies/load balancers
* Check the form data request body
* Check JSON data request body
* How about the file you are trying to upload?
* Check the AUTH request headers

See some examples:

## Examples by using HTTPIE

Post JSON

```bash
# Post to whatever route you wish. I will be using echo.
http localhost:1337/echo name="Something" hobby="something else"

X-My-Header:gecho --auth username:password
HTTP/1.1 200 OK
Content-Encoding: gzip
Content-Length: 250
Content-Type: text/plain; charset=utf-8
Date: Mon, 18 Jul 2022 07:59:33 GMT
Vary: Accept-Encoding

=========================
|--- REQUEST HEADERS ---|
=========================
Connection: keep-alive
Content-Type: application/json
X-My-Header: gecho
Content-Length: 51
User-Agent: HTTPie/2.4.0
Accept: application/json, */*;q=0.5
Authorization: Basic dXNlcm5hbWU6cGFzc3dvcmQ=


=========================
|---- REQUEST  BODY ----|
=========================
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

=========================
|--- REQUEST HEADERS ---|
=========================
X-My-Header: gecho
User-Agent: HTTPie/2.4.0
Connection: keep-alive
Content-Length: 196
Authorization: Basic dXNlcm5hbWU6cGFzc3dvcmQ=
Accept: */*
Content-Type: multipart/form-data; boundary=d2a6a71154154fde9698582d2d8b36a8


=========================
|---- REQUEST  BODY ----|
=========================
--d2a6a71154154fde9698582d2d8b36a8
Content-Disposition: form-data; name="file"; filename="test.txt"
Content-Type: text/plain

This is the file content

--d2a6a71154154fde9698582d2d8b36a8--
```

Example by using fetch from the frontend

```js
const req = await fetch("http://localhost:1337/echo?abc=def");

const text = await req.text();

console.log(text)
/*
=========================
|--- REQUEST DETAILS ---|
=========================
URI: /echo?abc=def

=========================
|--- REQUEST HEADERS ---|
=========================
Accept-Language: nb-NO,nb;q=0.9,no;q=0.8,nn;q=0.7,en-US;q=0.6,en;q=0.5,da;q=0.4,sv;q=0.3
Cache-Control: no-cache
Connection: keep-alive
Dnt: 1
Pragma: no-cache
Sec-Ch-Ua: ".Not/A)Brand";v="99", "Google Chrome";v="103", "Chromium";v="103"
User-Agent: Woohoo
*/
```

## Usage

```bash
HTTP_PORT=1337 gecho
```

Environment variables
* `HTTP_PORT` default `8080`

## Docker

The image is available at `bjarneo/gecho:latest`.
`docker pull bjarneo/gecho:latest`

```bash

# This example we override the HTTP_PORT to port 1337 as a show case
docker run -it -p 8080:1337 -e HTTP_PORT=1337 --rm --name gecho bjarneo/gecho:latest

# Now test it with curl
curl localhost:8080/echo
=========================
|--- REQUEST DETAILS ---|
=========================
URI: /echo

=========================
|--- REQUEST HEADERS ---|
=========================
Accept: */*
User-Agent: curl/7.79.1
```

## Kubernetes

I have added an example manifest file `gecho.k8s.yaml`.

```bash
kubectl apply -f gecho.k8s.yaml
```

## Build

```bash
go build -o bin/gecho
```

## Releases

Have a look at the releases page for binaries.

## Features?

Please create an issue

## License

See [LICENSE](LICENSE)
