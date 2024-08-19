

Overview
  This project implements a microservice in AWS.
  Single table design for DynamoDB using Amazon API Gateway endpoints and Lambda functions in Golang.
  AWS CLI is used to prototype most of the Lambda functions.

Tools
  www.serverless.com
  npm
  Node.JS
  GoLang
  
Setup
  go mod tidy
  go install golang.org/x/tools/cmd/goimports@latest
  npm install
  SET AWS_ACCESS_KEY_ID=<key>
  SET AWS_SECRET_ACCESS_KEY=<key>
  AWS_SESSION_TOKEN=<key>
  
Folders and Content
  AWSCLI
    CreateTable.bat <table name>  - creates a table in DynamoDB on AWS
    A bunch of .bat files implementing most CRUD operations.
  Source
    serverless.yaml - Configuration of endpoints
    A bunch of main.go files, which implements Lambdas in Golang
    
Build and deploy
  goimports -w Source
  sls deploy
  sls deploy --force  (Optionally)
  
  