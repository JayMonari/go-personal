package api

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"example.xyz/bank/internal/db"
	"example.xyz/bank/token"
	"github.com/gin-gonic/gin"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (s *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fromAcct, valid := s.validAccount(ctx, req.FromAccountID, req.Currency)
	if !valid {
		return
	}
	authPayload := ctx.MustGet(authPayloadKey).(*token.Payload)
	log.Println(fromAcct.Owner, authPayload.Username)
	if fromAcct.Owner != authPayload.Username {
		ctx.JSON(http.StatusForbidden, errors.New("from account doesn't belong to the authenticated user"))
		return
	}
	_, valid = s.validAccount(ctx, req.ToAccountID, req.Currency)
	if !valid {
		return
	}

	t, err := s.store.TransferTx(ctx, db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, t)
}

func (s *Server) validAccount(ctx *gin.Context, accountID int64, currency string) (db.Account, bool) {
	a, err := s.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return a, false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return a, false
	}
	if a.Currency != currency {
		ctx.JSON(http.StatusBadRequest, errorResponse(
			fmt.Errorf("account [%d] currency mismatch: %q vs %q",
				a.ID, a.Currency, currency)))
		return a, false
	}
	return a, true
}
