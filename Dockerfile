##
## The build go binary step
##
FROM golang:1.18-alpine as build

WORKDIR /

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

# CGO has to be disabled for alpine
ENV CGO_ENABLED=0

RUN go build -o gecho

RUN chmod +x gecho

##
## The final docker image step
##
FROM gcr.io/distroless/static-debian11

WORKDIR /

COPY --from=build /gecho /gecho

ENV HTTP_PORT=8080

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["./gecho"]
