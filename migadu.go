/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Mailbox struct {
	LocalPart             string //	Create/Update Only
	Domain                string // Read Only
	Address               string // Read Only
	Name                  string
	IsInternal            bool
	MaySend               bool
	MayReceive            bool
	MayAccessImap         bool
	MayAccessPop3         bool
	MayAccessManagesieve  bool
	PasswordMethod        string // Predefined Values
	Password              string //	Create/Update Only
	PasswordRecoveryEmail string
	SpamAction            string //	Predefined Values
	SpamAggressiveness    string // Predefined Values
	SenderDenylist        string
	SenderAllowlist       string
	RecipientDenylist     string
	AutorespondActive     bool
	AutorespondSubject    string
	AutorespondBody       string
	// TODO Figure out the type
	// AutorespondExpiresOn 	Date
	FooterActive    string
	FooterPlainBody string
	FooterHtmlBody  string
}

type Indentity struct {
	LocalPart            string // Read Only
	Domain               string // Read Only
	Address              string // Read Only
	Name                 string
	MaySend              bool
	MayReceive           bool
	MayAccessImap        bool
	MayAccessPop3        bool
	MayAccessManagesieve bool
	Password             string //	Create/Update Only
	FooterActive         string
	FooterPlainBody      string
	FooterHtmlBody       string
}

type Alias struct {
	LocalPart  string // Create/Read Only
	Domain     string // Read Only
	Address    string // Read Only
	IsInternal bool
	// TODO Figure this out
	// Destinations         Array or CSV String
}

type Rewrite struct {
	Domain        string // Read Only
	Name          string // Slug
	LocalPartRule string
	OrderNum      int64
	// TODO Figure this out
	// Destinations         Array or CSV String
}

const BASE_URL string = "https://api.migadu.com/v1/domains"

type MigaduClient struct {
	Username string
	ApiKey   string
}

func NewMigaduClient(username string, apiKey string) *MigaduClient {
	return &MigaduClient{
		Username: username,
		ApiKey:   apiKey,
	}
}

func (s *MigaduClient) sendRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(s.Username, s.ApiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

// func (s *MigaduClient) ListMailboxes() (*Todo, error) {
// 	url := fmt.Sprintf(baseURL+"/%s/todos/%d", s.Username, id)
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	bytes, err := s.doRequest(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var data Todo
// 	err = json.Unmarshal(bytes, &data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &data, nil
// }
