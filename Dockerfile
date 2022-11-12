FROM golang:1.19.3-buster
WORKDIR /usr/src/app/
USER ${USER}
COPY ./go.mod /usr/src/app/
COPY . /usr/src/app/
ENV GO111MODULE="on" \
  CGO_ENABLED="0"
RUN apt-get autoclean \
  && apt-get autoremove  \
  && apt-get update \
  && apt-get install -y \
  build-essential
RUN go mod tidy \
  && go mod download \
  && go build -o main .