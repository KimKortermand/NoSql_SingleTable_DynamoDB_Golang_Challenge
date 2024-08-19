REM 
REM Usage  : <this script> <tablename> <devicegroup ID> <device ID> <new MAC>
REM Example: <this script> kk-test1 dg#1 d#11 FF-FF-FF-FF-FF-11_new
REM 
aws dynamodb update-item ^
  --table-name %1 --key "{ \"PK\": {\"S\":\"%2\"}, \"SK\": {\"S\":\"%3\"} }" ^
  --update-expression "SET MAC = :mac" ^
  --expression-attribute-values  "{\":mac\": {\"S\": \"%4\"} }"