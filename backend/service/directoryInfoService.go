package service

import (
	"golang_grpc_testovoe/lib/directory"
	"os"
	"path/filepath"
	"strconv"
)

type DirectoryInfoService struct {
}

func NewDirectoryIndoService() *DirectoryInfoService {
	return &DirectoryInfoService{}
}
func (d DirectoryInfoService) GetInfoFromPath(path string) (directory.DirectoryResponse, error) {

	var fileCollection []directory.FileCollection
	var dirCollection []directory.DirectoryCollection

	items, err := os.ReadDir(path)
	if err != nil {
		return directory.DirectoryResponse{}, err

	}
	if len(items) == 0 {
		return directory.DirectoryResponse{}, nil
	}
	for _, i := range items {
		name := i.Name()
		if i.IsDir() {
			size, err := d.getDirectorySize(path)
			if err != nil {
				return directory.DirectoryResponse{}, err
			}

			dirCollection = append(dirCollection, directory.DirectoryCollection{
				Name: i.Name(),
				Size: strconv.Itoa(int(size)) + "bytes",
			})
		} else {
			fileInfo, _ := os.Stat(path + name)
			fileSize := fileInfo.Size()
			fileCollection = append(fileCollection, directory.FileCollection{
				Name: i.Name(),
				Size: strconv.Itoa(int(fileSize)) + "bytes",
			})
		}

	}
	return directory.DirectoryResponse{}, nil
}

func (d DirectoryInfoService) getDirectorySize(path string) (int64, error) {
	var size int64

	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			size += info.Size()
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return size, nil
}
