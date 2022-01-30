# IP Lookup API

## API Routes

|Route|Description|
|:--:|:--:|
|/search/ip/:ip|Searches IP address using VirusTotal and IPInfo.io APIs|
|/search/domain/:domain|Searches Domain using VirusTotal API|
|/search/file_hash/:file_hash|Searches file hash (md5, sha1, sha256) using VirusTotal API|

## Status Codes

|Code| Description|
|:---:|:-----------|
|200| Successful request|
|204|One of our data providers has limited our requests|
|400|One of our data providers was unable to process the request|
|406|An argument passed was not formed correctly. This typically occurs when an IP is not a valid IP address, a file hash is not a valid hash, etc.|