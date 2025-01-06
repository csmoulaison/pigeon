package main

import (
	"strings"
	"os"
	"strconv"
	"bufio"
	"net/http"
	"time"
)

const LetterDir = "./data/letters/"
const LetterFileExt = ".slowletter"
const LetterDateFmt = time.DateTime

type Letter struct {
	Id int
	Title string
	Created time.Time
	Sender string
	Recipient string
	Read bool
	Body string
}

func newLetterId() (int, error) {
    files, err := os.ReadDir(LetterDir)
    if err != nil {
		return 0, err
    }

    highestId := 0
    for _, f := range files {
	    fname := f.Name()
	    id := 0
		id, err = strconv.Atoi(fname[:strings.Index(fname, LetterFileExt)])
		if err != nil {
			return 0, err
		}

		if id > highestId {
			highestId = id
		}
    }
	return highestId + 1, nil
}

func (l *Letter) save() error {
	fname := LetterDir + strconv.Itoa(l.Id) + LetterFileExt
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString(l.Title + "\n")
	w.WriteString(l.Created.Format(LetterDateFmt) + "\n")
	w.WriteString(l.Sender + "\n")
	w.WriteString(l.Recipient + "\n")
	readString := "1"
	if !l.Read {
		readString = "0"
	}
	w.WriteString(readString + "\n")
	w.WriteString(l.Body)
	w.Flush()

	return nil
}

func loadLetter(id int) (Letter, error) {
	fname := LetterDir + strconv.Itoa(id) + LetterFileExt
	f, err := os.Open(fname)
	if err != nil {
		return Letter{}, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	l := Letter{}
	l.Id = id

	s.Scan()
	l.Title = s.Text()

	s.Scan()
	l.Created, err = time.Parse(LetterDateFmt, s.Text())
	if err != nil {
		return l, err
	}

	s.Scan()
	l.Sender = s.Text()

	s.Scan()
	l.Recipient = s.Text()

	s.Scan()
	l.Read = s.Text() == "1"

	for s.Scan() {
		l.Body += s.Text()
		l.Body += "\n"
	}

	if err = s.Err(); err != nil {
        return l, err
    }

	return l, nil
}

func lettersFromCache(w http.ResponseWriter, cache []int) []Letter {
	letters := []Letter{}
	for _, id := range cache {
		l, err := loadLetter(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		letters = append(letters, l)
	}
	return letters
}

