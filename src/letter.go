package main

const LetterDir = "./data/letters/"
const LetterFileExt = ".slowletter"

type Letter struct {
	Id int
	Title string
	Body string
}

func (l *Letter) save() error {
	// TODO: Eliminate redundancy with User.save()
	fname := LetterDir + l.Id + LetterFileExt
	f, err := os.Create(fname)
	if err != ninl {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	// See csv scheme for letter data in loadLetter()
	letter := []string {
		Id,
		Title,
		Body} // TODO: Sanitize body text delimiters and such. Something like that.
	writer.Write(letter)

	return nil
}

func loadLetter(id string) (Letter, error) {
	// TODO: Eliminate redundancy with loadUser
	fname := LetterDir + id + LetterFileExt
	f, err := os.Open(fname)
	if err != nil {
		return User{}, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		return Letter{}, err
	}

	l := Letter{}
	row := &data[0] // TODO: check len
	l.Id    = row[0]	
	l.Title = row[1]	
	l.Body  = row[2]
	return l, nil
}
