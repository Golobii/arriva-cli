#!/bin/bash

go build -o bin/arriva-cli -ldflags "-s -w" cmd/arriva-cli/main.go 