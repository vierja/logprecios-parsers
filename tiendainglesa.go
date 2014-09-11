package scraping

import (
	"errors"
	//"github.com/moovweb/gokogiri"
	"encoding/json"
	"fmt"
	"github.com/moovweb/gokogiri/html"
	"io/ioutil"
	"log"
	"net/http"
	net_url "net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type urlWrapper struct {
	Url string `json:"url"`
}

type imageJson struct {
	Urls []urlWrapper `json:"urls"`
}

var re_tiendainglesa = regexp.MustCompile(`https?://(www.)?tinglesa.com.uy.*`)

func TiendaInglesaUrl(url string) bool {
	return re_tiendainglesa.MatchString(url)
}

func getImage(productId int, result chan string) {
	url := fmt.Sprintf("http://www.tinglesa.com.uy/verCrearFotoPrueba.php?idArticulos=%d&pos=1&w=1200&h=900", productId)
	resp, err := http.Get(url)
	if err != nil {
		result <- ""
		return
	}

	var jsonData imageJson
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		result <- ""
		return
	}

	//result <- ""
	result <- jsonData.Urls[0].Url
}

func TiendaInglesa(url string) (data ProductData, err error) {
	// Find productId
	urlObj, err := net_url.Parse(url)
	if err != nil {
		err = &ScrapeError{INVALID_PRODUCT_URL, "Invalid url. Could not parse."}
		return
	}

	productIdList, present := urlObj.Query()["idarticulo"]

	if !present || len(productIdList) != 1 {
		err = &ScrapeError{INVALID_PRODUCT_URL, "Invalid url. Could not find idarticulo param."}
		return
	}

	productId, err := strconv.Atoi(productIdList[0])

	if err != nil {
		err = &ScrapeError{INVALID_PRODUCT_URL, "Invalid url. idarticulo param is not integer."}
		return
	}

	imgChan := make(chan string, 1)
	go getImage(productId, imgChan)

	resp, err := http.Get(url)

	if err != nil {
		log.Printf("Error getting site. ", err)
		return
	}

	if finalUrl := resp.Request.URL.String(); strings.Contains(finalUrl, "articulo_no_habilitado") {
		err = &ScrapeError{INVALID_PRODUCT_URL, "Product not found"}
		return
	}

	if resp.StatusCode != http.StatusOK {
		errorCode := int(resp.StatusCode/100) * 100
		err = &ScrapeError{errorCode, "Request error. Invalid StatusCode"}
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	doc, err := html.Parse(body, []byte("iso-8859-1"), nil, html.DefaultParseOption, html.DefaultEncodingBytes)
	defer doc.Free()

	if err != nil {
		err = errors.New("Error parsing website. ")
		return
	}
	// Find Title
	results, err := doc.Search("//h1[@class='titulo_producto_top']")
	if err != nil {
		err = errors.New("Error finding h1 tag")
		return
	}
	if len(results) == 0 {
		err = errors.New("No product title found")
		return
	}
	name := strings.TrimSpace(results[0].Content())

	// Find Price
	results, err = doc.Search("//div[@class='contendor_precio']//td[@class='precio']")
	if len(results) == 0 {
		err = errors.New("No price found")
		return
	}
	priceStr := results[0].Content()
	priceSplitList := strings.Fields(priceStr)
	price, err := strconv.ParseFloat(priceSplitList[len(priceSplitList)-1], 64)
	if err != nil {
		log.Printf("Error parsing price")
		return
	}

	// Find description
	results, err = doc.Search("//div[@class='contenido_descripcion']")
	var description string
	if err == nil && len(results) > 0 {
		description = strings.TrimSpace(results[0].Content())
	}

	// Find categories
	results, err = doc.Search("//div[@class='navegacion']/a")
	var categories []string
	if err == nil && len(results) > 1 {
		// Remove "Home" category.
		results = results[1:]
		categories = make([]string, len(results))
		for i := range results {
			categories[i] = strings.TrimSpace(results[i].Content())
		}
	}

	//Image Url
	imageUrl := <-imgChan

	data = ProductData{
		Name:        name,
		Url:         url,
		Price:       price,
		Description: description,
		Categories:  categories,
		Fetched:     time.Now().UTC().Format("2006-01-02T15:04Z"),
		ImageUrl:    imageUrl,
	}
	return
}
