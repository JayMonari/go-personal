package api

import (
	"database/sql"
	"errors"
	"net/http"

	"example.xyz/bank/internal/db"
	"example.xyz/bank/token"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createAccountRequest struct {
	Currency string `json:"currency" binding:"required,currency"`
}

func (s *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authPayloadKey).(*token.Payload)
	a, err := s.store.CreateAccount(ctx, db.CreateAccountParams{
		Owner:    authPayload.Username,
		Balance:  0,
		Currency: req.Currency,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, a)
}

type listAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (s *Server) listAccount(ctx *gin.Context) {
	var req listAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authPayloadKey).(*token.Payload)
	accts, err := s.store.ListAccounts(ctx, db.ListAccountsParams{
		Owner:  authPayload.Username,
		Offset: (req.PageID - 1) * req.PageSize,
		Limit:  req.PageSize,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, accts)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	a, err := s.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authPayloadKey).(*token.Payload)
	if a.Owner != authPayload.Username {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New(
			"account doesn't belong to the authenticated user")))
		return
	}
	ctx.JSON(http.StatusOK, a)
}
