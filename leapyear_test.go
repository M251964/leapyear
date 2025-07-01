package main

import (
	"testing"
)

// TestLeapYearBasicCommonYear は、一般的な平年（うるう年ではない年）のテストです。
// 例えば、2001年はうるう年ではありません。
func TestLeapYearBasicCommonYear(t *testing.T) {
	year := 2001
	expected := false
	actual := leapyear(year)
	if actual != expected {
		t.Errorf("leapyear(%d): expected %v, but got %v", year, expected, actual)
	}
}

// TestLeapYearDivisibleBy4 は、4で割り切れる一般的なうるう年のテストです。
// 例えば、1996年はうるう年です。
func TestLeapYearDivisibleBy4(t *testing.T) {
	year := 1996
	expected := true
	actual := leapyear(year)
	if actual != expected {
		t.Errorf("leapyear(%d): expected %v, but got %v", year, expected, actual)
	}
}

// TestLeapYearDivisibleBy100ButNot400 は、100で割り切れるが400で割り切れない平年のテストです。
// 例えば、1900年はうるう年ではありません。
func TestLeapYearDivisibleBy100ButNot400(t *testing.T) {
	year := 1900
	expected := false
	actual := leapyear(year)
	if actual != expected {
		t.Errorf("leapyear(%d): expected %v, but got %v", year, expected, actual)
	}
}

// TestLeapYearDivisibleBy400 は、400で割り切れるうるう年のテストです。
// 例えば、2000年はうるう年です。
func TestLeapYearDivisibleBy400(t *testing.T) {
	year := 2000     // 4で割り切れる かつ 100で割り切れる かつ 400で割り切れる年
	expected := true // うるう年なのでtrueを期待

	actual := leapyear(year) // leapyear関数を呼び出す

	if actual != expected {
		t.Errorf("leapyear(%d): expected %v, but got %v", year, expected, actual)
	}
}
