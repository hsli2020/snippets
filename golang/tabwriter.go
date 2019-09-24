package main
 
import (
	"fmt"
	"os"
	"text/tabwriter"
)
 
func main() {
	
	// initialize tabwriter
	w := new(tabwriter.Writer)
	
	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	
	defer w.Flush()
	
	fmt.Fprintf(w, "\n %s\t%s\t%s\t", "Col1", "Col2", "Col3")
	fmt.Fprintf(w, "\n %s\t%s\t%s\t", "----", "----", "----")
	
	for i := 0; i < 5; i++ {
		fmt.Fprintf(w, "\n %d\t%d\t%d\t", i, i+1, i+2)
	}
	
     //	Col1	Col2	Col3	
     //	----	----	----	
     //	0       1       2	
     //	1       2       3	
     //	2       3       4	
     //	3       4       5	
     // 4       5       6
}