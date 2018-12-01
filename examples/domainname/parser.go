package main

import (
	"github.com/shivamMg/rd"
)

const grammar = `
	domain      = subdomain | " "
	subdomain   = label { "." label }
	label       = letter [ [ ldh-str ] let-dig ]
	ldh-str     = let-dig-hyp | let-dig-hyp ldh-str
	let-dig-hyp = let-dig | "-"
	let-dig     = letter | digit

	letter = any one of the 52 alphabetic characters A through Z in
	         upper case and a through z in lower case
	digit  = any one of the ten digits 0 through 9
`

var (
	b            *rd.Builder
	enter        = b.Enter
	exit         = b.Exit
	match        = b.Match
	next         = b.Next
	add          = b.Add
	checkOrNotOK = b.CheckOrNotOK
)

func domain() (ok bool) {
	enter("domain")
	defer exit(&ok)

	return subdomain() || match(Whitespace)
}

func subdomain() (ok bool) {
	enter("subdomain")
	defer exit(&ok)

	for label() {
		if match(Dot) {
			continue
		}
		return true
	}
	return false
}

// label = letter [ [ ldh-str ] let-dig ]
// =>
// label = letter | letter let-dig | letter ldh-str let-dig
func label() (ok bool) {
	enter("label")
	defer exit(&ok)

	if checkOrNotOK(Dot, 2) {
		return letter()
	}
	if checkOrNotOK(Dot, 3) {
		return letter() && letdig()
	}
	return letter() && ldhstr() && letdig()
}

func ldhstr() (ok bool) {
	enter("ldhstr")
	defer exit(&ok)

	if !letdighyp() {
		return false
	}
	if checkOrNotOK(Dot, 2) {
		return true
	}
	return ldhstr()
}

func letdighyp() (ok bool) {
	enter("letdighyp")
	defer exit(&ok)

	return letdig() || match(Hyphen)
}

func letdig() (ok bool) {
	enter("letdig")
	defer exit(&ok)

	return letter() || digit()
}

func letter() (ok bool) {
	enter("letter")
	defer exit(&ok)

	token, ok := next()
	if !ok {
		return false
	}
	if rtoken, ok := token.(rune); ok && isLetter(rtoken) {
		add(string(rtoken))
		return true
	}
	return false
}

func digit() (ok bool) {
	enter("digit")
	defer exit(&ok)

	token, ok := next()
	if !ok {
		return false
	}
	if rtoken, ok := token.(rune); ok && isDigit(rtoken) {
		add(string(rtoken))
		return true
	}
	return false
}
