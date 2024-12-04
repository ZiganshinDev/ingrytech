FROM golang:1.22-alpine AS builder

ADD . /ingrytech
WORKDIR /ingrytech

RUN go build -o ./tmp/ingrytech  ingrytech/cmd/ingrytech/main.go

FROM jerray/goose:2.6.0 AS goose

FROM alpine:3.15

ENV MIGRATIONS_PATH=/var/lib/migrations
#PG_DSL in Ansible

COPY --from=builder /tmp/ingrytech /usr/bin/ingrytech
COPY --from=goose /bin/goose /usr/bin/goose
COPY ./migrations $MIGRATIONS_PATH

COPY /config /config

RUN echo '#! /bin/sh' > /usr/bin/entrypoint.sh
RUN echo 'goose -dir $MIGRATIONS_PATH postgres "$PG_DSL" up &&' >> /usr/bin/entrypoint.sh
RUN echo 'ingrytech' >> /usr/bin/entrypoint.sh
RUN chmod +x /usr/bin/entrypoint.sh 

EXPOSE 8085
ENTRYPOINT ["/ingrytech"]