#!/bin/bash

SIKRITKLAB_RECAPTCHA_SECRET="" \
SIKRITKLAB_DATABASE_PATH="./sikritklab.db" \
SIKRITKLAB_HOST=":8080" \
SIKRITKLAB_CORS_ORIGINS="*" \
SIKRITKLAB_CORS_METHODS="GET,POST" \
go run main.go

# GIN_MODE="release" if production