FROM golang:1.14

WORKDIR /go/src/app

COPY main.go /go/src/app/

EXPOSE 5000

RUN go get -d -v

RUN go install -v

ENTRYPOINT ["go", "run", "main.go"]



