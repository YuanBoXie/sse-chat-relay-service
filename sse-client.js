const EventSource = require('eventsource');

const url = 'http://localhost:8080/sse';

const eventSource = new EventSource(url);

// 处理消息事件
eventSource.onmessage = function(event) {
    console.log('New message:', event.data);
};

// 处理错误事件
eventSource.onerror = function(event) {
    console.error('Error:', event);
};

// 处理连接打开事件
eventSource.onopen = function() {
    console.log('Connection to SSE server opened.');
};
