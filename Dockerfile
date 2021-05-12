FROM golang:alpine

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /code

COPY ./go.mod .
RUN go mod download

COPY . .
RUN go build -o app .

CMD ["/code/app"]