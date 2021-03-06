// Copyright 2017 The ObjectHash-Proto Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package protohash

import (
	"crypto/sha256"
	"fmt"
	"math"
)

const hashLength int = sha256.Size

const (
	// Sorted alphabetically by value.
	boolIdentifier     = `b`
	mapIdentifier      = `d`
	floatIdentifier    = `f`
	intIdentifier      = `i`
	listIdentifier     = `l`
	nilIdentifier      = `n`
	byteIdentifier     = `r`
	unicodeIndentifier = `u`
)

func hash(t string, b []byte) ([]byte, error) {
	h := sha256.New()

	if _, err := h.Write([]byte(t)); err != nil {
		return nil, err
	}

	if _, err := h.Write(b); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

func hashBool(b bool) ([]byte, error) {
	bb := []byte(`0`)
	if b {
		bb = []byte(`1`)
	}
	return hash(boolIdentifier, bb)
}

func hashBytes(bs []byte) ([]byte, error) {
	return hash(byteIdentifier, bs)
}

func hashFloat(f float64) ([]byte, error) {
	var normalizedFloat string

	switch {
	case math.IsInf(f, 1):
		normalizedFloat = "Infinity"
	case math.IsInf(f, -1):
		normalizedFloat = "-Infinity"
	case math.IsNaN(f):
		normalizedFloat = "NaN"
	default:
		var err error
		normalizedFloat, err = floatNormalize(f)
		if err != nil {
			return nil, err
		}
	}

	return hash(floatIdentifier, []byte(normalizedFloat))
}

func hashInt64(i int64) ([]byte, error) {
	return hash(intIdentifier, []byte(fmt.Sprintf("%d", i)))
}

func hashNil() ([]byte, error) {
	return hash(nilIdentifier, []byte(``))
}

func hashUint64(i uint64) ([]byte, error) {
	return hash(intIdentifier, []byte(fmt.Sprintf("%d", i)))
}

func hashUnicode(s string) ([]byte, error) {
	return hash(unicodeIndentifier, []byte(s))
}
