package main

import (
"testing"
)

func TestOpEquals(t *testing.T) {
	result := OpEquals(5, 5)
	if result == false {
		t.Errorf("OpEquals(5, 5) = %t; want true", result)
	}
}

func TestFailOpEquals(t *testing.T) {
	result := OpEquals(5, 4)
	if result == true {
		t.Errorf("OpEquals(5, 4) = %t; want false", result)
	}
}

func TestOpNotEquals(t *testing.T) {
	result := OpNotEquals(5, 5)
	if result == true {
		t.Errorf("OpNotEquals(5, 5) = %t; want false", result)
	}
}

func TestFailOpNotEquals(t *testing.T) {
	result := OpNotEquals(5, 4)
	if result == false {
		t.Errorf("OpNotEquals(5, 4) = %t; want true", result)
	}
}

func TestOpThen(t *testing.T) {
	result := OpLess(5, 4)
	if result == true {
		t.Errorf("OpLessThen(5, 4) = %t; want false", result)
	}
}

func TestOpThen2(t *testing.T) {
	result := OpLess(5, 5)
	if result == true {
		t.Errorf("OpLessThen(5, 4) = %t; want false", result)
	}
}

func TestOpThen3(t *testing.T) {
	result := OpLess(5, 6)
	if result == false {
		t.Errorf("OpLessThen(5, 4) = %t; want true", result)
	}
}

// escreva os testes para OpGreater

// escreva os testes para OpLessOrEqual

// escreva os testes para OpGreaterOrEqual

// escreva os testes para PossoFazerUmSaque

// escreva os testes para StatusSaque
