package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	text := "新版党史仍维持中共历史决议，指文革不是任何意义上的革命或社会进步，而是一场由领导者错误发动，被反革命集团利用，给党、国家和各族人民带来严重灾难的内乱，留下了极其惨痛的教训"
	TextToSpeech(text, "speech.mp3", "zh")
}

func TextToSpeech(text, filename, lang string) error {
	apiURL := `http://translate.google.com/translate_tts?ie=UTF-8&total=1&idx=0&textlen=32&client=tw-ob&q=%s&tl=%s`
	response, err := http.Get(apiURL, url.QueryEscape(text), lang)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	output, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer output.Close()

	_, err = io.Copy(output, response.Body)
	return err
}
