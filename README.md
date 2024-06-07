> This project implements a chat relay service using Server-Sent Events (SSE). It leverages Python LangChain to handle GPT-based conversations, Golang to relay SSE messages, and a front-end to receive and display real-time updates. Ideal for applications requiring efficient real-time communication and updates.

# SSE 转发服务
SSE（Server-Sent Events）是一种服务器推送技术，允许服务器主动向客户端发送数据。与传统的轮询机制相比，SSE提供了一种更高效的方式来实现实时通信。以下是SSE的一些关键特点：
- 单向通信：SSE主要用于服务器向客户端推送数据，不支持客户端向服务器发送数据。
- 基于HTTP：SSE使用HTTP协议，客户端通过一个HTTP连接订阅服务器的事件流。
- 长连接：SSE使用长连接，一旦连接建立，服务器就可以在任何时候向客户端发送数据。
- 轻量级：SSE的实现相对简单，不需要复杂的协议或额外的库。
- 兼容性：大多数现代浏览器都支持SSE，但IE浏览器不支持。

在ChatGPT对话系统中，SSE可以用于实现实时更新。例如，当新的对话数据生成时，服务器可以立即通过SSE推送给客户端，客户端收到数据后可以立即更新UI，从而实现实时交互。这种机制可以提高用户体验，减少等待时间。

这个项目主要是基于python langchain实现了gpt对话，然后通过 golang 转发 sse，然后前端接收sse消息