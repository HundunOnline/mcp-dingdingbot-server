## 🚀 mcp-dingdingbot-server

An MCP server application that sends various types of messages to the DingDing group robot.

[English](#english) | [中文](#chinese)

<a name="english"></a>
## English

### Overview

This is an MCP (Message Control Protocol) server application that allows you to send various types of messages to DingDing group robots. It supports text, markdown, image, news, and template card messages, as well as file uploads.

### Features

- Text message support
- Markdown message support
- Image message support
- News message support
- Template card message support
- File upload support
- Signature verification for enhanced security

### Installation

#### Manual Installation
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
        "DINGDING_BOT_WEBHOOK_KEY": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxx",
        "DINGDING_BOT_SIGN_KEY": "SECxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
      }
    }
  }
}
```

## Environment Variables

- `DINGDING_BOT_WEBHOOK_KEY`: The webhook key for the DingDing Bot server. This is required.
- `DINGDING_BOT_SIGN_KEY`: The sign key for DingDing Bot signature verification. This is optional but recommended for enhanced security.

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
>
> DINGDING_BOT_SIGN_KEY is the signature key for enhanced security<br>
> When enabled in the DingDing robot security settings, you need to provide this key to authenticate requests.
> The signature verification uses HMAC-SHA256 algorithm with the timestamp and secret key.

<a name="chinese"></a>
## 中文

### 概述

这是一个 MCP（消息控制协议）服务器应用程序，允许您向钉钉群机器人发送各种类型的消息。它支持文本、Markdown、图片、图文和模板卡片消息，以及文件上传。

### 功能

- 文本消息支持
- Markdown 消息支持
- 图片消息支持
- 图文消息支持
- 模板卡片消息支持
- 文件上传支持
- 签名验证增强安全性

### 安装

#### 手动安装

```sh
# 克隆仓库并构建
$ git clone https://github.com/HundunOnline/mcp-dingdingbot-server.git
$ cd mcp-dingdingbot-server && make build
$ sudo ln -s $PWD/dist/mcp-dingdingbot-server_xxx_xxxx /usr/local/bin/mcp-dingdingbot-server

# "$PWD/dist/mcp-dingdingbot-server_xxx_xxxx" 替换为实际的二进制文件名

# 您也可以下载并使用预编译的发布二进制包。
```

### 配置

```json
{
  "mcpServers": {
    "mcp-dingdingbot-server": {
      "command": "mcp-dingdingbot-server",
      "env": {
        "DINGDING_BOT_WEBHOOK_KEY": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxx",
        "DINGDING_BOT_SIGN_KEY": "SECxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
      }
    }
  }
}
```

## 环境变量

- `DINGDING_BOT_WEBHOOK_KEY`: 钉钉机器人的 webhook 密钥。这是必需的。
- `DINGDING_BOT_SIGN_KEY`: 钉钉机器人签名验证的签名密钥。这是可选的，但建议用于增强安全性。

### 使用方法

- **send_text**

向钉钉群组发送文本消息

- **send_markdown**

向钉钉群组发送 markdown 消息

- **send_image**

向钉钉群组发送图片消息

- **send_news**

向钉钉群组发送图文消息，图文消息包括标题、描述、URL 和图片 URL

- **send_template_card**

向钉钉群组发送模板卡片消息

- **upload_file**

上传文件到钉钉

### 钉钉机器人

钉钉群机器人配置指南可参考：
https://open.dingtalk.com/document/robots/custom-robot-access

> DINGDING_BOT_WEBHOOK_KEY 是机器人的 webhook 密钥<br>例如：
> https://oapi.dingtalk.com/robot/send?access_token=693axxx6-7aoc-4bc4-97a0-0ec2sifa5aaa <br>
> "693axxx6-7aoc-4bc4-97a0-0ec2sifa5aaa" 是您自己的 DINGDING_BOT_WEBHOOK_KEY
>
> DINGDING_BOT_SIGN_KEY 是用于增强安全性的签名密钥<br>
> 当在钉钉机器人安全设置中启用时，您需要提供此密钥来验证请求。
> 签名验证使用 HMAC-SHA256 算法，结合时间戳和密钥。
