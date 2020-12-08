学习笔记


基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。

```
1运行命令：go run main.go
2使用：ctr+c 退出

go: downloading golang.org/x/sync v0.0.0-20181221193216-37e7f081c4d4
goroutine3: signal
goroutine1: service start
goroutine2: service stop
^Cgoroutine3: signal goroutine cancel
goroutine1: service start goroutine cancel
goroutine2: service stop goroutine cancel
err: goroutine3: signal hand exit% 
```
