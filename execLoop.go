package tools

import (
	"log"
	"time"

	"github.com/wxnacy/go-tools/signal"
)

type (
	ExecOnceFunc func(total, index int, item any) error
	FilterFunc   func(item any) (bool, error)
	OnlyPrint    bool
	Printf       func(format string, v ...any)
)

// 执行循环的公共方法
func ExecLoop(items []any, maxCount int, ExecFunc ExecOnceFunc, args ...any) error {
	var progressBar ProgressBar // 进度条
	uPrintf := log.Printf
	for _, arg := range args {
		switch val := arg.(type) {
		case Printf:
			uPrintf = val
		case ProgressBar:
			progressBar = val
		}
	}

	begin := time.Now()
	total := len(items)
	// 创建一个信道来接收终止信号
	stopChan := signal.NewStopChan()
	// 确保在函数退出时停止监听信号
	defer stopChan.Stop()

	var isOnlyPrint OnlyPrint

	// 检查 args 参数列表中是否有 FilterFunc 类型的方法，如果有对 items 列表进行过滤
	for _, arg := range args {
		switch val := arg.(type) {
		case FilterFunc:
			uPrintf("过滤前列表 %d 个项目", total)
			filteredItems := make([]any, 0)
			for _, item := range items {
				flag, err := val(item)
				if err != nil {
					return err
				}
				if flag {
					filteredItems = append(filteredItems, item)
				}
			}
			items = filteredItems
			total = len(items)
			uPrintf("过滤后剩余 %d 个项目", total)
		case OnlyPrint:
			isOnlyPrint = val
			if isOnlyPrint {
				uPrintf("只打印不执行")
			}
		}
	}
	if progressBar != nil {
		progressBar.Start(len(items))
		defer progressBar.Finish()
	}

	uPrintf("程序开始运行，按 CTRL+C 会在单次任务结束后停止...")
	for i, item := range items {
		// 检查是否收到了终止信号
		if stopChan.Check() {
			// 收到终止信号，退出循环
			break
		}

		if isOnlyPrint {
			uPrintf("[%2d/%d] item: %#v", i, total, item)
		} else {
			err := ExecFunc(total, i, item)
			if err != nil {
				return err
			}
		}
		// 更新进度条
		if progressBar != nil {
			progressBar.Increment()
		}

		// 再次检查是否收到了终止信号（在_run执行后）
		if stopChan.Check() {
			// 收到终止信号，退出循环
			break
		}

		// 运行个数符合退出条件
		if maxCount > 0 && i == maxCount-1 {
			uPrintf("运行%d个提前终结", maxCount)
			break
		}

	}
	uPrintf("循环任务耗时: %v", time.Since(begin))
	return nil
}
