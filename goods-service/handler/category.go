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

//获取分类列表
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

//获取子分类 不加载多级分类
func (g *GoodsServer) GetSubCategory(ctx context.Context, req *proto.CategoryListRequest) (*proto.SubCategoryListResponse, error) {
	subCategoryResp := proto.SubCategoryListResponse{}
	var category model.Category
	var subCategoryListResp []*proto.CategoryInfoResponse
	if result := global.DB.First(&category, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "该分类不存在")
	}

	subCategoryResp.Info = &proto.CategoryInfoResponse{
		Id:             category.ID,
		Name:           category.Name,
		ParentCategory: category.ParentCategoryID,
		Level:          category.Level,
		IsTab:          category.IsTab,
	}

	var subCategoryList []model.Category
	global.DB.Where(&model.Category{ParentCategoryID: req.Id}).Find(&subCategoryList)
	for _, subCategory := range subCategoryList {
		subCategoryListResp = append(subCategoryListResp, &proto.CategoryInfoResponse{
			Id:             subCategory.ID,
			Name:           subCategory.Name,
			ParentCategory: subCategory.ParentCategoryID,
			Level:          subCategory.Level,
			IsTab:          subCategory.IsTab,
		})
	}

	return &subCategoryResp, nil
}

//CreateCategory(context.Context, *CategoryInfoRequest) (*CategoryInfoResponse, error)
//DeleteCategory(context.Context, *DeleteCategoryRequest) (*emptypb.Empty, error)
//UpdateCategory(context.Context, *CategoryInfoRequest) (*emptypb.Empty, error)
