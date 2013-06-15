package scraping

import (
    "errors"
    "time"
)

type ProductData struct {
    Name        string
    Url         string
    Price       float64
    FetchedDate time.Time
}

func GetProductData(url string) (data ProductData, err error) {
    if DevotoUrl(url) {
        data, err = Devoto(url)
    } else if TiendaInglesaUrl(url) {
        data, err = TiendaInglesa(url)
    } else if MultiAhorroUrl(url) {
        data, err = MultiAhorro(url)
    } else {
        err = errors.New("Invalid url. Can't find parser.")
    }
    return

    if condition {

    }
}
