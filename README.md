# JudgeServer

[![](https://img.shields.io/badge/Version-v0.1-blue)](https://github.com/Cu1ii/JudgeServer) ![](https://img.shields.io/badge/go-1.19.3-brightgreen?logo=go)

这是 Online Judge 平台的**判题服务器模块**, 基于 LPOJ 的 Go 的重构版本, 将 LPOJ 的判题服务与判题调用合并了,  在 `Linux/Ubuntu20.04` 环境下实现

本程序基于调用 QingdaoU 平台提供的**判题核心**来对题目结果进行判断, 即该判题服务器 (JudgeServer) 负责接收用户的提交并将代码**编译、运行、比较，并返回判断情况**其中，代码运行的核心被单独分离在这个仓库[QingdaoU-Judger](https://github.com/QingdaoU/Judger), 考虑到判题的速度、短时间内可能需要大量判题，JudgeServer 可能需要考虑多线程、集群相关

后续会提供 **Docker 环境**作为运行环境

### 快速上手

```shell
git clone https://github.com/Cu1ii/JudgeServer.git
```

待补充....

### 使用了

[![](https://img.shields.io/badge/gorm-v1.24.2-%235698c3)](https://github.com/gin-gonic/gin) [![](https://img.shields.io/badge/logrus-v1.9.0-%23428675)](https://github.com/sirupsen/logrus) [![](https://img.shields.io/badge/ants-v2.6.0-%2315231b)](https://github.com/panjf2000/ants) [![](https://img.shields.io/badge/viper-%20v1.14.0-%23e2d849) ](https://github.com/spf13/viper)  [![](https://img.shields.io/badge/QingdaoU--judger-%20-%23e2d849)](https://github.com/QingdaoU/Judger)

### 版本日志

最新版本 `v0.1`

### 其他

关于判题核心请移步 [QingdaoU-Judger](https://github.com/QingdaoU/Judger)


### 待实现

- [x] 实现 SPJ
- [ ] 实现配置文件读取
- [ ] 完善日志
- [ ] 添加 Docker 环境

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