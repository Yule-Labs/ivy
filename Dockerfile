FROM golang:1.14.3-alpine3.11 as builder

ENV GIT_TERMINAL_PROMPT=1

RUN apk --update add git openssh && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*

WORKDIR /ivy
COPY go.mod go.sum ./

RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o bin/ivy cmd/ivy/main.go

FROM alpine:3.11

ENV TZ=Europe/Moscow
RUN apk --update add bash tzdata && \
    cp /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo ${TZ} > /etc/timezone

WORKDIR /usr/bin
COPY --from=builder /ivy/bin/ivy ./
COPY --from=builder /ivy/configs/ivy.conf.yml ./ivy.conf.yml
RUN chmod +x ./ivy

ENTRYPOINT ["/usr/bin/ivy"]
CMD ["-config", "./ivy-default.conf.yml"]