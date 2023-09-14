package _907

import "testing"

func initAddressMap() map[string][]string {
	return map[string][]string{
		"Avenida":  []string{"Av.", "Av"},
		"Rua":      []string{"R.", "R"},
		"Estrada":  []string{"Estr."},
		"Alameda":  []string{"Al"},
		"Travessa": []string{"Tv.", "Tv"},
		"Rodovia":  []string{"Rod."},
	}
}

func Test_isSameAddress(t *testing.T) {
	type args struct {
		m  map[string][]string
		s1 string
		s2 string
	}
	m := initAddressMap()
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case1",
			args{m,
				"Avenida Jose Vasconcelos 345, Zona Santa Engracia, 66267 San Pedro Garza Garcia, N.L.",
				"Av. Jose Vasconcelos 345",
			},
			true,
		},
		{
			"case1",
			args{m,
				"Av Jose Vasconcelos 345, Zona Santa Engracia, 66267 San Pedro Garza Garcia, N.L.",
				"Avenida Jose Vasconcelos 345",
			},
			true,
		},
		{
			"case1",
			args{m,
				"Av Jose Vasconcelos 345, Zona Santa Engracia, 66267 San Pedro Garza Garcia, N.L.",
				"av. jose vasconcelos 345",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameAddress(tt.args.m, tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isSameAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isMatchStr(t *testing.T) {
	type args struct {
		m  map[string][]string
		s1 string
		s2 string
	}
	m := initAddressMap()
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case1",
			args{m,
				"Avenida",
				"Av.",
			},
			true,
		},
		{
			"case2",
			args{m,
				"Av",
				"Avenida",
			},
			true,
		},
		{
			"case3",
			args{m,
				"Avenida",
				"Av.",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatchStr(tt.args.m, tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isMatchStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
