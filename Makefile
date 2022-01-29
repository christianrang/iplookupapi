# Define default for no args
all: apply

# Cleans local directory
lclean: 
	read -p "This will rm the bin and put directories. Press enter to continue or CTRL+C to cancel." tmp
	rm -r bin/
	rm -r outputs/

# Compiles Go code locally
compile:
	GOOS=linux GOARCH=amd64 go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o bin/lambda src/*

# GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/lambda src/*
init:
	terraform init

# Plans cloud deployment
plan: compile
	terraform plan

# Builds cloud architecture and pushes compiled lambda function to cloud
apply: compile
	terraform apply

# Destroys cloud architecture and cleans local directories
destroy: lclean
	terraform destroy

# Tests cloud deployment
test:
	mkdir outputs
	curl "$$(terraform output -raw api_url)/ping"
	curl "$$(terraform output -raw api_url)/search/ip/8.8.8.8" | jq >> outputs/ip.json
	curl "$$(terraform output -raw api_url)/search/domain/google.com" | jq >> outputs/domain.json
	curl "$$(terraform output -raw api_url)/search/file_hash/74768564ea2ac673e57e937f80c895c81d015e99a72544efa5a679d729c46d5f" | jq >> outputs/file_hash.json
