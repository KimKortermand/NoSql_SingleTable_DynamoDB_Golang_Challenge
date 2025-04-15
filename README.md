 
 
Overview<br/>
  - This project implements a microservice in AWS.<br/>
  - Single table design for DynamoDB using Amazon API Gateway endpoints and Lambda functions in Golang.<br/>
  - AWS CLI is used to prototype most of the Lambda functions.<br/>
<br/>
Tools<br/>
  - www.serverless.com<br/>
  - npm<br/>
  - Node.JS<br/>
  - GoLang<br/>
<br/>
<br/>
Folders and Content<br/>
   AWSCLI<br/>
    - CreateTable.bat <table name>  - creates a table in DynamoDB on AWS<br/>
    - A bunch of .bat files implementing most CRUD operations.<br/>
   Source<br/>
    - serverless.yaml - Configuration of endpoints<br/>
    - A bunch of main.go files, which implements Lambdas in Golang<br/>
<br/>
Setup<br/>
  - cd Source<br/>
  - go mod tidy<br/>
  - go install golang.org/x/tools/cmd/goimports@latest<br/>
  - npm install<br/>
  - SET AWS_ACCESS_KEY_ID=key<br/>
  - SET AWS_SECRET_ACCESS_KEY=key<br/>
  - AWS_SESSION_TOKEN=key<br/>
  - serverless plugin install -n serverless-go-plugin<br/>
<br/>
Build and deploy<br/>
  - goimports -w Source<br/>
  - sls deploy<br/>
  - sls deploy --force  (Optionally)<br/>
<br/>
<br/>
