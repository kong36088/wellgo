# Wellgo

Lite server framework

# 流程

### 框架流程

entry->components init->listen and serve

### 请求流程

Client->protocol handler->RPC handler->router->param validator->controller

### 协议栈

#### 运输协议
`HTTP` , `TPC`

#### 接口协议

`JSONRPC` , `SOAP`

更多协议可自行拓展

### 请求示例

#### HTTP + JSONRPC

##### request:

```bash
curl localhost:80/wellgo -iv -d '{"id":"1234","jsonrpc":2.0,"method":"app.api.testapi","param":{"a":"b","i":{"a":"c"},"float":1.2,"str":123}}'
```

##### response

```bash
{"id":"1234","jsonrpc":2,"result":{"a":"b"}}
```
