package tools

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// ZipFile 将单个文件压缩为zip文件
// src: 源文件路径
// dst: 目标zip文件路径
func ZipFile(src, dst string) (err error) {
	fw, err := os.Create(dst)
	defer fw.Close()
	if err != nil {
		return err
	}

	// 通过 fw 来创建 zip.Write
	zw := zip.NewWriter(fw)
	defer func() {
		// 检测一下是否成功关闭
		if err := zw.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	f1, err := os.Open(src)
	if err != nil {
		return err
	}

	_, fileName := filepath.Split(src)
	w1, err := zw.Create(fileName)
	if err != nil {
		return err
	}
	_, err = io.Copy(w1, f1)
	return err
}

// ZipDir 将整个目录压缩为zip文件，只保留最后一层目录名
// src: 源目录路径
// dst: 目标zip文件路径
func ZipDir(src, dst string) (err error) {
	// 创建准备写入的文件
	fw, err := os.Create(dst)
	defer fw.Close()
	if err != nil {
		return err
	}

	// 通过 fw 来创建 zip.Write
	zw := zip.NewWriter(fw)
	defer func() {
		// 检测一下是否成功关闭
		if err := zw.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// 获取源目录的最后一个目录名
	srcBase := filepath.Base(src)

	// 下面来将文件写入 zw ，因为有可能会有很多个目录及文件，所以递归处理
	return filepath.Walk(src, func(path string, fi os.FileInfo, errBack error) (err error) {
		if errBack != nil {
			return errBack
		}

		// 计算相对路径
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		// 如果是根目录，使用目录名
		if relPath == "." {
			relPath = srcBase
		} else {
			// 否则，使用根目录名加上相对路径
			relPath = filepath.Join(srcBase, relPath)
		}

		// 通过文件信息，创建 zip 的文件信息
		fh, err := zip.FileInfoHeader(fi)
		if err != nil {
			return
		}

		// 设置文件名
		fh.Name = relPath

		// 这步开始没有加，会发现解压的时候说它不是个目录
		if fi.IsDir() {
			fh.Name += "/"
		}

		// 写入文件信息，并返回一个 Write 结构
		w, err := zw.CreateHeader(fh)
		if err != nil {
			return
		}

		// 检测，如果不是标准文件就只写入头信息，不写入文件数据到 w
		// 如目录，也没有数据需要写
		if !fh.Mode().IsRegular() {
			return nil
		}

		// 打开要压缩的文件
		fr, err := os.Open(path)
		defer fr.Close()
		if err != nil {
			return
		}

		// 将打开的文件 Copy 到 w
		n, err := io.Copy(w, fr)
		if err != nil {
			return
		}
		// 输出压缩的内容
		fmt.Printf("成功压缩文件： %s, 共写入了 %d 个字符的数据\n", path, n)

		return nil
	})
}

// Unzip 解压zip文件到指定目录
// src: 源zip文件路径
// dst: 解压目标目录路径
func Unzip(src, dst string) error {
	// 打开zip文件
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	// 创建目标目录
	err = os.MkdirAll(dst, 0755)
	if err != nil {
		return err
	}

	// 遍历zip文件中的每个文件/目录
	for _, f := range r.File {
		// 构造解压后的文件路径
		fpath := filepath.Join(dst, f.Name)

		// 检查文件路径安全性，防止目录遍历漏洞
		if !strings.HasPrefix(fpath, filepath.Clean(dst)+string(os.PathSeparator)) {
			return fmt.Errorf("%s: 非法文件路径", f.Name)
		}

		if f.FileInfo().IsDir() {
			// 创建目录
			err = os.MkdirAll(fpath, f.Mode())
			if err != nil {
				return err
			}
		} else {
			// 创建文件
			err = os.MkdirAll(filepath.Dir(fpath), 0755)
			if err != nil {
				return err
			}

			// 打开zip中的文件
			rc, err := f.Open()
			if err != nil {
				return err
			}
			defer rc.Close()

			// 创建目标文件
			fw, err := os.Create(fpath)
			if err != nil {
				return err
			}
			defer fw.Close()

			// 设置文件权限
			err = os.Chmod(fpath, f.Mode())
			if err != nil {
				return err
			}

			// 复制文件内容
			_, err = io.Copy(fw, rc)
			if err != nil {
				return err
			}
		}
	}

	return nil
}