package main

import (
	"testing"
)

func TestLeapYear(t *testing.T) { // このテスト関数は名前がTestLeapYear01の場合もありますが、今回は統一します
	// まだ具体的なテストロジックは書かないので、最小限でOK
	// 例えば、以下のような簡単なアサーションを入れても良いでしょう
	expected := false
	actual := leapyear(2001) // leapyear関数を呼び出す
	if expected != actual {
		t.Errorf("Error: Expected %v, got %v", expected, actual)
	}
}
