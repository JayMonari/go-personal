package api

import (
	"os"
	"testing"
	"time"

	"example.xyz/bank/internal/db"
	"example.xyz/bank/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, s db.Store) *Server {
	svr, err := NewServer(
		util.Config{
			TokenSymmetricKey:   util.RandString(32),
			AccessTokenDuration: time.Minute,
		},
		s)
	require.NoError(t, err)
	return svr
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
