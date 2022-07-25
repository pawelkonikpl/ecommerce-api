#!/bin/bash

function cleanup {
    kill "$USERS_PID"
    kill "$PRODUCTS_PID"
    kill "$CART_PID"
}
trap cleanup EXIT

go build -o /tmp/srv-users ./users/server.go
go build -o /tmp/srv-products ./products/server.go
go build -o /tmp/srv-cart ./cart/server.go

/tmp/srv-users &
USERS_PID=$!

/tmp/srv-products &
PRODUCTS_PID=$!

/tmp/srv-cart &
CART_PID=$!

sleep 1

node web/gateway/index.js