package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Country struct {
	CountryName        CountryName       `json:"name"`
	TLD                []string          `json:"tld"`
	CCA2               string            `json:"cca2"`
	CCN3               string            `json:"ccn3"`
	CCA3               string            `json:"cca3"`
	Currency           []string          `json:"currency"`
	CallingCode        []string          `json:"callingCode"`
	Capital            string            `json:"capital"`
	AlternateSpellings []string          `json:"altSpellings"`
	Relevance          string            `json:"relevance"`
	Region             string            `json:"region"`
	Subregion          string            `json:"subregion"`
	NativeLanguage     string            `json:"nativeLanguage"`
	Languages          map[string]string `json:"languages"`
	Translations       map[string]string `json:"translations"`
	LatLng             [2]float64        `json:"latlng"`
	Demonym            string            `json:"demonym"`
	Borders            []string          `json:"borders"`
	Area               float64           `json:"area"`
}

type CountryName struct {
	Common   string            `json:"common"`
	Official string            `json:"official"`
	Native   CountryNameNative `json:"native"`
}

type CountryNameNative struct {
	Common   string `json:"common"`
	Official string `json:"official"`
}

func readJSONFromUrl(url string) ([]Country, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var countryList []Country
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()
	if err := json.Unmarshal(respByte, &countryList); err != nil {
		return nil, err
	}

	return countryList, nil
}

func main() {
    url := "https://raw.githubusercontent.com/mledoze/countries/master/dist/countries.json"
	countryList, err := readJSONFromUrl(url)
	if err != nil {
		panic(err)
	}

	for idx, row := range countryList {
		// skip header
		if idx == 0 {
			continue
		}

		if idx == 6 {
			break
		}

		fmt.Println(row.CountryName.Common)
	}
}

// Will Print:
// Åland Islands
// Albania
// Algeria
// American Samoa
// Andorra
