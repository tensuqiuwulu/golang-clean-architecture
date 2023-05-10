package middleware

import (
	"strings"

	"github.com/tensuqiuwulu/golang-clean-architecture/exception"
	"github.com/tensuqiuwulu/golang-clean-architecture/src/model"
	"github.com/tensuqiuwulu/golang-clean-architecture/src/repository"
	"github.com/tensuqiuwulu/golang-clean-architecture/utils"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Middleware interface {
	CheckInternalToken(next echo.HandlerFunc) echo.HandlerFunc
}

type middlewareImpl struct {
	DB              *gorm.DB
	OauthRepository repository.OauthRepository
}

func NewMiddleware(
	db *gorm.DB,
	oauthRepository repository.OauthRepository,
) Middleware {
	return &middlewareImpl{
		DB:              db,
		OauthRepository: oauthRepository,
	}
}

func (m *middlewareImpl) CheckInternalToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			exception.ErrorHandler(utils.LogStruct{Code: 401, Mssg: "Unauthorized", LogDetail: "Authorization header is empty"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			exception.ErrorHandler(utils.LogStruct{Code: 401, Mssg: "Unauthorized", LogDetail: "Invalid token format"})
		}

		tokenString := parts[1]

		tokenParse, _ := jwt.ParseWithClaims(tokenString, &model.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			return nil, nil
		})

		claims := tokenParse.Claims.(*model.TokenClaims)

		if claims.Id == "" {
			exception.ErrorHandler(utils.LogStruct{Code: 401, Mssg: "Unauthorized", LogDetail: "Invalid token"})
		}

		check, err := m.OauthRepository.CheckOauthAccessTokenByApiKey(m.DB, claims.Id)
		if err != nil {
			exception.ErrorHandler(utils.LogStruct{Code: 401, Mssg: "Unauthorized", LogDetail: "Invalid token"})
		}

		if !check {
			exception.ErrorHandler(utils.LogStruct{Code: 401, Mssg: "Unauthorized", LogDetail: "Invalid token"})
		}

		return next(c)
	}
}
