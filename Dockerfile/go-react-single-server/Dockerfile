# node base image
FROM node:16-alpine3.11 as build-node

# copy required files to /workdir
WORKDIR /workdir
COPY frontend/ .

# build frontend
RUN yarn install
RUN yarn build

# copy 
COPY .env /workdir/

# go base image
FROM golang:1.21-rc-alpine3.18 as build-go

# build go binary
ENV GOPATH ""
RUN go env -w GOPROXY=direct
RUN apk add git

ADD go.mod go.sum ./
RUN go mod download
ADD . .
RUN go build -o /main

FROM alpine:3.13
COPY --from=build-go /main /main
COPY --from=build-node /workdir/.env .env
COPY --from=build-node /workdir/build ./frontend/build

EXPOSE 8000
ENTRYPOINT [ "/main" ]
