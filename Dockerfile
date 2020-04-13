# Phase 1: App Build
FROM node:latest as builder

WORKDIR /usr/src/app
COPY ./deimos/package.json ./deimos/yarn.lock ./
RUN yarn
COPY ./deimos ./

# Finally, build the app! Production edition!
RUN yarn build

# Phase 2 : Server Build
FROM golang:latest
RUN mkdir /server 
ADD ./server /server/ 
WORKDIR /server 

# Copy in the compiled frontend app
COPY --from=builder /usr/src/app/build ../deimos/build

RUN go build -o main . 
CMD ["/server/main"]
