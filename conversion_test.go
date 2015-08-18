// Copyright 2014 The zhgo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nameconv

import (
	"testing"
)

func TestUnderscoreToCamelcase(t *testing.T) {
	path := "browse_by_set"
	method := UnderscoreToCamelcase(path, true)
	if method != "BrowseBySet" {
		t.Error("UnderscoreToCamelcase failure")
	}

	method = UnderscoreToCamelcase(path, false)
	if method != "browseBySet" {
		t.Error("UnderscoreToCamelcase failure")
	}
}

func TestCamelcaseToUnderscore(t *testing.T) {
	method := "BrowseBySet"
	path := CamelcaseToUnderscore(method)
	if path != "browse_by_set" {
		t.Errorf("CamelcaseToUnderscore failure: %#v", path)
	}
}

func TestCamelcaseToSlice(t *testing.T) {
	method := "BrowseBySet"

	sli := CamelcaseToSlice(method, true, -1)
	if len(sli) != 3 {
		t.Errorf("CamelcaseToSlice failure: %#v", sli)
	}
	if sli[0] != "browse" || sli[1] != "by" || sli[2] != "set" {
		t.Errorf("CamelcaseToSlice failure: %#v", sli)
	}

	sli = CamelcaseToSlice(method, true, 2)
	if len(sli) != 2 {
		t.Errorf("CamelcaseToSlice failure: %#v", sli)
	}
	if sli[0] != "browse" || sli[1] != "byset" {
		t.Errorf("CamelcaseToSlice failure: %#v", sli)
	}

	sli = CamelcaseToSlice(method, false, -1)
	if len(sli) != 3 {
		t.Errorf("CamelcaseToSlice failure: %#v", sli)
	}
	if sli[0] != "Browse" || sli[1] != "By" || sli[2] != "Set" {
		t.Errorf("CamelcaseToSlice failure: %#v", sli)
	}

	sli = CamelcaseToSlice(method, false, 2)
	if len(sli) != 2 {
		t.Errorf("CamelcaseToSlice failure: %#v", sli)
	}
	if sli[0] != "Browse" || sli[1] != "BySet" {
		t.Errorf("CamelcaseToSlice failure: %#v", sli)
	}
}
