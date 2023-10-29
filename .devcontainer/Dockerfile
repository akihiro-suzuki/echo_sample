FROM golang:1.20.10-alpine
RUN apk update && apk add git && apk add bash
RUN addgroup -g 1000 -S appgroup && adduser -u 1000 -S appuser -G appgroup -s /bin/bash
USER appuser
