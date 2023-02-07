package csv

import (
	_struct "Cap_Titles/struct"
	"encoding/csv"
	"os"
)

func ReadFiles(filePath1 string, filePath2 string) ([]_struct.DataBase, error) {
	file1, err := readCsvFile(filePath1, ',')
	if err != nil {
		return nil, err
	}

	file2, err := readCsvFile(filePath2, ',')
	if err != nil {
		return nil, err
	}

	for _, item := range file2 {
		file1 = append(file1, item)
	}

	return file1, nil
}

//readCsvFile Test csv file from Test given path
// the csv must contain only one column with the authors names
func readCsvFile(filePath string, separator rune) ([]_struct.DataBase, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer csvFile.Close()

	csvR := csv.NewReader(csvFile)
	csvR.Comma = separator

	csvData, err := csvR.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []_struct.DataBase
	for _, line := range csvData {
		data = append(data, _struct.DataBase{
			Id:       line[0],
			Original: line[1],
			Text:     "",
		})
	}

	return data, nil
}
