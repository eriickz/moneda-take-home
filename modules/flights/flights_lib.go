package flights

import "encoding/json"

func loadFlights(dbData []byte) ([]Flight, error) {
  var flights []Flight 

  if err := json.Unmarshal(dbData, &flights); err != nil {
    return nil, err
  }

  return flights, nil
}
