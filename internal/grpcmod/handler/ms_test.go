package handler_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/test"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/res"
	"github.com/stretchr/testify/require"
)

func TestGetOneDriveUser(t *testing.T) {
	token := os.Getenv("MS_ACCESS_TOKEN")
	req, err := http.NewRequest("GET", "https://graph.microsoft.com/v1.0/me", nil)
	require.NoError(t, err)
	req.Header.Add("Authorization", "Bearer "+token)
	client := http.Client{}
	r, err := client.Do(req)
	require.NoError(t, err)
	data := test.ReadJSONResponse(t, r, &res.MsUser{})
	require.NotNilf(t, data, "%v", data)

}
