#!/bin/bash

cd ../pkg/users && go run github.com/99designs/gqlgen generate
cd ../products && go run github.com/99designs/gqlgen generate
cd ../cart && go run github.com/99designs/gqlgen generate
