package handler_test

import (
	"net/http"
	"testing"

	"github.com/sociosarbis/grpc/boilerplate/internal/grpcmod"
	"github.com/sociosarbis/grpc/boilerplate/internal/mocks"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/test"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/res"
	"github.com/sociosarbis/grpc/boilerplate/proto"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUserDetail(t *testing.T) {
	t.Parallel()
	userGrpcClient := mocks.NewUserServiceClient(t)

	userGrpcClient.EXPECT().UserDetail(mock.Anything, mock.Anything).Return(&proto.UserDetailRes{
		Id:   1,
		Name: "guest",
	}, nil)

	grpcClient := &grpcmod.Client{
		User: userGrpcClient,
	}

	app := test.GetWebApp(t, test.Mock{
		GrpcClient: grpcClient,
	})

	r, err := app.Test(test.NewRequest(t, "/api/user/1").Get().Build())

	r2, _ := app.Test(test.NewRequest(t, "/api/user/1").Get().Build())

	require.NoError(t, err)

	require.Equal(t, r2.StatusCode, http.StatusTooManyRequests)

	data := test.ReadJSONResponse(t, r, &res.Response[proto.UserDetailRes]{})

	require.Equal(t, data.Data.Id, uint32(1))
}
