package scraping

import "testing"

func TestTiendaInglesaUrl(t *testing.T) {

	url := "http://www.tinglesa.com.uy/producto.php?idarticulo=86"
	isDevoto := TiendaInglesaUrl(url)
	if !isDevoto {
		t.Errorf("TiendaInglesaUrl(%s) = %t, want %t", url, isDevoto, true)
	}

	url = "http://tinglesa.com.uy/producto.php?idarticulo=86"
	isDevoto = TiendaInglesaUrl(url)
	if !isDevoto {
		t.Errorf("TiendaInglesaUrl(%s) = %t, want %t", url, isDevoto, true)
	}

	url = "www.tinglesa.com.uy/producto.php?idarticulo=86"
	isDevoto = TiendaInglesaUrl(url)
	if isDevoto {
		t.Errorf("TiendaInglesaUrl(%s) = %t, want %t", url, isDevoto, false)
	}

	url = "http://www.devoto.com.uy/producto.php?idarticulo=86"
	isDevoto = TiendaInglesaUrl(url)
	if isDevoto {
		t.Errorf("TiendaInglesaUrl(%s) = %t, want %t", url, isDevoto, false)
	}

}

func TestTiendaInglesaFetchProduct(t *testing.T) {
	url := "http://www.tinglesa.com.uy/producto.php?idarticulo=86"
	productData, err := TiendaInglesa(url)
	if err != nil {
		t.Errorf("TiendaInglesa(%s) returned error %v", url, err)
	}

	exceptedName := "Arroz Saman Parboiled 1kg"
	if productData.Name != exceptedName {
		t.Errorf("TiendaInglesa(%s).Name = %s, want %s", url, productData.Name, exceptedName)
	}
	var expectedPrice float64
	expectedPrice = 37.00
	if productData.Price != expectedPrice {
		t.Errorf("TiendaInglesa(%s).Price = %v, want %v", url, productData.Price, expectedPrice)
	}
}

func TestTiendaInglesaFetchProductError(t *testing.T) {
	url := "http://tinglesa.com.uy/"
	_, err := TiendaInglesa(url)
	if err == nil {
		t.Errorf("TiendaInglesa(%s) didnt error", url)
	}
}
