package main

// https://stackoverflow.com/questions/12486436/how-do-i-batch-sql-statements-with-package-database-sql

import (
	"fmt"
	"strings"
)

type ExampleRowStruct struct {
	Column1 int
	Column2 int
	Column3 int
}

func BulkInsert(unsavedRows []*ExampleRowStruct) error {
	valueStrings := make([]string, 0, len(unsavedRows))
	valueArgs := make([]interface{}, 0, len(unsavedRows)*3)
	for _, post := range unsavedRows {
		valueStrings = append(valueStrings, "(?, ?, ?)")
		valueArgs = append(valueArgs, post.Column1)
		valueArgs = append(valueArgs, post.Column2)
		valueArgs = append(valueArgs, post.Column3)
	}
	stmt := fmt.Sprintf("INSERT INTO my_sample_table (column1, column2, column3) VALUES %s",
		strings.Join(valueStrings, ","))
	//_, err := db.Exec(stmt, valueArgs...)
	//return err

	fmt.Println(stmt)
	fmt.Println(valueArgs...)
	return nil
}

func main() {
	rows := []*ExampleRowStruct{
		&ExampleRowStruct{1, 11, 111},
		&ExampleRowStruct{2, 22, 222},
		&ExampleRowStruct{3, 33, 333},
	}
	BulkInsert(rows)
}

// https://stackoverflow.com/questions/21108084/how-to-insert-multiple-data-at-once

func Example2() {
	data := []map[string]string{
		{"v1": "1", "v2": "1", "v3": "1"},
		{"v1": "2", "v2": "2", "v3": "2"},
		{"v1": "3", "v2": "3", "v3": "3"},
	}
	sqlStr := "INSERT INTO test(n1, n2, n3) VALUES "
	vals := []interface{}{}

	for _, row := range data {
		sqlStr += "(?, ?, ?),"
		vals = append(vals, row["v1"], row["v2"], row["v3"])
	}
	/*
		//trim the last ,
		sqlStr = sqlStr[0:len(sqlStr)-2]

		//prepare the statement
		stmt, _ := db.Prepare(sqlStr)

		//format all vals at once
		res, _ := stmt.Exec(vals...)
	*/
	fmt.Println(sqlStr)
	fmt.Println(vals)
}

/*
func (repo *repo) CreateBalancesForAsset(ctx context.Context, wallets []*Wallet, asset *SimpleAsset) (error) {
  valueStrings := []string{}
  valueArgs := []interface{}{}
  for _, w := range wallets {
    valueStrings = append(valueStrings, "(?, ?, ?, ?)")

    valueArgs = append(valueArgs, w.Address)
    valueArgs = append(valueArgs, asset.Symbol)
    valueArgs = append(valueArgs, asset.Identify)
    valueArgs = append(valueArgs, asset.Decimal)
  }
  smt := `INSERT INTO balances(address, symbol, identify, decimal)
    VALUES %s ON CONFLICT (address, symbol) DO UPDATE SET address = excluded.address`
  smt = fmt.Sprintf(smt, strings.Join(valueStrings, ","))
  fmt.Println("smttt:", smt)
  tx := repo.db.Begin()
  err := tx.Exec(smt, valueArgs...).Error
  if err != nil {
    tx.Rollback()
    return err
  }
  return tx.Commit().Error
}
*/

// bulk/insert.go

import (
    "strconv"
    "strings"
)

type ValueExtractor = func(int) []interface{}

func Generate(tableName string, columns []string, numRows int, postgres bool, valueExtractor ValueExtractor) (string, []interface{}) {
    numCols := len(columns)
    var queryBuilder strings.Builder
    queryBuilder.WriteString("INSERT INTO ")
    queryBuilder.WriteString(tableName)
    queryBuilder.WriteString("(")
    for i, column := range columns {
        queryBuilder.WriteString("\"")
        queryBuilder.WriteString(column)
        queryBuilder.WriteString("\"")
        if i < numCols-1 {
            queryBuilder.WriteString(",")
        }
    }
    queryBuilder.WriteString(") VALUES ")
    var valueArgs []interface{}
    valueArgs = make([]interface{}, 0, numRows*numCols)
    for rowIndex := 0; rowIndex < numRows; rowIndex++ {
        queryBuilder.WriteString("(")
        for colIndex := 0; colIndex < numCols; colIndex++ {
            if postgres {
                queryBuilder.WriteString("$")
                queryBuilder.WriteString(strconv.Itoa(rowIndex*numCols + colIndex + 1))
            } else {
                queryBuilder.WriteString("?")
            }
            if colIndex < numCols-1 {
                queryBuilder.WriteString(",")
            }
        }
        queryBuilder.WriteString(")")
        if rowIndex < numRows-1 {
            queryBuilder.WriteString(",")
        }
        valueArgs = append(valueArgs, valueExtractor(rowIndex)...)
    }
    return queryBuilder.String(), valueArgs
}

// bulk/insert_test.go

import (
    "fmt"
    "strconv"
)

func valueExtractor(index int) []interface{} {
    return []interface{}{
        "trx-" + strconv.Itoa(index),
        "name-" + strconv.Itoa(index),
        index,
    }
}

func ExampleGeneratePostgres() {
    query, valueArgs := Generate("tbl_persons", []string{"transaction_id", "name", "age"}, 3, true, valueExtractor)
    fmt.Println(query)
    fmt.Println(valueArgs)
    // Output:
    // INSERT INTO tbl_persons("transaction_id","name","age") VALUES ($1,$2,$3),($4,$5,$6),($7,$8,$9)
    // [[trx-0 name-0 0] [trx-1 name-1 1] [trx-2 name-2 2]]
}

func ExampleGenerateOthers() {
    query, valueArgs := Generate("tbl_persons", []string{"transaction_id", "name", "age"}, 3, false, valueExtractor)
    fmt.Println(query)
    fmt.Println(valueArgs)
    // Output:
    // INSERT INTO tbl_persons("transaction_id","name","age") VALUES (?,?,?),(?,?,?),(?,?,?)
    // [[trx-0 name-0 0] [trx-1 name-1 1] [trx-2 name-2 2]]
}
