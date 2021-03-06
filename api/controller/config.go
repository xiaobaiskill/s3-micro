package controller

import (
	"golang.org/x/net/context"

	"github.com/lxlxw/s3-micro/api/service"

	pb "github.com/lxlxw/s3-micro/proto"
)

// Updates S3 config info.
func (h ApiService) UpdateConfigStoreInfo(ctx context.Context, r *pb.UpdateConfigStoreInfoRequest) (*pb.UpdateConfigStoreInfoResponse, error) {
	res := service.UpdateConfigStoreInfo(r)
	return &res, nil
}

// Gets S3 config info.
func (h ApiService) GetConfigStoreInfo(ctx context.Context, r *pb.GetConfigStoreInfoRequest) (*pb.GetConfigStoreInfoResponse, error) {
	res := service.GetConfigStoreInfo(r)
	return &res, nil
}
