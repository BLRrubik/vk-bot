FROM golang:1.20.4-alpine

ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o vk-bot ./cmd/main.go

CMD ["./vk-bot"]