REM 
REM Usage  : <this script> <tablename> <devicegroup ID>
REM Example: <this script> kk-test1 dg#1
REM 
aws dynamodb put-item ^
  --table-name %1 ^
  --item "{ \"PK\": {\"S\":\"%2\"}, \"SK\": {\"S\":\"%2\"}, \"Name\": {\"S\":\"%2\"}, \"EntityType\": {\"S\":\"devicegroup\"} }"