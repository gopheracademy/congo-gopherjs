// +build js

package main

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/gopherjs/gopherjs/js"
	"github.com/shurcooL/go/gopherjs_http/jsutil"

	"honnef.co/go/js/dom"
	"honnef.co/go/js/xhr"
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

func GetSpeakers(event dom.Event) {
	event.PreventDefault()
	if event.(*dom.MouseEvent).Button != 0 {
		return
	}

	go func() {
		req := xhr.NewRequest("GET", "http://arrakis:18080/api/tenants/1/events/1/speakers")
		err := req.Send(url.Values{}.Encode())
		if err != nil {
			println(err.Error())
			return
		}

		speakers := document.GetElementByID("speakers").(dom.HTMLElement)

		var s []Speaker
		err = json.Unmarshal([]byte(req.ResponseText), &s)
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Printf("%+v", s)
		var c string
		for _, speak := range s {
			c = c + fmt.Sprintf("<span>%s</span><br/>", *speak.Bio)
		}

		speakers.SetInnerHTML(c)
	}()
}

func main() {
	js.Global.Set("GetSpeakers", jsutil.Wrap(GetSpeakers))
}
