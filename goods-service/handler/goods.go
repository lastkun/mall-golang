package handler

import (
	"mall/goods-service/proto"
)

type GoodsServer struct {
	//嵌入未实现的接口
	proto.UnimplementedGoodsServer
}

//商品接口实现
//func (g *GoodsServer) GoodsList(context.Context, *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error) {
//
//}
//BatchGetGoods(context.Context, *BatchGoodsIdInfo) (*GoodsListResponse, error)
//CreateGoods(context.Context, *CreateGoodsInfo) (*GoodsInfoResponse, error)
//DeleteGoods(context.Context, *DeleteGoodsInfo) (*emptypb.Empty, error)
//UpdateGoods(context.Context, *CreateGoodsInfo) (*emptypb.Empty, error)
//GetGoodsDetail(context.Context, *GoodInfoRequest) (*GoodsInfoResponse, error)
