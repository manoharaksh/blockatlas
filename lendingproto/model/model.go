package model

type (
	// LendingProvider static info about the lending provider, such as name and asset classes supported.
	LendingProvider struct {
		ID     string              `json:"id"`
		Info   LendingProviderInfo `json:"info"`
		Assets []AssetClass        `json:"assets"`
	}

	// LendingProviderInfo basic information about a lending provider.
	LendingProviderInfo struct {
		ID          string `json:"id"`
		Description string `json:"description"`
		Image       string `json:"image"`
		Website     string `json:"website"`
	}

	// AssetClass Info about an asset that can be lent
	AssetClass struct {
		Symbol      string `json:"symbol"`
		Chain       string `json:"chain"`
		Description string `json:"description"`
		// YieldFrequency the period of yield computation in seconds, e.g. 86400 for daily.
		YieldFrequency int64 `json:"yield_freq"`
		// Terms Predefined lending term periods, like [7, 30.5, 180].
		Terms []Term `json:"terms"`
	}

	// Term length of a predefined term, in days
	Term float64

	// RatesRequest Rates API request
	RatesRequest struct {
		Assets []string `json:"assets"`
	}

	// RatesResponse Rates API response
	RatesResponse struct {
		Provider string       `json:"provider"`
		Rates    LendingRates `json:"rates"`
	}

	// LendingTermAPR Asset yield APR, for an asset for a term.  E.g. {30, 1.45}
	LendingTermAPR struct {
		Term `json:"term"`
		APR  float64 `json:"apr"`
	}

	// LendingAssetRates Asset yield rates, for an asset for one or more periods.  E.g. [{7, 0.9}, {30, 1.45}]
	LendingAssetRates struct {
		Asset     string           `json:"asset"`
		TermRates []LendingTermAPR `json:"term_rates"`
		// MaxAPR the rate of the term with the highest rate
		MaxAPR float64 `json:"max_apr"`
	}

	// LendingRates List of yield rates, for multiple assets.
	LendingRates []LendingAssetRates

	// AccountRequest Account API request
	AccountRequest struct {
		Address string   `json:"address"`
		Assets  []string `json:"assets"`
	}

	// AccountResponse Account API response, contracts
	AccountResponse struct {
		Contracts AccountLendingContracts `json:"contracts"`
	}

	// Time Second-granular UNIX time
	Time int32

	// LendingContract Describes a lending contract, of a user, of an asset.
	LendingContract struct {
		Asset             string  `json:"asset"`
		Term              Term    `json:"term"`
		StartAmount       string  `json:"start_amount"`
		CurrentAmount     string  `json:"current_amount"`
		EndAmountEstimate string  `json:"end_amount_estimate"`
		CurrentAPR        float64 `json:"current_apr"`
		StartTime         Time    `json:"start_time"`
		CurrentTime       Time    `json:"current_time"`
		EndTime           Time    `json:"end_time"`
	}

	// AccountLendingContracts Contracts of an address
	AccountLendingContracts struct {
		Address   string            `json:"address"`
		Contracts []LendingContract `json:"contracts"`
	}
)