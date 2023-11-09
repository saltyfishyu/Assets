# Assets

适用于本地部署的src资产管理工具

## Todo

- 代码更新中
- 前端

## api
### /assets/all

```
GET /assets/all HTTP/1.1
Host: 127.0.0.1:7070

HTTP/1.1 200 OK
Access-Control-Allow-Credentials: true
Access-Control-Allow-Headers: *
Access-Control-Allow-Methods: POST, GET, OPTIONS, PUT, DELETE, UPDATE
Access-Control-Allow-Origin: *
Access-Control-Expose-Headers: Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type
Content-Type: application/json; charset=utf-8
Date: Wed, 08 Nov 2023 08:02:21 GMT
Content-Length: 160

{"data":[{"ID":1,"CreatedAt":"2023-11-08T15:46:43.755142+08:00","UpdatedAt":"2023-11-08T15:46:43.755142+08:00","DeletedAt":null,"content":"京东"}],"status":0}
```

### /assets/add

```
POST /assets/add HTTP/1.1
Host: 127.0.0.1:7070
Content-Type: application/json
Content-Length: 21

{"content": "淘宝"}

HTTP/1.1 200 OK
Access-Control-Allow-Credentials: true
Access-Control-Allow-Headers: *
Access-Control-Allow-Methods: POST, GET, OPTIONS, PUT, DELETE, UPDATE
Access-Control-Allow-Origin: *
Access-Control-Expose-Headers: Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type
Content-Type: application/json; charset=utf-8
Date: Wed, 08 Nov 2023 09:25:03 GMT
Content-Length: 42

{"msg":"新增资产组成功","status":0}
```

### /assetsbelong/add

```
POST /assetsbelong/add HTTP/1.1
Host: 127.0.0.1:7070
Content-Type: application/json
Content-Length: 34

{"aid": 2,"domain":["taobao.com"]}

HTTP/1.1 200 OK
Access-Control-Allow-Credentials: true
Access-Control-Allow-Headers: *
Access-Control-Allow-Methods: POST, GET, OPTIONS, PUT, DELETE, UPDATE
Access-Control-Allow-Origin: *
Access-Control-Expose-Headers: Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type
Content-Type: application/json; charset=utf-8
Date: Wed, 08 Nov 2023 09:25:11 GMT
Content-Length: 39

{"msg":"新增资产成功","status":0}
```