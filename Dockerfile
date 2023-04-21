FROM golang:1.20.3-alpine3.16 AS build


WORKDIR /CANY
COPY ./public ./public
COPY ./statserv.go ./statserv.go
COPY ./go.mod ./go.mod
RUN go build -o serve .

FROM alpine:3.17.3

EXPOSE 80

COPY --from=build /CANY /CANY 


CMD ["/CANY/serve", "--static=/CANY/public"]

