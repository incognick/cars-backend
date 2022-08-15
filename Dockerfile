FROM golang:1.19-alpine as builder


WORKDIR /build

COPY *.go .
COPY go.mod .
COPY go.sum .
COPY db db

RUN CGO_ENABLED=0 GOOS=linux go build -o cars-backend .

FROM alpine:3.16
COPY --from=builder /build/cars-backend .
COPY db/migrations /db/migrations 

ENTRYPOINT [ "./cars-backend" ]

