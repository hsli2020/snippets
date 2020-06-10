package gostrutils

import "unicode"

// CamelCaseToUnderscore converts CamelCase to camel_case
func CamelCaseToUnderscore(str string) string {
	var result string
	for idx, ch := range str {
		// first letter will just be lowered
		if idx == 0 {
			result = string(unicode.ToLower(ch))
			continue
		}

		// anywhere else
		if unicode.IsUpper(ch) {
			result = result + "_" + string(unicode.ToLower(ch))
			continue
		}

		// nothing to see here, just accept it
		result += string(ch)
	}

	return result
}

// CamelCaseToJavascriptCase convert CamelCase to camelCase
func CamelCaseToJavascriptCase(str string) string {
	var result string
	for idx, ch := range str {
		if idx == 0 {
			result = string(unicode.ToLower(ch))
			continue
		}

		result += string(ch)
	}
	return result
}

package gostrutils

import (
	"errors"
	"strings"
	"unicode/utf8"
)

// IsEmptyChars return true if after cleaning given chars the string is empty
func IsEmptyChars(s string, chars []rune) bool {
	toMapSet := make(map[rune]bool)
	for _, ch := range chars {
		toMapSet[ch] = true
	}

	result := strings.TrimFunc(s, func(ch rune) bool {
		_, found := toMapSet[ch]
		return found
	})

	return result == ""
}

// IsEmpty returns true if a string with whitespace only was provided or an
// empty string
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

package gostrutils

import "database/sql"

// ToSQLString convert string to sql.NullString
func ToSQLString(str string) sql.NullString {
	return sql.NullString{Valid: str != "", String: str}
}

// ToSQLStringNotNullable convert string to sql.NullString not nullable
func ToSQLStringNotNullable(str string) sql.NullString {
	return sql.NullString{Valid: true, String: str}
}

package gostrutils

import (
	"math"
	"unicode/utf8"
)

// SmsCalculateSmsFragments calculate based on a string how many SMS messages will it take to deliver a massage
func SmsCalculateSmsFragments(message string) uint16 {
	mbLen := utf8.RuneCountInString(message)
	byeLen := len(message)
	maxMessageLength := 160
	maxMultiMessageLength := 153

	if mbLen != byeLen {
		maxMessageLength = 70
		maxMultiMessageLength = 67
	}

	if mbLen > maxMessageLength {
		return uint16(math.Ceil(float64(mbLen) / float64(maxMultiMessageLength)))
	}

	return 1
}

package gostrutils

// GetStringIndexInSlice returns the index of the slice if string was found or -1 if not
func GetStringIndexInSlice(list []string, needle string) int {
	f := func(i int) bool {
		return list[i] == needle
	}

	return SliceIndex(len(list), f)
}

// IsStringInSlice looks for a string inside a slice and return true if it exists
func IsStringInSlice(list []string, needle string) bool {
	return GetStringIndexInSlice(list, needle) > -1
}

package gostrutils

// DefaultEllipse is the default char for marking abbreviation
const DefaultEllipse = "…"

// Abbreviate takes 'str' and based on length returns an abbreviate string
// with marks that represents middle of words.
//
// str - The original string
// startAt - Where to start to abbreviate inside a string
// maxLen - The maximum length for a string. It must be at least 4 chatrs
//
//
func Abbreviate(str string, startAt, maxLen int, abbrvChar string) string {
	if str == "" {
		return ""
	}

	if maxLen == 0 {
		return ""
	}

	if abbrvChar == "" {
		abbrvChar = DefaultEllipse
	}

	abbrvCharLen := len(abbrvChar)
	length := len(str)
	if length <= 4 {
		return str
	}

	if maxLen <= abbrvCharLen {
		return str[0:maxLen]
	}

	if length <= maxLen {
		return str
	}

	if startAt > length {
		startAt = length
	}

	if length-startAt < (maxLen - abbrvCharLen) {
		startAt = length - (maxLen - abbrvCharLen)
	}

	if startAt <= 4 {
		return str[0:maxLen-abbrvCharLen-1] + abbrvChar + str[length-1:]
	}
	if (startAt + maxLen - abbrvCharLen) < length {
		abrevStr := Abbreviate(str[startAt:length], (maxLen - abbrvCharLen + 1), maxLen-abbrvCharLen-1, abbrvChar)
		return str[0:1] + abbrvChar + abrevStr
	}
	return str[0:1] + abbrvChar + str[(length-(maxLen-abbrvCharLen)+1):length]
}

