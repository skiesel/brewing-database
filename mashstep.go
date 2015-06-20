package main

type MashStep struct {
	Name         string  `xml:"NAME" json:"NAME"`
	Version      int64   `xml:"VERSION" json:"VERSION"`
	Type         string  `xml:"TYPE" json:"TYPE"`
	InfuseAmount float64 `xml:"INFUSE_AMOUNT" json:"INFUSE_AMOUNT"`
	StepTemp     float64 `xml:"STEP_TEMP" json:"STEP_TEMP"`
	StepTime     float64 `xml:"STEP_TIME" json:"STEP_TIME"`
	RampTime     float64 `xml:"RAMP_TIME" json:"RAMP_TIME"`
	EndTemp      float64 `xml:"END_TEMP" json:"END_TEMP"`
}
