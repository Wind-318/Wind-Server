package algorithm

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// 压缩    source: 源文件夹位置    target：目标位置
func Zip(source string, target string) error {
	// 删除旧文件
	os.RemoveAll(target)

	// 创建 zip 文件
	zipfile, err := os.Create(target)
	// 错误处理
	if err != nil {
		return err
	}
	defer zipfile.Close()

	// 打开 zip 文件
	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	// 遍历路径信息
	filepath.Walk(source, func(path string, info os.FileInfo, _ error) error {
		// 如果是源路径，提前进行下一个遍历
		if path == source {
			return nil
		}

		// 获取文件头信息
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			fmt.Println(err)
			return err
		}
		header.Name = strings.TrimPrefix(path, source+`\`)

		// 判断文件是不是文件夹
		if info.IsDir() {
			header.Name += `/`
		} else {
			// 设置 zip 文件压缩算法
			header.Method = zip.Deflate
		}

		// 创建压缩包头部信息
		writer, err := archive.CreateHeader(header)
		if err != nil {
			fmt.Println(err)
			return err
		}
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}
		return nil
	})

	return nil
}
