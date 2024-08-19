REM 
REM Usage  : <this script> <tablename> <MAC>
REM Example: <this script> kk-test4  FF-FF-FF-FF-FF-11
REM 
aws dynamodb query ^
  --table-name %1 ^
  --index-name EntityType-MAC-index ^
  --key-condition-expression "EntityType = :pk AND MAC = :sk" ^
  --expression-attribute-values  "{\":pk\": { \"S\": \"device\" }, \":sk\": { \"S\": \"%2\" }}"
  
