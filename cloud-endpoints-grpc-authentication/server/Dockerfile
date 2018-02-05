FROM golang:1.9

WORKDIR /go/src/app
ADD . .
RUN go get .

EXPOSE 50051

CMD ["go", "run", "main.go", "echo.pb.go"]
