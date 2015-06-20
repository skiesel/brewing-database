package main

type Misc struct {
	Name           string  `xml:"NAME" json:"NAME"`
	Version        int64   `xml:"VERSION" json:"VERSION"`
	Type           string  `xml:"TYPE" json:"TYPE"`
	Use            string  `xml:"USE" json:"USE"`
	Time           float64 `xml:"TIME" json:"TIME"`
	Amount         float64 `xml:"AMOUNT" json:"AMOUNT"`
	AmountIsWeight bool    `xml:"AMOUNT_IS_WEIGHT" json:"AMOUNT_IS_WEIGHT"`
	UseFor         string  `xml:"USE_FOR" json:"USE_FOR"`
	Notes          string  `xml:"NOTES" json:"NOTES"`
}
