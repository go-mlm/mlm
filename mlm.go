package mlm

import ()

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