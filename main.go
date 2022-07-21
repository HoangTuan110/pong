package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gocolly/colly"
	"github.com/labstack/echo/v4"
)

var spf = fmt.Sprintf

func getHostName(c echo.Context) string {
	return c.Request().Host
}

// GET /:url
func checkURL(c echo.Context) error {
	url := "https://" + c.Param("url")
	response := ""
	col := colly.NewCollector()
	col.SetRequestTimeout(30 * time.Second)

	col.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	col.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
		response = spf("No + Something went wrong (%s)", err)
	})

	col.OnResponse(func(r *colly.Response) {
		if r.StatusCode != 200 {
			response = "No"
			return
		}
		response = "Yes"
	})

	col.Visit(url)

	return c.String(http.StatusOK, response)
}

// GET /
func index(c echo.Context) error {
	return c.String(http.StatusOK, spf(`
PONG - SELF-HOSTED 'IS IT UP?' SERVICE
================

TL;DR:
	$ curl %s/<URL>

EXAMPLE:
	$ curl %s/google.com
	Yes
`, getHostName(c), getHostName(c)))
}

func main() {
	e := echo.New()
	e.GET("/", index)
	e.GET("/:url", checkURL)
	e.Logger.Fatal(e.Start(":1323"))
}
