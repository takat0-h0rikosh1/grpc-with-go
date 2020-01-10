FROM golang:1.13-buster as builder

COPY ./ ./app/

RUN cd ./app && go get
RUN cd ./app && go build -o /main

FROM debian:stretch-slim

COPY --from=builder /main .

ENTRYPOINT ["./main"]