package handler

//品牌分类接口实现
CategoryBrandList(context.Context, *CategoryBrandFilterRequest) (*CategoryBrandListResponse, error)
GetCategoryBrandList(context.Context, *CategoryInfoRequest) (*BrandListResponse, error)
CreateCategoryBrand(context.Context, *CategoryBrandRequest) (*CategoryBrandResponse, error)
DeleteCategoryBrand(context.Context, *CategoryBrandRequest) (*emptypb.Empty, error)
UpdateCategoryBrand(context.Context, *CategoryBrandRequest) (*emptypb.Empty, error)
