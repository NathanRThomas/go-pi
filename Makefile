
# go params
GOCMD=go
GOBUILD=$(GOCMD) build -buildvcs=false
GOPATH=/usr/local/bin
DIR=$(shell pwd)

# normal entry points
build:
	clear
	@echo "building..."
	@$(GOBUILD) -o $(GOPATH)/gopi ./cmd/
	
update:
	clear
	@echo "updating dependencies..."
	@go get -u -t ./...
	@go mod tidy 

#export CONFIG=/var/www/go/api/testing/config.json
#export TEMPLATES=/var/www/go/api/templates/

