package dbutil

import(
    "fmt"
    "strings"
)

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
