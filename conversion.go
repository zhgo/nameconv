// Copyright 2014 The zhgo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package console

import (
	"strings"
	"unicode"
)

// for example: transfer browse_by_set to BrowseBySet
func UnderscoreToCamelcase(str string) string {
	return LowerToCamelcase(str, "_")
}

// for example: transfer BrowseBySet to browse_by_set
func CamelcaseToUnderscore(str string) string {
	return CamelcaseToLower(str, "_")
}

func LowerToCamelcase(str string, sp string) string {
	var method string
	sli := strings.Split(str, sp)
	for _, v := range sli {
		method += strings.Title(v)
	}
	return method
}

func CamelcaseToLower(str string, sp string) string {
	return strings.Join(CamelcaseToLowerSlice(str, -1), sp)
}

func CamelcaseToLowerSlice(str string, limit int) []string {
	var words []string
	l := 0
	i := 1

	for s := str; s != ""; s = s[l:] {
		l = strings.IndexFunc(s[1:], unicode.IsUpper) + 1
		if l < 1 {
			l = len(s)
		}
		words = append(words, strings.ToLower(s[:l]))

		if limit > 0 && limit >= i {
			words = append(words, strings.ToLower(s))
			break
		}

		i++
	}

	return words
}