package gostrutils

import (
	"regexp"
	"strconv"
	"strings"
)

// StrToInt64 convert a string to int64
func StrToInt64(str string, def int64) int64 {
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return def
	}
	return result
}

// StrToUInt64 convert string to uint64
func StrToUInt64(str string, def uint64) uint64 {
	result, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return def
	}
	return result
}

// UInt64Join takes a list of uint64 and join them with a separator
// based on https://play.golang.org/p/KpdkVS1B4s
func UInt64Join(list []uint64, sep string) string {
	length := len(list)

	if length == 0 {
		return ""
	}

	if length == 1 {
		return strconv.FormatUint(list[0], 10)
	}

	s := ""

	for i, item := range list {
		s += strconv.FormatUint(item, 10)

		if i < length-1 {
			s += sep
		}
	}

	return s
}

// Int64Join takes a list of int64 and join them with a separator
// based on https://play.golang.org/p/KpdkVS1B4s
func Int64Join(list []int64, sep string) string {
	length := len(list)

	if length == 0 {
		return ""
	}

	if length == 1 {
		return strconv.FormatInt(list[0], 10)
	}

	s := ""

	for i, item := range list {
		s += strconv.FormatInt(item, 10)

		if i < length-1 {
			s += sep
		}
	}

	return s
}

// Uin64Split get a string with separator and convert it to a slice of uint64
func Uin64Split(data, sep string) []uint64 {
	fields := strings.Split(data, sep)
	result := make([]uint64, len(fields))

	for i, elem := range fields {
		result[i], _ = strconv.ParseUint(elem, 10, 64)
	}

	return result
}

// In64Split get a string with separator and convert it to a slice of int64
func In64Split(data, sep string) []int64 {
	fields := strings.Split(data, sep)
	result := make([]int64, len(fields))

	for i, elem := range fields {
		result[i], _ = strconv.ParseInt(elem, 10, 64)
	}

	return result
}

// ToFloat32Default convert string to float32 without errors. If error returns, defaultValue is set instead.
func ToFloat32Default(field string, defaultValue float32) float32 {
	result, err := strconv.ParseFloat(field, 32)
	if err != nil {
		return defaultValue
	}

	return float32(result)
}

// ToFloat32 convert string to float32 without errors!
func ToFloat32(field string) float32 {
	return ToFloat32Default(field, 0.0)
}

// ToFloat6Default convert string to float64 without errors. If error returns, default is set instead.
func ToFloat6Default(field string, defaultValue float64) float64 {
	result, err := strconv.ParseFloat(field, 64)
	if err != nil {
		return defaultValue
	}

	return result
}

// ToFloat64 convert string to float64 without errors!
func ToFloat64(field string) float64 {
	return ToFloat6Default(field, 0.0)
}

// IsUInteger returns true if a string is unsigned integer
func IsUInteger(txt string) bool {
	re := regexp.MustCompile("^[0-9]+$")
	return re.MatchString(txt)
}

// IsInteger returns true if a string is an integer
func IsInteger(txt string) bool {
	re := regexp.MustCompile("^-?[0-9]+$")
	return re.MatchString(txt)
}

// IsUFloat returns true if a string is unsigned floating point
func IsUFloat(txt string) bool {
	re := regexp.MustCompile(`^[0-9]+\.[0-9]+$`)
	return re.MatchString(txt)
}

// IsFloat returns true if a given text is a floating point
func IsFloat(txt string) bool {
	re := regexp.MustCompile(`^-?[0-9]+\.[0-9]+$`)
	return re.MatchString(txt)
}

// IsUNumber returns true if a given string is unsigned integer or float
func IsUNumber(txt string) bool {
	return IsUInteger(txt) || IsUFloat(txt)
}

// IsNumber returns true if a given string is integer or float
func IsNumber(txt string) bool {
	return IsInteger(txt) || IsFloat(txt)
}

package gostrutils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unicode/utf16"
	"unicode/utf8"
)

// BOM Types
const (
	BOMNone = iota
	BOMBE
	BOMLE
)

