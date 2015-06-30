package frontend

import (
	"fmt"
	"github.com/skiesel/brewing-database/MySQL"
	"html/template"
	"net/http"
)

type xmlElement struct {
	Field string
	Description string
	Format string
	Required bool
}

var (
	beerXML = map[string][]xmlElement{}
)

func BeerXML(w http.ResponseWriter, r *http.Request) {
	if len(beerXML) == 0 {
		readBeerXML()
	}

	t, err := template.ParseFiles("FrontEnd/Templates/beerxml.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, beerXML)
}

func readBeerXML() {
	rows := database.GetBeerXMLRows()
	for rows.Next() {
		elem := xmlElement{}
		var elemType string
		var req int64
		err := rows.Scan(&elemType, &elem.Field, &elem.Description, &elem.Format, &req)
		if err != nil {
			fmt.Println(err)
			continue
		}
		elem.Required = req == 1
		beerXML[elemType] = append(beerXML[elemType], elem)
	}
}