FROM golang:1.18.4-alpine

WORKDIR /app

COPY . ./

RUN go mod download

ENV PORT ${PORT}
ENV PG_DB ${PG_DB}
ENV PG_PASSWORD ${PG_PASSWORD}
ENV PG_PORT ${PG_PORT}
ENV PG_USER ${PG_USER}

EXPOSE ${PORT}
RUN ls -la /app/cmd/cart
RUN cd ./cmd/cart && go build -o service

CMD [ "sh", "/app/cmd/service" ]