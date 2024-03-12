//Global functions to configure key functionalities in the project or
//shared in some modules through the code. 
package global

import (
	"io"
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
)

func GetJwtConfig() echojwt.Config {
  jwtSecret := os.Getenv("JWT_SECRET")

  return echojwt.Config{
    SigningKey: []byte(jwtSecret),
    ContextKey: "username",
  }
}

func CreateToken(username string) (string, error) {
  jwtIssuer := os.Getenv("JWT_ISSUER")
  jwtSecret := os.Getenv("JWT_SECRET")

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "iss": jwtIssuer,
    "sub": username,
  })

  tokenString, err := token.SignedString([]byte(jwtSecret))

  if err != nil {
    return "", err
  }

  return tokenString, nil
}

func LoadDB() []byte {
  jsonFile, err := os.Open("flight_data.json") 

  if err != nil {
    panic(err)
  }

  defer jsonFile.Close()

  data, err := io.ReadAll(jsonFile)

  if err != nil {
    panic(err)
  }

  return data
}
