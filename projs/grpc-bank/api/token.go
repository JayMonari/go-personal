package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type accessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type accessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func (s *Server) renewAccessToken(ctx *gin.Context) {
	var req accessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	refload, err := s.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	sess, err := s.store.GetSession(ctx, refload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	switch {
	case sess.IsBlocked:
		ctx.JSON(http.StatusUnauthorized,
			errorResponse(errors.New("blocked session")))
		return
	case sess.Username != refload.Username:
		ctx.JSON(http.StatusUnauthorized,
			errorResponse(errors.New("user session mismatch")))
		return
	case sess.RefreshToken != req.RefreshToken:
		ctx.JSON(http.StatusUnauthorized,
			errorResponse(errors.New("refresh token mismatch")))
		return
	case time.Now().After(sess.ExpiresAt):
		ctx.JSON(http.StatusUnauthorized,
			errorResponse(errors.New("expired session")))
		return
	}

	accessTok, accPayload, err :=
		s.tokenMaker.CreateToken(sess.Username, s.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accessTokenResponse{
		AccessToken:          accessTok,
		AccessTokenExpiresAt: accPayload.ExpiredAt,
	})
}
