package handler

import (
	"context"
	"crypto/sha512"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"strings"
	"time"

	"github.com/anaskhan96/go-password-encoder"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"mall/mall-user/global"
	"mall/mall-user/model"
	"mall/mall-user/proto"
)

type UserServer struct {
}

//model->response
//对于没有默认值的字段 不能直接赋nil
func ModelToResponse(user model.User) proto.UserInfoResponse {
	userInfoResp := proto.UserInfoResponse{
		Id:       user.ID,
		Password: user.Password,
		Nickname: user.NickName,
		Gender:   int32(user.Gender),
		Role:     int32(user.Role),
		Mobile:   user.Mobile,
	}
	if user.Birthday != nil {
		userInfoResp.Birthday = int32(user.Birthday.Unix())
	}
	return userInfoResp
}

//分页通用方法
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

//获取用户列表
func (s *UserServer) GetUserList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	var users []model.User
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	resp := &proto.UserListResponse{}
	resp.Total = int32(result.RowsAffected)
	global.DB.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&users)

	for _, user := range users {
		userInfoResp := ModelToResponse(user)
		resp.Data = append(resp.Data, &userInfoResp)
	}

	return resp, nil
}

//通过电话号码查询
func (s *UserServer) GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	//没有查询到
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "该用户还未注册，请检查电话号码是否正确或立即注册")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	resp := ModelToResponse(user)
	return &resp, nil
}

//通过id查询
func (s *UserServer) GetUserById(ctx context.Context, req *proto.IDRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.First(&user, req.Id)
	//没有查询到
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	resp := ModelToResponse(user)
	return &resp, nil
}

//添加一个用户(注册)
func (s *UserServer) AddUser(ctx context.Context, req *proto.AddUserRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected > 0 {
		return nil, status.Errorf(codes.AlreadyExists, "手机号已被注册，请直接登录")
	}
	user.Mobile = req.Mobile
	user.NickName = req.Nickname

	//密码md5加密 格式：$加密方式$盐$加密后的密码
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode(req.Password, options)
	user.Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)

	result = global.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "新增user出现错误"+result.Error.Error())
	}

	resp := ModelToResponse(user)
	return &resp, nil
}

//更新用户信息
func (s *UserServer) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*emptypb.Empty, error) {
	var user model.User
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	//int64转为time类型
	birthday := time.Unix(int64(req.Birthday), 0)
	user.NickName = req.Nickname
	user.Birthday = &birthday
	user.Gender = int(req.Gender)

	result = global.DB.Save(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &emptypb.Empty{}, nil
}

//密码校验
func (s *UserServer) CheckPassword(ctx context.Context, req *proto.CheckPasswordRequest) (*proto.CheckResponse, error) {
	options := &password.Options{16, 100, 32, sha512.New}
	passwordSli := strings.Split(req.EncryptedPwd, "$")
	check := password.Verify(req.Password, passwordSli[2], passwordSli[3], options)
	return &proto.CheckResponse{Success: check}, nil
}
