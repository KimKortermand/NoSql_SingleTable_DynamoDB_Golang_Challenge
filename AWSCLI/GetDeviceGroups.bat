REM 
REM Usage  : <this script> <tablename>
REM Example: <this script> kk-test4
REM 
aws dynamodb query ^
  --table-name %1  ^
  --index-name EntityType-index ^
  --key-condition-expression "EntityType = :pk" --expression-attribute-values  "{\":pk\": { \"S\": \"devicegroup\" }}"