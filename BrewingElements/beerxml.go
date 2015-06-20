package elements

import (
	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
	"encoding/xml"
	"fmt"
	// "io"
	"bytes"
	"strings"
)

type RecipeList struct {
	Recipes []Recipe `xml:"RECIPE" json:"RECIPE"`
}

type FermentableList struct {
	Fermentables []Fermentable `xml:"FERMENTABLE" json:"FERMENTABLE"`
}

type HopList struct {
	Hops []Hop `xml:"HOP" json:"HOP"`
}

type MiscList struct {
	Miscs []Misc `xml:"MISC" json:"MISC"`
}

type YeastList struct {
	Yeasts []Yeast `xml:"YEAST" json:"YEAST"`
}

type WaterList struct {
	Waters []Water `xml:"WATER" json:"WATER"`
}

type EquipmentList struct {
	Equipmements []Equipment `xml:"EQUIPMENT" json:"EQUIPMENT"`
}

type MashProfileList struct {
	MashProfiles []MashProfile `xml:"MASH" json:"MASH"`
}

type StyleList struct {
	Styles []Style `xml:"STYLE" json:"STYLE"`
}

func Parse(beerXML []byte) {
	peek := string(beerXML[:256])

	reader := bytes.NewReader(beerXML)

	if strings.Contains(peek, "RECIPES") {
		fmt.Println("\tRecipe")
		parsed := RecipeList{}
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = charset.NewReader
		err := decoder.Decode(&parsed)

		if nil != err {
			fmt.Println("Error decoding from XML", err)
			return
		}

		for i := range parsed.Recipes {
			fmt.Println("\t\t" + parsed.Recipes[i].Name)
		}

	} else if strings.Contains(peek, "FERMENTABLE") {
		fmt.Println("\tFermentable")
		parsed := FermentableList{}
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = charset.NewReader
		err := decoder.Decode(&parsed)

		if nil != err {
			fmt.Println("Error decoding from XML", err)
			return
		}

		for i := range parsed.Fermentables {
			fmt.Println("\t\t" + parsed.Fermentables[i].Name)
		}
	} else if strings.Contains(peek, "HOPS") {
		fmt.Println("\tHop")
		parsed := HopList{}
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = charset.NewReader
		err := decoder.Decode(&parsed)

		if nil != err {
			fmt.Println("Error decoding from XML", err)
			return
		}

		for i := range parsed.Hops {
			fmt.Println("\t\t" + parsed.Hops[i].Name)
		}
	} else if strings.Contains(peek, "MISCS") {
		fmt.Println("\tMisc")
		parsed := MiscList{}
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = charset.NewReader
		err := decoder.Decode(&parsed)

		if nil != err {
			fmt.Println("Error decoding from XML", err)
			return
		}

		for i := range parsed.Miscs {
			fmt.Println("\t\t" + parsed.Miscs[i].Name)
		}
	} else if strings.Contains(peek, "YEASTS") {
		fmt.Println("\tYeast")
		parsed := YeastList{}
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = charset.NewReader
		err := decoder.Decode(&parsed)

		if nil != err {
			fmt.Println("Error decoding from XML", err)
			return
		}

		for i := range parsed.Yeasts {
			fmt.Println("\t\t" + parsed.Yeasts[i].Name)
		}
	} else if strings.Contains(peek, "WATERS") {
		fmt.Println("\tWater")
		parsed := WaterList{}
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = charset.NewReader
		err := decoder.Decode(&parsed)

		if nil != err {
			fmt.Println("Error decoding from XML", err)
			return
		}

		for i := range parsed.Waters {
			fmt.Println("\t\t" + parsed.Waters[i].Name)
		}
	} else if strings.Contains(peek, "EQUIPMENTS") {
		fmt.Println("\tEquipment")
		parsed := EquipmentList{}
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = charset.NewReader
		err := decoder.Decode(&parsed)

		if nil != err {
			fmt.Println("Error decoding from XML", err)
			return
		}

		for i := range parsed.Equipmements {
			fmt.Println("\t\t" + parsed.Equipmements[i].Name)
		}
	} else if strings.Contains(peek, "MASHS") {
		fmt.Println("\tMash")
		parsed := MashProfileList{}
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = charset.NewReader
		err := decoder.Decode(&parsed)

		if nil != err {
			fmt.Println("Error decoding from XML", err)
			return
		}

		for i := range parsed.MashProfiles {
			fmt.Println("\t\t" + parsed.MashProfiles[i].Name)
		}
	} else if strings.Contains(peek, "STYLES") {
		fmt.Println("\tStyle")
		parsed := StyleList{}
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = charset.NewReader
		err := decoder.Decode(&parsed)

		if nil != err {
			fmt.Println("Error decoding from XML", err)
			return
		}

		for i := range parsed.Styles {
			fmt.Println("\t\t" + parsed.Styles[i].Name)
		}
	}
}

/*

*/
