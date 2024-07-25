FROM golang:1.22.0 as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ../*.go ./
RUN CGO_ENABLED=0 go build -o /docker-gs-ping
EXPOSE 8080
# Run
CMD ["/docker-gs-ping"]