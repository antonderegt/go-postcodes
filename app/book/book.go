package book

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber"
)

type Address struct {
	Street string `query:"street"`
	Number string `query:"num"`
	City   string `query:"city"`
}

func ReturnAddress(c *fiber.Ctx) {
	var address = new(Address)
	if err := c.QueryParser(address); err != nil {
		c.Send(err)
	}
	c.Send("Address: ", address.Street, " ", address.Number, ", ", address.City)
}

func ConsumeAPI(c *fiber.Ctx) {
	response, err := http.Get("http://77.248.22.231:7070/search/amsterdam?format=json&countrycodes=NL&limit=1")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(responseData))
	c.Send(string(responseData))
}
