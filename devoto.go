package scraping

import (
	"errors"
	"github.com/moovweb/gokogiri"
	"io/ioutil"
	"log"
	"net/http"
	net_url "net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var re_devoto = regexp.MustCompile(`https?://(www.)?devoto.com.uy.*`)

func DevotoUrl(url string) bool {
	return re_devoto.MatchString(url)
}

func Devoto(url string) (data ProductData, err error) {

	urlObj, err := net_url.Parse(url)
	if err != nil {
		err = &ScrapeError{INVALID_PRODUCT_URL, "Invalid url. Could not parse."}
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		err = &ScrapeError{SITE_ERROR, "Error connecting to site."}
		return
	}
	if resp.StatusCode != http.StatusOK {
		errorCode := int(resp.StatusCode/100) * 100
		err = &ScrapeError{errorCode, "Request error. Invalid StatusCode."}
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	doc, err := gokogiri.ParseHtml(body)
	defer doc.Free()
	if err != nil {
		err = errors.New("Error parsing website. ")
		return
	}
	results, err := doc.Search("//h1")
	if err != nil {
		err = errors.New("Error finding h1 tag")
		return
	}
	if len(results) == 0 {
		err = errors.New("No product title found")
		return
	}
	name := results[0].Content()
	results, err = doc.Search("//div[@id='ProductPrice']")
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

	//Image Url
	results, err = doc.Search("//img[@id='producto_img']/@src")
	var imageUrl string
	if err == nil && len(results) == 1 {
		imageUrl = results[0].Content()
		imageUri, err := net_url.Parse(imageUrl)
		if err == nil {
			imageUrl = urlObj.ResolveReference(imageUri).String()
		}
	}

	data = ProductData{
		Name:     name,
		Url:      url,
		Price:    price,
		Fetched:  time.Now().UTC().Format("2006-01-02T15:04Z"),
		ImageUrl: imageUrl,
	}
	return
}
