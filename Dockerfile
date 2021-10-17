FROM golang:alpine AS build

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /code

COPY ./go.mod .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp

#CMD ["/code/app"]


#FROM golang:alpine AS build
#ENV GO111MODULE=on \
#    GOPROXY=https://goproxy.cn,direct
#
#RUN mkdir /app
#COPY . /app
#WORKDIR /app
#RUN CGO_ENABLED=0 GOOS=linux go build -o myapp

###
FROM scratch as final
COPY --from=build /code .
CMD ["/myapp"]