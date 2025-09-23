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
	err := os.MkdirAll(subDir, 0755)
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
		testDir + "/":           false, // 目录应该以 / 结尾
		testDir + "/file1.txt":  false,
		testDir + "/subdir/":    false, // 子目录也应该以 / 结尾
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
	err := os.MkdirAll(subDir, 0755)
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