#Version of golang we want to run container on!
FROM golang:1.15-alpine AS building

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o bin/api .
CMD [ "/app/bin/api" ]






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


