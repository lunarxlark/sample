package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
)

type (
	User struct {
		Name  string `json: "name"`
		Email string `json: "email"`
	}
)

type (
	lang struct {
		Name        string `db: "name"`
		Version     string `db: "version"`
		ReleaseDate string `db: "release_date"`
	}

	langJSON struct {
		Name        string `json: "name"`
		Version     string `json: "version"`
		ReleaseDate string `json: "release_date"`
	}

	resData struct {
		Langs []lang `json: "langs"`
	}
)

var (
	tablename = "language"
	seq       = 1
	conn, _   = dbr.Open("mysql", "tech_admin:password@tcp(tech_db:3306)/tech_link", nil)
	sess      = conn.NewSession(nil)
)

func main() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		message := "Hello, World!!!"
		return c.String(http.StatusOK, message)
	})

	e.GET("/langs", func(c echo.Context) error {
		var l []lang

		sess.Select("*").From(tablename).Load(&l)
		response := new(resData)
		response.Langs = l
		return c.JSON(http.StatusOK, response)
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func getGoVersion() (*goquery.Selection, error) {
	url := "https://github.com/golang/go/releases"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("a").Each(func(_ int, _s *goquery.Selection) {
		fmt.Println(_s.Find("a"))
	})
}
