#!/usr/bin/env bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
docker build -t meetdocker/http-authservice .
docker push meetdocker/http-authservice
