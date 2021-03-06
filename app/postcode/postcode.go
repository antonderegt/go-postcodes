package postcode

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
)

type QueryAddress struct {
	Street   string `query:"street"`
	Number   string `query:"num"`
	City     string `query:"city"`
	Postcode string `query:"postcode"`
}

type LatLon struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

type AddressDetails struct {
	Number      string `json:"house_number"`
	Road        string `json:"road"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	Postcode    string `json:"postcode"`
	CountryCode string `json:"country_code"`
}

type Address struct {
	Address AddressDetails `json:"address"`
}

func GetQueryAddress(c *fiber.Ctx) *QueryAddress {
	address := new(QueryAddress)
	if err := c.QueryParser(address); err != nil {
		c.Status(500).Send(err)
	}
	return address
}

func ReturnAddress(c *fiber.Ctx) {
	address := GetQueryAddress(c)

	// Send response
	if err := c.JSON(fiber.Map{
		"street":   address.Street,
		"number":   address.Number,
		"city":     address.City,
		"postcode": address.Postcode,
	}); err != nil {
		c.Status(500).Send(err)
		return
	}
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		println("Error loading .env file")
	}

	return os.Getenv(key)
}

func GetLatLon(c *fiber.Ctx) {
	// Extract address from query
	address := GetQueryAddress(c)

	// GET request to nominatim
	query := "street=" + address.Number + "%20" + address.Street + "&city=" + address.City + "&postalcode=" + address.Postcode
	api := goDotEnvVariable("API_ADDRESS")
	queryAddress := api + "/search?" + query + "&format=json&limit=1"

	var res []LatLon
	err := getJson(queryAddress, &res)
	if err != nil {
		c.Status(500).Send(err)
	}
	if len(res) == 0 {
		c.Status(500).Send("No results")
		return
	}

	// Send response
	if err := c.JSON(fiber.Map{
		"lat": res[0].Lat,
		"lon": res[0].Lon,
	}); err != nil {
		c.Status(500).Send(err)
		return
	}
}

func GetFullAddress(c *fiber.Ctx) {
	// Extract address from query
	address := GetQueryAddress(c)

	// GET request to nominatim
	query := "street=" + address.Number + "%20" + address.Street + "&city=" + address.City + "&postalcode=" + address.Postcode
	api := goDotEnvVariable("API_ADDRESS")
	queryAddress := api + "/search?" + query + "&format=json&limit=1&addressdetails=1"

	var res []Address
	err := getJson(queryAddress, &res)
	if err != nil {
		c.Status(500).Send(err)
		return
	}
	if len(res) == 0 {
		c.Status(500).Send("No results")
		return
	}

	// Send response
	if err := c.JSON(fiber.Map{
		"postcode": res[0].Address.Postcode,
		"street":   res[0].Address.Road,
		"number":   res[0].Address.Number,
		"city":     res[0].Address.City,
		"state":    res[0].Address.State,
	}); err != nil {
		c.Status(500).Send(err)
		return
	}
}
