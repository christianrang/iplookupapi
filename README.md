# IP Lookup API

## Status Codes

|Code| Description|
|:---:|:-----------|
|200| Successful request|
|406|An argument passed was not formed correctly. This typically occurs when an IP is not a valid IP address, a file hash is not a valid hash, etc.|
|400|One of our data providers was unable to process the request|
|204|One of our data providers has limited our requests|