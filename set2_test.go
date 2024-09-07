package rice

import "testing"

func TestPrintAllElements(t *testing.T) {
	type E int
	type args struct {
		s *Set[E]
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				s: &Set[E]{m: map[E]struct{}{1: {}, 2: {}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintAllElements(tt.args.s)
		})
	}
}
