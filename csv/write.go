package csv

import (
	_struct "Cap_Titles/struct"
	"encoding/csv"
	"os"
	"path/filepath"
)

//WriteCSV exports csv to given folder, with a given name
func WriteCSV(fileName string, folderName string, decisions []_struct.DataBase) error {
	var rows [][]string

	rows = append(rows, generateHeaders())

	for _, decision := range decisions {
		rows = append(rows, generateRow(decision))
	}

	cf, err := createFile(folderName + "/" + fileName + ".csv")
	if err != nil {
		return err
	}

	defer cf.Close()

	w := csv.NewWriter(cf)

	err = w.WriteAll(rows)
	if err != nil {
		return err
	}

	return nil
}

// create csv file from operating system
func createFile(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

// generate the necessary headers for csv file
func generateHeaders() []string {
	return []string{
		"ID",
		"Original",
		"UpperCase",
	}
}

// returns []string that compose the row in the csv file
func generateRow(result _struct.DataBase) []string {
	return []string{
		result.Id,
		result.Original,
		result.Text,
	}
}
