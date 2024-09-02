FROM golang:1.23-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=1  \
    GOOS=linux 

RUN apk add alpine-sdk

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Install Compile plug

RUN apk -U add ca-certificates
RUN apk update && apk upgrade && apk add pkgconf git bash build-base sudo
RUN git clone https://github.com/edenhill/librdkafka.git && cd librdkafka && ./configure --prefix /usr && make && make install


COPY . ./
RUN go build -tags musl --ldflags "-extldflags -static" -o lily .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/lily .

EXPOSE 8080

CMD ["./lily"]