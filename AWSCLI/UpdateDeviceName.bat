REM 
REM Usage  : <this script> <tablename> <devicegroup ID> <device ID> <new name>
REM Example: <this script> kk-test1 dg#1 d#11 d#11_new
REM 
aws dynamodb update-item ^
  --table-name %1 --key "{ \"PK\": {\"S\":\"%2\"}, \"SK\": {\"S\":\"%3\"} }" ^
  --update-expression "SET #Name = :name" ^
  --expression-attribute-values  "{\":name\": {\"S\": \"%4\"} }" ^
  --expression-attribute-names "{\"#Name\":\"Name\"}"