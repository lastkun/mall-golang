package handler

//轮播图接口实现
BannerList(context.Context, *emptypb.Empty) (*BannerListResponse, error)
CreateBanner(context.Context, *BannerRequest) (*BannerResponse, error)
DeleteBanner(context.Context, *BannerRequest) (*emptypb.Empty, error)
UpdateBanner(context.Context, *BannerRequest) (*emptypb.Empty, error)
