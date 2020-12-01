#Version of golang we want to run container on!
FROM golang:1.15-alpine AS building

RUN mkdir /cal_server
COPY . /cal_server
WORKDIR /cal_server
RUN go build -o bin/processing-server .

EXPOSE 8081

CMD [ "/cal_server/bin/processing-server" ]


# Environment variables for building GOALNG image
# 설명::
    #GO111MODULE    ->
    #CGO_ENABLED    ->
    #GOOS           ->
    #GOARCH         ->
#ENV GO111MODULE=on \
#    CGO_ENABLED=0 \
#    GOOS=linux \
#    GOARCH=amd64


