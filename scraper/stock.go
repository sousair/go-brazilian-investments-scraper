package scraper

type StockData struct {
	// Pregão
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
	// Cotação
	Price float64 `json:"price"`
	// Patrimônio Liquido
	Equity float64 `json:"equity"`
	// Ativos
	TotalAssets float64 `json:"total_assets"`
	// Ativo Circulante
	TotalCurrentAssets float64 `json:"total_current_assets"`
	// Divida Bruta
	GrossDebt float64 `json:"gross_debt"`
	// Disponibilidades
	Disponibilities float64 `json:"disponibilities"`
	// Divida Liquida
	NetDebt float64 `json:"net_debt"`
	// Valor de Mercado
	MarketCap float64 `json:"market_cap"`
	// Valor de Firma
	EV float64 `json:"ev"`
	// Total de papéis
	TotalShares int `json:"total_shares"`
	// Free Float
	FreeFloat float64 `json:"free_float"`

	ValuationIndicators ValuationIndicators `json:"valuation_indicators"`
	DebtIndicators      DebtIndicators      `json:"debt_indicators"`
	EficiencyIndicators EficiencyIndicators `json:"eficiency_indicators"`
	ProfitIndicators    ProfitIndicators    `json:"profit_indicators"`
	GrowthIndicators    GrowthIndicators    `json:"growth_indicators"`
}

type ValuationIndicators struct {
	// DY
	DividendYield float64 `json:"dividend_yield"`
	// P/L
	PLRatio float64 `json:"pl_ratio"`
	// PEG RATIO
	PEGRatio float64 `json:"peg_ratio"`
	// P/VP
	PVpRatio float64 `json:"pvp_ratio"`
	// EV/EBITDA
	EVEbitdaRatio float64 `json:"ev_ebitda_ratio"`
	// EV/EBIT
	EVEbitRatio float64 `json:"ev_ebit_ratio"`
	// P/EBITDA
	PEBitdaRatio float64 `json:"p_ebitda_ratio"`
	// P/EBIT
	PEBitRatio float64 `json:"p_ebit_ratio"`
	// VPA
	VPA float64 `json:"vpa"`
	// P/ATIVO
	PAtivoRatio float64 `json:"p_ativo_ratio"`
	// LPA
	LPA float64 `json:"lpa"`
	// P/SR
	PSRRatio float64 `json:"p_sr_ratio"`
	// P/CAP. GIRO
	PCapGiroRatio float64 `json:"p_cap_giro_ratio"`
	// P/ATIVO CIRC. LIQ.
	PAtivoCircLiqRatio float64 `json:"p_ativo_circ_liq_ratio"`
}

type DebtIndicators struct {
	// Div. Liq.
	DivLiq float64 `json:"div_liq"`
	// Div. Bruta
	DivBruta float64 `json:"div_bruta"`
	// Div. Liq. / PL
	DivLiqPLRatio float64 `json:"div_liq_pl_ratio"`
	// Div. Liq. / EBITDA
	DivLiqEbitdaRatio float64 `json:"div_liq_ebitda_ratio"`
	// Div. Liq. / EBIT
	DivLiqEbitRatio float64 `json:"div_liq_ebit_ratio"`
	// PL / ATIVOS
	PLAtivosRatio float64 `json:"pl_ativos_ratio"`
	// Passivos / Ativos
	PassivosAtivosRatio float64 `json:"passivos_ativos_ratio"`
	// Liq. Corrente
	LiqCorRatio float64 `json:"liq_cor_ratio"`
}

type EficiencyIndicators struct {
	// Margem Bruta
	MBruta float64 `json:"m_bruta"`
	// Margem EBITDA
	MEbitda float64 `json:"m_ebitda"`
	// Margem EBIT
	MEbit float64 `json:"m_ebit"`
	// Margem Liquida
	MLiq float64 `json:"m_liq"`
}

type ProfitIndicators struct {
	// ROE
	ROE float64 `json:"roe"`
	// ROA
	ROA float64 `json:"roa"`
	// ROIC
	ROIC float64 `json:"roic"`
	// Giro Ativos
	GiroAtivos float64 `json:"giro_ativos"`
}

type GrowthIndicators struct {
	// Crescimento Receita Líquida 5 anos
	GARGer5Anos float64 `json:"gar_ger_5_anos"`
	// Crescimento EBITDA 5 anos
	GARLuc5Anos float64 `json:"gar_luc_5_anos"`
}
