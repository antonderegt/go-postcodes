package postcode

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gofiber/fiber"
)

type QueryAddress struct {
	Street string `query:"street"`
	Number string `query:"num"`
	City   string `query:"city"`
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

func ReturnAddress(c *fiber.Ctx) {
	var address = new(QueryAddress)
	if err := c.QueryParser(address); err != nil {
		c.Send(err)
	}
	c.Send("Address: ", address.Street, " ", address.Number, ", ", address.City)
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
	var address = new(QueryAddress)
	if err := c.QueryParser(address); err != nil {
		c.Send(err)
	}

	// GET request to nominatim
	var query = address.Street + " " + address.Number + ", " + address.City
	var queryAddress = "http://77.248.22.231:7070/search/" + query + "?format=json&countrycodes=NL&limit=1"
	var res []LatLon
	getJson(queryAddress, &res)

	// Send response
	var response = `{"lat": "` + res[0].Lat + `", "lon": "` + res[0].Lon + `"}`
	c.Send(response)
}

func GetPostcode(c *fiber.Ctx) {
	// Extract address from query
	var address = new(QueryAddress)
	if err := c.QueryParser(address); err != nil {
		c.Send(err)
	}

	// GET request to nominatim
	var query = address.Street + " " + address.Number + ", " + address.City
	var queryAddress = "http://77.248.22.231:7070/search/" + query + "?format=json&addressdetails=1&countrycodes=NL&limit=1"
	var res []Address
	getJson(queryAddress, &res)

	// Send response
	var response = `{"postcode": "` + res[0].Address.Postcode + `"}`
	c.Send(response)
}
