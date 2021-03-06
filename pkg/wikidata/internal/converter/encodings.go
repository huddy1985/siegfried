// Copyright 2020 Ross Spencer, Richard Lehane. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

// Registration of different file-format signature sequence encodings
// that we might discover e.g. from sources such as Wikidata.

package converter

import (
	"strings"
)

const (
	UnknownEncoding = iota // UnknownEncoding provides us with a default to work with.
	HexEncoding            // HexEncoding describes magic numbers written in plain-hexadecimal.
	PronomEncoding         // PronomEncoding describe PRONOM based file format signatures.
	PerlEncoding           // PerlEncoding describe PERL regular expression encoded signatures.
	ASCIIEncoding          // ASCIIEncoding encoded patterns are those written entirely in plain ASCII.
	GUIDEncoding           // GUIDEncoding are globally unique identifiers.
)

const hexadecimal = "hexadecimal"
const guid = "globally unique identifier"
const ascii = "ascii"
const perl = "perl compatible regular expressions 2"
const pronom = "pronom internal signature"

// LookupEncoding will return a best-guess encoding type for a supplied
// encoding string.
func LookupEncoding(encoding string) int {
	encoding = strings.ToLower(encoding)
	switch encoding {
	case hexadecimal:
		return HexEncoding
	case pronom:
		return PronomEncoding
	case perl:
		return PerlEncoding
	case ascii:
		return ASCIIEncoding
	case guid:
		return GUIDEncoding
	}
	return UnknownEncoding
}

// ReverseEncoding can provide a human readable string for us if we
// ever need it, e.g. if we need to debug this module.
func ReverseEncoding(encoding int) string {
	switch encoding {
	case HexEncoding:
		return hexadecimal
	case PronomEncoding:
		return pronom
	case PerlEncoding:
		return perl
	case ASCIIEncoding:
		return ascii
	case GUIDEncoding:
		return guid
	}
	return "Unknown encoding"
}
