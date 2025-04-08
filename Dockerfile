FROM golang:1.23.0

RUN mkdir /taskAPI

WORKDIR /taskAPI

COPY ./ ./

RUN go env -w GO111MODULE=on

RUN go mod download

RUN go build ./cmd/main.go

CMD ["./main"]