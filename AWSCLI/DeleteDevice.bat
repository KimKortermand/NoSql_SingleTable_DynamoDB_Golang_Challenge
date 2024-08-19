REM 
REM Usage  : <this script> <tablename> <devicegroup ID> <device ID>
REM Script :               %1          %2               %3         
REM Example: <this script> kk-test4    dg#1             d#11       
REM 
aws dynamodb delete-item ^
  --table-name %1 ^
  --key "{ \"PK\": {\"S\":\"%2\"}, \"SK\": {\"S\":\"%3\"} }"
  