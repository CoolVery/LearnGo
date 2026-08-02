package main

import "testing"

type TestValueLenght struct {
	in int
	out string
}

type TestValueDeleteVowels struct {
	in string
	out string
}

type TestValueGetUTFLength struct {
	in []byte
	out int
}

type TestValueGetUTF struct {
	in []byte
	out int
    outError bool
}

var testValueLenght = []TestValueLenght {
	{-5, "negative"},
	{0, "zero"},
	{6, "short"},
	{32, "long"},
	{10000, "very long"},
}

var testValueDeleteVowels = []TestValueDeleteVowels {
	{"eeeeef", "f"},
	{"aie", ""},
	{"done", "dn"},
	{"pyp", "pyp"},
}

var testValueGetUTF = []TestValueGetUTF {
	{[]byte("Hello"), 5, false},
	{[]byte("Привет"), 6, false},
    {[]byte(""), 0, false},
    {[]byte{0xD0}, 0, true},
}

func TestPrintHello(t *testing.T) {
    got := PrintHello("Igor")
    expected := "Hello, !"

    if got != expected {
    t.Fatalf(`PrintHello("Igor") = %q, want %q`, got, expected)
    }
}

func TestSum(t *testing.T) {
	got := Sum(2, 2)
	expected := 4
	
	if got != expected {
		t.Fatalf("Сумма рабоатет не правильно. Ждали %d, получили %d", expected, got )
	}
}

func TestLengh(t *testing.T) {
	for id, test := range testValueLenght {
		totalOut := Length(test.in)
		if totalOut != test.out {
			t.Errorf("%d - Функция длины работает не корректно Ждали %s, получили %s", id, test.out, totalOut )
		}
	}
}

func TestMultiply(t *testing.T) {
	got := Multiply(5, 0)
	expected := 0

	if got != expected {
		t.Fatalf("При умножении должно быть 0. Ждали %d, получили %d", expected, got )
	}
}

func TestDeleteVowels(t *testing.T) {
	for id, test := range testValueDeleteVowels {
		totalOut := DeleteVowels(test.in)
		if totalOut != test.out {
			t.Errorf("%d - Функция удаления работает не корректно Ждали %s, получили %s", id, test.out, totalOut )
		}
	}
}



func TestGetUTFLength(t *testing.T) {
	for id, test := range testValueGetUTF {
		totalOut, err := GetUTFLength(test.in)
        if totalOut != test.out {
			t.Errorf("%d - Функция подсчета не корректно Ждали %d, получили %d", id, test.out, totalOut )	
        }
        if err == nil && test.outError {
			t.Errorf("%d - Функция подсчета не вернула ошибку: ErrInvalidUTF8", id)
        }
	}
}
