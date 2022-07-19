FROM golang:1.18.4-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

ENV PORT ${PORT}
ENV PG_DB ${PG_DB}
ENV PG_PASSWORD ${PG_PASSWORD}
ENV PG_PORT ${PG_PORT}
ENV PG_USER ${PG_USER}

EXPOSE ${PORT}

RUN go build -o /service
CMD [ "/service" ]