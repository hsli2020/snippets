package main

import ("fmt" "strings")

type CodeLines struct {
    lines []string
}

func NewCodeLines() *CodeLines {
    return &CodeLines{lines: make([]string, 0)}
}

func (c *CodeLines) Push(s string, a...interface{}) {
    c.lines = append(c.lines, fmt.Sprintf(s, a...))
}

func (c *CodeLines) ToString() string {
    return strings.Join(c.lines, "\n")
}

func main() {
    strfmt := fmt.Sprintf
    lines := NewCodeLines()
    lines.Push("<Request>")
    lines.Push(strfmt("<Weight>%d</Weight>", 23))
    lines.Push("<Height>%d</Height>", 45)
    lines.Push("</Request>")
    fmt.Println(lines.ToString())
}
// ------------------------------------------------------------
package main

import ("fmt" "strings" "time")

func FmtDateTime(format string, t time.Time) string {
    // Y-m-d H:i:s => 2006-01-02 15:04:05
    r := strings.NewReplacer(
        "Y", "2006",
        "m", "01",
        "d", "02",
        "H", "15",
        "i", "04",
        "s", "05",
    )
    format = r.Replace(format)
    return t.Format(format)
}

func main() {
    fmt.Println(FmtDateTime("Y-m-d H:i:s", time.Now()))
}
// ------------------------------------------------------------
package main

import ("fmt" "os" "error")

func main() {
    // To check if a file doesn't exist, 

    if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
      // path/to/whatever does not exist
    }

    // In the above example we are not checking if err != nil because os.IsNotExist(nil) == false.

    // To check if a file exists

    if _, err := os.Stat("/path/to/whatever"); err == nil {
      // path/to/whatever exists
    }
}

// Exists reports whether the named file or directory exists.

// this code returns true, even if the file does not exist, for example 
// when Stat() returns permission denied.
func Exists(name string) bool {
    if _, err := os.Stat(name); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}

// good
func Exists(name string) bool {
    _, err := os.Stat(name)
    return !os.IsNotExist(err)
}

// better
func Exists(name string) (bool, error) {
    err := os.Stat(name)
    if os.IsNotExist(err) {
        return false, nil
    }
    return err != nil, err
}
// ------------------------------------------------------------
// array_keys() in php
func Keys(v interface{}) ([]string, error) {
    rv := reflect.ValueOf(v)
    if rv.Kind() != reflect.Map {
        return nil, errors.New("not a map")
    }
    t := rv.Type()
    if t.Key().Kind() != reflect.String {
        return nil, errors.New("not string key")
    }
    var result []string
    for _, kv := range rv.MapKeys() {
        result = append(result, kv.String())
    }
    return result, nil
}
// ------------------------------------------------------------
package main

import("fmt" "strings")

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
//*/
    fmt.Println(InsertSql("mytable", columns, items));
}
// ------------------------------------------------------------
package main

import (
    "fmt"
    "log"
    "os"
    "time"
    "strings"
    "reflect"
    "errors"
    "path/filepath"
)

func pr(v ...interface{}) {
    //fmt.Printf("%q\n", v...)
    //fmt.Printf("%+v\n", v...)
    fmt.Printf("%#v\n", v...)
}

// archiveFiles("e:/amazon/*.xml")
func archiveFiles(pattern string) {
    files, err := filepath.Glob(pattern)
    checkError(err)
    //fmt.Println(files)

    for _, fname := range files {
        dir, file := filepath.Split(fname)
        fileTime, _ := getFileTime(fname)

        newDir := dir + "archive\\" + fileTime.Format("2006-01-02");
        os.MkdirAll(newDir, 0777)

        newFile := newDir + "\\" + file

        if time.Now().Sub(fileTime) > 15*24*time.Hour {
            //os.Rename(fname, newFile)
            fmt.Println(fname, newFile)
        }
    }
}

func checkError(err error) {
    if err != nil {
        log.Fatal(err)
        //panic(err)
    }
}

func checkErr(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}

func getFileTime(name string) (mtime time.Time, err error) {
    fi, err := os.Stat(name)
    if err != nil {
        return
    }
    mtime = fi.ModTime()
    return
}

// fmt.Println(getPartNum("DH-1234-ABC"))
func getPartNum(sku string) string {
    const sep = "-"

    arr := strings.Split(sku, sep);
    if (len(arr) > 1) {
        arr = arr[1:]
    }
    return strings.Join(arr, sep)
}
// ------------------------------------------------------------
// Inspired by http://talks.golang.org/2014/go4java.slide#23

package main

import "fmt"
//import "github.com/davecgh/go-spew/spew"

type Person struct {
	name string
}

func (p *Person) Name() string {
	if p == nil {
		return "Anonymous"
	} else {
		return p.name
	}
}

func NewPerson(name string) *Person {
	var p = new(Person)
	p.name = name
	return p
}

func main() {
	var p1, p2 *Person

	p1 = NewPerson("Nemo")

	fmt.Println(p1.Name(), p2, p2.Name())  // p2==nil
}
// ------------------------------------------------------------
func check(err error) {
    if err != nil {
        panic(err)
    }
}

func ensureDir(dir string) {
    err := os.MkdirAll(dir, 0755)
    check(err)
}

func copyFile(src, dst string) {
    dat, err := ioutil.ReadFile(src)
    check(err)
    err = ioutil.WriteFile(dst, dat, 0644)
    check(err)
}

func sha1Sum(s string) string {
    h := sha1.New()
    h.Write([]byte(s))
    b := h.Sum(nil)
    return fmt.Sprintf("%x", b)
}

func mustReadFile(path string) string {
    bytes, err := ioutil.ReadFile(path)
    check(err)
    return string(bytes)
}

func debug(msg string) {
    if os.Getenv("DEBUG") == "1" {
        fmt.Fprintln(os.Stderr, msg)
    }
}

func pipe(bin string, arg []string, src string) []byte {
    cmd := exec.Command(bin, arg...)
    in, err := cmd.StdinPipe();         check(err)
    out, err := cmd.StdoutPipe();       check(err)
    err = cmd.Start();                  check(err)
    _, err = in.Write([]byte(src));     check(err)
    err = in.Close();                   check(err)
    bytes, err := ioutil.ReadAll(out);  check(err)
    err = cmd.Wait();                   check(err)
    return bytes
}
