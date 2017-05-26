// Code generated by gocc; DO NOT EDIT.

// This file is dual licensed under CC0 and The gonum license.
//
// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Copyright ©2017 Robin Eklind.
// This file is made available under a Creative Commons CC0 1.0
// Universal Public Domain Dedication.

package errors

import (
	"bytes"
	"fmt"

	"gonum.org/v1/gonum/graph/formats/dot/internal/token"
)

type ErrorSymbol interface {
}

type Error struct {
	Err            error
	ErrorToken     *token.Token
	ErrorSymbols   []ErrorSymbol
	ExpectedTokens []string
	StackTop       int
}

func (E *Error) String() string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "Error")
	if E.Err != nil {
		fmt.Fprintf(w, " %s\n", E.Err)
	} else {
		fmt.Fprintf(w, "\n")
	}
	fmt.Fprintf(w, "Token: type=%d, lit=%s\n", E.ErrorToken.Type, E.ErrorToken.Lit)
	fmt.Fprintf(w, "Pos: offset=%d, line=%d, column=%d\n", E.ErrorToken.Pos.Offset, E.ErrorToken.Pos.Line, E.ErrorToken.Pos.Column)
	fmt.Fprintf(w, "Expected one of: ")
	for _, sym := range E.ExpectedTokens {
		fmt.Fprintf(w, "%s ", sym)
	}
	fmt.Fprintf(w, "ErrorSymbol:\n")
	for _, sym := range E.ErrorSymbols {
		fmt.Fprintf(w, "%v\n", sym)
	}
	return w.String()
}

func (e *Error) Error() string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "Error in S%d: %s, %s", e.StackTop, token.TokMap.TokenString(e.ErrorToken), e.ErrorToken.Pos.String())
	if e.Err != nil {
		fmt.Fprintf(w, e.Err.Error())
	} else {
		fmt.Fprintf(w, ", expected one of: ")
		for _, expected := range e.ExpectedTokens {
			fmt.Fprintf(w, "%s ", expected)
		}
	}
	return w.String()
}