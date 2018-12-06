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

    foreach($data as $row) {
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

    columnList := make([]string, 0)
    updateList := make([]string, 0)

    for _, col := range columns {
        columnList = append(columnList, "`" + col + "`")
        updateList = append(updateList, "`" + col + "`=VALUE(`" + col + "`)")
    }

    columnStr := strings.Join(columnList, ", ")
    updateStr := "\nON DUPLICATE KEY UPDATE\n"
    updateStr += strings.Join(updateList, ", ")

    valueList := make([]string, 0)
    for _, row := range data {
        valueRow := make([]string, 0)
        for _, val := range row {
            valueRow = append(valueRow, "'" + val + "'")
        }
        valueList = append(valueList, "(" + strings.Join(valueRow, ", ") + ")")
    }
    valueStr := strings.Join(valueList, ",\n")

    return "INSERT INTO `" + table + "` (" + columnStr + ") VALUES\n" + valueStr + updateStr;
}

func main() {
    columns := []string{"sku", "qty", "price"}

    items := make([]map[string]string, 0)

    var item map[string]string

    item = make(map[string]string)
    item["sku"] = "SKU-ABC-1"
    item["qty"] = "1"
    item["price"] = "11.11"
    items = append(items, item)

    item = make(map[string]string)
    item["sku"] = "SKU-ABC-2"
    item["qty"] = "2"
    item["price"] = "22.22"
    items = append(items, item)

    item = make(map[string]string)
    item["sku"] = "SKU-ABC-3"
    item["qty"] = "3"
    item["price"] = "33.33"
    items = append(items, item)

    fmt.Println(InsertSql("mytable", columns, items));
}
