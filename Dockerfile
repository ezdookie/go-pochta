FROM golang:1.16-alpine

WORKDIR /usr/src/app
COPY . .

RUN go build

CMD ["./pochta-api"]