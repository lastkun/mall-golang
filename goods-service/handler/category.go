package handler

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"mall/goods-service/global"
	"mall/goods-service/model"
	"mall/goods-service/proto"
)

func (g *GoodsServer) GetAllCategorysList(ctx context.Context, req *emptypb.Empty) (*proto.CategoryListResponse, error) {
	categoryListResp := proto.CategoryListResponse{}
	var total int64
	result := global.DB.Model(&model.Category{}).Count(&total)
	if result.Error != nil {
		return nil, status.Errorf(codes.InvalidArgument, "查询分类总数出错")
	}
	categoryListResp.Total = int32(total)
	//加载三级类目 使用gorm的预加载Preload
	var categories []model.Category
	result = global.DB.Where("level=?", 1).Preload("SubCategory.SubCategory").Find(&categories)
	if result.Error != nil {
		return nil, status.Errorf(codes.InvalidArgument, "查询三级分类出错")
	}

	//转为Json格式
	bytes, _ := json.Marshal(&categories)
	categoryListResp.JsonData = string(bytes)

	return &categoryListResp, nil
}
func (g *GoodsServer) GetSubCategory(context.Context, *proto.CategoryListRequest) (*proto.SubCategoryListResponse, error) {

}

//CreateCategory(context.Context, *CategoryInfoRequest) (*CategoryInfoResponse, error)
//DeleteCategory(context.Context, *DeleteCategoryRequest) (*emptypb.Empty, error)
//UpdateCategory(context.Context, *CategoryInfoRequest) (*emptypb.Empty, error)
