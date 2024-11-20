# Use the official Go image as the build stage.
FROM golang:1.22 AS firststage

WORKDIR /build/

COPY . /build

ENV CGO_ENABLED=0

RUN go get
# Build the Go application with optimizations.
RUN go build -o echo
# Create a lightweight final stage to run the application.
FROM alpine:latest

WORKDIR /app/

RUN apk update

RUN apk upgrade

RUN apk add ca-certificates && update-ca-certificates

RUN apk add --no-cache tzdata gcompat

ENV TZ=Asia/Jakarta

COPY --from=firststage /build/pkg/rsa/private.pem /app/pkg/rsa/private.pem
COPY --from=firststage /build/pkg/rsa/public.pem /app/pkg/rsa/public.pem

COPY --from=firststage /build/echo .

CMD ["./echo"]