## 🚀 mcp-dingdingbot-server

An MCP server application that sends various types of messages to the DingDing group robot.

### Installation

### Manual Installation
```sh
# clone the repo and build
$ git clone https://github.com/HundunOnline/mcp-dingdingbot-server.git
$ cd mcp-dingdingbot-server && make build
$ sudo ln -s $PWD/dist/mcp-dingdingbot-server_xxx_xxxx /usr/local/bin/mcp-dingdingbot-server

# "$PWD/dist/mcp-dingdingbot-server_xxx_xxxx" replace with the actual binary file name

#You can also download and use the pre-compiled release binary package.
```

### Configuration

```json
{
  "mcpServers": {
    "mcp-dingdingbot-server": {
      "command": "mcp-dingdingbot-server",
      "env": {
        "DINGDING_BOT_WEBHOOK_KEY": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxx"
      }
    }
  }
}
```

### Usage

- **send_text**

Send a text message to DingDing group

- **send_markdown**

Send a markdown message to DingDing group

- **send_image**

Send an image message to DingDing group

- **send_news**

Send a news message to DingDing group, a news includes title, description, url, picurl

- **send_template_card**

Send a template card message to DingDing group

- **upload_file**

Upload a file to DingDing

### Samples

```prompt

> prompt: 给我在钉钉发送一条文本消息，消息内容为：这是一条测试消息
> prompt: 给我在钉钉发送一条markdown消息，消息内容为：# 这是一条测试 Markdown 消息
> prompt: 给我在钉钉发送一条图文消息，图文标题为：这是一条图文消息，图文描述为：这是一条图文消息，图文链接为：https://github.com/HundunOnline，图文图片为：https://img-blog.csdnimg.cn/fcc22710385e4edabccf2451d5f64a99.jpeg

> Send me a text message on DingDing with the content: This is a test message.
> Send me a Markdown message on DingDing with the content: # This is a test Markdown message
> Send me a graphic message on DingDing with the title: This is a graphic message, the description: This is a graphic message, the link: https://github.com/HundunOnline, and the image: https://img-blog.csdnimg.cn/fcc22710385e4edabccf2451d5f64a99.jpeg


```

### DingDing Robot

DingDing group robot configuration guide can be referred to:
https://open.dingtalk.com/document/robots/custom-robot-access

> DINGDING_BOT_WEBHOOK_KEY is the robot webhook key<br>For example：
> https://oapi.dingtalk.com/robot/send?access_token=693axxx6-7aoc-4bc4-97a0-0ec2sifa5aaa <br>
> "693axxx6-7aoc-4bc4-97a0-0ec2sifa5aaa" is your own DINGDING_BOT_WEBHOOK_KEY
