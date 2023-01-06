FROM golang:latest

RUN mkdir /build
WORKDIR /build

COPY . /build

RUN export GO111MODULE=ON

RUN cd /build/main && go build

EXPOSE 3003

ENTRYPOINT [ "/build/main/main" ]
