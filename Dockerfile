FROM golang:1.13.6-alpine3.11 as builder

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o /go/app -v .

# ********* Final image ********* 

FROM alpine:3.11

WORKDIR /go
COPY --from=builder /go/app .
COPY wait_for_redis.sh .

ENTRYPOINT [ "/go/wait_for_redis.sh", "/go/app" ]