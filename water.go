package main

type Water struct {
	Name        string  `xml:"NAME" json:"NAME"`
	Version     int64   `xml:"VERSION" json:"VERSION"`
	Amount      float64 `xml:"AMOUNT" json:"AMOUNT"`
	Calcium     float64 `xml:"CALCIUM" json:"CALCIUM"`
	Bicarbonate float64 `xml:"BICARBONATE" json:"BICARBONATE"`
	Sulfate     float64 `xml:"SULFATE" json:"SULFATE"`
	Chloride    float64 `xml:"CHLORIDE" json:"CHLORIDE"`
	Sodium      float64 `xml:"SODIUM" json:"SODIUM"`
	Magnesium   float64 `xml:"MAGNESIUM" json:"MAGNESIUM"`
	PH          float64 `xml:"PH" json:"PH"`
	Notes       string  `xml:"NOTES" json:"NOTES"`
}
