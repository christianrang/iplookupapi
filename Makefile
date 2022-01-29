# Cleans local directory
lclean: 

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
	curl "$$(terraform output -raw api_url)/ping"
	curl "$$(terraform output -raw api_url)/ip/8.8.8.8"
