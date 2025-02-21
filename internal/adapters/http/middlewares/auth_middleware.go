package middlewares

import (
	"cashflow-go/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			cookie, err := c.Cookie("token")
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing token"})
			}
			tokenString := cookie.Value

			token, err := utils.CheckJWT(tokenString)

			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid claims"})
			}

			c.Set("user_id", uint(claims["user_id"].(float64)))

			return next(c)
		}
	}
}
