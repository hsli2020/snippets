/*
  Wrap long lines on rough column boundaries at spaces
  Working example: https://play.golang.org/p/3u0X6NyMua
  Based on algo from RosettaCode, which is nifty
  https://www.rosettacode.org/wiki/Word_wrap#Go
*/

package main

import (
	"fmt"
	"strings"
)

// Wraps text at the specified column lineWidth on word breaks

func word_wrap(text string, lineWidth int) string {
	words := strings.Fields(strings.TrimSpace(text))
	if len(words) == 0 {
		return text
	}
	wrapped := words[0]
	spaceLeft := lineWidth - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			wrapped += "\n" + word
			spaceLeft = lineWidth - len(word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}

	return wrapped

}

func main() {

	longTextStr :=
		`
		<!-- BEGIN CONTENT -->
		British voters just shattered political convention in a stunning repudiation of the ruling establishment. Donald Trump is betting America is about to do the same.</p></div><div class="zn-body__paragraph">Voters in the UK did more than reject the European Union and topple their pro-EU Prime Minister David Cameron in a referendum Thursday.</div><div class="ad ad--epic ad--mobile" data-ad-text="show"><div id="ad_rect_atf_02" class="ad-ad_rect_atf_02 ad-refresh-default"></div></div><ul class="cn cn-list-hierarchical-xs cn--idx-4 cn-zoneAdContainer"></ul><div class="zn-body__paragraph">They also set off a cascade of events that could spark global economic chaos, remake the Western world, reverberate through November's presidential election and challenge U.S. security for years to come.</div>
		<!-- END CONTENT -->
		`

	fmt.Printf("Original: [%v] \n", longTextStr)
	fmt.Println("--------------------------------")

	wrapped := word_wrap(longTextStr, 70)

	// Some minimal html fixups
	// Note: this can introduce newlines inside class attributes, but that's perfectly
	// valid html (nb: http://stackoverflow.com/a/14928606)

	r := strings.NewReplacer("<p>", "\n<p>", "<!--", "\n<!--", "-->", "-->\n", "<br", "\n<br")
	fmt.Println(r.Replace(wrapped))

}
