# syntax=docker/dockerfile:1.7
FROM golang:1.26-bookworm AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /out/musu-marketer .

FROM gcr.io/distroless/static-debian12:latest
COPY --from=build /out/musu-marketer /usr/local/bin/musu-marketer
WORKDIR /app
EXPOSE 8081
ENTRYPOINT ["/usr/local/bin/musu-marketer"]
