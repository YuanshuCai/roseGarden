package rosegarden

import (
	"testing"
)

// If wanna raun single test use the shell command
// $ go test ./... -run functionName
// $ go test ./... -run TestAddNormalItem
func TestAddNormalItem(t *testing.T) {
	testName(t, "foo", 1, 12, "foo")
}

func TestAddSpecialItem(t *testing.T) {
	testName(t, "Backstage passes to a TAFKAL80ETC concert", 1, 12, "Backstage passes to a TAFKAL80ETC concert")
	testName(t, "Sulfuras, Hand of Ragnaros", 1, 12, "Sulfuras, Hand of Ragnaros")
	testName(t, "Aged Brie", 1, 12, "Aged Brie")
}
func TestNormalItemQualityNormalDecreases(t *testing.T) {
	testQuality(t, "foo", 3, 8, 7)
}
func TestNormalItemQuality2XDecreases(t *testing.T) {
	testQuality(t, "foo", 0, 8, 6)
}

func TestNormalItemQualityQualityLowBound(t *testing.T) {
	testQuality(t, "foo", 0, 0, 0)
}

func TestAgedBrieSellInDecreases(t *testing.T) {
	testSellIn(t, "Aged Brie", 3, 9, 2)
}
func TestAgedBrieQualityIncreases(t *testing.T) {
	testQuality(t, "Aged Brie", 3, 9, 10)
}

func TestSulfurasQualityUnchange(t *testing.T) {
	testQuality(t, "Sulfuras, Hand of Ragnaros", 2, 49, 49)
}
func TestSulfurasQualityUnchangEvenOver50(t *testing.T) {
	testQuality(t, "Sulfuras, Hand of Ragnaros", 2, 51, 51)
}

func TestSulfurasNeverSell(t *testing.T) {
	testSellIn(t, "Sulfuras, Hand of Ragnaros", 0, 51, 0)
}

func TestBackstageQualityIncreasesAndSellInDecreases(t *testing.T) {
	testQuality(t, "Backstage passes to a TAFKAL80ETC concert", 12, 20, 21)
	testSellIn(t, "Backstage passes to a TAFKAL80ETC concert", 12, 20, 11)
}
func TestBackstageQualityIncreases2XAnd3x(t *testing.T) {
	testQuality(t, "Backstage passes to a TAFKAL80ETC concert", 11, 20, 22)
	testQuality(t, "Backstage passes to a TAFKAL80ETC concert", 5, 20, 23)
}

func TestBackstageQualityInQualityHighBound(t *testing.T) {
	testQuality(t, "Backstage passes to a TAFKAL80ETC concert", 3, 49, 50)
}

func TestBackstageQualityDropAfterExpire(t *testing.T) {
	testQuality(t, "Backstage passes to a TAFKAL80ETC concert", 0, 23, 0)
}

// General testing for 3 features
func testName(t *testing.T, name string, startSellIn int, startQuality int, expectedName string) {

	items := []*Item{
		{
			Name:    name,
			SellIn:  startSellIn,
			Quality: startQuality},
	}
	UpdateQuality(items)
	if items[0].Name != expectedName {
		t.Errorf("Name: Expected %s but got %s ", expectedName, items[0].Name)
	}
}
func testSellIn(t *testing.T, name string, startSellIn int, startQuality int, expectedSellIn int) {
	items := []*Item{
		{
			Name:    name,
			SellIn:  startSellIn,
			Quality: startQuality},
	}
	UpdateQuality(items)
	if items[0].SellIn != expectedSellIn {
		t.Errorf("sellIn: Expected %d but got %d ", expectedSellIn, items[0].SellIn)
	}
}
func testQuality(t *testing.T, name string, startSellIn int, startQuality int, expectedQuality int) {
	items := []*Item{
		{
			Name:    name,
			SellIn:  startSellIn,
			Quality: startQuality},
	}
	UpdateQuality(items)
	if items[0].Quality != expectedQuality {
		t.Errorf("sellIn: Expected %d but got %d ", expectedQuality, items[0].Quality)
	}
}
