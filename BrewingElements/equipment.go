package elements

type Equipment struct {
	Name            string  `xml:"NAME" json:"NAME"`
	Version         int64   `xml:"VERSION" json:"VERSION"`
	BoilSize        float64 `xml:"BOIL_SIZE" json:"BOIL_SIZE"`
	BatchSize       float64 `xml:"BATCH_SIZE" json:"BATCH_SIZE"`
	TunVolume       float64 `xml:"TUN_VOLUME" json:"TUN_VOLUME"`
	TunWeight       float64 `xml:"TUN_WEIGHT" json:"TUN_WEIGHT"`
	TunSpecificHeat float64 `xml:"TUN_SPECIFIC_HEAT" json:"TUN_SPECIFIC_HEAT"`
	TopUpWater      float64 `xml:"TOP_UP_WATER" json:"TOP_UP_WATER"`
	TrubChillerLoss float64 `xml:"TRUB_CHILLER_LOSS" json:"TRUB_CHILLER_LOSS"`
	EvapRate        float64 `xml:"EVAP_RATE" json:"EVAP_RATE"`
	Sulfate         float64 `xml:"BOIL_TIME" json:"BOIL_TIME"`
	CalcBoilVolume  bool    `xml:"CALC_BOIL_VOLUME" json:"CALC_BOIL_VOLUME"`
	LauterDeadspace float64 `xml:"LAUTER_DEADSPACE" json:"LAUTER_DEADSPACE"`
	TopUpKettle     float64 `xml:"TOP_UP_KETTLE" json:"TOP_UP_KETTLE"`
	HopUtilization  float64 `xml:"HOP_UTILIZATION" json:"HOP_UTILIZATION"`
	Notes           string  `xml:"NOTES" json:"NOTES"`

	/* Extensions */

	DisplayBoilSize        string `xml:"DISPLAY_BOIL_SIZE" json:"DISPLAY_BOIL_SIZE"`
	DisplayBatchSize       string `xml:"DISPLAY_BATCH_SIZE" json:"DISPLAY_BATCH_SIZE"`
	DisplayTunVolume       string `xml:"DISPLAY_TUN_VOLUME" json:"DISPLAY_TUN_VOLUME"`
	DisplayTunWeight       string `xml:"DISPLAY_TUN_WEIGHT" json:"DISPLAY_TUN_WEIGHT"`
	DisplayTopUpWater      string `xml:"DISPLAY_TOP_UP_WATER" json:"DISPLAY_TOP_UP_WATER"`
	DisplayTrubChillerLoss string `xml:"DISPLAY_TRUB_CHILLER_LOSS" json:"DISPLAY_TRUB_CHILLER_LOSS"`
	DisplayLauterDeadspace string `xml:"DISPLAY_LAUTER_DEADSPACE" json:"DISPLAY_LAUTER_DEADSPACE"`
	DisplayTopUpKettle     string `xml:"DISPLAY_TOP_UP_KETTLE" json:"DISPLAY_TOP_UP_KETTLE"`
}
