package main

import (
	"io/ioutil"
	"os"
	"strings"
	"encoding/csv"
	"strconv"
)

const UserDir = "./data/users/"
const UserFileExt = ".slowuser"

type User struct {
	Handle string
	Password string
	Rolodex []string
	DisplayName string
	Email string
	NotifyByEmail bool
	MailboxCache []int
	SentCache []int
}

func (u *User) save() error {
	fname := UserDir + u.Handle + UserFileExt
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	// See CSV schema for user data in loadUser()
	// Conversion from bool to string for NotifyByEmail field
	notifyString := "0"
	if u.NotifyByEmail {
		notifyString = "1"
	}
	// D = Detail
	userDetails := []string{
		"D",
		u.Handle, 
		u.Password, 
		u.DisplayName, 
		u.Email, 
		notifyString}
	writer.Write(userDetails)

	// C = Contact
	for _, handle := range u.Rolodex {
		contact := []string{
			"C",
			handle}
		writer.Write(contact)
	}

	// R = Mailbox (received) letter cached id
	for _, index := range u.MailboxCache {
		mailboxCache := []string{
			"R",
			handle}
		writer.Write(mailboxCache)
	}

	// S = Sent letter cached id
	for _, handle := range u.SentCache {
		sentCache := []string{
			"S",
			handle}
		writer.Write(sentCache)
	}

	return nil
}

func loadUser(handle string) (User, error) {
	fname := UserDir + handle + UserFileExt
	f, err := os.Open(fname)
	if err != nil {
		return User{}, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		return User{}, err
	}

	// There are multiple types of line, determined by the indicator at row[0].
	// 1. D,HANDLE,PASSWORD,DISPLAYNAME,EMAIL,NOTIFYBYEMAIL("0"||"1") - Detail
	// 2. C,HANDLE - Contact
	// 2. R,LETTERID - Mailbox Cache (received letters)
	// 2. S,LETTERID - Sent Cache (sent letters)
	u := User{}
	for _, row := range data {
		switch row[0] {
		case "D": // Detail
			u.Handle      = row[1]
			u.Password    = row[2]
			u.DisplayName = row[3]
			u.Email       = row[4]
			// Conversion from string to bool for NotifyByEmail
			if row[5] == "0" {
				u.NotifyByEmail = false
			} else {
				u.NotifyByEmail = true
			}
		case "C": // Contact
			u.Rolodex = append(u.Rolodex, row[1])
		case "R": // Mailbox (received) letters cached id
			u.MailboxCache = append(u.MailboxCache, atoi(row[1])
		case "S": // Sent letters cached id
			u.SentCache = append(u.SentCache, atoi(row[1])
		}
	}
	return u, nil
}

func getUserList() ([]User, error) {
	files, err := ioutil.ReadDir(UserDir)
	if err != nil {
		return nil, err
	}

	users := []User{}
	for _, f := range files {
		fname := f.Name()
		u, err := loadUser(fname[:strings.Index(fname, ".slowuser")])
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
