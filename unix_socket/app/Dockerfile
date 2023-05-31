FROM golang:1.20-alpine

ADD . /workdir
WORKDIR /workdir
RUN go build -o app app.go

CMD [ "./app" ]