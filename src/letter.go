package main

import (
	"os"
	"encoding/csv"
	"strconv"
)

const LetterDir = "./data/letters/"
const LetterFileExt = ".slowletter"

type Letter struct {
	Id int
	Title string
	Body string
}

func (l *Letter) save() error {
	// TODO: Eliminate redundancy with User.save()
	fname := LetterDir + strconv.Itoa(l.Id) + LetterFileExt
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	// See csv scheme for letter data in loadLetter()
	letter := []string {
		strconv.Itoa(l.Id),
		l.Title,
		l.Body} // TODO: Sanitize body text delimiters and such. Something like that.
	writer.Write(letter)

	return nil
}

func loadLetter(id int) (Letter, error) {
	// TODO: Eliminate redundancy with loadUser
	fname := LetterDir + strconv.Itoa(id) + LetterFileExt
	f, err := os.Open(fname)
	if err != nil {
		return Letter{}, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		return Letter{}, err
	}

	l := Letter{}
	row := data[0] // TODO: check len
	l.Id    = id
	l.Title = row[1]	
	l.Body  = row[2]
	return l, nil
}
