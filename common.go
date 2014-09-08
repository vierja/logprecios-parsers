package scraping

import (
	"errors"
	"fmt"
)

type ProductData struct {
	Name        string   `json:"name"`
	Url         string   `json:"url"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`
	Price       float64  `json:"price"`
	Fetched     string   `json:"fetched"`
	ImageUrl    string   `json:"image_url"`
}

const INVALID_PRODUCT_URL int = 300
const SITE_ERROR int = 500
const CLIENT_ERROR int = 400

type ScrapeError struct {
	Arg  int
	prob string
}

func (e *ScrapeError) Error() string {
	return fmt.Sprintf("%d - %s", e.Arg, e.prob)
}

func GetProductData(url string) (data ProductData, err error) {
	if DevotoUrl(url) {
		data, err = Devoto(url)
	} else if TiendaInglesaUrl(url) {
		data, err = TiendaInglesa(url)
	} else {
		err = errors.New("Invalid url. Can't find parser.")
	}
	return
}
