rows, _ := db.Query("SELECT ...") // Note: Ignoring errors for brevity
cols, _ := rows.Columns()

for rows.Next() {
    // Create a slice of interface{}'s to represent each column,
    // and a second slice to contain pointers to each item in the columns slice.
    columns := make([]interface{}, len(cols))
    columnPointers := make([]interface{}, len(cols))
    for i, _ := range columns {
        columnPointers[i] = &columns[i]
    }

    // Scan the result into the column pointers...
    if err := rows.Scan(columnPointers...); err != nil {
        return err
    }

    // Create our map, and retrieve the value for each column from the pointers slice,
    // storing it in the map with the name of the column as the key.
    m := make(map[string]interface{})
    for i, colName := range cols {
        val := columnPointers[i].(*interface{})
        m[colName] = *val
    }

    // Outputs: map[columnName:value columnName2:value2 columnName3:value3 ...] 
    fmt.Print(m)
}

////////////////////////////////////////////////////////////////////////////////

db, _ := sql.Open("mysql", "user=demo dbname=demo password=test")

rows, _ := db.Query("select * from sites")

columns, _ := rows.Columns()
count := len(columns)
values := make([]interface{}, count)
valuePtrs := make([]interface{}, count)

final_result := map[int]map[string]string{}
result_id := 0
for rows.Next() {
    for i, _ := range columns {
        valuePtrs[i] = &values[i]
    }
    rows.Scan(valuePtrs...)

    tmp_struct := map[string]string{}

    for i, col := range columns {
        var v interface{}
        val := values[i]
        b, ok := val.([]byte)
        if (ok) {
            v = string(b)
        } else {
            v = val
        }
        tmp_struct[col] = fmt.Sprintf("%s",v)
    }

    final_result[ result_id ] = tmp_struct
    result_id++
}

fmt.Println(final_result)
