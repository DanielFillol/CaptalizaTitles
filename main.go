package main

import (
	"Cap_Titles/capitalizeFunctions"
	"Cap_Titles/csv"
	_struct "Cap_Titles/struct"
	"fmt"
)

const (
	term2            = "csv/data/ti_term_2.csv"
	alternativeLabel = "csv/data/ti_alternative_label_2.csv"
)

func main() {
	dataEntry, err := csv.ReadFiles(term2, alternativeLabel)
	if err != nil {
		fmt.Println(err)
	}

	var dataExit []_struct.DataBase
	for _, dt := range dataEntry {
		spiltWord, err := capitalizeFunctions.SplitWord(dt.Original)
		if err != nil {
			abr := capitalizeFunctions.CapAbbreviations(dt)
			dataExit = append(dataExit, abr)
		} else {
			abr := capitalizeFunctions.CapitalizeWord(spiltWord)
			dataExit = append(dataExit, _struct.DataBase{Id: dt.Id, Original: dt.Original, Text: abr})
		}
	}

	err = csv.WriteCSV("capitalized", "Results", dataExit)
	if err != nil {
		fmt.Println(err)
	}

}
