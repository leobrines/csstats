# build stage
FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build

# final stage
FROM scratch

COPY --from=builder /app/csstats /app/

EXPOSE 8080

ENTRYPOINT ["/app/csstats"]

