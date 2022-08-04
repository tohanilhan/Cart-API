# build stage #0
# define base image
FROM golang:1.19-alpine3.15 AS stage0
# create work directory
RUN mkdir /prog
# switch to work directory
WORKDIR /prog
# copy all files
ADD . .
# download dependencies
RUN go mod download
# build proglication
RUN CGO_ENABLED=0 GOOS=linux GODEBUG=netdns=1 GOARCH=amd64 go build -o prog .
# build stage #1
FROM alpine:latest AS stage1
# create work directory
RUN mkdir /prog
# switch to work directory
WORKDIR /prog
# copy proglication artifacts to current directory
COPY --from=stage0 /prog/prog .
# expose port
EXPOSE 8080
# run proglication
CMD ["sh", "-c", "./prog"]
