package handler

GetAllCategorysList(context.Context, *emptypb.Empty) (*CategoryListResponse, error)
GetSubCategory(context.Context, *CategoryListRequest) (*SubCategoryListResponse, error)
CreateCategory(context.Context, *CategoryInfoRequest) (*CategoryInfoResponse, error)
DeleteCategory(context.Context, *DeleteCategoryRequest) (*emptypb.Empty, error)
UpdateCategory(context.Context, *CategoryInfoRequest) (*emptypb.Empty, error)
