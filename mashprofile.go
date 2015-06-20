package main

type MashProfile struct {
	Name            string     `xml:"NAME" json:"NAME"`
	Version         int64      `xml:"VERSION" json:"VERSION"`
	GrainTemp       string     `xml:"GRAIN_TEMP" json:"GRAIN_TEMP"`
	MashSteps       []MashStep `xml:"MASH_STEPS>MASH_STEP" json:"MASH_STEPS>MASH_STEP"`
	Notes           string     `xml:"NOTES" json:"NOTES"`
	TunTemp         float64    `xml:"TUN_TEMP" json:"TUN_TEMP"`
	SpargeTemp      float64    `xml:"SPARGE_TEMP" json:"SPARGE_TEMP"`
	PH              float64    `xml:"PH" json:"PH"`
	TunWeight       float64    `xml:"TUN_WEIGHT" json:"TUN_WEIGHT"`
	TunSpecificHeat float64    `xml:"TUN_SPECIFIC_HEAT" json:"TUN_SPECIFIC_HEAT"`
	EquipAdjust     bool       `xml:"EQUIP_ADJUST" json:"EQUIP_ADJUST"`
}