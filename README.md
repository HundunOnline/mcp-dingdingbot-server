## 🚀 mcp-wecombot-server
[![smithery badge](https://smithery.ai/badge/@gotoolkits/mcp-wecombot-server)](https://smithery.ai/server/@gotoolkits/mcp-wecombot-server)

An MCP server application that sends various types of messages to the WeCom group robot.

### Installation

### Installing via Smithery

To install mcp-wecombot-server for Claude Desktop automatically via [Smithery](https://smithery.ai/server/@gotoolkits/mcp-wecombot-server):

```bash
npx -y @smithery/cli install @gotoolkits/mcp-wecombot-server --client claude-desktop
```

### Manual Installation
```sh
# clone the repo and build
$ git clone https://github.com/gotoolkits/mcp-wecombot-server.git
$ cd mcp-wecombot-server && make build
$ sudo ln -s $PWD/dist/mcp-wecombot-server_xxx_xxxx /usr/local/bin/mcp-wecombot-server

# "$PWD/dist/mcp-wecombot-server_xxx_xxxx" replace with the actual binary file name

#You can also download and use the pre-compiled release binary package.
```

### Configuration

```json
{
  "mcpServers": {
    "mcp-wecombot-server": {
      "command": "mcp-wecombot-server",
      "env": {
        "WECOM_BOT_WEBHOOK_KEY": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxx"
      }
    }
  }
}
```

### Usage

- **send_text**

Send a text message to WeCom group

- **send_markdown**

Send a markdown message to WeCom group

- **send_image**

Send an image message to WeCom group

- **send_news**

Send a news message to WeCom group,a news include title,description,url,picurl

- **send_template_card**

Send a template card message to WeCom group

- **upload_file**

Upload a file to WeCom

### Samples

```prompt

> prompt: 给我在WeCom发送一条文本消息，消息内容为：这是一条测试消息
> prompt: 给我在WeCom发送一条markdown消息，消息内容为：# 这是一条测试 Markdown 消息
> prompt: 给我在WeCom发送一条图文消息，图文标题为：这是一条图文消息，图文描述为：这是一条图文消息，图文链接为：https://github.com/gotoolkits，图文图片为：https://img-blog.csdnimg.cn/fcc22710385e4edabccf2451d5f64a99.jpeg

> Send me a text message on WeCom with the content: This is a test message.
> Send me a Markdown message on WeCom with the content: # This is a test Markdown message
> Send me a graphic message on WeCom with the title: This is a graphic message, the description: This is a graphic message, the link: https://github.com/gotoolkits, and the image: https://img-blog.csdnimg.cn/fcc22710385e4edabccf2451d5f64a99.jpeg


```

### WeCom Robot

WeCom group robot configuration guide can be referred to:
https://developer.work.weixin.qq.com/document/path/91770

> WECOM_BOT_WEBHOOK_KEY is the robot webhook key<br>For example：
> https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=693axxx6-7aoc-4bc4-97a0-0ec2sifa5aaa <br>
> “693axxx6-7aoc-4bc4-97a0-0ec2sifa5aaa” is your own WECOM_BOT_WEBHOOK_KEY
