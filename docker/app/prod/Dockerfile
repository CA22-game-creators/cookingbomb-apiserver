FROM golang:1.16-alpine3.13 as builder

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
WORKDIR /go/src/github.com/CA22-game-creators/cookingbomb-apiserver

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o app

FROM alpine:3.13
WORKDIR /root/
COPY --from=builder /go/src/github.com/CA22-game-creators/cookingbomb-apiserver .

RUN adduser -D myuser && chown -R myuser ./app
USER myuser

EXPOSE 8080
ENTRYPOINT ["./app"]
