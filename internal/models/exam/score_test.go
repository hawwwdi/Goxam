package exam

import "testing"

func TestCalcScore(t *testing.T) {
	cases :=[]struct {
		noq, noa, noi, forfeit int
		want string
	}{
		{10, 5, 3, 0, "50"},
		{10, 5, 3, 3, "40"},
	}
	for _, c := range cases {
		res := CalcScore(c.noq, c.noa, c.noi, c.forfeit)
		if res != c.want {
			t.Errorf("%v != %v", res, c.want)
		}
	}
}