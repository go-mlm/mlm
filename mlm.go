package mlm

import (
	"encoding/json"
	"fmt"

	"github.com/zenazn/money"
	"github.com/zenazn/money/currency"
)

type (
	Mlm struct {
		Name string
	}
)

func NewMlm(name string) *Mlm {
	return &Mlm{
		Name: name,
	}
}

func (mlm *Mlm) ToJson() ([]byte, error) {
	buf, err := json.Marshal(mlm)
	return buf, err
}

func (mlm *Mlm) R() {
	m1 := money.FromMinorUnits(123, currency.USD)
	fmt.Printf("%#+v\n", m1)
}
