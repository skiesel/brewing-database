package main

type Style struct {
	Name           string  `xml:"NAME" json:"NAME"`
	Category       string  `xml:"CATEGORY" json:"CATEGORY"`
	Version        int64   `xml:"VERSION" json:"VERSION"`
	CategoryNumber string  `xml:"CATEGORY_NUMBER" json:"CATEGORY_NUMBER"`
	StyleLetter    string  `xml:"STYLE_LETTER" json:"STYLE_LETTER"`
	StyleGuide     string  `xml:"STYLE_GUIDE" json:"STYLE_GUIDE"`
	Type           string  `xml:"TYPE" json:"TYPE"`
	OGMin          float64 `xml:"OG_MIN" json:"OG_MIN"`
	OGMax          float64 `xml:"OG_MAX" json:"OG_MAX"`
	FGMin          float64 `xml:"FG_MIN" json:"FG_MIN"`
	FGMax          float64 `xml:"FG_MAX" json:"FG_MAX"`
	IBUMin         float64 `xml:"IBU_MIN" json:"IBU_MIN"`
	IBUMax         float64 `xml:"IBU_MAX" json:"IBU_MAX"`
	ColorMin       float64 `xml:"COLOR_MIN" json:"COLOR_MIN"`
	ColorMax       float64 `xml:"COLOR_MAX" json:"COLOR_MAX"`
	CarbMin        float64 `xml:"CARB_MIN" json:"CARB_MIN"`
	CarbMax        float64 `xml:"CARB_MAX" json:"CARB_MAX"`
	ABVMin         float64 `xml:"ABV_MIN" json:"ABV_MIN"`
	ABVMax         float64 `xml:"ABV_MAX" json:"ABV_MAX"`
	Notes          string  `xml:"NOTES" json:"NOTES"`
	Profile        string  `xml:"PROFILE" json:"PROFILE"`
	Ingredients    string  `xml:"INGREDIENTS" json:"INGREDIENTS"`
	Examples       string  `xml:"EXAMPLES" json:"EXAMPLES"`
}
