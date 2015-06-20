package elements

type Hop struct {
	Name          string  `xml:"NAME" json:"NAME"`
	Version       int64   `xml:"VERSION" json:"VERSION"`
	Origin        string  `xml:"ORIGIN" json:"ORIGIN"`
	Alpha         float64 `xml:"ALPHA" json:"ALPHA"`
	Amount        float64 `xml:"AMOUNT" json:"AMOUNT"`
	Use           string  `xml:"USE" json:"USE"`
	Time          float64 `xml:"TIME" json:"TIME"`
	Notes         string  `xml:"NOTES" json:"NOTES"`
	Type          string  `xml:"TYPE" json:"TYPE"`
	Form          string  `xml:"FORM" json:"FORM"`
	Beta          float64 `xml:"BETA" json:"BETA"`
	HSI           float64 `xml:"HSI" json:"HSI"`
	Substitutes   string  `xml:"SUBSTITUTES" json:"SUBSTITUTES"`
	Humulene      float64 `xml:"HUMULENE" json:"HUMULENE"`
	Caryophyllene float64 `xml:"CARYOPHYLLENE" json:"CARYOPHYLLENE"`
	Cohumulone    float64 `xml:"COHUMULONE" json:"COHUMULONE"`
	Myrcene       float64 `xml:"MYRCENE" json:"MYRCENE"`

	/* Extensions */

	DisplayAmount string `xml:"DISPLAY_AMOUNT" json:"DISPLAY_AMOUNT"`
	Inventory     string `xml:"INVENTORY" json:"INVENTORY"`
	DisplayTime   string `xml:"DISPLAY_TIME" json:"DISPLAY_TIME"`
}
