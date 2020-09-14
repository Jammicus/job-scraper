FROM golang:latest

# Allow us to use go modules.
ENV GOPATH ""

COPY . .

RUN go test ./... -v
RUN go build cmd/main.go
RUN chmod +x main

CMD ["./main"]