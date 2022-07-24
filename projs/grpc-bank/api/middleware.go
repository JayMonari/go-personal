package api

import (
	"errors"
	"net/http"
	"strings"

	"example.xyz/bank/token"
	"github.com/gin-gonic/gin"
)

const (
	authorization  = "Authorization"
	typeBearer     = "bearer"
	authPayloadKey = "authorization_payload"
)

func authMiddleware(maker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(authorization)
		if len(authHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(
				errors.New("Authorization header is not provided")))
			return
		}
		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(
				errors.New("invalid Authorization header format")))
			return
		}
		if strings.ToLower(fields[0]) != typeBearer {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(
				errors.New("unsupported authorization type")))
			return
		}
		p, err := maker.VerifyToken(fields[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		ctx.Set(authPayloadKey, p)
		ctx.Next()
	}
}
