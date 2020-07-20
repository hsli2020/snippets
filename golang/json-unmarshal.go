package main

// https://www.calhoun.io/how-to-parse-json-that-varies-between-an-array-or-a-single-item-with-go/

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	bytes1 := []byte(` 
	
	
	  {
            "id": 1,
            "name": "House Stark",
            "created_at": "2016-02-28T08:07:00Z"
	   }`)
	var ctr CreateTagsResponse
	err := json.Unmarshal(bytes1, &ctr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", ctr)


	// Whitespace for days
	bytes2 := []byte(` 
	
	
	[
	  {
            "id": 1,
            "name": "House Stark",
            "created_at": "2016-02-28T08:07:00Z"
           },{
             "id": 2,
             "name": "House Lannister",
             "created_at": "2016-02-28T08:10:00Z"
           }
        ]`)
	//var ctr CreateTagsResponse
	err = json.Unmarshal(bytes2, &ctr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", ctr)
}

// CreateTagsResponse is the data returned from a CreateTags call.
type CreateTagsResponse struct {
	Tags []Tag
}

// UnmarshalJSON implements json.Unmarshaler
func (ctr *CreateTagsResponse) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return fmt.Errorf("no bytes to unmarshal")
	}
	// See if we can guess based on the first character
	switch b[0] {
	case '{':
		return ctr.unmarshalSingle(b)
	case '[':
		return ctr.unmarshalMany(b)
	}
	// TODO: Figure out what do we do here
	return nil
}

func (ctr *CreateTagsResponse) unmarshalSingle(b []byte) error {
	var t Tag
	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}
	ctr.Tags = []Tag{t}
	return nil
}

func (ctr *CreateTagsResponse) unmarshalMany(b []byte) error {
	var tags []Tag
	err := json.Unmarshal(b, &tags)
	if err != nil {
		return err
	}
	ctr.Tags = tags
	return nil
}

// Tag can be applied to subscribers to help filter and customize your mailing
// list actions. Eg you might tag a subscriber "beginner" and send them
// beginner-oriented emails, or you might tag them as interested in a paid
// course so they get information about future sales.
type Tag struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
