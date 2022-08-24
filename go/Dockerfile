FROM golang:1.16-alpine

ENV GOPATH=/go

COPY . "/go/src/github.com/user/api"
WORKDIR "/go/src/github.com/user/api"

RUN go mod download

EXPOSE 8080

# Run the executable
CMD ["go", "run", "./cmd/api/main.go"]