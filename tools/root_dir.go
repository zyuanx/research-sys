package tools

import (
	"os"
	"path/filepath"
)

func GetWorkingDirPath() string {
	searchDirectory, err := os.Getwd()
	if err != nil {
		return ""
	}
	// 从当前路径往上找，第一个拥有go.mod文件的目录认为是项目的根路径
	for searchDirectory != "" {
		goModPath := filepath.Join(searchDirectory, "go.mod")
		stat, err := os.Stat(goModPath)
		if err == nil && stat != nil && !stat.IsDir() {
			break
		}
		searchDirectory = filepath.Dir(searchDirectory)
	}
	return searchDirectory
}
