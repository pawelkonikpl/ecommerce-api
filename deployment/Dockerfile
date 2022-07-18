FROM golang:1.18.4-alpine

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./

RUN go get -d -v ./...
RUN go install -v ./...
COPY ./users ./

RUN go build -o /users-service

CMD [ "/users-service" ]