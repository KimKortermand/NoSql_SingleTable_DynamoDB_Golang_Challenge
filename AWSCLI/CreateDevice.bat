REM 
REM Usage  : <this script> <tablename> <devicegroup ID> <device ID> <device MAC> <device Status>
REM Script :               %1          %2               %3          %4           %5
REM Example: <this script> kk-test4    dg#1             d#11        FF-FF-FF-FF-FF-11 Offline
REM 
aws dynamodb put-item ^
  --table-name %1 ^
  --item "{ \"PK\": {\"S\":\"%2\"}, \"SK\": {\"S\":\"%3\"}, \"Name\": {\"S\":\"%3\"}, \"EntityType\": {\"S\":\"device\"}, \"MAC\": {\"S\":\"%4\"}, \"Status\": {\"S\":\"%5\"} }"
