package brazilianinvestimentsscrapper

import (
	"fmt"

	"github.com/gocolly/colly"
)

type StockData struct {
	// Pregão
	Symbol string
	Name   string
	// Cotação
	Price float64
	// Patrimônio Liquido
	PL float64
	// Ativos
	TotalAssets float64
	// Ativo Circulante
	TotalCurrentAssets float64
	// Divida Bruta
	GrossDebt float64
	// Disponibilidades
	Disponibilities float64
	// Divida Liquida
	NetDebt float64
	// Valor de Mercado
	MarketCap float64
	// Valor de Firma
	EV float64
	// Total de papéis
	TotalShares int
	// Free Float
	FreeFloat float64

	ValuationIndicators ValuationIndicators
	DebtIndicators      DebtIndicators
	EficiencyIndicators EficiencyIndicators
	ProfitIndicators    ProfitIndicators
	GrowthIndicators    GrowthIndicators
}

type ValuationIndicators struct {
	// DY
	DividendYield float64
	// P/L
	PLRatio float64
	// PEG RATIO
	PEGRatio float64
	// P/VP
	PVpRatio float64
	// EV/EBITDA
	EVEbitdaRatio float64
	// EV/EBIT
	EVEbitRatio float64
	// P/EBITDA
	PEBitdaRatio float64
	// P/EBIT
	PEBitRatio float64
	// VPA
	VPA float64
	// P/ATIVO
	PAtivoRatio float64
	// LPA
	LPA float64
	// P/SR
	PSRRatio float64
	// P/CAP. GIRO
	PCapGiroRatio float64
	// P/ATIVO CIRC. LIQ.
	PAtivoCircLiqRatio float64
}

type DebtIndicators struct {
	// Div. Liq.
	DivLiq float64
	// Div. Bruta
	DivBruta float64
	// Div. Liq. / PL
	DivLiqPLRatio float64
	// Div. Liq. / EBITDA
	DivLiqEbitdaRatio float64
	// Div. Liq. / EBIT
	DivLiqEbitRatio float64
	// PL / ATIVOS
	PLAtivosRatio float64
	// Passivos / Ativos
	PassivosAtivosRatio float64
	// Liq. Corrente
	LiqCorRatio float64
}

type EficiencyIndicators struct {
	// Margem Bruta
	MBruta float64
	// Margem EBITDA
	MEbitda float64
	// Margem EBIT
	MEbit float64
	// Margem Liquida
	MLiq float64
}

type ProfitIndicators struct {
	// ROE
	ROE float64
	// ROA
	ROA float64
	// ROIC
	ROIC float64
	// Giro Ativos
	GiroAtivos float64
}

type GrowthIndicators struct {
	// Crescimento Receita Líquida 5 anos
	GARGer5Anos float64
	// Crescimento EBITDA 5 anos
	GARLuc5Anos float64
}

var (
	statusInvestBaseURL = "https://statusinvest.com.br/acoes/%s"
)

func GetStockData(symbol string) *StockData {
	stockData := &StockData{}

	stockData.Symbol = symbol

	c := colly.NewCollector()

	fieldsToScrap := scrapper.GetStockScrapFields()

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
