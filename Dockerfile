FROM debian
FROM golang

RUN mkdir /temp

COPY /api /go/src/ISAN-api/api
COPY /isan /go/src/ISAN-api/isan
COPY /queries /go/src/ISAN-api/queries
COPY /main.go /go/src/ISAN-api
COPY /go.mod /go/src/ISAN-api
COPY /go.sum /go/src/ISAN-api

RUN go build -C /go/src/ISAN-api/ -o /temp/frank

EXPOSE 22
EXPOSE 8080


ENTRYPOINT [ "/temp/frank" ]