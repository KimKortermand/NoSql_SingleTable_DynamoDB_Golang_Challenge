REM 
REM Usage  : <this script> <tablename> <devicegroup ID>
REM Example: <this script> kk-test1    dg#1
REM 
aws dynamodb query ^
  --table-name %1 --key-condition-expression "PK = :pk"    ^
  --expression-attribute-values  "{\":pk\": { \"S\": \"%2\" }}"