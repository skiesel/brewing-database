package elements

type Fermentable struct {
	Name           string  `xml:"NAME" json:"NAME"`
	Version        int64   `xml:"VERSION" json:"VERSION"`
	Type           string  `xml:"TYPE" json:"TYPE"`
	Amount         float64 `xml:"AMOUNT" json:"AMOUNT"`
	Yield          float64 `xml:"YIELD" json:"YIELD"`
	Color          float64 `xml:"COLOR" json:"COLOR"`
	AddAfterBoil   bool    `xml:"ADD_AFTER_BOIL" json:"ADD_AFTER_BOIL"`
	Origin         string  `xml:"ORIGIN" json:"ORIGIN"`
	Supplier       string  `xml:"SUPPLIER" json:"SUPPLIER"`
	Notes          string  `xml:"NOTES" json:"NOTES"`
	CoarseFineDiff float64 `xml:"COARSE_FINE_DIFF" json:"COARSE_FINE_DIFF"`
	Moisture       float64 `xml:"MOISTURE" json:"MOISTURE"`
	DiastaticPower float64 `xml:"DIASTATIC_POWER" json:"DIASTATIC_POWER"`
	Protein        float64 `xml:"PROTEIN" json:"PROTEIN"`
	MaxInBatch     float64 `xml:"MAX_IN_BATCH" json:"MAX_IN_BATCH"`
	RecommendMash  bool    `xml:"RECOMMEND_MASH" json:"RECOMMEND_MASH"`
	IBUGalPerLB    string  `xml:"IBU_GAL_PER_LB" json:"IBU_GAL_PER_LB"`

	/* Extensions */

	DisplayAmount string `xml:"DISPLAY_AMOUNT" json:"DISPLAY_AMOUNT"`
	Potential     string `xml:"POTENTIAL" json:"POTENTIAL"`
	Inventory     string `xml:"INVENTORY" json:"INVENTORY"`
	DisplayColor  string `xml:"DISPLAY_COLOR" json:"DISPLAY_COLOR"`
}
