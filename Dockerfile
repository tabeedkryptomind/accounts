FROM golang:1.17.8 as dev

WORKDIR /api

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux go build -a -o accounts .


FROM alpine:3

WORKDIR /app

COPY --from=dev /api/accounts ./

EXPOSE 8080

CMD ["./accounts", "serve"]