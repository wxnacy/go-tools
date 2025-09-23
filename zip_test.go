package tools

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestZipDir(t *testing.T) {
	// 创建测试目录结构
	testDir := "test_dir"
	subDir := filepath.Join(testDir, "subdir")

	// 清理之前的测试数据
	os.RemoveAll(testDir)
	os.Remove("test.zip")

	// 创建测试目录和文件
	err := os.MkdirAll(subDir, 0o755)
	if err != nil {
		t.Fatal(err)
	}

	// 创建测试文件
	file1, err := os.Create(filepath.Join(testDir, "file1.txt"))
	if err != nil {
		t.Fatal(err)
	}
	file1.WriteString("content1")
	file1.Close()

	file2, err := os.Create(filepath.Join(subDir, "file2.txt"))
	if err != nil {
		t.Fatal(err)
	}
	file2.WriteString("content2")
	file2.Close()

	// 执行压缩
	err = ZipDir(testDir, "test.zip")
	if err != nil {
		t.Fatal(err)
	}

	// 检查压缩文件
	r, err := zip.OpenReader("test.zip")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	// 验证压缩文件中的路径是否正确
	expectedFiles := map[string]bool{
		testDir + "/":                 false, // 目录应该以 / 结尾
		testDir + "/file1.txt":        false,
		testDir + "/subdir/":          false, // 子目录也应该以 / 结尾
		testDir + "/subdir/file2.txt": false,
	}

	for _, f := range r.File {
		expectedFiles[f.Name] = true
	}

	// 检查所有预期的文件是否都存在
	for name, found := range expectedFiles {
		if !found {
			t.Errorf("Expected file %s not found in zip", name)
		}
	}

	// 验证文件内容
	for _, f := range r.File {
		if f.FileInfo().IsDir() {
			continue
		}

		rc, err := f.Open()
		if err != nil {
			t.Fatal(err)
		}

		content, err := io.ReadAll(rc)
		rc.Close()
		if err != nil {
			t.Fatal(err)
		}

		switch f.Name {
		case testDir + "/file1.txt":
			if string(content) != "content1" {
				t.Errorf("file1.txt content mismatch: got %s, want %s", string(content), "content1")
			}
		case testDir + "/subdir/file2.txt":
			if string(content) != "content2" {
				t.Errorf("file2.txt content mismatch: got %s, want %s", string(content), "content2")
			}
		}
	}

	// 清理测试数据
	os.RemoveAll(testDir)
	os.Remove("test.zip")
}

