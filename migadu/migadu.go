/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package migadu

import (
	"bytes"
	"encoding/json"
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

func (s *MigaduClient) sendRequest(method string, url string, payload any, response any) error {

	j, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(j))

	req.SetBasicAuth(s.Username, s.ApiKey)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if 200 != resp.StatusCode {
		return fmt.Errorf("%s", body)
	}

	return json.Unmarshal(body, response)
}

type ListMailboxesResult struct {
	Mailboxes []Mailbox `json:mailboxes`
}

func (s *MigaduClient) ListMailboxes(domain string) (ListMailboxesResult, error) {
	url := fmt.Sprintf(BASE_URL+"/%s/mailboxes", domain)
	var data ListMailboxesResult
	err := s.sendRequest("GET", url, nil, &data)
	return data, err
}

func (s *MigaduClient) ShowMailbox(domain string, localPart string) (Mailbox, error) {
	url := fmt.Sprintf(BASE_URL+"/%s/mailboxes/%s", domain, localPart)
	var data Mailbox
	err := s.sendRequest("GET", url, nil, &data)
	return data, err
}

func (s *MigaduClient) CreateMailbox(domain string) ([]Mailbox, error) {
	url := fmt.Sprintf(BASE_URL+"/%s/mailboxes", domain)
	var data []Mailbox
	err := s.sendRequest("POST", url, nil, &data)
	return data, err
}

func (s *MigaduClient) UpdateMailbox(domain string, localPart string) ([]Mailbox, error) {
	url := fmt.Sprintf(BASE_URL+"/%s/mailboxes/%s", domain, localPart)
	var data []Mailbox
	err := s.sendRequest("PUT", url, nil, &data)
	return data, err
}

func (s *MigaduClient) DeleteMailbox(domain string, localPart string) ([]Mailbox, error) {
	url := fmt.Sprintf(BASE_URL+"/%s/mailboxes/%s", domain, localPart)
	var data []Mailbox
	err := s.sendRequest("DELETE", url, nil, &data)
	return data, err
}
