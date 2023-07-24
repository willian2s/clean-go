#!/bin/bash

# Execute the tests
go test -coverprofile coverage/cover.out ./...

# Generate the coverage report
go tool cover -html=coverage/cover.out -o coverage/index.html