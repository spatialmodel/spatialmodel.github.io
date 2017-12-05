#!/bin/bash

# This script compiles a program for different systems

env GOOS=linux GOARCH=amd64 go build -v process_emissions.go
mv process_emissions process_emissions_linux

env GOOS=darwin GOARCH=amd64 go build -v process_emissions.go
mv process_emissions process_emissions_mac

env GOOS=windows GOARCH=amd64 go build -v process_emissions.go
mv process_emissions process_emissions_windows.exe
