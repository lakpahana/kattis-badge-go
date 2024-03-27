ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/app
COPY . .
WORKDIR /usr/src/app/db
RUN go mod download && go mod verify
WORKDIR /usr/src/app/api
RUN go mod download && go mod verify
# COPY go.mod go.sum ./ db/ 

WORKDIR /usr/src/app/api
#COPY . .
RUN go build -v -o /run-app .


FROM debian:bookworm

COPY --from=builder /run-app /usr/local/bin/
CMD ["run-app"]