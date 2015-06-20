package main

type Yeast struct {
	Name           string  `xml:"NAME" json:"NAME"`
	Version        int64   `xml:"VERSION" json:"VERSION"`
	Type           string  `xml:"TYPE" json:"TYPE"`
	Form           string  `xml:"FORM" json:"FORM"`
	Amount         float64 `xml:"AMOUNT" json:"AMOUNT"`
	AmountIsWeight bool    `xml:"AMOUNT_IS_WEIGHT" json:"AMOUNT_IS_WEIGHT"`
	Laboratory     string  `xml:"LABORATORY" json:"LABORATORY"`
	ProductID      string  `xml:"PRODUCT_ID" json:"PRODUCT_ID"`
	MinTemperature float64 `xml:"MIN_TEMPERATURE" json:"MIN_TEMPERATURE"`
	MaxTemperature float64 `xml:"MAX_TEMPERATURE" json:"MAX_TEMPERATURE"`
	Flocculation   string  `xml:"FLOCCULATION" json:"FLOCCULATION"`
	Attenutation   float64 `xml:"ATTENUATION" json:"ATTENUATION"`
	Notes          string  `xml:"NOTES" json:"NOTES"`
	BestFor        string  `xml:"BEST_FOR" json:"BEST_FOR"`
	TimesCultured  int64   `xml:"TIMES_CULTURED" json:"TIMES_CULTURED"`
	MaxReuse       int64   `xml:"MAX_REUSE" json:"MAX_REUSE"`
	AddToSecondary bool    `xml:"ADD_TO_SECONDARY" json:"ADD_TO_SECONDARY"`
}