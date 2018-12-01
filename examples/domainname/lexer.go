package main

import (
	"bufio"
	"io"
)

const eof = rune(0)

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{bufio.NewReader(r)}
}

func (s *Scanner) read() (rune, error) {
	r, _, err := s.r.ReadRune()
	if err == io.EOF {
		return eof, nil
	}
	if err != nil {
		return eof, err
	}
	return r, nil
}

func (s *Scanner) unread() {
	s.r.UnreadRune()
}

func (s *Scanner) Scan() (token Token, lit rune, err error) {
	r, err := s.read()
	if err != nil {
		return 0, 0, err
	}
	switch r {
	case ' ':
		return Whitespace, r, nil
	case '.':
		return Dot, r, nil
	case '-':
		return Hyphen, r, nil
	case eof:
		return EOF, r, nil
	}
	if isLetter(r) {
		return Letter, r, nil
	}
	if isDigit(r) {
		return Digit, r, nil
	}
	return Illegal, r, nil
}
