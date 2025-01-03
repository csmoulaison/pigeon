package main

import (
	"io/ioutil"
	"os"
	"strings"
	"encoding/csv"
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

	// See csv schema for user details in loadUser()
	notifyString := "0"
	if u.NotifyByEmail {
		notifyString = "1"
	}
	userDetails := []string{
		"DETAILS",
		u.Handle, 
		u.Password, 
		u.DisplayName, 
		u.Email, 
		notifyString}
	writer.Write(userDetails)

	for _, handle := range u.Rolodex {
		contact := []string{
			"CONTACT",
			handle}
		writer.Write(contact)
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

	// There are two types of line, determined by the indicator at row[0], 
	// which can be either DETAILS or CONTACT:
	// 1. DETAILS,HANDLE,PASSWORD,DISPLAYNAME,EMAIL,NOTIFYBYEMAIL("0"||"1")
	// 2. CONTACT,HANDLE
	u := User{}
	for _, row := range data {
		switch row[0] {
		case "DETAILS":
			u.Handle      = row[1]
			u.Password    = row[2]
			u.DisplayName = row[3]
			u.Email       = row[4]
			if row[5] == "0" {
				u.NotifyByEmail = false
			} else {
				u.NotifyByEmail = true
			}
		case "CONTACT":
			u.Rolodex = append(u.Rolodex, row[1])
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
