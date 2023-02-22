package models

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/henriSedjame/passwords/src/messages"
	"io/fs"
	"io/ioutil"
)

type Password struct {
	Label string
	Value string
}

type Passwords []Password

func (p *Passwords) Add(name, value string) error {
	if p.ExistsByName(name) {
		return errors.New(fmt.Sprintf(messages.AlreadyExist, name))
	}

	item := Password{
		Label: name,
		Value: base64.StdEncoding.EncodeToString([]byte(value)),
	}
	*p = append(*p, item)
	return nil
}

func (p *Passwords) Update(name, value string) error {

	for i, pass := range *p {
		if pass.Label == name {
			(*p)[i] = Password{
				Label: name,
				Value: base64.StdEncoding.EncodeToString([]byte(value)),
			}
			return nil
		}
	}

	return errors.New(fmt.Sprintf(messages.PasswordNotFound, name))
}

func (p *Passwords) Find(name string) (string, error) {
	for _, password := range *p {
		if password.Label == name {
			decodeString, err := base64.URLEncoding.DecodeString(password.Value)
			if err != nil {
				return "", nil
			}
			return string(decodeString), nil
		}
	}
	return "", errors.New(fmt.Sprintf(messages.PasswordNotFound, name))
}

func (p *Passwords) Delete(name string) error {

	for i, password := range *p {
		if password.Label == name {
			*p = append((*p)[:i], (*p)[i+1:]...)
			return nil
		}
	}
	return errors.New(fmt.Sprintf(messages.PasswordNotFound, name))
}

func (p *Passwords) Store(filename string) error {
	if data, err := json.Marshal(p); err != nil {
		return err
	} else {
		return ioutil.WriteFile(filename, data, fs.ModePerm)
	}
}

func (p *Passwords) Load(filename string) error {
	if data, err := ioutil.ReadFile(filename); err != nil {
		return err
	} else {
		if err := json.Unmarshal(data, p); err != nil {
			return err
		}
	}
	return nil
}

func (p *Passwords) ExistsByName(name string) bool {
	for _, password := range *p {
		if password.Label == name {
			return true
		}
	}
	return false
}
