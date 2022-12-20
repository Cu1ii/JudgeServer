# JudgeServer

[![](https://img.shields.io/badge/Version-v0.2-blue)](https://github.com/Cu1ii/JudgeServer) ![](https://img.shields.io/badge/go-1.19.3-brightgreen?logo=go)

这是 Online Judge 平台的**判题服务器模块**, 基于 LPOJ 的 Go 的重构版本, 将 LPOJ 的判题服务与判题调用合并了,  在 `Linux/Ubuntu20.04` 环境下实现

本程序基于调用 QingdaoU 平台提供的**判题核心**来对题目结果进行判断, 即该判题服务器 (JudgeServer) 负责接收用户的提交并将代码**编译、运行、比较，并返回判断情况**其中，代码运行的核心被单独分离在这个仓库[QingdaoU-Judger](https://github.com/QingdaoU/Judger), 考虑到判题的速度、短时间内可能需要大量判题，JudgeServer 可能需要考虑多线程、集群相关

后续会提供 **Docker 环境**作为运行环境

### 快速上手

**环境 Ubuntu20.04**

```shell
git clone https://github.com/Cu1ii/JudgeServer.git
```

如果不使用 **Docker**:

- 修改配置文件 `resources/config/judge-environment.yaml` 将**用户的提交内容目录**以及**题目数据目录**修改为自己的

- 将 `main.go` 该部分进行修改

  ```go
  // 如果需要通过配置文件读取就取消掉注释
  	//err = set.ReadSection("mysql", &global.MySQLSetting)
  	//if err != nil {
  	//	return err
  	//}
  	// 不采用 docker 设置环境变量就注释掉
  	GetMySQLConfigByEnv()
  ```



- 需要在 `judger/qdu_core_go.go` 文件中修改`proc_args := []string{"/usr/lib/judger/libjudger.so"}`  此处为

  `proc_args := []string{"echo yourpassword | sudo -S /usr/lib/judger/libjudger.so"}` 或者使用 **root** 权限来执行

- 需要将 `internal/judge/single_judge.go ` 中的 `JudgeC, JudgeCPP` 等函数的沙箱权限给关闭掉 不然判题核心会运行错误

使用 **Docker**

- 进入根目录, 使用 `Dockerfile` 创建镜像 `docker build -t image-name .`

-  `resources/config/judge-environment.yaml` 中的**用户的提交内容目录**以及**题目数据目录**已经写好, 只需要在启动 **Docker** 容器时

​		将目录挂载到本地就可以了,  **用户的提交内容目录**以及**题目数据目录** 均在`/home/cu1/XOJ `下所以只需要对该目录进行挂载即可

- MySQL 配置, 通过启动 Docker 容器时指定环境变量, 在初始化时会自动读取环境变量进行数据库链接

- 启动 **Docker(参考)**

  ```she
  docker run -it --name 容器名称 -v 要挂载的本地卷:/home/cu1/XOJ  \
  -e mysql-user=your_username -e mysql-pwd=your_pwd -e mysql-host=localhost \
  -e mysql-port=3306 -e mysql-dbname=judge_backend 镜像名称 /bin/bash
  ```

**建议使用 Docker 运行**: 可以使用 **QingdaoU**  的沙箱限制运行时权限, 相比于本地裸奔会安全

### 使用了

[![](https://img.shields.io/badge/gorm-v1.24.2-%235698c3)](https://github.com/gin-gonic/gin) [![](https://img.shields.io/badge/logrus-v1.9.0-%23428675)](https://github.com/sirupsen/logrus) [![](https://img.shields.io/badge/ants-v2.6.0-%2315231b)](https://github.com/panjf2000/ants) [![](https://img.shields.io/badge/viper-%20v1.14.0-%23e2d849)](https://github.com/spf13/viper)  [![](https://img.shields.io/badge/QingdaoU--judger-%20-%23e2d849)](https://github.com/QingdaoU/Judger)

### 版本日志

最新版本 `v0.2`

### 其他

关于判题核心请移步 [QingdaoU-Judger](https://github.com/QingdaoU/Judger)


### 待实现

- [x] 实现 SPJ
- [x] 实现配置文件读取
- [x] 完善日志
- [x] 添加 Docker 环境

### Contributions

欢迎大家提 Issue, PR

### *License*

```
MIT License

Copyright (c) 2022 cu1

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```