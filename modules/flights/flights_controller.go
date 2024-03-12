package flights

import (
	"net/http"
	"slices"

	"github.com/labstack/echo/v4"
)

func GetFlightData(c echo.Context, dbData []byte) error {
  var req FlightsGetRequest

  if err := c.Bind(&req); err != nil {
    return c.String(http.StatusInternalServerError, "Error parsing the request data.")
  }

  flights, err := loadFlights(dbData)

  if err != nil {
    return c.String(http.StatusInternalServerError, "Error getting flights data from database.")
  }

  foundFlights := []Flight{}

  for _, flight := range flights {
    if flight.Hex       == req.Hex &&
    flight.RegNumber    == req.RegNumber &&
    flight.AirlineIata  == req.AirlineIata && 
    flight.AirlineIcao  == req.AirlineIcao && 
    flight.Flag         == req.Flag &&
    flight.FlightNumber == req.FlightNumber {
      foundFlights = append(foundFlights, flight)
    }
  } 

  return c.JSON(http.StatusOK, foundFlights) 
}

func SearchFlightInfo(c echo.Context, dbData []byte) error {
  var req FlightSearchRequest 

  if err := c.Bind(&req); err != nil {
    return c.String(http.StatusInternalServerError, "Error parsing the request data.")
  }

  flights, err := loadFlights(dbData)

  if err != nil {
    return c.String(http.StatusInternalServerError, "Error getting flights data from database.")
  }
  
  foundFlightIndex := slices.IndexFunc(flights, func(flight Flight) bool {
    return flight.FlightIcao == req.FlightIcao && flight.FlightIata == req.FlightIata   
  })

  if foundFlightIndex == -1 {
    return c.JSON(http.StatusOK, []Flight{})
  }

  return c.JSON(http.StatusOK, flights[foundFlightIndex]) 
}
