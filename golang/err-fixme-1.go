package main

import  (
   "fmt"
)

//var cells []interface{}=[]interface{}{ "x",124,true }
var cells = []interface{}{ "x", 124, true }

func main() {
   for i:=0; i<len(cells); i++ {
      fmt.Printf("Item %d ", i);
    //switch( cells[i].(type) ) {
      switch  cells[i].(type)   {
         case int:
            fmt.Printf("int :  %d\n",cells[i].(int))
            break;
         case string:
            fmt.Printf("string : %s\n",cells[i].(string))
            break;
         case bool:
            fmt.Printf("bool : %t\n",cells[i].(bool))
            break;
      }
   }
}
