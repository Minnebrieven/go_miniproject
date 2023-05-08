package middlewares

import "github.com/labstack/echo/v4"

func IsSameUser(e echo.Context, userID float64) bool {
	isAdmin := ExtractTokenIsAdmin(e)
	if isAdmin {
		return true
	}

	isSameUser := ExtractTokenUserID(e) == userID
	return isSameUser
}
