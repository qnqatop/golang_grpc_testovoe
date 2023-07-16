package controller

import (
	"context"
	"golang_grpc_testovoe/backend/service"
	"golang_grpc_testovoe/lib/directory"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DirectoryInfoController struct {
	directory.UnimplementedDirectoryInfoServiceServer
	directoryInfoService *service.DirectoryInfoService
}

func NewDirectoryInfoController(directoryInfoService *service.DirectoryInfoService) *DirectoryInfoController {
	return &DirectoryInfoController{directoryInfoService: directoryInfoService}
}
func (d *DirectoryInfoController) GetInformationForPath(ctx context.Context, request *directory.DirectoryRequest) (*directory.DirectoryResponse, error) {
	if request.Path == "" {
		return &directory.DirectoryResponse{}, status.Error(codes.InvalidArgument, "path is empty")
	}

	path := request.Path

	response, err := d.directoryInfoService.GetInfoFromPath(path)

	if err != nil {
		return &directory.DirectoryResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &response, nil
}
