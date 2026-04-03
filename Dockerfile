FROM golang:1.26-alpine3.22 AS go
WORKDIR /app
COPY go.mod go.sum main.go ./
RUN go mod download \
&& go build -o main /app/main.go

FROM alpine:3.23
WORKDIR /app
COPY --from=go /app/main .
USER 1001
CMD [ "/app/main" ]
