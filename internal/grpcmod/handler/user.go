package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/sociosarbis/grpc/boilerplate/internal/dal/dao"
	"github.com/sociosarbis/grpc/boilerplate/internal/jwtgo"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/logger"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/slice"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/res"
	"github.com/sociosarbis/grpc/boilerplate/proto"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
		Id:     uint32(user.ID),
		Name:   user.Name,
		Groups: slice.Map(user.Groups, toGroupProto),
	}, nil
}

func (u User) toUserLoginRes(user dao.User) (*proto.UserLoginRes, error) {

	token, err := u.jwtManager.Generate(jwtgo.User{
		ID:    uint32(user.ID),
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

func (u User) Login(ctx context.Context, req *proto.UserLoginReq) (*proto.UserLoginRes, error) {
	var user dao.User
	err := u.db.Where("user.name = ?", req.Name).Preload("User").First(&user).Error
	if err != nil {
		logger.Err(err, "User.Login")
		return nil, errgo.Wrap(err, "User.Login")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		logger.Err(err, "incorrect password")
		return nil, errgo.Wrap(err, "User.Login")
	}
	return u.toUserLoginRes(user)
}

func (u User) LoginMs(ctx context.Context, req *proto.UserMsLoginReq) (*proto.UserLoginRes, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://graph.microsoft.com/v1.0/me", nil)
	if err != nil {
		return nil, errgo.Wrap(err, "http.NewRequestWithContext")
	}
	request.Header.Add("Authorization", "Bearer "+req.Token)
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, errgo.Wrap(err, "client.Do")
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errgo.Wrap(err, "io.ReadAll")
	}
	msUser := res.MsUser{}
	err = json.Unmarshal(data, &msUser)
	if err != nil {
		return nil, errgo.Wrap(err, "json.Unmarshal")
	}
	msUserDao := dao.MsUser{}
	db := u.db.WithContext(ctx)
	var user dao.User
	err = db.Where("ms_user.id = ?", msUser.Id).Preload("User").First(&msUserDao).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, errgo.Wrap(err, "find ms_user")
		}
		user.Name = msUser.DisplayName
		user.Email = msUser.Mail
		err = db.Transaction(func(tx *gorm.DB) error {
			err = tx.Create(&user).Error
			if err != nil {
				return errgo.Wrap(err, "CreateUser")
			}
			msUserDao.ID = msUser.Id
			msUserDao.User = dao.User{
				Model: gorm.Model{
					ID: user.ID,
				},
			}
			err = tx.Create(&msUserDao).Error
			if err != nil {
				return errgo.Wrap(err, "CreateMsUser")
			}
			return nil
		})
		if err != nil {
			return nil, errgo.Wrap(err, "db.Transaction")
		}
	} else {
		user = msUserDao.User
	}
	return u.toUserLoginRes(user)
}
