package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type R1 struct {
	Page    int
	Fruits  []string
}

type R2 struct {
	Page int 		 `json:"page"`
	Fruits []string  `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	
	intB,_ := json.Marshal(1)
	fmt.Println(string(intB))
	
	fltB,_ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	
	strB,_ := json.Marshal("gopher")
	fmt.Println(string(strB))
	
	slcD := []string{"a","b","c"}
	slcB,_ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	
	mapD := map[string]int{"a":5, "l":7}
	mapB,_ := json.Marshal(mapD)
	fmt.Println(string(mapB))
	
	r1D := &R1 {
		Page: 1,
		Fruits: []string{"a", "b", "c"}
	}
	r1B,_ := json.Marshal(r1D)
	fmt.Println(string(r1B))
	
	r2D := &R2 {
		Page: 1,
		Fruists: []string{"c","b","c"}
	}
	r2B,_ := json.Marshal(r2D)
	fmt.Println(string(r2B))
	
	byt := []byte(`{"num":6.0, "strs":["a","b"]}`)
	
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	
	num := dat["num"].(float64)
	fmt.Println(num)
	
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)
	
	str := `{"page": 1, "fruits": ["a","b"]}`
	res := &R2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
	
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"a":5, "b":7}
	enc.Encode(d)
}