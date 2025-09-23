package tools

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// 前景 背景 颜色
// ---------------------------------------
// 30  40  黑色
// 31  41  红色
// 32  42  绿色
// 33  43  黄色
// 34  44  蓝色
// 35  45  紫红色
// 36  46  青蓝色
// 37  47  白色
//
// 代码 意义
// -------------------------
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见

const (
	TextBlack = iota + 30
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextMagenta
	TextCyan
	TextWhite
)

const (
	BackgroundBlack = iota + 40
	BackgroundRed
	BackgroundGreen
	BackgroundYellow
	BackgroundBlue
	BackgroundMagenta
	BackgroundCyan
	BackgroundWhite
)

func colorize(msg string, conf, bg, text int) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, conf, bg, text, msg, 0x1B)
}

func green(msg string) string {
	return colorize(msg, 0, 0, TextGreen)
}

func yellow(msg string) string {
	return colorize(msg, 0, 0, TextYellow)
}

func red(msg string) string {
	return colorize(msg, 0, 0, TextRed)
}

// GradientProgressBar 带颜色渐变的进度条结构体
type GradientProgressBar struct {
	total       int
	current     int
	width       int
	description string
	startTime   time.Time
}

// NewGradientProgressBar 创建一个新的渐变进度条实例
func NewGradientProgressBar(description string, width int) *GradientProgressBar {
	return &GradientProgressBar{
		width:       width,
		description: description,
		startTime:   time.Now(),
	}
}

// Start 启动进度条，传入总任务数量
func (p *GradientProgressBar) Start(total int) error {
	p.total = total
	p.current = 0
	p.startTime = time.Now()
	p.render()
	return nil
}

// Increment 增加进度，每次调用增加1个单位
func (p *GradientProgressBar) Increment() {
	p.current++
	if p.current > p.total {
		p.current = p.total
	}
	p.render()
}

// SetProgress 设置当前进度值
func (p *GradientProgressBar) SetProgress(current int) {
	p.current = current
	if p.current > p.total {
		p.current = p.total
	}
	if p.current < 0 {
		p.current = 0
	}
	p.render()
}

// Finish 完成进度条显示
func (p *GradientProgressBar) Finish() {
	p.current = p.total
	p.render()
	fmt.Println() // 换行，以便后续输出不会覆盖进度条
}

// GetProgress 获取当前进度百分比 (0.0-1.0)
func (p *GradientProgressBar) GetProgress() float64 {
	if p.total == 0 {
		return 0.0
	}
	return float64(p.current) / float64(p.total)
}

// render 渲染进度条
func (p *GradientProgressBar) render() {
	// 计算进度百分比
	percentage := p.GetProgress()
	
	// 计算已完成的格子数
	filledWidth := int(math.Round(percentage * float64(p.width)))
	
	// 创建进度条字符串
	var bar strings.Builder
	
	// 添加描述信息
	bar.WriteString(p.description)
	bar.WriteString(" [")
	
	// 添加已完成部分（带颜色渐变）
	for i := 0; i < filledWidth; i++ {
		// 计算颜色渐变（从绿色到黄色再到红色）
		ratio := float64(i) / float64(p.width)
		coloredChar := p.getGradientColor(ratio, "█")
		bar.WriteString(coloredChar)
	}
	
	// 添加未完成部分
	for i := filledWidth; i < p.width; i++ {
		bar.WriteString(" ")
	}
	
	bar.WriteString("] ")
	
	// 添加百分比
	bar.WriteString(fmt.Sprintf("%3.0f%%", percentage*100))
	
	// 添加计数信息
	bar.WriteString(fmt.Sprintf(" (%d/%d)", p.current, p.total))
	
	// 添加耗时信息
	elapsed := time.Since(p.startTime).Seconds()
	bar.WriteString(fmt.Sprintf(" %.1fs", elapsed))
	
	// 输出进度条（覆盖上一行）
	fmt.Printf("\r%s", bar.String())
}

// getGradientColor 根据进度比例获取渐变颜色
func (p *GradientProgressBar) getGradientColor(ratio float64, char string) string {
	// 根据进度比例计算颜色
	// 0.0-0.5: 绿色到黄色
	// 0.5-1.0: 黄色到红色
	
	if ratio <= 0.5 {
		// 绿色到黄色的过渡
		greenRatio := ratio * 2 // 0.0-1.0
		// 绿色保持最大值，红色从0增加到最大值
		redValue := int(5 * greenRatio) // 简化的颜色计算
		switch redValue {
		case 0:
			return green(char)
		case 1, 2:
			return yellow(char)
		default:
			return yellow(char)
		}
	} else {
		// 黄色到红色的过渡
		yellowRatio := (ratio - 0.5) * 2 // 0.0-1.0
		// 红色保持最大值，绿色从最大值减少到0
		greenValue := int(5 * (1 - yellowRatio)) // 简化的颜色计算
		switch greenValue {
		case 0:
			return red(char)
		case 1, 2:
			return yellow(char)
		default:
			return yellow(char)
		}
	}
}