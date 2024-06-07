package main

import (
    "github.com/gin-gonic/gin"
    "io"
    "log"
    "net/http"
)

func main() {
    r := gin.Default()

    r.GET("/sse", func(c *gin.Context) {
        // 创建一个HTTP请求以获取Python应用的SSE数据
        resp, err := http.Get("http://localhost:5000/sse")
        if err != nil {
            log.Printf("Failed to connect to SSE server: %v", err)
            c.String(http.StatusInternalServerError, "Failed to connect to SSE server")
            return
        }
        defer resp.Body.Close()

        // 设置响应头以支持SSE
        c.Writer.Header().Set("Content-Type", "text/event-stream")
        c.Writer.Header().Set("Cache-Control", "no-cache")
        c.Writer.Header().Set("Connection", "keep-alive")

        // 将Python应用的SSE数据转发到Gin的响应中
        buffer := make([]byte, 1024)
        for {
            n, err := resp.Body.Read(buffer)
            if err != nil {
                if err == io.EOF {
                    break
                }
                log.Printf("Error reading SSE data: %v", err)
                break
            }
            if n > 0 {
                if _, err := c.Writer.Write(buffer[:n]); err != nil {
                    log.Printf("Error writing SSE data: %v", err)
                    break
                }
                c.Writer.(http.Flusher).Flush()
            }
        }
    })

    r.Run(":8080")
}
