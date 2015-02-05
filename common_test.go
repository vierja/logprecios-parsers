package scraping

import "testing"

func TestCommonParser(t *testing.T) {
	url := "http://devoto.com.uy/aproduct.aspx?528"
	data, err := GetProductData(url)
	if err != nil {
		t.Errorf("GetProductData(%s) resulted in error.", url)
	}

	exceptedName := "LECHE FRESCA ENTERA CONAPROLE SACHET 1LT."
	if data.Name != exceptedName {
		t.Errorf("GetProductData(%s).Name = %s, want %s", url, data.Name, exceptedName)
	}
}
