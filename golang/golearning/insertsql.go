package main

import(
    "fmt"
    "strings"
)

/*
function genInsertSql($table, $columns, $data)
{
    $columnList = '`' . implode('`, `', $columns) . '`';

    $query = "INSERT INTO `$table` ($columnList) VALUES\n";

    $values = array();

    foreach ($data as $row) {
        foreach($row as &$val) {
            $val = addslashes($val);
        }
        $values[] = "('" . implode("', '", $row). "')";
    }

    $update = implode(', ',
        array_map(function($name) {
            return "`$name`=VALUES(`$name`)";
        }, $columns)
    );

    return $query . implode(",\n", $values) . "\nON DUPLICATE KEY UPDATE " . $update . ';';
}
*/

func InsertSql(table string, columns []string, data []map[string]string) string {

    columnStr := strings.Join(columns, "`, `")

    updateList := make([]string, len(columns))
    for i, col := range columns {
        updateList[i] = fmt.Sprintf("`%s`=VALUE(`%s`)", col, col)
    }
    updateStr := strings.Join(updateList, ",\n")

    valueList := make([]string, 0)
    for _, row := range data {
        valueRow := make([]string, 0)
        //for _, val := range row { // WRONG
        for _, col := range columns {
            valueRow = append(valueRow, "'" + row[col] + "'")
        }
        valueList = append(valueList, "(" + strings.Join(valueRow, ", ") + ")")
    }
    valueStr := strings.Join(valueList, ",\n")

    //return "INSERT INTO `" + table + "` (" + columnStr + ") VALUES\n" + valueStr + updateStr;
    return fmt.Sprintf("INSERT INTO `%s` (`%s`) VALUES\n%s\nON DUPLICATE KEY UPDATE\n%s",
        table, columnStr, valueStr, updateStr);
}

func main() {
    columns := []string{"sku", "qty", "price"}

    items := make([]map[string]string, 0)
//*
    for i:=1; i<10; i++ {
        items = append(items, map[string]string{
            "sku":   fmt.Sprintf("SKU-ABC-%d", i),
            "qty":   fmt.Sprintf("%d", i),
            "price": fmt.Sprintf("%d.%d", i*11, i*11),
        })
    }
//*/ 
/*
    items = append(items, map[string]string{
        "sku":   "SKU-ABC-1",
        "qty":   "1",
        "price": "11.11",
    })

    items = append(items, map[string]string{
        "sku":   "SKU-ABC-2",
        "qty":   "2",
        "price": "22.22",
    })

    items = append(items, map[string]string{
        "sku":   "SKU-ABC-3",
        "qty":   "3",
        "price": "33.33",
    })
//*/
    fmt.Println(InsertSql("mytable", columns, items));
}
