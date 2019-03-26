package main

 import (
         "fmt"
 )

 func chunkSplit(body string, limit int, end string) string {

         var charSlice []rune

         // push characters to slice
         for _, char := range body {
                 charSlice = append(charSlice, char)
         }

         var result string = ""

         for len(charSlice) >= 1 {
                 // convert slice/array back to string
                 // but insert end at specified limit

                 result = result + string(charSlice[:limit]) + end

                 // discard the elements that were copied over to result
                 charSlice = charSlice[limit:]

                 // change the limit
                 // to cater for the last few words in
                 // charSlice
                 if len(charSlice) < limit {
                         limit = len(charSlice)
                 }

         }

         return result

 }

 func main() {
         before := "this is a long string that needs to be chunked into smaller chunks"

         // chunk after 30 characters and append with newline
         // for RFC 2045, change limit from 30 to 76
         after := chunkSplit(before, 30, "\n")

         fmt.Println("Before :\n", before)
         fmt.Println("After :\n", after)
 }