package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"example.xyz/bank/internal/db"
	mockdb "example.xyz/bank/internal/db/mock"
	"example.xyz/bank/internal/util"
	"example.xyz/bank/token"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAccountAPI(t *testing.T) {
	t.Parallel()
	user, _ := randomUser(t)
	acct := randAccount(user.Username)

	tt := map[string]struct {
		accountID     int64
		buildStubs    func(s *mockdb.MockStore)
		setupAuth     func(t *testing.T, req *http.Request, maker token.Maker)
		checkResponse func(t *testing.T, rec *httptest.ResponseRecorder)
	}{
		"OK": {
			accountID: acct.ID,
			buildStubs: func(s *mockdb.MockStore) {
				s.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(acct.ID)).
					Return(acct, nil)
			},
			setupAuth: func(t *testing.T, req *http.Request, maker token.Maker) {
				addAuthorization(t, req, maker, "Bearer", user.Username, time.Minute)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, rec.Code)
				requireBodyMatchAccount(t, rec.Body, acct)
			},
		},
		"Unauthorized User": {
			accountID: acct.ID,
			buildStubs: func(s *mockdb.MockStore) {
				s.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(acct.ID)).
					Return(acct, nil)
			},
			setupAuth: func(t *testing.T, req *http.Request, maker token.Maker) {
				addAuthorization(t, req, maker, "Bearer", "UNAUTHORIZED", time.Minute)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, rec.Code)
			},
		},
		"No Authorization": {
			accountID: acct.ID,
			buildStubs: func(s *mockdb.MockStore) {
				s.EXPECT().
					GetAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			setupAuth: func(t *testing.T, req *http.Request, maker token.Maker) {
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, rec.Code)
			},
		},
		"Not Found": {
			accountID: acct.ID,
			buildStubs: func(s *mockdb.MockStore) {
				s.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(acct.ID)).
					Return(db.Account{}, sql.ErrNoRows)
			},
			setupAuth: func(t *testing.T, req *http.Request, maker token.Maker) {
				addAuthorization(t, req, maker, "Bearer", user.Username, time.Minute)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, rec.Code)
			},
		},
		"Internal Server Error": {
			accountID: acct.ID,
			buildStubs: func(s *mockdb.MockStore) {
				s.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(acct.ID)).
					Return(db.Account{}, sql.ErrConnDone)
			},
			setupAuth: func(t *testing.T, req *http.Request, maker token.Maker) {
				addAuthorization(t, req, maker, "Bearer", user.Username, time.Minute)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, rec.Code)
			},
		},
		"Invalid ID": {
			accountID: 0,
			buildStubs: func(s *mockdb.MockStore) {
				s.EXPECT().
					GetAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			setupAuth: func(t *testing.T, req *http.Request, maker token.Maker) {
				addAuthorization(t, req, maker, "Bearer", user.Username, time.Minute)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, rec.Code)
			},
		},
	}
	ctrl := gomock.NewController(t)
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)
			svr := newTestServer(t, store)
			rec := httptest.NewRecorder()
			url := fmt.Sprintf("/accounts/%d", tc.accountID)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, req, svr.tokenMaker)
			svr.router.ServeHTTP(rec, req)
			tc.checkResponse(t, rec)
		})
	}
}

func randAccount(owner string) db.Account {
	return db.Account{
		ID:       util.RandInt(1, 1000),
		Owner:    owner,
		Balance:  util.RandBalance(),
		Currency: util.RandCurrency(),
	}
}

func requireBodyMatchAccount(t *testing.T, body io.Reader, a db.Account) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var got db.Account
	require.NoError(t, json.Unmarshal(data, &got))
	require.Equal(t, a, got)
}
