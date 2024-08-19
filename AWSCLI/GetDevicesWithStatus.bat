REM 
REM Usage  : <this script> <tablename> <status>
REM Example: <this script> kk-test4  Online
REM 
aws dynamodb query ^
  --table-name %1 ^
  --index-name EntityType-Status-index ^
  --key-condition-expression "EntityType = :pk AND #Status = :sk" ^
  --expression-attribute-values  "{\":pk\": { \"S\": \"device\" }, \":sk\": { \"S\": \"%2\" }}" ^
  --expression-attribute-names "{\"#Status\":\"Status\"}"
