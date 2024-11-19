# thrifturl 

旨在实现一套类似于 curl， gRPCurl一样的命令行工具，用来调试 thrift rpc 接口

设想用法：
```shell
thrifturl \
--thrift-file foo.thrift \
--host localhost \
--port 9090 \
--method FooMethod \
--params '{
  "arg1": 1,
  "arg2": "2",
  "arg3": {
    "arg4": 4
  }
}'
```

设想架构：

thrifturl 应该由3大模块构成
### thrift 解析器
用于解析用户指定的thrift文件，词法分析，句法分析

### json 解析器
用于解析并且序列化用户指定的param参数，类型校验，以及反序列化响应

### dial调用模块
负责将序列化的结果发送给服务器，并且接收响应


