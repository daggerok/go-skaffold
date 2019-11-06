FROM golang:1.13.4-alpine3.10 as build
COPY main.go .
RUN go build -o /app main.go

FROM golang:1.13.4-alpine3.10
COPY --from=build /app .
ENTRYPOINT ["./app"]
