package handler

//品牌接口实现
BrandList(context.Context, *BrandFilterRequest) (*BrandListResponse, error)
CreateBrand(context.Context, *BrandRequest) (*BrandInfoResponse, error)
DeleteBrand(context.Context, *BrandRequest) (*emptypb.Empty, error)
UpdateBrand(context.Context, *BrandRequest) (*emptypb.Empty, error)
