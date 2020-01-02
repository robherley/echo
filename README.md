# test-server

[![Build Status](https://drone.robherley.xyz/api/badges/robherley/test-server/status.svg)](https://drone.robherley.xyz/robherley/test-server)

docker image that contains a simple golang server for quick & dirty debugging. container also has some useful tools preinstalled like curl, wget, jq, vim, dig, etc

example response:

`GET` `http://<server-endpoint>/`

```
{
  "Body": {},
  "ContentLength": 0,
  "Form": null,
  "Header": {
    "Accept": [
      "*/*"
    ],
    "User-Agent": [
      "curl/7.64.1"
    ]
  },
  "Host": "localhost:8080",
  "Method": "GET",
  "MultiPartForm": null,
  "PostForm": null,
  "Proto": "HTTP/1.1",
  "RemoteAddr": "[::1]:49695",
  "RequestURI": "/",
  "RequestedAt": "2019-11-05T11:01:31.914031-05:00",
  "TLS": null,
  "Trailer": null,
  "TransferEncoding": null,
  "URL": "/"
}
```
