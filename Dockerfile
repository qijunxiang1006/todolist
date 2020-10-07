FROM node:14.5 AS Web_Build
COPY webapp /webapp
WORKDIR /webapp
RUN npm install && npm run-script build

FROM golang AS GO_BUILD
COPY server /server
WORKDIR /server
RUN go env -w GOPROXY="https://goproxy.cn,direct"
RUN go build -o ./server

FROM ubuntu:15.10
COPY --from=Web_Build /webapp/build* /webapp/
COPY --from=GO_BUILD /server/server /
WORKDIR /
CMD ./server