// EncodeUTF16 get a utf8 string and translate it into a slice of bytes
func EncodeUTF16(s string, bigEndian, addBom bool) []byte {
	r := []rune(s)
	iresult := utf16.Encode(r)
	bytes := []byte{}
	if addBom {
		if bigEndian {
			bytes = append(bytes, 254, 255)
		} else {
			bytes = append(bytes, 255, 254)
		}
	}
	for _, i := range iresult {
		temp := make([]byte, 2)
		if bigEndian {
			binary.BigEndian.PutUint16(temp, i)
		} else {
			binary.LittleEndian.PutUint16(temp, i)
		}
		bytes = append(bytes, temp...)
	}
	return bytes
}

// DecodeUTF16 get a slice of bytes and decode it to UTF-8
func DecodeUTF16(b []byte) (string, error) {

	if len(b)%2 != 0 {
		return "", fmt.Errorf("Must have even length byte slice")
	}

	bom := UTF16Bom(b)
	if bom < 0 {
		return "", fmt.Errorf("Buffer is too small")
	}

	u16s := make([]uint16, 1)
	ret := &bytes.Buffer{}
	b8buf := make([]byte, 4)
	lb := len(b)

	// if there is BOM, we start at 2, otherwise at the beginning
	start := 2
	if bom == BOMNone {
		start = 0
	}
	for i := start; i < lb; i += 2 {
		//assuming bom is big endian if 0 returned
		if bom == BOMNone || bom == BOMBE {
			u16s[0] = uint16(b[i+1]) + (uint16(b[i]) << 8)
		} else if bom == BOMLE {
			u16s[0] = uint16(b[i]) + (uint16(b[i+1]) << 8)
		}
		r := utf16.Decode(u16s)
		n := utf8.EncodeRune(b8buf, r[0])
		_, err := ret.Write([]byte(string(b8buf[:n])))
		if err != nil {
			return "", err
		}
	}

	return ret.String(), nil
}

// UTF16Bom returns 0 for no BOM, 1 for Big Endian and 2 for little endian
// it will return -1 if b is too small for having BOM
func UTF16Bom(b []byte) int8 {
	if len(b) < 2 {
		return -1
	}

	if b[0] == 0xFE && b[1] == 0xFF {
		return BOMBE
	}

	if b[0] == 0xFF && b[1] == 0xFE {
		return BOMLE
	}

	return BOMNone
}

// HexToUTF16Runes takes a hex based string and converts it into UTF16 runes
//
// Such string looks like:
//
//    "\x00H\x00e\x00l\x00l\x00o\x00 \x00W\x00o\x00r\x00l\x00d"
//    "\x05\xe9\x05\xdc\x05\xd5\x05\xdd\x00 \x05\xe2\x05\xd5\x05\xdc\x05\xdd"
//
// Result of the first string is (big endian):
//   [U+0048 'H' U+0065 'e' U+006C 'l' U+006C 'l' U+006F 'o' U+0020 ' ' U+0057 'W' U+006F 'o' U+0072 'r' U+006C 'l' U+0064 'd'] Hello World
//
// Result of the second string is (big endian):
//   [U+05E9 'ש' U+05DC 'ל' U+05D5 'ו' U+05DD 'ם' U+0020 ' ' U+05E2 'ע' U+05D5 'ו' U+05DC 'ל' U+05DD 'ם'] שלום עולם
func HexToUTF16Runes(s string, bigEndian bool) []rune {
	var chars []byte
	var position int
	length := len(s)

	// extract bytes from string
	for {
		if position >= length {
			break
		}
		if s[position] == '\\' {
			position++
			chars = append(chars, s[position+1], s[position+2])
			position += 2
			continue
		}

		chars = append(chars, s[position])
		position++
		continue
	}

	var runes []rune
	position = 0
	length = len(chars)

	// convert bytes two runes
	for {
		if position >= length-1 {
			break
		}
		aByte := []byte{chars[position], chars[position+1]}
		var aRune uint16
		if bigEndian {
			aRune = binary.BigEndian.Uint16(aByte)
		} else {
			aRune = binary.LittleEndian.Uint16(aByte)
		}
		unicodeChar := utf16.Decode([]uint16{aRune})
		runes = append(runes, unicodeChar...)
		position += 2
	}

	return runes
}

