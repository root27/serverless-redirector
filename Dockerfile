FROM golang:1.22 as builder

WORKDIR /src/application

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./redirector.out .





FROM gcr.io/distroless/static

COPY --from=builder /src/application/redirector.out /myapp

ENTRYPOINT ["/myapp"]

