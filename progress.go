package tools

// ProgressBar 进度条接口，定义进度条的基本操作
// 此接口设计为通用接口，方便后续更换其他进度条实现
type ProgressBar interface {
	// Start 启动进度条，传入总任务数量
	Start(total int) error

	// Increment 增加进度，每次调用增加1个单位
	Increment()

	// SetProgress 设置当前进度值
	SetProgress(current int)

	// Finish 完成进度条显示
	Finish()

	// GetProgress 获取当前进度百分比 (0.0-1.0)
	GetProgress() float64
}
