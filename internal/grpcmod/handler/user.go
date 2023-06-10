package handler

import (
	"context"

	"gorm.io/gorm"

	"github.com/sociosarbis/grpc/boilerplate/internal/dal/dao"
	"github.com/sociosarbis/grpc/boilerplate/internal/jwtgo"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/logger"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/slice"
	"github.com/sociosarbis/grpc/boilerplate/proto"
)

type User struct {
	db         *gorm.DB
	jwtManager *jwtgo.JWTManager
}

func NewUser(db *gorm.DB, jwtManager *jwtgo.JWTManager) User {
	return User{
		db,
		jwtManager,
	}
}

func toRoleProto(role dao.Role) *proto.Role {
	return &proto.Role{
		Id:   role.ID,
		Name: role.Name,
	}
}

func toGroupProto(group dao.Group) *proto.Group {
	return &proto.Group{
		Id:       group.ID,
		Name:     group.Name,
		Children: slice.Map(group.Children, toGroupProto),
		Roles:    slice.Map(group.Roles, toRoleProto),
	}
}

func (u User) Detail(ctx context.Context, req *proto.UserDetailReq) (*proto.UserDetailRes, error) {
	var user dao.User

	err := u.db.Where("user.id = ?", req.Id).
		Preload("Groups", "group.parentId is null").
		Preload("Groups.Children").
		Preload("Groups.Roles").
		First(&user).Error

	if err != nil {
		logger.Err(err, "server.Detail")
		return nil, errgo.Wrap(err, "server.Detail")
	}

	return &proto.UserDetailRes{
		Id:     user.ID,
		Name:   user.Name,
		Groups: slice.Map(user.Groups, toGroupProto),
	}, nil
}

func (u User) Login(ctx context.Context, req *proto.UserLoginReq) (*proto.UserLoginRes, error) {
	var user dao.User
	err := u.db.Where("user.name = ?", req.Name).First(&user).Error
	if err != nil {
		logger.Err(err, "User.Login")
		return nil, errgo.Wrap(err, "User.Login")
	}
	if user.Password != req.Password {
		logger.Err(err, "incorrect password")
		return nil, errgo.Wrap(err, "User.Login")
	}
	token, err := u.jwtManager.Generate(jwtgo.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
	if err != nil {
		logger.Err(err, "jwtManager.Generate")
		return nil, errgo.Wrap(err, "jwtManager.Generate")
	}
	return &proto.UserLoginRes{
		Code:  "0",
		Msg:   "ok",
		Token: &token,
	}, nil
}
