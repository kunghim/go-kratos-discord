# go-kratos-discord
go-kratos-discord 是一个基于 [kratos](https://github.com/go-kratos/kratos) 框架的 discord 机器人，可以将 discord 机器人和自己的服务器进行关联，实现 discord 机器人和服务器的双向通信。
其中与 discord 机器人的通信使用的是 [discordgo](https://github.com/bwmarrin/discordgo) 库。

## 安装
```
# 克隆项目
git clone https://github.com/kunghim/go-kratos-discord.git
# 安装依赖
make init
```
## 修改配置文件
```
# Create a template project
server: 
  # discord 配置
  discord:
    # 用户 token
    user_token: xxx
    # 机器人 token
    bot_token: xxx
    # 服务器 id
    server_id: "101010"
    # 频道 id
    channel_id: "010101"
```

## 运行项目
需要科学上网，否则无法连接到 discord 机器人
```
server -conf ./configs
```
启动成功

![img.png](doc/image/start_success.png)

在频道内发送和修改消息
![img.png](doc/image/send.png)

服务器打印消息

![img_1.png](doc/image/print.png)
## 项目拓展
项目内置新消息（CreateMessage）和修改（UpdateMessage）消息事件函数，可根据自己的需求在 internal/service/discord.go 中进行拓展
```
// NewMessage 新消息事件（删除消息）
func NewMessage(s *discordgo.Session, m *discordgo.MessageDelete) {
    // TODO
}
```
新的消息事件函数需要在 internal/server/discord.go 中注册
```
// RegisterEvent 注册事件
func NewDiscordService(c *conf.Server, discordService *service.DiscordService) *DiscordServer {
	session, err := discordgo.New(fmt.Sprintf("Bot %s", c.Discord.BotToken))
	if err != nil {
		panic(err)
	}
	session.AddHandler(discordService.CreateMessage)
	session.AddHandler(discordService.UpdateMessage)
	// 注册新的消息事件
	session.AddHandler(discordService.Something)
	
	return &DiscordServer{
		discordService: discordService,
		session:        session,
	}
}
```
## 说明文档
[说明文档（如何创建机器人，获取用户 token，服务器 id，频道 id 登）](https://www.yuque.com/pugongyingbushimeng/dsghkz/xvmzagiuqtriers6?singleDoc)

---

## 感谢
- [kratos](https://github.com/go-kratos/kratos)
- [discordgo](https://github.com/bwmarrin/discordgo)

---
### 任何疑问欢迎联系，加群讨论
![img.png](doc/image/contact.png)