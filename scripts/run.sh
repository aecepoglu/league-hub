#!/bin/sh

go build

export REDIS_URI=localhost:6379
./league-hub
