```go
package models

import "time"

// Fundamental represents a company's complete financial profile with 30+ metrics
type Fundamental struct {
	// Identifiers
	Ticker      string `json:"ticker" bson:"ticker"`
	CIK         string `json:"cik" bson:"cik"`
	CompanyName string `json:"company_name" bson:"company_name"`
	Sector      string `json:"sector" bson:"sector"`
	Industry    string `json:"industry" bson:"industry"`

	// Period Info
	FiscalYear    int       `json:"fiscal_year" bson:"fiscal_year"`
	FiscalQuarter int       `json:"fiscal_quarter" bson:"fiscal_quarter"`
	ReportDate    time.Time `json:"report_date" bson:"report_date"`
	FiledDate     time.Time `json:"filed_date" bson:"filed_date"`
	AccessionNum  string    `json:"accession_num" bson:"accession_num"`

	// Income Statement Metrics
	Revenue              float64 `json:"revenue" bson:"revenue"`
	RevenueGrowth        float64 `json:"revenue_growth" bson:"revenue_growth"`                    // YoY %
	RevenueGrowthQoQ     float64 `json:"revenue_growth_qoq" bson:"revenue_growth_qoq"`             // QoQ %
	GrossProfit          float64 `json:"gross_profit" bson:"gross_profit"`
	GrossMargin          float64 `json:"gross_margin" bson:"gross_margin"`                         // % of revenue
	OperatingIncome      float64 `json:"operating_income" bson:"operating_income"`
	OperatingMargin      float64 `json:"operating_margin" bson:"operating_margin"`                 // % of revenue
	EBITDA               float64 `json:"ebitda" bson:"ebitda"`
	EBITDAMargin         float64 `json:"ebitda_margin" bson:"ebitda_margin"`                       // % of revenue
	EBITDAGrowth         float64 `json:"ebitda_growth" bson:"ebitda_growth"`                       // YoY %
	NetIncome            float64 `json:"net_income" bson:"net_income"`
	NetIncomeGrowth      float64 `json:"net_income_growth" bson:"net_income_growth"`               // YoY %
	NetMargin            float64 `json:"net_margin" bson:"net_margin"`                             // % of revenue
	EPS                  float64 `json:"eps" bson:"eps"`                                           // Earnings Per Share
	EPSGrowth            float64 `json:"eps_growth" bson:"eps_growth"`                             // YoY %
	ForwardEPS           float64 `json:"forward_eps" bson:"forward_eps"`                           // ML-predicted
	DilutedEPS           float64 `json:"diluted_eps" bson:"diluted_eps"`

	// Expense Metrics
	CostOfGoodsSold      float64 `json:"cogs" bson:"cogs"`
	ResearchDevelopment  float64 `json:"rd_expenses" bson:"rd_expenses"`
	RDAsPercentRevenue   float64 `json:"rd_percent_revenue" bson:"rd_percent_revenue"`
	SellingGeneral       float64 `json:"sg_a_expenses" bson:"sg_a_expenses"`
	SGAAsPercentRevenue  float64 `json:"sga_percent_revenue" bson:"sga_percent_revenue"`
	InterestExpense      float64 `json:"interest_expense" bson:"interest_expense"`
	TaxExpense           float64 `json:"tax_expense" bson:"tax_expense"`
	EffectiveTaxRate     float64 `json:"effective_tax_rate" bson:"effective_tax_rate"`             // %

	// Cash Flow Metrics
	OperatingCashFlow    float64 `json:"operating_cash_flow" bson:"operating_cash_flow"`
	FreeCashFlow         float64 `json:"free_cash_flow" bson:"free_cash_flow"`                     // OCF - CapEx
	FreeCashFlowMargin   float64 `json:"fcf_margin" bson:"fcf_margin"`                             // FCF / Revenue %
	FreeCashFlowGrowth   float64 `json:"fcf_growth" bson:"fcf_growth"`                             // YoY %
	CapitalExpenditure   float64 `json:"capex" bson:"capex"`
	CapexAsPercentRevenue float64 `json:"capex_percent_revenue" bson:"capex_percent_revenue"`

	// Balance Sheet Metrics
	TotalAssets          float64 `json:"total_assets" bson:"total_assets"`
	CurrentAssets        float64 `json:"current_assets" bson:"current_assets"`
	Cash                 float64 `json:"cash" bson:"cash"`
	AccountsReceivable   float64 `json:"accounts_receivable" bson:"accounts_receivable"`
	Inventory            float64 `json:"inventory" bson:"inventory"`
	TotalLiabilities     float64 `json:"total_liabilities" bson:"total_liabilities"`
	CurrentLiabilities   float64 `json:"current_liabilities" bson:"current_liabilities"`
	LongTermDebt         float64 `json:"long_term_debt" bson:"long_term_debt"`
	ShortTermDebt        float64 `json:"short_term_debt" bson:"short_term_debt"`
	TotalDebt            float64 `json:"total_debt" bson:"total_debt"`                             // LT + ST debt
	ShareholdersEquity   float64 `json:"shareholders_equity" bson:"shareholders_equity"`
	RetainedEarnings     float64 `json:"retained_earnings" bson:"retained_earnings"`

	// Liquidity Ratios
	CurrentRatio         float64 `json:"current_ratio" bson:"current_ratio"`                       // CA / CL
	QuickRatio           float64 `json:"quick_ratio" bson:"quick_ratio"`                           // (CA - Inv) / CL
	CashRatio            float64 `json:"cash_ratio" bson:"cash_ratio"`                             // Cash / CL
	WorkingCapital       float64 `json:"working_capital" bson:"working_capital"`                   // CA - CL

	// Profitability Ratios
	ROE                  float64 `json:"roe" bson:"roe"`                                           // Net Income / SE %
	ROA                  float64 `json:"roa" bson:"roa"`                                           // Net Income / TA %
	ROIC                 float64 `json:"roic" bson:"roic"`                                         // Return on Invested Capital %
	ReturnOnCapital      float64 `json:"return_on_capital" bson:"return_on_capital"`

	// Valuation Ratios
	PERatio              float64 `json:"pe_ratio" bson:"pe_ratio"`                                 // Price / EPS
	ForwardPERatio       float64 `json:"forward_pe_ratio" bson:"forward_pe_ratio"`                 // Price / Forward EPS
	PBRatio              float64 `json:"pb_ratio" bson:"pb_ratio"`                                 // Price / Book Value
	PSRatio              float64 `json:"ps_ratio" bson:"ps_ratio"`                                 // Price / Sales
	PCFRatio             float64 `json:"pcf_ratio" bson:"pcf_ratio"`                               // Price / Operating CF
	PEGRatio             float64 `json:"peg_ratio" bson:"peg_ratio"`                               // PE / Earnings Growth %
	EvToRevenue          float64 `json:"ev_to_revenue" bson:"ev_to_revenue"`
	EvToEBITDA           float64 `json:"ev_to_ebitda" bson:"ev_to_ebitda"`
	FCFYield             float64 `json:"fcf_yield" bson:"fcf_yield"`                               // FCF / Market Cap %

	// Leverage & Solvency Ratios
	DebtToEquity         float64 `json:"debt_to_equity" bson:"debt_to_equity"`                     // TD / SE
	DebtToAssets         float64 `json:"debt_to_assets" bson:"debt_to_assets"`                     // TD / TA
	EquityMultiplier     float64 `json:"equity_multiplier" bson:"equity_multiplier"`               // TA / SE
	InterestCoverage     float64 `json:"interest_coverage" bson:"interest_coverage"`               // EBIT / IE
	DebtServiceCoverage  float64 `json:"debt_service_coverage" bson:"debt_service_coverage"`       // OCF / Debt
	NetDebt              float64 `json:"net_debt" bson:"net_debt"`                                 // TD - Cash

	// Efficiency Ratios
	AssetTurnover        float64 `json:"asset_turnover" bson:"asset_turnover"`                     // Revenue / TA
	InventoryTurnover    float64 `json:"inventory_turnover" bson:"inventory_turnover"`             // COGS / Inv
	ReceivablesTurnover  float64 `json:"receivables_turnover" bson:"receivables_turnover"`         // Revenue / AR
	DaysInventory        float64 `json:"days_inventory" bson:"days_inventory"`                     // 365 / Inv Turnover
	DaysSalesOut         float64 `json:"days_sales_out" bson:"days_sales_out"`                     // 365 / Rec Turnover

	// Valuation Estimates (from ML model)
	FairValue            float64 `json:"fair_value" bson:"fair_value"`                             // ResMLP prediction
	FairValueRange       struct {
		Low  float64 `json:"low" bson:"low"`
		High float64 `json:"high" bson:"high"`
	} `json:"fair_value_range" bson:"fair_value_range"`
	ValuationConfidence  float64 `json:"valuation_confidence" bson:"valuation_confidence"`         // 0-1

	// Dividend & Share Info
	DividendPerShare     float64 `json:"dividend_per_share" bson:"dividend_per_share"`
	DividendYield        float64 `json:"dividend_yield" bson:"dividend_yield"`                     // %
	PayoutRatio          float64 `json:"payout_ratio" bson:"payout_ratio"`                         // Div / NI %
	SharesOutstanding    float64 `json:"shares_outstanding" bson:"shares_outstanding"`             // millions
	MarketCap            float64 `json:"market_cap" bson:"market_cap"`

	// Metadata
	LastUpdated time.Time `json:"last_updated" bson:"last_updated"`
	DataQuality string    `json:"data_quality" bson:"data_quality"`                                // "high", "medium", "low"
}

// YFinanceData represents real-time market data from yfinance
type YFinanceData struct {
	Ticker          string    `json:"ticker" bson:"ticker"`
	Date            time.Time `json:"date" bson:"date"`
	Price           float64   `json:"price" bson:"price"`
	Volume          int64     `json:"volume" bson:"volume"`
	MarketCap       float64   `json:"market_cap" bson:"market_cap"`
	Dividend        float64   `json:"dividend" bson:"dividend"`
	DividendYield   float64   `json:"dividend_yield" bson:"dividend_yield"`
	Beta            float64   `json:"beta" bson:"beta"`
	PERatio         float64   `json:"pe_ratio" bson:"pe_ratio"`
	EPS             float64   `json:"eps" bson:"eps"`
	FiftyTwoWeekHigh float64  `json:"52_week_high" bson:"52_week_high"`
	FiftyTwoWeekLow float64   `json:"52_week_low" bson:"52_week_low"`
	MovingAverage50 float64   `json:"moving_avg_50" bson:"moving_avg_50"`
	MovingAverage200 float64  `json:"moving_avg_200" bson:"moving_avg_200"`
	LastUpdated     time.Time `json:"last_updated" bson:"last_updated"`
}

// PlaidHolding represents user portfolio holding synced via Plaid
type PlaidHolding struct {
	UserID           string    `json:"user_id" bson:"user_id"`
	SecurityID       string    `json:"security_id" bson:"security_id"`
	Ticker           string    `json:"ticker" bson:"ticker"`
	Quantity         float64   `json:"quantity" bson:"quantity"`
	CostBasis        float64   `json:"cost_basis" bson:"cost_basis"`
	CostBasisPerUnit float64   `json:"cost_basis_per_unit" bson:"cost_basis_per_unit"`
	CurrentValue     float64   `json:"current_value" bson:"current_value"`
	CurrentPrice     float64   `json:"current_price" bson:"current_price"`
	GainLoss         float64   `json:"gain_loss" bson:"gain_loss"`                           // Current - Cost
	GainLossPercent  float64   `json:"gain_loss_percent" bson:"gain_loss_percent"`           // %
	PortfolioWeight  float64   `json:"portfolio_weight" bson:"portfolio_weight"`             // % of total
	LastUpdated      time.Time `json:"last_updated" bson:"last_updated"`
}

// PortfolioMetrics aggregates user portfolio-level analytics
type PortfolioMetrics struct {
	UserID                  string                 `json:"user_id" bson:"user_id"`
	TotalValue              float64                `json:"total_value" bson:"total_value"`
	TotalCostBasis          float64                `json:"total_cost_basis" bson:"total_cost_basis"`
	TotalGainLoss           float64                `json:"total_gain_loss" bson:"total_gain_loss"`
	TotalGainLossPercent    float64                `json:"total_gain_loss_percent" bson:"total_gain_loss_percent"`
	DiversificationScore    float64                `json:"diversification_score" bson:"diversification_score"`     // 0-100
	SectorExposure          map[string]float64     `json:"sector_exposure" bson:"sector_exposure"`                 // % per sector
	PortfolioWeightedPE     float64                `json:"portfolio_weighted_pe" bson:"portfolio_weighted_pe"`
	PortfolioWeightedFCFYield float64              `json:"portfolio_fcf_yield" bson:"portfolio_fcf_yield"`
	PortfolioValuationGap   float64                `json:"portfolio_valuation_gap" bson:"portfolio_valuation_gap"` // Fair vs Market
	LastUpdated             time.Time              `json:"last_updated" bson:"last_updated"`
}

// ComparisonMetrics for valuation gap analysis
type ComparisonMetrics struct {
	Ticker              string  `json:"ticker" bson:"ticker"`
	CurrentPrice        float64 `json:"current_price" bson:"current_price"`
	FairValue           float64 `json:"fair_value" bson:"fair_value"`
	ValuationGap        float64 `json:"valuation_gap" bson:"valuation_gap"`                      // Fair - Current
	ValuationGapPercent float64 `json:"valuation_gap_percent" bson:"valuation_gap_percent"`     // %
	Signal              string  `json:"signal" bson:"signal"`                                    // "overvalued", "fair", "undervalued"
	Confidence          float64 `json:"confidence" bson:"confidence"`                            // 0-1
}
```