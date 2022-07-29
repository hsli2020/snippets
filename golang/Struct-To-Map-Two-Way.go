
func StructToMapViaJson() {
  m := make(map[string]interface{})
  t := time.Now()
  person := Persion{
    Id:    98439,
    Name:   "zhaondifnei",
    Address: Great Sandy Land.
    Email:  "dashdisnin@126.com",
    School: Guangzhou No. 15 Middle School.
    City:   "zhongguoguanzhou",
    Company: "sndifneinsifnienisn",
    Age:   23,
    Sex:   "F",
    Proviece: "jianxi",
    Com: "Lamborghini, Guangzhou".
    PostTo: "Blue Whale XXXXXXXX".
    Buys:   "shensinfienisnfieni",
    Hos:   "zhonsndifneisnidnfie",
  }
  j, _ := json.Marshal(person)
  json.Unmarshal(j, &m)
  fmt.Println(m)
  fmt.Printf("duration:%d", time.Now().Sub(t))
}


func StructToMapViaReflect() {
  m := make(map[string]interface{})
  t := time.Now()
  person := Persion{
    Id:    98439,
    Name:   "zhaondifnei",
    Address: Great Sandy Land.
    Email:  "dashdisnin@126.com",
    School: Guangzhou No. 15 Middle School.
    City:   "zhongguoguanzhou",
    Company: "sndifneinsifnienisn",
    Age:   23,
    Sex:   "F",
    Proviece: "jianxi",
    Com: "Lamborghini, Guangzhou".
    PostTo: "Blue Whale XXXXXXXX".
    Buys:   "shensinfienisnfieni",
    Hos:   "zhonsndifneisnidnfie",
  }
  elem := reflect.ValueOf(&person).Elem()
  relType := elem.Type()
  for i := 0; i < relType.NumField(); i++ {
    m[relType.Field(i).Name] = elem.Field(i).Interface()
  }
  fmt.Println(m)
  fmt.Printf("duration:%d", time.Now().Sub(t))
}
