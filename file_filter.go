package tools

import (
	"fmt"
	"io"
	"os"
)

func NewFileFilter(root string, handlerFn func(paths []string) error) *filefilter {
	return &filefilter{
		root:        root,
		writer:      os.Stdout,
		handlerFunc: handlerFn,
		confirmFunc: func(s string) (bool, error) {
			var input string
			fmt.Printf("%s(y/N): ", s)
			fmt.Scanln(&input)
			if input == "y" {
				return true, nil
			} else {
				return false, nil
			}
		},
		confirmText: "是否确认",
	}
}

type ConfirmFunc func(confirm string) (bool, error)
type FileFilterFunc func(string, os.FileInfo, error) (bool, error)
type FileHandlerFunc func([]string) error

type filefilter struct {
	root        string
	withHide    bool // 带有隐藏文件
	isRecursion bool // 是否递归
	isConfirm   bool
	writer      io.Writer       //输出 writer
	filterFunc  FileFilterFunc  //过滤方法
	handlerFunc FileHandlerFunc //处理结果方法
	confirmFunc ConfirmFunc     // 确认提示方法
	confirmText string          // 提示文案
}

func (f *filefilter) WithHide() *filefilter {
	f.withHide = true
	return f
}

func (f *filefilter) EnableRecursion() *filefilter {
	f.isRecursion = true
	return f
}

func (f *filefilter) EnableConfirm() *filefilter {
	f.isConfirm = true
	return f
}

func (f *filefilter) SetPrintWriter(w io.Writer) *filefilter {
	f.writer = w
	return f
}

func (f *filefilter) SetFilter(fn func(path string, info os.FileInfo, err error) (bool, error)) *filefilter {
	f.filterFunc = fn
	return f
}

func (f *filefilter) SetConfirm(fn ConfirmFunc) *filefilter {
	f.confirmFunc = fn
	return f
}

func (f *filefilter) SetConfirmText(s string) *filefilter {
	f.confirmText = s
	return f
}

func (f *filefilter) Run() error {
	paths := make([]string, 0)
	var err error
	var flag bool
	err = FileList(f.root, f.withHide, f.isRecursion, func(p string, info os.FileInfo, err error) error {
		if f.filterFunc != nil {
			flag, err = f.filterFunc(p, info, err)
			if err != nil {
				return err
			}
		} else {
			flag = true
		}
		if flag {
			paths = append(paths, p)
		}
		return nil
	})
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		// fmt.Fprintln(f.writer, "没有找到符合要求的地址")
		return nil
	}

	if f.isConfirm {
		fmt.Fprintln(f.writer, "找到如下地址:")
		for _, p := range paths {
			fmt.Fprintln(f.writer, p)
		}
		flag, err = f.confirmFunc(f.confirmText)
	}
	if flag {
		err = f.handlerFunc(paths)
		if err != nil {
			return err
		}
	}
	return nil
}
