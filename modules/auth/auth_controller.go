package auth

import (
	"moneda/evaluation/global"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SignIn(c echo.Context) error {
  var req AuthRequest 

  if err := c.Bind(&req); err != nil {
    return c.String(http.StatusInternalServerError, "Error parsing the request.")  
  }

  if req.Username != "admin" || req.Password == "" {
    return c.String(http.StatusBadRequest, "Invalid data provided.")  
  } 

  token, err := global.CreateToken(req.Username)

  if err != nil {
    return c.String(http.StatusInternalServerError, "Error creating the token.")
  }

  res := AuthResponse{
    Token: token,
  }
  
  return c.JSON(http.StatusOK, res)
}
