#!/bin/bash

SIKRITKLAB_RECAPTCHA_SECRET="" \
SIKRITKLAB_DATABASE_PATH="./sikritklab.db" \
SIKRITKLAB_HOST=":8080" \
SIKRITKLAB_CORS="true" \
SIKRITKLAB_CORS_ORIGINS="http://localhost:8080" \
SIKRITKLAB_CORS_METHODS="GET,POST,PUT,HEAD" \
SIKRITKLAB_CORS_ALLOW_HEADERS="Origin,Authorization,Accept,Content-Length,Content-Type" \
SIKRITKLAB_CORS_EXPOSE_HEADERS="Content-Length" \
go run main.go

# GIN_MODE="release" if production