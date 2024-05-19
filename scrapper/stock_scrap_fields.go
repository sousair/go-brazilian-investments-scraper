package scrapper

import brazilianinvestimentsscrapper "github.com/sousair/go-brazilian-investments-scraper"

type scrapField struct {
	Selector  string
	FieldType string
	Fn        func(field string, stockData *brazilianinvestimentsscrapper.StockData) error
}

var (
	stockDataToScrap = []scrapField{
		// Stock Data
		{
			nameSelector,
			"string",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				stockData.Name = field
				return nil
			},
		},
		{
			priceSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.Price = value
				return nil
			},
		},
		{
			equitySelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.Equity = value
				return nil
			},
		},
		{
			totalAssetsSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.TotalAssets = value
				return nil
			},
		},
		{
			totalCurrentAssetsSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.TotalCurrentAssets = value
				return nil
			},
		},
		{
			grossDebtSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.GrossDebt = value
				return nil
			},
		},
		{
			disponibilitiesSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.Disponibilities = value
				return nil
			},
		},
		{
			netDebtSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.NetDebt = value
				return nil
			},
		},
		{
			marketCapSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.MarketCap = value
				return nil
			},
		},
		{
			evSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.EV = value
				return nil
			},
		},
		{
			totalSharesSelector,
			"int",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseIntField(field)
				if err != nil {
					return err
				}
				stockData.TotalShares = value
				return nil
			},
		},
		{
			freeFloatSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.FreeFloat = value
				return nil
			},
		},
	}

	valuationIndicatorsToScrap = []scrapField{
		// Valuation Indicators
		{
			dividendYieldSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.ValuationIndicators.DividendYield = value
				return nil
			},
		},
		{
			plRatioSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.ValuationIndicators.PLRatio = value
				return nil
			},
		},
		{
			pegRatioSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.ValuationIndicators.PEGRatio = value
				return nil
			},
		},
		{
			pvpRatioSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				stockData.ValuationIndicators.PVpRatio = value
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			evEbitdaRatioSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				stockData.ValuationIndicators.EVEbitdaRatio = value
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			evEbitRatioSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				stockData.ValuationIndicators.EVEbitRatio = value
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			pEbitdaRatioSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				stockData.ValuationIndicators.PEBitdaRatio = value
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			pEbitRatioSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				stockData.ValuationIndicators.PEBitRatio = value
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			vpaSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				stockData.ValuationIndicators.VPA = value
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			pAtivoRatioSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				stockData.ValuationIndicators.PAtivoRatio = value
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			lpaSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				stockData.ValuationIndicators.LPA = value
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			pSRRatioSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				stockData.ValuationIndicators.PSRRatio = value
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			pCapGiroRatioSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				stockData.ValuationIndicators.PCapGiroRatio = value
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			pAtivoCircLiqRatioSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				stockData.ValuationIndicators.PAtivoCircLiqRatio = value
				if err != nil {
					return err
				}
				return nil
			},
		},
	}

	debtIndicatorsToScrap = []scrapField{
		// Debt Indicators
		{
			divLiqSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.DebtIndicators.DivLiq = value
				return nil
			},
		},
		{
			divBrutaSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.DebtIndicators.DivBruta = value
				return nil
			},
		},
		{
			divLiqPLRatioSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.DebtIndicators.DivLiqPLRatio = value
				return nil
			},
		},
		{
			divLiqEbitdaRatioSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.DebtIndicators.DivLiqEbitdaRatio = value
				return nil
			},
		},
		{
			divLiqEbitRatioSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.DebtIndicators.DivLiqEbitRatio = value
				return nil
			},
		},
		{
			plAtivosRatioSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.DebtIndicators.PLAtivosRatio = value
				return nil
			},
		},
		{
			passivosAtivosRatioSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.DebtIndicators.PassivosAtivosRatio = value
				return nil
			},
		},
		{
			liqCorRatioSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.DebtIndicators.LiqCorRatio = value
				return nil
			},
		},
	}

	eficiencyIndicatorsToScrap = []scrapField{
		// Eficiency Indicators
		{
			mBrutaSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.EficiencyIndicators.MBruta = value
				return nil
			},
		},
		{
			mEbitdaSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.EficiencyIndicators.MEbitda = value
				return nil
			},
		},
		{
			mEbitSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.EficiencyIndicators.MEbit = value
				return nil
			},
		},
		{
			mliqSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.EficiencyIndicators.MLiq = value
				return nil
			},
		},
	}

	profitIndicatorsToScrap = []scrapField{
		// Profit Indicators
		{
			roeSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.ProfitIndicators.ROE = value
				return nil
			},
		},
		{
			roaSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.ProfitIndicators.ROA = value
				return nil
			},
		},
		{
			roicSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.ProfitIndicators.ROIC = value
				return nil
			},
		},
		{
			giroAtivosSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.ProfitIndicators.GiroAtivos = value
				return nil
			},
		},
	}

	growthIndicatorsToScrap = []scrapField{
		// Growth Indicators
		{
			garGer5AnosSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.GrowthIndicators.GARGer5Anos = value
				return nil
			},
		},
		{
			garLuc5AnosSelector,
			"float64",
			func(field string, stockData *brazilianinvestimentsscrapper.StockData) error {
				value, err := parseFloatField(field)
				if err != nil {
					return err
				}
				stockData.GrowthIndicators.GARLuc5Anos = value
				return nil
			},
		},
	}

	fieldsToScrap = append(
		stockDataToScrap,
		append(
			valuationIndicatorsToScrap,
			append(
				debtIndicatorsToScrap,
				append(
					eficiencyIndicatorsToScrap,
					append(
						profitIndicatorsToScrap,
						growthIndicatorsToScrap...,
					)...,
				)...,
			)...,
		)...,
	)
)

func GetStockScrapFields() []scrapField {
	return fieldsToScrap
}
