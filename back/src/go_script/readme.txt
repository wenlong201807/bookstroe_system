



参考：https://blog.csdn.net/netdxy/article/details/79431415

一般部署脚本应该具有构建、启动、停止、回滚已经查看记录日志等功能，以下分别将这些功能以单个脚本的形式给出，当然也可以写成Makefile 的形式。


单个部署脚本的形式，在一个目录下建立如下文件：
1.bin # 目录，用于存放每次 build 之后存放的二进制文件
2.app.log # 用来记录的日志文件
3.log.sh # 实时查看日志
4.build.sh # 构建
5.run.sh # 启动某一次编译版本
6.start.sh # 启动最新版本，并且备份之前前一次运行的版本
7.shutdown.sh # 停止
8.rollback.sh # 回滚到上一版本


本例中的 GOPATH=”/go”