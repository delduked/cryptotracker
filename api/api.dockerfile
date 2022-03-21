FROM golang:alpine3.14 AS stage1
WORKDIR /app
COPY . /app
RUN go mod download
RUN go get -v
RUN go build -o /app/main

FROM alpine:latest
WORKDIR /app
COPY --from=stage1 /app/main /app/main
COPY --from=stage1 /app/.env /app/.env
EXPOSE 8080
CMD ["./main"]
