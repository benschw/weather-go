FROM golang:1.12.5-alpine3.9 as build-env

RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN apk add --update --no-cache ca-certificates git
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/server

FROM scratch
COPY --from=build-env /go/bin/server /go/bin/server
ENTRYPOINT ["/go/bin/server"]
