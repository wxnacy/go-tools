package signal

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// StopChan 用于接收终止信号的信道
type StopChan chan os.Signal

// NewStopChan 创建并初始化一个信号监听信道
func NewStopChan() StopChan {
	stopChan := make(StopChan, 1)
	// 监听 SIGINT 和 SIGTERM 信号
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
	return stopChan
}

// Stop 停止监听信号
func (s StopChan) Stop() {
	signal.Stop(s)
}

// Check 检查是否收到了终止信号
// 如果收到信号则返回true，否则返回false
func (s StopChan) Check() bool {
	select {
	case <-s:
		// 收到终止信号
		fmt.Println("收到终止信号，程序退出...")
		return true
	default:
		// 没有收到信号，继续执行
		return false
	}
}