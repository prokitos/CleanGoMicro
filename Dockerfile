FROM golang:alpine AS builder

WORKDIR /build
ADD go.mod .
COPY . .
COPY /config /tempConfig
COPY test.db /tempDatabase/test.db
# нижняя строка нужна только если будет работа с sqlLite3. без неё не работает. она по идее устанавливает GCC и настраивает его?
RUN apk add build-base 
RUN go build -o . cmd/main.go

FROM alpine
WORKDIR /build
COPY --from=builder /build/main /build/main
COPY --from=builder /tempConfig /build/config
COPY --from=builder /tempDatabase/test.db /build
CMD ["./main"]