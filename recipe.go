package main

type Recipe struct {
	Name               string        `xml:"NAME" json:"NAME"`
	Version            int64         `xml:"VERSION" json:"VERSION"`
	Type               string        `xml:"TYPE" json:"TYPE"`
	Style              Style         `xml:"STYLE" json:"STYLE"`
	Equipment          Equipment     `xml:"EQUIPMENT" json:"EQUIPMENT"`
	Brewer             string        `xml:"BREWER" json:"BREWER"`
	AsstBrewer         string        `xml:"ASST_BREWER" json:"ASST_BREWER"`
	BatchSize          float64       `xml:"BATCH_SIZE" json:"BATCH_SIZE"`
	BoilSize           float64       `xml:"BOIL_SIZE" json:"BOIL_SIZE"`
	BoilTime           float64       `xml:"BOIL_TIME" json:"BOIL_TIME"`
	Efficiency         float64       `xml:"EFFICIENCY" json:"EFFICIENCY"`
	Hops               []Hop         `xml:"HOPS>HOP" json:"HOPS>HOP"`
	Fermentables       []Fermentable `xml:"FERMENTABLES>FERMENTABLE" json:"FERMENTABLES>FERMENTABLE"`
	Miscs              []Misc        `xml:"MISCS>MISC" json:"MISCS>MISC"`
	Yeasts             []Yeast       `xml:"YEASTS>YEAST" json:"YEASTS>YEAST"`
	Waters             []Water       `xml:"WATERS>WATER" json:"WATERS>WATER"`
	Mash               MashProfile   `xml:"MASH" json:"MASH"`
	Notes              string        `xml:"NOTES" json:"NOTES"`
	TasteNotes         string        `xml:"TASTE_NOTES" json:"TASTE_NOTES"`
	TasteRating        float64       `xml:"TASTE_RATING" json:"TASTE_RATING"`
	OG                 float64       `xml:"OG" json:"OG"`
	FG                 float64       `xml:"FG" json:"FG"`
	FermentationStages int64         `xml:"FERMENTATION_STAGES" json:"FERMENTATION_STAGES"`
	PrimaryAge         float64       `xml:"PRIMARY_AGE" json:"PRIMARY_AGE"`
	PrimaryTemp        float64       `xml:"PRIMARY_TEMP" json:"PRIMARY_TEMP"`
	SecondaryAge       float64       `xml:"SECONDARY_AGE" json:"SECONDARY_AGE"`
	SecondaryTemp      float64       `xml:"SECONDARY_TEMP" json:"SECONDARY_TEMP"`
	TertiaryAge        float64       `xml:"TERTIARY_AGE" json:"TERTIARY_AGE"`
	TertiaryTemp       float64       `xml:"TERTIARY_TEMP" json:"TERTIARY_TEMP"`
	Age                float64       `xml:"AGE" json:"AGE"`
	AgeTemp            float64       `xml:"AGE_TEMP" json:"AGE_TEMP"`
	Date               string        `xml:"DATE" json:"DATE"`
	Carbonation        float64       `xml:"CARBONATION" json:"CARBONATION"`
	ForceCarbonation   bool          `xml:"FORCED_CARBONATION" json:"FORCED_CARBONATION"`
	PrimingSugarName   string        `xml:"PRIMING_SUGAR_NAME" json:"PRIMING_SUGAR_NAME"`
	CarbonationTemp    float64       `xml:"CARBONATION_TEMP" json:"CARBONATION_TEMP"`
	PrimingSugarEquiv  float64       `xml:"PRIMING_SUGAR_EQUIV" json:"PRIMING_SUGAR_EQUIV"`
	KegPrimingFactor   float64       `xml:"KEG_PRIMING_FACTOR" json:"KEG_PRIMING_FACTOR"`

	/* Extensions */

	EstOG                string  `xml:"EST_OG" json:"EST_OG"`
	EstFg                string  `xml:"EST_FG" json:"EST_FG"`
	EstColor             string  `xml:"EST_COLOR" json:"EST_COLOR"`
	IBU                  float64 `xml:"IBU" json:"IBU"`
	IBUMethod            string  `xml:"IBU_METHOD" json:"IBU_METHOD"`
	EstABV               float64 `xml:"EST_ABV" json:"EST_ABV"`
	ABV                  float64 `xml:"ABV" json:"ABV"`
	ActualEfficiency     float64 `xml:"ACTUAL_EFFICIENCY" json:"ACTUAL_EFFICIENCY"`
	Calories             string  `xml:"CALORIES" json:"CALORIES"`
	DisplayBatchSize     string  `xml:"DISPLAY_BATCH_SIZE" json:"DISPLAY_BATCH_SIZE"`
	DisplayBoilSize      string  `xml:"DISPLAY_BOIL_SIZE" json:"DISPLAY_BOIL_SIZE"`
	DisplayOG            string  `xml:"DISPLAY_OG" json:"DISPLAY_OG"`
	DisplayFG            string  `xml:"DISPLAY_FG" json:"DISPLAY_FG"`
	DisplayPrimaryTemp   string  `xml:"DISPLAY_PRIMARY_TEMP" json:"DISPLAY_PRIMARY_TEMP"`
	DisplaySecondaryTemp string  `xml:"DISPLAY_SECONDARY_TEMP" json:"DISPLAY_SECONDARY_TEMP"`
	DisplayTertiaryTemp  string  `xml:"DISPLAY_TERTIARY_TEMP" json:"DISPLAY_TERTIARY_TEMP"`
	DisplayAgeTemp       string  `xml:"DISPLAY_AGE_TEMP" json:"DISPLAY_AGE_TEMP"`
	CarbonationUsed      string  `xml:"CARBONATION_USED" json:"CARBONATION_USED"`
	DisplayCarbTemp      string  `xml:"DISPLAY_CARB_TEMP" json:"DISPLAY_CARB_TEMP"`
}
