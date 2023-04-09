package handler_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
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

	r, err := app.Test(httptest.NewRequest("GET", "/api/user/1", http.NoBody))

	r2, _ := app.Test(httptest.NewRequest("GET", "/api/user/1", http.NoBody))

	require.NoError(t, err)

	require.Equal(t, r2.StatusCode, http.StatusTooManyRequests)

	detailWebRes := res.Response[proto.UserDetailRes]{}

	data, err := io.ReadAll(r.Body)

	err = json.Unmarshal(data, &detailWebRes)

	require.NoError(t, err)

	require.Equal(t, detailWebRes.Data.Id, uint32(1))
}
