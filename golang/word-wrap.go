package main

 import (
         "fmt"
         "strings"
 )

 // Wraps text at the specified number of columns

 func word_wrap(s string, limit int) string {

         if strings.TrimSpace(s) == "" {
                 return s
         }

         // convert string to slice
         strSlice := strings.Fields(s)

         var result string = ""

         for len(strSlice) >= 1 {
                 // convert slice/array back to string
                 // but insert \r\n at specified limit

                 result = result + strings.Join(strSlice[:limit], " ") + "\r\n"

                 // discard the elements that were copied over to result
                 strSlice = strSlice[limit:]

                 // change the limit
                 // to cater for the last few words in
                 //
                 if len(strSlice) < limit {
                         limit = len(strSlice)
                 }

         }

         return result

 }

 func main() {

         str := "Here is a nice text string consisting of eleven words."
         fmt.Printf("Original : [%v] \n", str)
         // limit to 3 words
         fmt.Println(word_wrap(str, 3))

         strWithDoubleSpaces := "This code example will attempt to word  warp this text string - 'Here    is a    nice   text string consisting of eleven words.'"
         fmt.Printf("Original : [%v] \n", strWithDoubleSpaces)
         // limit to 7 words
         fmt.Println(word_wrap(strWithDoubleSpaces, 7))

         unicodeStrWithDoubleSpaces := "Here    is a    nice   text string consisting of eleven words and some unicode characters such as ????."
         fmt.Printf("Original : [%v] \n", unicodeStrWithDoubleSpaces)
         // limit to 5 words
         fmt.Println(word_wrap(unicodeStrWithDoubleSpaces, 5))

         longTextStr := "KUALA LUMPUR: The FBM KLCI and key Asian markets were in the red at the midday on Monday but off the day’s worst on some buying support at lower levels while China’s markets and the yuan managed to find a steadier footing. At 12.30pm, the KLCI was down 9.43 point or 0.58% to 1,619.12 with 24 out of the 30 stocks in the red. Turnover was 1.32 billion shares valued at RM855.53mil. However, the broader market was much weaker with decliners beating advancers more than eight to one with 819 losers to 105 gainers while 203 counters were unchanged. China stocks opened sharply lower on Monday but managed to erase much of the losses by midday, with real estate shares reversing initial declines on data showing continuous recovery in the property sector, Reuters report. Hong Kong shares dropped, with investors' risk appetites soured by resumed declines in oil prices as well as Friday's tumbles for equities prices in US and European markets. US light crude oil fell 23 cents to US$29.19 and Brent lost 27 cents to US$28.67. Petronas Gas was the top loser, down 38 sen to RM21.08 and erased 1.25 points from the KLCI, Petronas Dagangan fell 18 sen to RM23.72 but Petronas Chemicals rose three sen to RM7.03. SapuraKencana shed two sen to RM1.66. Refiners Shell and Petron Malaysia were among the top losers, Shell was down 29 sen to RM5.23 and  Petron lost 26 sen to RM6.17.Crude palm oil for third-month delivery rose RM26 to RM2,490. Genting Plantations added 14 sen to RM10.36 while PPB Group rose six sen to RM15.78. IOI Corp was flat at RM4.32, KL Kepong fell 18 sen to RM22.68 and Sime Darby three sen to RM7.09. "

         fmt.Printf("Original : [%v] \n", longTextStr)
         fmt.Println("--------------------------------")

         // limit to 11 words per line
         fmt.Println(word_wrap(longTextStr, 11))

 }