package service

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"golang_grpc_testovoe/lib/directory"
	"golang_grpc_testovoe/server/config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type DirectoryInfoService struct {
}

func NewDirectoryIndoService() *DirectoryInfoService {
	return &DirectoryInfoService{}
}
func (d DirectoryInfoService) GetInfoFromPath(path string) (directory.DirectoryResponse, error) {

	var fileCollection []*directory.FileCollection
	var dirCollection []*directory.DirCollection
	var resp directory.DirectoryResponse

	flag, val := d.checkCachePath(path)
	if flag {
		_ = json.Unmarshal([]byte(val), &resp)
		return resp, nil
	}
	items, err := os.ReadDir(path)
	if err != nil {
		return directory.DirectoryResponse{}, err

	}
	if len(items) == 0 {
		return directory.DirectoryResponse{}, err
	}
	for _, i := range items {
		name := i.Name()
		if i.IsDir() {
			size, err := d.getDirectorySize(path + "/" + i.Name())
			if err != nil {
				return directory.DirectoryResponse{}, err
			}

			dirCollection = append(dirCollection, &directory.DirCollection{
				Name: i.Name(),
				Size: strconv.Itoa(int(size)) + " bytes",
			})
		} else {
			fileInfo, _ := os.Stat(path + name)
			fileSize := fileInfo.Size()
			fileCollection = append(fileCollection, &directory.FileCollection{
				Name: i.Name(),
				Size: strconv.Itoa(int(fileSize)) + " bytes",
			})
		}

	}
	resp = directory.DirectoryResponse{
		Files:       fileCollection,
		Directories: dirCollection,
	}
	err = d.newCachePath(path, &resp)
	if err != nil {
		return directory.DirectoryResponse{}, err
	}

	return resp, nil

}

func (d DirectoryInfoService) newCachePath(path string, resp *directory.DirectoryResponse) error {
	ctx := context.Background()
	rdb := config.ApplicationConfig.GetRDb()

	val, _ := rdb.Exists(ctx, path).Result()
	if val == 1 {
		return status.Error(codes.AlreadyExists, "already sent")
	} else {
		respJSON, _ := json.Marshal(resp)
		er := rdb.Set(ctx, path, respJSON, 300*time.Second).Err()
		if er != nil {
			return er
		}
		return nil
	}
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

func (d DirectoryInfoService) checkCachePath(path string) (bool, string) {
	ctx := context.Background()
	rdb := config.ApplicationConfig.GetRDb()

	val, err := rdb.Get(ctx, path).Result()
	if err == redis.Nil {
		return false, ""
	}
	return true, val
}
