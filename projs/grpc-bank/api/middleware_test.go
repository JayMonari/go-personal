package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"example.xyz/bank/token"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func addAuthorization(
	t *testing.T,
	req *http.Request,
	maker token.Maker,
	authType string,
	username string,
	d time.Duration,
) {
	tok, err := maker.CreateToken(username, d)
	require.NoError(t, err)
	authHeader := fmt.Sprintf("%s %s", authType, tok)
	req.Header.Set("Authorization", authHeader)
}

func TestAuthMiddleware(t *testing.T) {
	t.Parallel()
	tt := map[string]struct {
		setupAuth     func(t *testing.T, req *http.Request, maker token.Maker)
		checkResponse func(t *testing.T, rr *httptest.ResponseRecorder)
	}{
		"OK": {
			setupAuth: func(t *testing.T, req *http.Request, maker token.Maker) {
				addAuthorization(t, req, maker, "Bearer", "user", time.Minute)
			},
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, rr.Code)
			},
		},
		"No Authorization": {
			setupAuth: func(t *testing.T, req *http.Request, maker token.Maker) {
			},
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, rr.Code)
			},
		},
		"Unsupported Authorization": {
			setupAuth: func(t *testing.T, req *http.Request, maker token.Maker) {
				addAuthorization(t, req, maker, "UNSUPPORTED", "user", time.Minute)
			},
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, rr.Code)
			},
		},
		"Invalid Authorization Format": {
			setupAuth: func(t *testing.T, req *http.Request, maker token.Maker) {
				addAuthorization(t, req, maker, "", "user", time.Minute)
			},
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, rr.Code)
			},
		},
		"Expired Token": {
			setupAuth: func(t *testing.T, req *http.Request, maker token.Maker) {
				addAuthorization(t, req, maker, "Bearer", "user", -time.Minute)
			},
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, rr.Code)
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			svr := newTestServer(t, nil)
			authPath := "/auth"
			svr.router.GET(authPath, authMiddleware(svr.tokenMaker), func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{})
			})

			rr := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, authPath, nil)
			require.NoError(t, err)

			tc.setupAuth(t, req, svr.tokenMaker)
			svr.router.ServeHTTP(rr, req)
			tc.checkResponse(t, rr)
		})
	}
}
