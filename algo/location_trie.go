package algo

import (
	"encoding/json"
	"io/ioutil"
)

type CountryIdTrie struct {
	node [10]*CountryIdTrie
	code string
	id   int
}

type CountryCodeTrie struct {
	node [26]*CountryCodeTrie
	code string
	id   int
}

type Country struct {
	CountryId   int    `json:"countryId"`
	CountryCode string `json:"countryCode"`
}

func InitLocationTrie(dataPath string) (*CountryIdTrie, *CountryCodeTrie, error) {
	fileData, err := ioutil.ReadFile(dataPath)
	if err != nil {
		return nil, nil, err
	}
	countryList := make([]*Country, 0)
	err = json.Unmarshal(fileData, &countryList)
	if err != nil {
		return nil, nil, err
	}
	rootCountryIdTrie := new(CountryIdTrie)
	rootCountryCodeTrie := new(CountryCodeTrie)
	for _, c := range countryList {
		tempCountryCodeTrie := rootCountryCodeTrie
		for _, str := range []byte(c.CountryCode) {
			index := int(str - 'A')
			if tempCountryCodeTrie.node[index] == nil {
				tempCountryCodeTrie.node[index] = new(CountryCodeTrie)
			}
			tempCountryCodeTrie = tempCountryCodeTrie.node[index]
		}
		tempCountryCodeTrie.code = c.CountryCode
		tempCountryCodeTrie.id = c.CountryId
		tempCountryIdTrie := rootCountryIdTrie
		countryId := c.CountryId
		for countryId != 0 {
			index := countryId % 10
			if tempCountryIdTrie.node[index] == nil {
				tempCountryIdTrie.node[index] = new(CountryIdTrie)
			}
			tempCountryIdTrie = tempCountryIdTrie.node[index]
			countryId = countryId / 10
		}
		tempCountryIdTrie.id = c.CountryId
		tempCountryIdTrie.code = c.CountryCode
	}
	return rootCountryIdTrie, rootCountryCodeTrie, nil
}

func (c *CountryIdTrie) ParseCountryIdToCode(countryId int) (code string, ok bool) {
	tempCountryIdTrie := c
	tempCountryId := countryId
	for countryId != 0 {
		index := countryId % 10
		if tempCountryIdTrie.node[index] == nil {
			return
		}
		tempCountryIdTrie = tempCountryIdTrie.node[index]
		countryId = countryId / 10
	}
	if tempCountryIdTrie.id == tempCountryId {
		code = tempCountryIdTrie.code
		ok = true
	}
	return
}

func (c *CountryCodeTrie) ParseCountryCodeToId(countryCode string) (id int, ok bool) {
	tempCountryCodeTrie := c
	for _, str := range []byte(countryCode) {
		index := int(str - 'A')
		if tempCountryCodeTrie.node[index] == nil {
			return
		}
		tempCountryCodeTrie = tempCountryCodeTrie.node[index]
	}
	if tempCountryCodeTrie.code == countryCode {
		id = tempCountryCodeTrie.id
		ok = true
	}
	return
}
