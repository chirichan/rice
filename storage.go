package rice

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Storage interface {
	Save(name string, data []byte) error
	Append(name string, data any) error
}

type FileStorage struct {
	BasePath string
	Date     string
}

func NewFileStorage(basePath, date string) *FileStorage {
	return &FileStorage{BasePath: basePath, Date: date}
}

func (fs *FileStorage) Save(name string, data []byte) error {

	if len(fs.Date) != 8 {
		return errors.New("日期格式不对")
	}

	var (
		err      error
		path     string = filepath.Join(fs.BasePath, fs.Date[0:4], fs.Date[4:6])
		filePath string = filepath.Join(path, name)
	)

	if !isExist(path) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	err = os.WriteFile(filePath, data, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (fs *FileStorage) Append(name string, data any) error {

	if len(fs.Date) < 6 {
		return errors.New("日期格式不对")
	}

	var (
		err      error
		fileName string = fs.Date[4:6] + ".json"
		path     string = filepath.Join(fs.BasePath, fs.Date[0:4])
		filePath string = filepath.Join(path, fileName)
	)

	// 如果路径不存在，创建路径
	if !isExist(path) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// 如果文件不存在，创建文件并写入
	if !isExist(filePath) {

		var datas []any
		datas = append(datas, data)

		marshalBytes, err := json.MarshalIndent(datas, "", "\t")
		if err != nil {
			return err
		}

		err = os.WriteFile(filePath, marshalBytes, os.ModePerm)
		if err != nil {
			return err
		}

		return nil
	}

	b1, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// 如果文件里没有任何内容，写入
	if len(b1) == 0 {

		var datas []any
		datas = append(datas, data)

		marshalBytes, err := json.MarshalIndent(datas, "", "\t")
		if err != nil {
			return err
		}

		err = os.WriteFile(filePath, marshalBytes, os.ModePerm)
		if err != nil {
			return err
		}

		return nil
	}

	var b2 []any

	err = json.Unmarshal(b1, &b2)
	if err != nil {
		return err
	}

	b2 = append(b2, data)

	b3, err := json.MarshalIndent(b2, "", "\t")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, b3, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}
