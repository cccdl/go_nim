# 云信 SDK for GO  !
go_nim 是用 GO 语言实现的网易云信的服务端 API 封装，

## 特性
* 针对其他github上的包，实现发送自定义消息，github上的包只发其他消息
* 持续更新中...

###更新日志

- 1.0.0 发送自定义消息

### 已实现功能
- [ ] 通信服务
  - [x] 发送自定义消息

## 使用方法

#### 安装:

`go get -u github.com/cccdl/go_nim`

#### 导入:

`import "github.com/cccdl/go_nim"`

#### 使用:

> 发送自定义消息：
```go
msg := &nim.Message{}
c := nim.CreateImClient("", "")
opt := &nim.ImSendMessageOption{
    MsgSenderNoSense:   0,
    MsgReceiverNoSense: 1,
}

err := c.SendCustomMessage("test1", "test2", msg, opt)
if err != nil {
t.Error(err)
}
```

## 文档
[云信api文档](https://doc.yunxin.163.com/docs/TM5MzM5Njk/jk3MzY2MTI?platformId=60353)

## 问题
[提交 Issue](https://github.com/cccdl/go_nim/issues)，不符合指南的问题可能会立即关闭。

## License
go_nim 使用[MIT](https://opensource.org/licenses/MIT)开源协议


