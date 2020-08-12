package strsim

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Compare_Special(t *testing.T) {

	for _, v := range []testCase{
		{arg1: "", arg2: "", got: 1},
		{arg1: "1", arg2: "", got: 0},
		{arg1: "", arg2: "1", got: 0},
	} {
		for _, o := range []Option{
			Default(),
			Jaro(),
			DiceCoefficient(),
			Hamming(),
		} {
			sim := Compare(v.arg1, v.arg2, o)
			assert.Equal(t, v.got, sim)
		}
	}
}

type bestTest struct {
	best []string
	key  string
	need string
}

func Test_FindBestMatchOne(t *testing.T) {
	for _, d := range []bestTest{
		{best: []string{"朝辞白帝彩云间", "千里江陵一日还", "两岸猿声啼不住", "轻舟已过万重山"}, key: "千里还", need: "千里江陵一日还"},
	} {
		for _, o := range []Option{
			DiceCoefficient(1),
			Jaro(),
			Default(),
		} {
			m := FindBestMatchOne(d.key, d.best, o)
			assert.Equal(t, m.S, d.need)
		}
	}
}
