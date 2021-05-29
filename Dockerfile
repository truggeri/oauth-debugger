# -----------------------------------------------------------------------------
# Go Build
# 
FROM golang:1.16 as go-build
ENV CGO_ENABLED 0
ARG VCS_REF

RUN mkdir /service
COPY go.* /service/
WORKDIR /service
RUN go mod download

COPY functions/*.go /service/functions/
COPY cmd/server/main.go /service/cmd/server/main.go
RUN go build -ldflags "-X main.build=${VCS_REF}" cmd/server/main.go

# -----------------------------------------------------------------------------
# Node Frontend Build
# 
FROM node:14.16.0 as node-build

RUN mkdir /work
COPY frontend/src/*.js /work/
COPY *.json /work/
WORKDIR /work
RUN npm install

RUN mkdir -p /static/js
RUN mkdir -p /static/css
COPY frontend/src /work/
RUN npm run build

# -----------------------------------------------------------------------------
# Hugo Frontend Build
# 
FROM klakegg/hugo:0.83.1 as hugo-build
COPY --from=node-build /static/js/* /src/static/js/
COPY --from=node-build /static/css/* /src/static/css/
COPY frontend /src
RUN hugo

# -----------------------------------------------------------------------------
# Running Application
# 
FROM alpine:3.13 as main
ARG BUILD_DATE
ARG VCS_REF

RUN apk add --no-cache bash

COPY --from=go-build /service/main /service/main
COPY templates /service/serverless_function_source_code
COPY --from=hugo-build /src/public /service/public
WORKDIR /service
EXPOSE 8090
CMD ["./main"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.title="oauthdebugger" \
      org.opencontainers.image.authors="Thomas Ruggeri <truggeri@gmail.com>" \
      org.opencontainers.image.source="https://github.com/truggeri/oauthdebugger" \
      org.opencontainers.image.revision="${VCS_REF}" \
      org.opencontainers.image.vendor=""
