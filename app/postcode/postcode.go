package postcode

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gofiber/fiber"
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

func GetLatLon(c *fiber.Ctx) {
	// Extract address from query
	address := GetQueryAddress(c)

	// GET request to nominatim
	query := address.Street + " " + address.Number + ", " + address.City + ", " + address.Postcode
	queryAddress := "http://77.248.22.231:7070/search/" + query + "?format=json&countrycodes=NL&limit=1"

	var res []LatLon
	err := getJson(queryAddress, &res)
	if err != nil {
		c.Status(500).Send(err)
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

func GetPostcode(c *fiber.Ctx) {
	// Extract address from query
	address := GetQueryAddress(c)

	// GET request to nominatim
	query := address.Street + " " + address.Number + ", " + address.City + ", postcode=" + address.Postcode
	queryAddress := "http://77.248.22.231:7070/search/" + query + "?format=json&addressdetails=1&countrycodes=NL&limit=1"

	var res []Address
	err := getJson(queryAddress, &res)
	if err != nil {
		c.Status(500).Send(err)
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
