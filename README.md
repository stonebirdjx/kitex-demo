# kitex-demo
kitex演示项目

```bash
# thrift 生成golang代码
kitex -module github.com/stonebirdjx/kitex-demo -service kitex-server idl/echo.thrift
```

运行服务端

```bash
bash build.sh
bash output/bootstrap.sh
```

运行客户端

```bash
go run ./client/main.go
2023/09/13 21:32:41 Response({Message:my request})
```

