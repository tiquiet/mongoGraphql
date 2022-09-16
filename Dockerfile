FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o mongoGraph ./cmd/main.go
CMD ["./mongoGraph"]