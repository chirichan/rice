package rice

import (
	"os"
	"path/filepath"
)

// PathExists 路径是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// PathIsDir 路径是否是文件夹
func PathIsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// PathIsFile 路径是否是文件
func PathIsFile(path string) bool {
	return !PathIsDir(path)
}

// FileExists 文件是否存在
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// WriteFileSafe 安全的写入文件，如果文件已经存在则跳过写入
func WriteFileSafe(name string, data []byte) error {
	if FileExists(name) {
		return nil
	}
	dir := filepath.Dir(name)
	if !PathExists(dir) {
		if err := os.MkdirAll(dir, 0644); err != nil {
			return err
		}
	}
	return os.WriteFile(name, data, 0644)
}

// WriteFile 写入文件
func WriteFile(name string, data []byte) error {
	dir := filepath.Dir(name)
	if !PathExists(dir) {
		if err := os.MkdirAll(dir, 0644); err != nil {
			return err
		}
	}
	return os.WriteFile(name, data, 0644)
}

// AppendFile 往文件末尾追加内容
func AppendFile(name string, data []byte) error {
	if FileExists(name) {
		f, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		_, err = f.Write(data)
		if err1 := f.Close(); err1 != nil && err == nil {
			err = err1
		}
		return err
	}
	dir := filepath.Dir(name)
	if !PathExists(dir) {
		if err := os.MkdirAll(dir, 0644); err != nil {
			return err
		}
	}
	return os.WriteFile(name, data, 0644)
}
