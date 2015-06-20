package database

import (
	"database/sql"
	"fmt"
	"github.com/skiesel/brewing-database/BrewingElements"
)

func parseRecipes(rows *sql.Rows) []elements.Recipe {
	recipes := []elements.Recipe{}
	for rows.Next() {
		var recipe elements.Recipe
		err := rows.Scan(&recipe)
		if err != nil {
			fmt.Println(err)
			continue
		}

		var id int64
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println(err)
			continue
		}

		query := fmt.Sprintf("SELECT * FROM StyleLookup INNER JOIN Style ON StyleLookup.StyleID=Style.ID WHERE StyleLookup.RecipeID=%d", id)
		styleRows, err := db.Query(query)
		styles := parseStyles(styleRows)
		if len(styles) > 0 {
			recipe.Style = styles[0]
		}

		query = fmt.Sprintf("SELECT * FROM EquipmentLookup INNER JOIN Equipment ON EquipmentLookup.EquipmentID=Equipment.ID WHERE EquipmentLookup.RecipeID=%d", id)
		equipmentRows, err := db.Query(query)
		equipment := parseEqupiment(equipmentRows)
		if len(equipment) > 0 {
			recipe.Equipment = equipment[0]
		}

		query = fmt.Sprintf("SELECT * FROM MashProfileLookup INNER JOIN MashProfile ON MashProfileLookup.MashProfileID=MashProfile.ID WHERE MashProfileLookup.RecipeID=%d", id)
		mashProfileRows, err := db.Query(query)
		profiles := parseMashProfiles(mashProfileRows)
		if len(profiles) > 0 {
			recipe.Mash = profiles[0]
		}

		query = fmt.Sprintf("SELECT * FROM HopLookup INNER JOIN Hop ON HopLookup.HopID=Hop.ID WHERE HopLookup.RecipeID=%d", id)
		hopRows, err := db.Query(query)
		recipe.Hops = parseHops(hopRows)

		query = fmt.Sprintf("SELECT * FROM FermentableLookup INNER JOIN Fermentable ON FermentableLookup.FermentableID=Fermentable.ID WHERE FermentableLookup.RecipeID=%d", id)
		fermentalRows, err := db.Query(query)
		recipe.Fermentables = parseFermentables(fermentalRows)

		query = fmt.Sprintf("SELECT * FROM MiscLookup INNER JOIN Misc ON MiscLookup.MiscID=Misc.ID WHERE MiscLookup.RecipeID=%d", id)
		miscRows, err := db.Query(query)
		recipe.Miscs = parseMiscs(miscRows)

		query = fmt.Sprintf("SELECT * FROM YeastLookup INNER JOIN Yeast ON YeastLookup.YeastID=Yeast.ID WHERE YeastLookup.RecipeID=%d", id)
		yeastRows, err := db.Query(query)
		recipe.Yeasts = parseYeasts(yeastRows)

		query = fmt.Sprintf("SELECT * FROM WaterLookup INNER JOIN Water ON WaterLookup.WaterID=Water.ID WHERE WaterLookup.RecipeID=%d", id)
		waterRows, err := db.Query(query)
		recipe.Waters = parseWaters(waterRows)

		recipes = append(recipes, recipe)
	}
	return recipes
}

func parseMashProfiles(rows *sql.Rows) []elements.MashProfile {
	profiles := []elements.MashProfile{}
	for rows.Next() {
		var profile elements.MashProfile
		err := rows.Scan(&profile)
		if err != nil {
			fmt.Println(err)
			continue
		}
		var id int64
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println(err)
			continue
		}

		query := fmt.Sprintf("SELECT * FROM MashStepLookup INNER JOIN MashStep ON MashStepLookup.MashStepID=MashStep.ID WHERE MashStepLookup.MashProfileID=%d", id)
		mashSteps, err := db.Query(query)

		profile.MashSteps = parseMashSteps(mashSteps)

		profiles = append(profiles, profile)
	}
	return profiles
}

func parseMashSteps(rows *sql.Rows) []elements.MashStep {
	steps := []elements.MashStep{}
	for rows.Next() {
		var step elements.MashStep
		err := rows.Scan(&step)
		if err != nil {
			fmt.Println(err)
			continue
		}
		steps = append(steps, step)
	}
	return steps
}

func parseYeasts(rows *sql.Rows) []elements.Yeast {
	yeasts := []elements.Yeast{}
	for rows.Next() {
		var yeast elements.Yeast
		err := rows.Scan(&yeast)
		if err != nil {
			fmt.Println(err)
			continue
		}
		yeasts = append(yeasts, yeast)
	}
	return yeasts
}

func parseEqupiment(rows *sql.Rows) []elements.Equipment {
	equipment := []elements.Equipment{}
	for rows.Next() {
		var equip elements.Equipment
		err := rows.Scan(&equip)
		if err != nil {
			fmt.Println(err)
			continue
		}
		equipment = append(equipment, equip)
	}
	return equipment
}

func parseMiscs(rows *sql.Rows) []elements.Misc {
	miscs := []elements.Misc{}
	for rows.Next() {
		var misc elements.Misc
		err := rows.Scan(&misc)
		if err != nil {
			fmt.Println(err)
			continue
		}
		miscs = append(miscs, misc)
	}
	return miscs
}

func parseStyles(rows *sql.Rows) []elements.Style {
	styles := []elements.Style{}
	for rows.Next() {
		var style elements.Style
		err := rows.Scan(&style)
		if err != nil {
			fmt.Println(err)
			continue
		}
		styles = append(styles, style)
	}
	return styles
}

func parseWaters(rows *sql.Rows) []elements.Water {
	waters := []elements.Water{}
	for rows.Next() {
		var water elements.Water
		err := rows.Scan(&water)
		if err != nil {
			fmt.Println(err)
			continue
		}
		waters = append(waters, water)
	}
	return waters
}

func parseFermentables(rows *sql.Rows) []elements.Fermentable {
	feremntables := []elements.Fermentable{}
	for rows.Next() {
		var ferementable elements.Fermentable
		err := rows.Scan(&ferementable)
		if err != nil {
			fmt.Println(err)
			continue
		}
		feremntables = append(feremntables, ferementable)
	}
	return feremntables
}

func parseHops(rows *sql.Rows) []elements.Hop {
	hops := []elements.Hop{}
	for rows.Next() {
		var hop elements.Hop
		err := rows.Scan(&hop)
		if err != nil {
			fmt.Println(err)
			continue
		}
		hops = append(hops, hop)
	}
	return hops
}
