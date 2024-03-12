package flights

type Flight struct {
  Hex          string `json:"hex"`
	RegNumber    string `json:"reg_number"`
	Icao         string `json:"aircraft_icao"`
	Flag         string `json:"flag"`
	Lat          string `json:"lat"`
	Lng          string `json:"lng"`
	Alt          string `json:"alt"`
	Dir          string `json:"dir"`
	Speed        string `json:"speed"`
	VSpeed       string `json:"v_speed"`
	Squawk       string `json:"squawk"`
	AirlineIcao  string `json:"airline_icao"`
	AirlineIata  string `json:"airline_iata"`
	FlightNumber string `json:"flight_number"`
	FlightIcao   string `json:"flight_icao"`
	FlightIata   string `json:"flight_iata"`
  DepIcao      string `json:"dep_icao,omitempty"`
	DepIata      string `json:"dep_iata,omitempty"`
	ArrIcao      string `json:"arr_icao,omitempty"`
	ArrIata      string `json:"arr_iata,omitempty"`
	Duration     string `json:"duration"`
	Updated      string `json:"updated"`
	Status       string `json:"status"`
}

type FlightSearchRequest struct {
  FlightIcao  string  `json:"flight_icao"`
  FlightIata  string  `json:"flight_iata"`
}

type FlightsGetRequest struct {
  Bbox         string `json:"bbox,omitempty"`
  Zoom         string `json:"zoom,omitempty"` 
	Hex          string `json:"hex,omitempty"`
  RegNumber    string `json:"reg_number,omitempty"`
	AirlineIcao  string `json:"airline_icao,omitempty"`
	AirlineIata  string `json:"airline_iata,omitempty"`
	Flag         string `json:"flag,omitempty"`
	FlightNumber string `json:"flight_number,omitempty"`
	DepIcao      string `json:"dep_icao,omitempty"`
	DepIata      string `json:"dep_iata,omitempty"`
	ArrIcao      string `json:"arr_icao,omitempty"`
	ArrIata      string `json:"arr_iata,omitempty"`
}
