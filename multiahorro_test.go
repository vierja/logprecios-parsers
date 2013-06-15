package scraping

import "testing"

func TestMultiAhorroUrl(t *testing.T) {

    url := "http://www.multiahorro.com.uy/Product.aspx?p=94285"
    isDevoto := MultiAhorroUrl(url)
    if !isDevoto {
        t.Errorf("MultiAhorroUrl(%s) = %t, want %t", url, isDevoto, true)
    }

    url = "http://multiahorro.com.uy/Product.aspx?p=94285"
    isDevoto = MultiAhorroUrl(url)
    if !isDevoto {
        t.Errorf("MultiAhorroUrl(%s) = %t, want %t", url, isDevoto, true)
    }

    url = "www.multiahorro.com.uy/Product.aspx?p=94285"
    isDevoto = MultiAhorroUrl(url)
    if isDevoto {
        t.Errorf("MultiAhorroUrl(%s) = %t, want %t", url, isDevoto, false)
    }

    url = "http://www.devoto.com.uy/producto.php?idarticulo=86"
    isDevoto = MultiAhorroUrl(url)
    if isDevoto {
        t.Errorf("MultiAhorroUrl(%s) = %t, want %t", url, isDevoto, false)
    }

}

func TestMultiAhorroFetchProduct(t *testing.T) {
    url := "http://www.multiahorro.com.uy/Product.aspx?p=94285"
    productData, err := MultiAhorro(url)
    if err != nil {
        t.Errorf("MultiAhorro(%s) returned error %v", url, err)
    }

    exceptedName := "LECHE CONAP. FRESCA  ENTERA 1L"
    if productData.Name != exceptedName {
        t.Errorf("MultiAhorro(%s).Name = '%s', want '%s'", url, productData.Name, exceptedName)
    }
    var expectedPrice float64
    expectedPrice = 16
    if productData.Price != expectedPrice {
        t.Errorf("MultiAhorro(%s).Price = %v, want %v", url, productData.Price, expectedPrice)
    }
}

func TestMultiAhorroFetchProductError(t *testing.T) {
    url := "http://tinglesa.com.uy/"
    _, err := MultiAhorro(url)
    if err == nil {
        t.Errorf("TiendaInglesa(%s) didnt error", url)
    }
}
