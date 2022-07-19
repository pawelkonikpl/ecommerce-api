#!/bin/bash

function cleanup {
    kill "$USERS_PID"
    kill "$PRODUCTS_PID"
}
trap cleanup EXIT

go build -o /tmp/srv-users ./users
go build -o /tmp/srv-products ./products

/tmp/srv-users &
USERS_PID=$!

/tmp/srv-products &
PRODUCTS_PID=$!

sleep 1

node web/gateway/index.js