package snow

import (
	"errors"

	calc "github.com/Ricky-fight/go_learn/2/calc"
)

type Snow map[string]func(float64, float64) (float64, error)

func init() {
	var snow Snow
	snow["add"] = calc.Add

}
func (s Snow) Exec(cmd string, x, y float64) (float64, error) {
	for k, f := range s {
		if cmd == k {
			if f == nil {
				return 0.0, errors.New("function " + cmd + " doesn't exist")
			}
			rst, err := f(x, y)
			if err != nil {
				return 0.0, err
			}
			return rst, nil
		}
	}
	return 0.0, errors.New("command " + cmd + " doesn't exist")
}
