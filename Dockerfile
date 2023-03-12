FROM golang:1.20 as builder
WORKDIR /usr/src/app
COPY ./src .
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM gcr.io/distroless/static:latest
WORKDIR /opt
COPY --from=builder /usr/src/app/server /opt/server
ENTRYPOINT [ "/opt/server" ]
