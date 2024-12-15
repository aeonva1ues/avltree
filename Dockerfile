FROM golang:latest

RUN go version
ENV GOPATH=/

COPY . .

RUN go mod download
RUN go build -o avltree ./cmd/main.go

CMD ["./avltree"]
