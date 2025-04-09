package service

import (
	"fmt"
	"io"
	"net/http"
	"nlm/context"
	"nlm/utils"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 允许来自 localhost 的请求
		return strings.Contains(r.Host, "localhost")
	},
}

// CreateLog 创建日志文件
func CreateLog(ctx context.PipelineContext, moduleName string) (*os.File, error) {
	// 创建 log 目录
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.MkdirAll("logs", 0755)
	}

	// 创建 log 文件
	file, err := os.Create(fmt.Sprintf("logs/pipeline-%s-%s.log", ctx.Id, moduleName))
	if err != nil {
		return nil, err
	}

	return file, nil
}

// StreamLog 通过 WebSocket 流式传输日志
func StreamLog(c *gin.Context, pipelineId, moduleName string) {
	// 升级 HTTP 连接为 WebSocket
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Printf("WebSocket upgrade failed: %v\n", err)
		return
	}
	defer ws.Close()

	// 打开日志文件
	logFile, err := os.Open(fmt.Sprintf("logs/pipeline-%s-%s.log", pipelineId, moduleName))
	if err != nil {
		fmt.Printf("Failed to open log file: %v\n", err)
		ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Error: %v", err)))
		return
	}
	defer logFile.Close()

	fileSize := int64(0)

	// 创建缓冲区
	buffer := make([]byte, 1024)
	for {
		println("Checking log file...")
		// 检查文件是否有新内容
		newFileInfo, err := logFile.Stat()
		if err != nil {
			fmt.Printf("Failed to get new file info: %v\n", err)
			return
		}

		if newFileInfo.Size() > fileSize {
			// 定位到上次读取的位置
			logFile.Seek(fileSize, 0)

			// 读取新内容
			for {
				n, err := logFile.Read(buffer)
				if err == io.EOF {
					break
				}
				if err != nil {
					fmt.Printf("Failed to read log file: %v\n", err)
					return
				}

				// 发送新内容到 WebSocket
				err = ws.WriteMessage(websocket.TextMessage, buffer[:n])
				if err != nil {
					fmt.Printf("Failed to write to WebSocket: %v\n", err)
					return
				}
			}

			// 更新文件大小
			fileSize = newFileInfo.Size()
		}

		// 检查 WebSocket 心跳
		_, p, err := ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("WebSocket connection closed: %v\n", err)
			}
			break
		}
		if string(p) != "pong" {
			fmt.Printf("WebSocket connection closed: invalid heartbeat message: %s\n", string(p))
			break
		}

		// 短暂休眠，避免过于频繁的检查
		time.Sleep(1 * time.Second)
	}
}

// CleanLogs 清理 30 天前的日志
func CleanLogs() error {
	return utils.CleanOutdatedFiles("logs")
}
