---
name: Accelerator 程序问题
about: "提交一个 Accelerator 的程序问题报告。"
---

<!--
除非特殊情况，请完整填写所有问题。不按模板发的 issue 将直接被关闭。
如果你遇到的问题不是 Accelerator 的 bug，比如你不清楚如何配置，请在 https://github.com/vdonkey/accelerator/discussions 进行讨论。
-->

## 你正在使用哪个版本的 Accelerator？

<!-- 如果服务端和客户端使用了不同版本，请注明 -->


## 你的使用场景是什么？

<!-- 比如使用 Chrome 通过 Socks/VMess 代理观看 YouTube 视频 -->


## 你看到的异常现象是什么？

<!-- 请描述具体现象，比如访问超时、TLS 证书错误等 -->


## 你期待看到的正常表现是怎样的？



## 请附上你的配置

<!-- 提交 issue 前，请隐去服务器域名或 IP 地址 -->

**服务端配置：**

```javascript
// 在这里附上服务器端配置文件

```

**客户端配置：**

```javascript
// 在这里附上客户端配置

```

## 请附上出错时软件输出的错误日志

<!-- 在 Linux 中，日志通常在 `/var/log/accelerator/error.log` 文件中 -->

**服务器端错误日志：**

```javascript
// 在这里附上服务器端日志

```

**客户端错误日志：**

```javascript
// 在这里附上客户端日志

```

## 请附上访问日志

<!-- 在 Linux 中，访问日志通常在 `/var/log/accelerator/access.log` 文件中 -->

```javascript
// 在这里附上服务器端日志

```

## 其它相关的配置文件（如 Nginx）和相关日志



## 如果 Accelerator 无法启动，请附上 `--test` 命令的输出



## 如果 Accelerator 服务运行异常，请附上 journal 日志

<!-- 通常的命令为 `journalctl -u accelerator` -->

<!-- 请预览一下你填写的内容并整理好格式后，再提交 -->
