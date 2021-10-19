package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"mall/goods-service/global"
	"mall/goods-service/model"

	"mall/goods-service/proto"
)

type GoodsServer struct {
	//嵌入未实现的接口
	proto.UnimplementedGoodsServer
}

//商品接口实现
func (g *GoodsServer) GoodsList(ctx context.Context, req *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error) {
	goodsListResponse := &proto.GoodsListResponse{}
	db := global.DB.Model(model.Goods{})
	if req.KeyWords != "" {
		db.Where("name like ?", "%"+req.KeyWords+"%")
	}
	if req.IsHot {
		db.Where("is_hot = true")
	}
	if req.IsNew {
		db.Where("is_new = true")
	}
	if req.PriceMin > 0 {
		db.Where("shop_price >= ?", req.PriceMin)
	}
	if req.PriceMax > 0 {
		db.Where("shop_price <= ?", req.PriceMax)
	}
	if req.Brand > 0 {
		db.Where("brand_id = ?", req.Brand)
	}
	if req.TopCategory > 0 {
		var category model.Category
		result := global.DB.Find(&category, req.TopCategory)
		if result.RowsAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "商品分类不存在")
		}
		if category.Level == 1 {

		}
		var sqls string
		if category.Level == 1 {
			sqls = fmt.Sprintf("select id from category where parent_category_id in (select id from category WHERE parent_category_id=%d)", req.TopCategory)
		} else if category.Level == 2 {
			sqls = fmt.Sprintf("select id from category WHERE parent_category_id=%d", req.TopCategory)
		} else if category.Level == 3 {
			sqls = fmt.Sprintf("select id from category WHERE id=%d", req.TopCategory)
		}

		db.Where("category_id in ?", sqls)
	}
	//db.Find()

	return goodsListResponse, nil
}
func (g *GoodsServer) BatchGetGoods(ctx context.Context, req *proto.BatchGoodsIdInfo) (*proto.GoodsListResponse, error) {

}
func (g *GoodsServer) CreateGoods(ctx context.Context, req *proto.CreateGoodsInfo) (*proto.GoodsInfoResponse, error) {

}
func (g *GoodsServer) DeleteGoods(ctx context.Context, req *proto.DeleteGoodsInfo) (*emptypb.Empty, error) {

}
func (g *GoodsServer) UpdateGoods(ctx context.Context, req *proto.CreateGoodsInfo) (*emptypb.Empty, error) {

}
func (g *GoodsServer) GetGoodsDetail(ctx context.Context, req *proto.GoodInfoRequest) (*proto.GoodsInfoResponse, error) {

}
