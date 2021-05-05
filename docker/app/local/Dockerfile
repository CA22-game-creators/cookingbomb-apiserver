FROM golang:1.16-alpine3.13

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
WORKDIR /go/src/github.com/CA22-game-creators/cookingbomb-apiserver

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build .

RUN go get github.com/pilu/fresh

EXPOSE 8080
CMD ["fresh"]
