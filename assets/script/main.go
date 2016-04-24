// +build js

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document().(dom.HTMLDocument)

type Speaker struct {
	// speaker bio
	Bio *string `json:"bio,omitempty" xml:"bio,omitempty"`
	// first name
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// github handle
	Github *string `json:"github,omitempty" xml:"github,omitempty"`
	// ID of record
	ID *int `json:"id,omitempty" xml:"id,omitempty"`
	// url of speaker image
	ImageURL *string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// last name
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// linkedin url
	Linkedin *string `json:"linkedin,omitempty" xml:"linkedin,omitempty"`
	// twitter handle - no @
	Twitter *string `json:"twitter,omitempty" xml:"twitter,omitempty"`
}

func main() {
	speakers, err := getSpeakers()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(speakers)
	speaks := document.GetElementByID("speakers").(dom.HTMLElement)
	var c string
	for _, speak := range speakers {
		c = c + fmt.Sprintf("<span>%s</span><br/>", *speak.Bio)
	}

	speaks.SetInnerHTML(c)
}

func getSpeakers() ([]Speaker, error) {
	resp, err := http.Get("http://arrakis:18080/api/tenants/1/events/1/speakers")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("did not get acceptable status code: %v body: %q", resp.Status, body)
	}
	var s []Speaker
	err = json.NewDecoder(resp.Body).Decode(&s)
	return s, err
}
