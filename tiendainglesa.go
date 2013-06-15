package scraping

import (
    "errors"
    "github.com/moovweb/gokogiri"
    "io/ioutil"
    "log"
    "net/http"
    "regexp"
    "strconv"
    "strings"
    "time"
)

var re_tiendainglesa = regexp.MustCompile(`https?://(www.)?tinglesa.com.uy.*`)

func TiendaInglesaUrl(url string) bool {
    return re_tiendainglesa.MatchString(url)
}

func TiendaInglesa(url string) (data ProductData, err error) {
    resp, err := http.Get(url)
    if err != nil {
        log.Printf("Error getting site. ", err)
        return
    }
    if resp.StatusCode != http.StatusOK {
        err = errors.New("Request error. Invalid StatusCode")
        return
    }
    body, _ := ioutil.ReadAll(resp.Body)
    doc, err := gokogiri.ParseHtml(body)
    defer doc.Free()
    if err != nil {
        err = errors.New("Error parsing website. ")
        return
    }
    results, err := doc.Search("//h2")
    if err != nil {
        err = errors.New("Error finding h1 tag")
        return
    }
    if len(results) == 0 {
        err = errors.New("No product title found")
        return
    }
    name := results[0].Content()
    results, err = doc.Search("//div[@class='precio_producto']/span")
    if len(results) == 0 {
        err = errors.New("No price found")
        return
    }
    priceStr := results[0].Content()
    priceSplitList := strings.Split(priceStr, " ")
    price, err := strconv.ParseFloat(priceSplitList[len(priceSplitList)-1], 64)
    if err != nil {
        log.Printf("Error parsing price")
        return
    }
    data = ProductData{
        Name:        name,
        Url:         url,
        FetchedDate: time.Now(),
        Price:       price,
    }
    return
}
