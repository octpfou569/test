FROM golang:alpine

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /go/src/hjkim-portfolio
COPY main .

RUN apk update && apk add --no-cacahe git

EXPOSE $PORT

ENTRYPOINT ["./main"]