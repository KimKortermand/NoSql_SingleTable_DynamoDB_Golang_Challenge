


This project implements a microservice in AWS.
It is build single table design for DynamoDB using Amazon API Gateway endpoints and Lambda functions in Golang.
AWS CLI is used to prototype most of the Lambda functions.
Lambdas are written in Golang.

Tools
  www.serverless.com
  npm
  Node.JS
  GoLang
  
Setup
  go mod tidy
  npm install
  SET AWS_ACCESS_KEY_ID=<key>
  SET AWS_SECRET_ACCESS_KEY=<key>
  AWS_SESSION_TOKEN=<key>
  
Build and deploy
  sls deploy
  sls deploy --force  (Optionally)
  
Folders and Content
  AWSCLI contains bat files:
    CreateTable.bat <table name>  - creates a table in DynamoDB on AWS
    plus a bunch of files implementing most CRUD operations.
  Source contains configuration of endpoints and Lambdas impl in Golang
    
  