func TestUnzip(t *testing.T) {
	// 创建测试目录结构
	testDir := "test_unzip"
	subDir := filepath.Join(testDir, "subdir")

	// 清理之前的测试数据
	os.RemoveAll(testDir)
	os.Remove("test_unzip.zip")

	// 创建测试目录和文件
	err := os.MkdirAll(subDir, 0o755)
	if err != nil {
		t.Fatal(err)
	}

	// 创建测试文件
	file1, err := os.Create(filepath.Join(testDir, "file1.txt"))
	if err != nil {
		t.Fatal(err)
	}
	file1.WriteString("content1")
	file1.Close()

	file2, err := os.Create(filepath.Join(subDir, "file2.txt"))
	if err != nil {
		t.Fatal(err)
	}
	file2.WriteString("content2")
	file2.Close()

	// 执行压缩
	err = ZipDir(testDir, "test_unzip.zip")
	if err != nil {
		t.Fatal(err)
	}

	// 执行解压
	extractDir := "extracted"
	err = Unzip("test_unzip.zip", extractDir)
	if err != nil {
		t.Fatal(err)
	}

	// 验证解压后的文件结构和内容
	// 检查根目录文件
	content1, err := os.ReadFile(filepath.Join(extractDir, testDir, "file1.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if string(content1) != "content1" {
		t.Errorf("file1.txt content mismatch: got %s, want %s", string(content1), "content1")
	}

	// 检查子目录文件
	content2, err := os.ReadFile(filepath.Join(extractDir, testDir, "subdir", "file2.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if string(content2) != "content2" {
		t.Errorf("file2.txt content mismatch: got %s, want %s", string(content2), "content2")
	}

	// 清理测试数据
	os.RemoveAll(testDir)
	os.Remove("test_unzip.zip")
	os.RemoveAll(extractDir)
}

func TestZip(t *testing.T) {
	// 测试Zip函数处理目录的情况
	testDir := "test_zip_dir"
	subDir := filepath.Join(testDir, "subdir")

	// 清理之前的测试数据
	os.RemoveAll(testDir)
	os.Remove("test_zip_dir.zip")

	// 创建测试目录和文件
	err := os.MkdirAll(subDir, 0o755)
	if err != nil {
		t.Fatal(err)
	}

	// 创建测试文件
	file1, err := os.Create(filepath.Join(testDir, "file1.txt"))
	if err != nil {
		t.Fatal(err)
	}
	file1.WriteString("content1")
	file1.Close()

	file2, err := os.Create(filepath.Join(subDir, "file2.txt"))
	if err != nil {
		t.Fatal(err)
	}
	file2.WriteString("content2")
	file2.Close()

	// 使用Zip函数压缩目录
	err = Zip(testDir, "test_zip_dir.zip")
	if err != nil {
		t.Fatal(err)
	}

	// 检查压缩文件
	r, err := zip.OpenReader("test_zip_dir.zip")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	// 验证压缩文件中的路径是否正确
	expectedFiles := map[string]bool{
		testDir + "/":                 false, // 目录应该以 / 结尾
		testDir + "/file1.txt":        false,
		testDir + "/subdir/":          false, // 子目录也应该以 / 结尾
		testDir + "/subdir/file2.txt": false,
	}

	for _, f := range r.File {
		expectedFiles[f.Name] = true
	}

	// 检查所有预期的文件是否都存在
	for name, found := range expectedFiles {
		if !found {
			t.Errorf("Expected file %s not found in zip", name)
		}
	}

	// 清理测试数据
	os.RemoveAll(testDir)
	os.Remove("test_zip_dir.zip")

	// 测试Zip函数处理单个文件的情况
	testFile := "test_file.txt"

	// 创建测试文件
	file, err := os.Create(testFile)
	if err != nil {
		t.Fatal(err)
	}
	file.WriteString("single file content")
	file.Close()

	// 使用Zip函数压缩单个文件
	err = Zip(testFile, "test_file.zip")
	if err != nil {
		t.Fatal(err)
	}

	// 检查压缩文件
	r, err = zip.OpenReader("test_file.zip")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	// 验证压缩文件中是否包含正确的文件
	if len(r.File) != 1 {
		t.Errorf("Expected 1 file in zip, got %d", len(r.File))
	} else if r.File[0].Name != testFile {
		t.Errorf("Expected file name %s, got %s", testFile, r.File[0].Name)
	}

	// 验证文件内容
	rc, err := r.File[0].Open()
	if err != nil {
		t.Fatal(err)
	}
	content, err := io.ReadAll(rc)
	rc.Close()
	if err != nil {
		t.Fatal(err)
	}
	if string(content) != "single file content" {
		t.Errorf("File content mismatch: got %s, want %s", string(content), "single file content")
	}

	// 清理测试数据
	os.Remove(testFile)
	os.Remove("test_file.zip")
}

// simpleProgressBar 简单的进度条实现，用于测试
type simpleProgressBar struct {
	total   int
	current int
}

func (p *simpleProgressBar) Start(total int) error {
	p.total = total
	p.current = 0
	return nil
}

func (p *simpleProgressBar) Increment() {
	p.current++
}

func (p *simpleProgressBar) SetProgress(current int) {
	p.current = current
}

func (p *simpleProgressBar) Finish() {
	p.current = p.total
}

func (p *simpleProgressBar) GetProgress() float64 {
	if p.total == 0 {
		return 0.0
	}
	return float64(p.current) / float64(p.total)
}

func TestZipWithProgressBar(t *testing.T) {
	// 测试带进度条的ZipDir函数
	testDir := "test_progress_dir"
	subDir := filepath.Join(testDir, "subdir")

	// 清理之前的测试数据
	os.RemoveAll(testDir)
	os.Remove("test_progress.zip")

	// 创建测试目录和文件
	err := os.MkdirAll(subDir, 0o755)
	if err != nil {
		t.Fatal(err)
	}

	// 创建测试文件
	file1, err := os.Create(filepath.Join(testDir, "file1.txt"))
	if err != nil {
		t.Fatal(err)
	}
	file1.WriteString("content1")
	file1.Close()

	file2, err := os.Create(filepath.Join(subDir, "file2.txt"))
	if err != nil {
		t.Fatal(err)
	}
	file2.WriteString("content2")
	file2.Close()

	// 创建进度条实例
	progressBar := &simpleProgressBar{}

	// 执行带进度条的压缩
	err = ZipDir(testDir, "test_progress.zip", progressBar)
	if err != nil {
		t.Fatal(err)
	}

	// 验证进度条是否完成
	if progressBar.current != progressBar.total {
		t.Errorf("Progress bar not finished: current=%d, total=%d", progressBar.current, progressBar.total)
	}

	// 清理测试数据
	os.RemoveAll(testDir)
	os.Remove("test_progress.zip")
}
