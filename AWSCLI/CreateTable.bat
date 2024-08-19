REM 
REM Usage  : <this script> <tablename>
REM Example: <this script> kk-test1
REM 
aws dynamodb create-table --table-name %1 ^
--attribute-definitions AttributeName=PK,AttributeType=S AttributeName=SK,AttributeType=S AttributeName=EntityType,AttributeType=S AttributeName=Status,AttributeType=S AttributeName=MAC,AttributeType=S ^
--key-schema AttributeName=PK,KeyType=HASH AttributeName=SK,KeyType=RANGE ^
--provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 ^
--global-secondary-indexes "[ { \"IndexName\": \"EntityType-MAC-index\", \"KeySchema\": [{\"AttributeName\":\"EntityType\",\"KeyType\":\"HASH\"}, {\"AttributeName\":\"MAC\",\"KeyType\":\"RANGE\"}], \"Projection\":{ \"ProjectionType\":\"ALL\" }, \"ProvisionedThroughput\": { \"ReadCapacityUnits\": 1, \"WriteCapacityUnits\": 1 } } ,  { \"IndexName\": \"EntityType-Status-index\", \"KeySchema\": [{\"AttributeName\":\"EntityType\",\"KeyType\":\"HASH\"}, {\"AttributeName\":\"Status\",\"KeyType\":\"RANGE\"}], \"Projection\":{ \"ProjectionType\":\"ALL\" }, \"ProvisionedThroughput\": { \"ReadCapacityUnits\": 1, \"WriteCapacityUnits\": 1 } } ,  { \"IndexName\": \"EntityType-index\", \"KeySchema\": [{\"AttributeName\":\"EntityType\",\"KeyType\":\"HASH\"}], \"Projection\":{ \"ProjectionType\":\"ALL\" }, \"ProvisionedThroughput\": { \"ReadCapacityUnits\": 1, \"WriteCapacityUnits\": 1 } } ]"
