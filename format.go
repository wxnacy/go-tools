package tools

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"text/template"
	"time"
)

// format size to string
func FormatSize(size int64) string {
	suffixs := []string{"B", "K", "M", "G", "T"}
	scale := float64(1024)
	for i, suffix := range suffixs {
		suffix_size := math.Pow(scale, float64(i+1))
		if size < int64(suffix_size) {
			beichu := math.Pow(scale, float64(i))
			if beichu == 0 {
				beichu = 1
			}
			res := float64(size) / beichu
			if float64(int64(res)) == res {
				return fmt.Sprintf("%d%s", int64(res), suffix)
			} else {
				return fmt.Sprintf("%.2f%s", res, suffix)
			}
		}
	}
	return fmt.Sprintf("%dB", size)
}

// format float digit
func FormatFloat(f float64, digit int16) float64 {
	fmt_str := fmt.Sprintf("%%.%df", digit)
	res, _ := strconv.ParseFloat(fmt.Sprintf(fmt_str, f), 64)
	return res
}

// format template
func FormatTemplate(tpl string, i interface{}) string {
	tmpl, _ := template.New(IDGen()).Parse(tpl)
	buf := new(strings.Builder)
	_ = tmpl.Execute(buf, i)
	return buf.String()
}

func FormatTimeDT(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
