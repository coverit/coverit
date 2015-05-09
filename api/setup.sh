#!/bin/bash

export MONGO_URL='mongodb://localhost'
export MONGO_DB='coverit'

go run server.go
