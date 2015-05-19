package main

import (
	"bufio"
	"io"
)

// EqualFunc ...
type EqualFunc func(s1 string, s2 string) bool

// UniqScanner ...
type UniqScanner struct {
	scanner   *bufio.Scanner
	isInitial bool
	isFinal   bool
	prev      string
	this      string
	token     string
	repeats   int
	counter   int
	equal     EqualFunc
}

// NewScanner return a new UniqScanner to read from r.
func NewScanner(r io.Reader) *UniqScanner {
	return &UniqScanner{
		scanner:   bufio.NewScanner(r),
		isInitial: true,
		isFinal:   true,
		equal: func(s1 string, s2 string) bool {
			return s1 == s2
		},
	}
}

// Scan advances the Scanner to the next token.
func (u *UniqScanner) Scan() bool {
	if u.isInitial {
		u.isInitial = false
		if u.scanner.Scan() {
			u.prev = u.scanner.Text()
			u.token = u.prev
			return true
		}
		return false
	}

	for u.scanner.Scan() {
		u.this = u.scanner.Text()
		if !u.equal(u.prev, u.this) {
			u.prev = u.this
			u.token = u.prev
			return true
		}
		u.repeats++
	}

	return false
}

// ScanCount ...
func (u *UniqScanner) ScanCount() bool {
	u.counter = 1
	if u.isInitial {
		u.isInitial = false
		if u.scanner.Scan() {
			u.prev = u.scanner.Text()
		}
	}

	for u.scanner.Scan() {
		u.this = u.scanner.Text()
		if !u.equal(u.prev, u.this) {
			u.token = u.prev
			u.prev = u.this
			u.counter += u.repeats
			u.repeats = 0
			return true
		}
		u.repeats++
	}

	if u.isFinal {
		u.isFinal = false
		u.token = u.prev
		return true
	}

	return false
}

// ScanDuplicate advances the Scanner to the next token.
func (u *UniqScanner) ScanDuplicate() bool {
	if u.isInitial {
		u.isInitial = false
		if u.scanner.Scan() {
			u.prev = u.scanner.Text()
		}
	}

	for u.scanner.Scan() {
		u.this = u.scanner.Text()
		if !u.equal(u.prev, u.this) {
			if u.repeats != 0 {
				u.token = u.prev
				u.prev = u.this
				u.repeats = 0
				return true
			}
			u.prev = u.this
			u.repeats = 0
		} else {
			u.repeats++
		}
	}

	if u.isFinal {
		u.isFinal = false
		if u.repeats != 0 {
			u.token = u.prev
			return true
		}
	}

	return false
}

// ScanUnique advances the Scanner to the next token.
func (u *UniqScanner) ScanUnique() bool {
	if u.isInitial {
		u.isInitial = false
		if u.scanner.Scan() {
			u.prev = u.scanner.Text()
		}
	}

	for u.scanner.Scan() {
		u.this = u.scanner.Text()
		if !u.equal(u.prev, u.this) {
			if u.repeats == 0 {
				u.token = u.prev
				u.prev = u.this
				return true
			}
			u.prev = u.this
			u.repeats = 0
		} else {
			u.repeats++
		}
	}

	if u.isFinal {
		u.isFinal = false
		if u.repeats == 0 {
			u.token = u.prev
			return true
		}
	}

	return false
}

// Equal sets the equal function for the UniqScanner. If called, it must be
// called before Scan/ScanDuplicate/ScanUnique.
func (u *UniqScanner) Equal(equal EqualFunc) {
	u.equal = equal
}

// Text return the most recent token generated by a call to
// Scan/ScanDuplicate/ScanUnique.
func (u *UniqScanner) Text() string {
	return u.token
}

// Count ...
func (u *UniqScanner) Count() int {
	return u.counter
}
