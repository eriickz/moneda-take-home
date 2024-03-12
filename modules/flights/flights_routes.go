package flights

import (
	"moneda/evaluation/global"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RegisterFlightsRoutes(apiGroup *echo.Group, dbData []byte) {
  flightsGroup := apiGroup.Group("/flights")

  flightsGroup.Use(echojwt.WithConfig(global.GetJwtConfig()))

  flightsGroup.POST("/get-flight-data", func(c echo.Context) error {
    return GetFlightData(c, dbData)
  })
  flightsGroup.POST("/search-flight-info", func(c echo.Context) error {
    return SearchFlightInfo(c, dbData)
  })
}
