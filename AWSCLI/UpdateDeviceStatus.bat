REM 
REM Usage  : <this script> <tablename> <devicegroup ID> <device ID> <new status>
REM Example: <this script> kk-test1 dg#1 d#11 Online
REM 
aws dynamodb update-item ^
  --table-name %1 --key "{ \"PK\": {\"S\":\"%2\"}, \"SK\": {\"S\":\"%3\"} }" ^
  --update-expression "SET #Status = :status" ^
  --expression-attribute-values  "{\":status\": {\"S\": \"%4\"} }" ^
  --expression-attribute-names "{\"#Status\":\"Status\"}"