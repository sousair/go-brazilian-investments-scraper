package scraper

var (
	// Stock Data
	nameSelector               = "#main-header > div:nth-child(2) > div > div:nth-child(1) > h1 > small"
	priceSelector              = "#main-2 > div:nth-child(4) > div > div.pb-3.pb-md-5 > div > div.info.special.w-100.w-md-33.w-lg-20 > div > div:nth-child(1) > strong"
	equitySelector             = "#company-section > div:nth-child(1) > div > div.top-info.info-3.sm.d-flex.justify-between.mb-3 > div:nth-child(1) > div > div > strong"
	totalAssetsSelector        = "#company-section > div:nth-child(1) > div > div.top-info.info-3.sm.d-flex.justify-between.mb-3 > div:nth-child(2) > div > div > strong"
	totalCurrentAssetsSelector = "#company-section > div:nth-child(1) > div > div.top-info.info-3.sm.d-flex.justify-between.mb-3 > div:nth-child(3) > div > div > div > strong"
	grossDebtSelector          = "#company-section > div:nth-child(1) > div > div.top-info.info-3.sm.d-flex.justify-between.mb-3 > div:nth-child(4) > div > div > strong"
	disponibilitiesSelector    = "#company-section > div:nth-child(1) > div > div.top-info.info-3.sm.d-flex.justify-between.mb-3 > div:nth-child(5) > div > div > strong"
	netDebtSelector            = "#company-section > div:nth-child(1) > div > div.top-info.info-3.sm.d-flex.justify-between.mb-3 > div:nth-child(6) > div > div > strong"
	marketCapSelector          = "#company-section > div:nth-child(1) > div > div.top-info.info-3.sm.d-flex.justify-between.mb-3 > div:nth-child(7) > div > div > strong"
	evSelector                 = "#company-section > div:nth-child(1) > div > div.top-info.info-3.sm.d-flex.justify-between.mb-3 > div:nth-child(8) > div > div > strong"
	totalSharesSelector        = "#company-section > div:nth-child(1) > div > div.top-info.info-3.sm.d-flex.justify-between.mb-3 > div:nth-child(9) > div > div > strong"
	freeFloatSelector          = "#company-section > div:nth-child(1) > div > div.top-info.info-3.sm.d-flex.justify-between.mb-3 > div:nth-child(11) > div > div > strong"

	indicatorSectionTodaySelector = "#indicators-section > div.indicator-today-container > div"
	valuationIndicatorsSelector   = indicatorSectionTodaySelector + " > div:nth-child(1) > div"
	debtIndicatorsSelector        = indicatorSectionTodaySelector + " > div:nth-child(2) > div"
	eficiencyIndicatorsSelector   = indicatorSectionTodaySelector + " > div:nth-child(3) > div"
	profitIndicatorsSelector      = indicatorSectionTodaySelector + " > div:nth-child(4) > div"
	growthIndicatorsSelector      = indicatorSectionTodaySelector + " > div:nth-child(5) > div"

	// Valuation Indicators
	dividendYieldSelector      = valuationIndicatorsSelector + " > div:nth-child(1) > div > div > strong"
	plRatioSelector            = valuationIndicatorsSelector + " > div:nth-child(2) > div > div > strong"
	pegRatioSelector           = valuationIndicatorsSelector + " > div:nth-child(3) > div > div > strong"
	pvpRatioSelector           = valuationIndicatorsSelector + " > div:nth-child(4) > div > div > strong"
	evEbitdaRatioSelector      = valuationIndicatorsSelector + " > div:nth-child(5) > div > div > strong"
	evEbitRatioSelector        = valuationIndicatorsSelector + " > div:nth-child(6) > div > div > strong"
	pEbitdaRatioSelector       = valuationIndicatorsSelector + " > div:nth-child(7) > div > div > strong"
	pEbitRatioSelector         = valuationIndicatorsSelector + " > div:nth-child(8) > div > div > strong"
	vpaSelector                = valuationIndicatorsSelector + " > div:nth-child(9) > div > div > strong"
	pAtivoRatioSelector        = valuationIndicatorsSelector + " > div:nth-child(10) > div > div > strong"
	lpaSelector                = valuationIndicatorsSelector + " > div:nth-child(11) > div > div > strong"
	pSRRatioSelector           = valuationIndicatorsSelector + " > div:nth-child(12) > div > div > strong"
	pCapGiroRatioSelector      = valuationIndicatorsSelector + " > div:nth-child(13) > div > div > strong"
	pAtivoCircLiqRatioSelector = valuationIndicatorsSelector + " > div:nth-child(14) > div > div > strong"

	// Debt Indicators
	divLiqSelector              = netDebtSelector
	divBrutaSelector            = grossDebtSelector
	divLiqPLRatioSelector       = debtIndicatorsSelector + " > div:nth-child(1) > div > div > strong"
	divLiqEbitdaRatioSelector   = debtIndicatorsSelector + " > div:nth-child(2) > div > div > strong"
	divLiqEbitRatioSelector     = debtIndicatorsSelector + " > div:nth-child(3) > div > div > strong"
	plAtivosRatioSelector       = debtIndicatorsSelector + " > div:nth-child(4) > div > div > strong"
	passivosAtivosRatioSelector = debtIndicatorsSelector + " > div:nth-child(5) > div > div > strong"
	liqCorRatioSelector         = debtIndicatorsSelector + " > div:nth-child(6) > div > div > strong"

	// Eficiency Indicators
	mBrutaSelector  = eficiencyIndicatorsSelector + " > div:nth-child(1) > div > div > strong"
	mEbitdaSelector = eficiencyIndicatorsSelector + " > div:nth-child(2) > div > div > strong"
	mEbitSelector   = eficiencyIndicatorsSelector + " > div:nth-child(3) > div > div > strong"
	mliqSelector    = eficiencyIndicatorsSelector + " > div:nth-child(4) > div > div > strong"

	// Profit Indicators
	roeSelector        = profitIndicatorsSelector + " > div:nth-child(1) > div > div > strong"
	roaSelector        = profitIndicatorsSelector + " > div:nth-child(2) > div > div > strong"
	roicSelector       = profitIndicatorsSelector + " > div:nth-child(3) > div > div > strong"
	giroAtivosSelector = profitIndicatorsSelector + " > div:nth-child(4) > div > div > strong"

	// Growth Indicators
	garGer5AnosSelector = growthIndicatorsSelector + " > div:nth-child(1) > div > div > strong"
	garLuc5AnosSelector = growthIndicatorsSelector + " > div:nth-child(2) > div > div > strong"
)
