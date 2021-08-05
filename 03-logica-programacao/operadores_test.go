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

func TestStatusSaque(t *testing.T) {
	type args struct {
		saldo int
		saque int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "SALDO_SUFICIENTE",
			args: args{100, 50},
			want: "SALDO_SUFICIENTE",
		},
		{
			name: "SALDO_SUFICIENTE_IRA_ZERAR_O_SALDO",
			args: args{100, 100},
			want: "SALDO_SUFICIENTE_IRA_ZERAR_O_SALDO",
		},
		{
			name: "SEM_SALDO",
			args: args{100, 200},
			want: "SEM_SALDO",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StatusSaque(tt.args.saldo, tt.args.saque); got != tt.want {
				t.Errorf("StatusSaque() = %v, want %v", got, tt.want)
			}
		})
	}
}