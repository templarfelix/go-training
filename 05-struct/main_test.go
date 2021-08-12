package main

import (
	"reflect"
	"testing"
)

func Test_criarAnimal(t *testing.T) {
	type args struct {
		nome string
		peso float64
	}
	tests := []struct {
		name string
		args args
		want animal
	}{
		// TODO: Add test cases.
		{
			name: "criar animal fernando",
			args: args{nome: "fernando", peso: 92},
			want: animal{
				nome: "fernando",
				peso: 92,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := criarAnimal(tt.args.nome, tt.args.peso); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("criarAnimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_criarCliente(t *testing.T) {
	type args struct {
		nome   string
		peso   float64
		animal animal
	}
	tests := []struct {
		name string
		args args
		want cliente
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := criarCliente(tt.args.nome, tt.args.peso, tt.args.animal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("criarCliente() = %v, want %v", got, tt.want)
			}
		})
	}
}
