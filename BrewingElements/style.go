package elements

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

	/* Extensions */

	DisplayOGMin    string `xml:"DISPLAY_OG_MIN" json:"DISPLAY_OG_MIN"`
	DisplayOGMax    string `xml:"DISPLAY_OG_MAX" json:"DISPLAY_OG_MAX"`
	DisplayFGMin    string `xml:"DISPLAY_FG_MIN" json:"DISPLAY_FG_MIN"`
	DisplayFGMax    string `xml:"DISPLAY_FG_MAX" json:"DISPLAY_FG_MAX"`
	DisplayColorMin string `xml:"DISPLAY_COLOR_MIN" json:"DISPLAY_COLOR_MIN"`
	DisplayColorMax string `xml:"DISPLAY_COLOR_MAX" json:"DISPLAY_COLOR_MAX"`
	OGRange         string `xml:"OG_RANGE" json:"OG_RANGE"`
	FGRange         string `xml:"FG_RANGE" json:"FG_RANGE"`
	IBURange        string `xml:"IBU_RANGE" json:"IBU_RANGE"`
	CarbRange       string `xml:"CARB_RANGE" json:"CARB_RANGE"`
	ColorRange      string `xml:"COLOR_RANGE" json:"COLOR_RANGE"`
	ABVRange        string `xml:"ABV_RANGE" json:"ABV_RANGE"`
}
