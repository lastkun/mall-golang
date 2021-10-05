package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"mall/goods-service/global"
	"mall/goods-service/model"
	"mall/goods-service/proto"
)

//品牌接口实现
//查询品牌列表
func (g *GoodsServer) BrandList(ctx context.Context, req *proto.BrandFilterRequest) (*proto.BrandListResponse, error) {
	brandListResp := proto.BrandListResponse{}
	var brandData []*proto.BrandInfoResponse

	//获取总数
	var total int64
	result := global.DB.Model(&model.Brands{}).Count(&total)
	if result.Error != nil {
		return nil, result.Error
	}
	brandListResp.Total = int32(total)

	//分页查询
	var brandList []model.Brands
	result = global.DB.Scopes(Paginate(req.Pages, req.PagePerNums)).Find(&brandList)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, brand := range brandList {
		brandData = append(brandData, &proto.BrandInfoResponse{
			Id:   brand.ID,
			Name: brand.Name,
			Logo: brand.Logo,
		})
	}

	brandListResp.Data = brandData
	return &brandListResp, nil
}
func (g *GoodsServer) CreateBrand(ctx context.Context, req *proto.BrandRequest) (*proto.BrandInfoResponse, error) {
	result := global.DB.Where("name=?", req.Name).First(&model.Brands{})
	if result.RowsAffected != 0 {
		return nil, status.Errorf(codes.InvalidArgument, "存在同名品牌，无法新增")
	}

	brand := &model.Brands{
		Name: req.Name,
		Logo: req.Logo,
	}

	global.DB.Save(&brand)

	return &proto.BrandInfoResponse{Id: brand.ID}, nil
}
func (g *GoodsServer) DeleteBrand(ctx context.Context, req *proto.BrandRequest) (*emptypb.Empty, error) {
	result := global.DB.Delete(&model.Brands{}, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "不存在该品牌，无法删除")
	}

	return &emptypb.Empty{}, nil
}

func (g *GoodsServer) UpdateBrand(ctx context.Context, req *proto.BrandRequest) (*emptypb.Empty, error) {
	brand := model.Brands{}
	result := global.DB.Where("id=?", req.Id).First(brand)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "不存在该品牌，无法修改")
	}

	if req.Name != "" {
		brand.Name = req.Name
	}

	if req.Logo != "" {
		brand.Logo = req.Logo
	}

	global.DB.Save(&brand)
	return &emptypb.Empty{}, nil
}
