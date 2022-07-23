package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.xyz/bank/internal/db"
	mockdb "example.xyz/bank/internal/db/mock"
	"example.xyz/bank/internal/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAccountAPI(t *testing.T) {
	t.Parallel()
	acct := randAccount()

	tt := map[string]struct {
		accountID     int64
		buildStubs    func(s *mockdb.MockStore)
		checkResponse func(t *testing.T, rec *httptest.ResponseRecorder)
	}{
		"OK": {
			accountID: acct.ID,
			buildStubs: func(s *mockdb.MockStore) {
				s.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(acct.ID)).
					Return(acct, nil)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, rec.Code)
				requireBodyMatchAccount(t, rec.Body, acct)
			},
		},
		"Not Found": {
			accountID: acct.ID,
			buildStubs: func(s *mockdb.MockStore) {
				s.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(acct.ID)).
					Return(db.Account{}, sql.ErrNoRows)
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
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, rec.Code)
				requireBodyMatchAccount(t, rec.Body, acct)
			},
		},
		"Invalid ID": {
			accountID: 0,
			buildStubs: func(s *mockdb.MockStore) {
				s.EXPECT().
					GetAccount(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, rec.Code)
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)
			svr := NewServer(store)
			rec := httptest.NewRecorder()
			url := fmt.Sprintf("/accounts/%d", tc.accountID)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			svr.router.ServeHTTP(rec, req)
			tc.checkResponse(t, rec)
		})
	}

}

func randAccount() db.Account {
	return db.Account{
		ID:       util.RandInt(1, 1000),
		Owner:    util.RandOwner(),
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
