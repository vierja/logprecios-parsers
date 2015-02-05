package scraping

import "testing"

func TestDevotoUrl(t *testing.T) {

	url := "http://devoto.com.uy/aproduct.aspx?528"
	isDevoto := DevotoUrl(url)
	if !isDevoto {
		t.Errorf("DevotoUrl(%s) = %t, want %t", url, isDevoto, true)
	}

	url = "http://devoto.com.uy/aproduct.aspx?528"
	isDevoto = DevotoUrl(url)
	if !isDevoto {
		t.Errorf("DevotoUrl(%s) = %t, want %t", url, isDevoto, true)
	}

	url = "www.devoto.com.uy/aproduct.aspx?528"
	isDevoto = DevotoUrl(url)
	if isDevoto {
		t.Errorf("DevotoUrl(%s) = %t, want %t", url, isDevoto, false)
	}

	url = "http://www.tinglesa.com.uy/producto.php?idarticulo=86"
	isDevoto = DevotoUrl(url)
	if isDevoto {
		t.Errorf("DevotoUrl(%s) = %t, want %t", url, isDevoto, false)
	}

}

func TestDevotoFetchProduct(t *testing.T) {
	url := "http://devoto.com.uy/aproduct.aspx?528"
	productData, err := Devoto(url)
	if err != nil {
		t.Errorf("Devoto(%s) returned error %v", url, err)
	}

	exceptedName := "LECHE FRESCA ENTERA CONAPROLE SACHET 1LT."
	if productData.Name != exceptedName {
		t.Errorf("Devoto(%s).Name = %s, want %s", url, productData.Name, exceptedName)
	}
	var expectedPrice float64
	expectedPrice = 19
	if productData.Price != expectedPrice {
		t.Errorf("Devoto(%s).Price = %v, want %v", url, productData.Price, expectedPrice)
	}
}

func TestDevotoFetchProductError(t *testing.T) {
	url := "http://devoto.com.uy/"
	_, err := Devoto(url)
	if err == nil {
		t.Errorf("Devoto(%s) didnt error", url)
	}
}
