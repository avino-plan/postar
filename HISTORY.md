## ✒ 历史版本的特性介绍 (Features in old versions)

### v0.1.2-alpha
> 此版本发布于 2020-07-29
* 增加 JsonRPC 远程调用接口
* 增加 gRPC 远程调用接口
* TCP 连接优雅关闭机制
* 支持异步发送邮件

### v0.1.1-alpha
> 此版本发布于 2020-07-25
* 修复 wg.Add 调用位置导致的 Wait 失效问题，Linux 下尤其明显

### v0.1.0-alpha
> 此版本发布于 2020-07-24
* 基础的 HTTP API 提供邮件发送功能
* 支持 PLAIN 和 HTML 两种邮件格式