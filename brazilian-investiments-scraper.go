package brazilianinvestimentsscraper

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/sousair/go-brazilian-investments-scraper/scraper"
)

var (
	statusInvestBaseURL = "https://statusinvest.com.br/acoes/%s"
)

func GetStockData(symbol string) *scraper.StockData {
	stockData := &scraper.StockData{}

	stockData.Symbol = symbol

	c := colly.NewCollector()

	fieldsToScrap := scraper.GetStockScrapFields()

	for _, field := range fieldsToScrap {
		c.OnHTML(field.Selector, func(e *colly.HTMLElement) {
			err := field.Fn(e.Text, stockData)
			if err != nil {
				fmt.Println(err)
			}
		})
	}

	url := fmt.Sprintf(statusInvestBaseURL, symbol)
	c.Visit(url)
	c.Wait()

	return stockData
}
