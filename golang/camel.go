// Copyright (c) 2016 Danilo Bürger <info@danilobuerger.de>

package phrase

import (
	"bytes"
	"strings"
)

// ToCamel converts a lower case string with separators to a camel case string.
func ToCamel(s, sep string) string {
	if sep == "" {
		return strings.Title(s)
	}

	sa := strings.Split(s, sep)

	var buf bytes.Buffer
	for _, sv := range sa {
		buf.WriteString(strings.Title(sv))
	}

	return buf.String()
}
