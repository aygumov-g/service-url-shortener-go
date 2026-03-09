package shortcode

import (
	"github.com/speps/go-hashids/v2"
)

type encoder struct {
	h *hashids.HashID
}

func NewEncoder(alphabet string, secret string) (*encoder, error) {
	data := hashids.NewData()

	data.Salt = secret
	data.Alphabet = alphabet
	data.MinLength = 8

	h, err := hashids.NewWithData(data)
	if err != nil {
		return nil, err
	}

	return &encoder{h: h}, nil
}

func (e *encoder) Encode(id int64) (string, error) {
	return e.h.EncodeInt64([]int64{id})
}

func (e *encoder) Decode(code string) (int64, error) {
	res, err := e.h.DecodeInt64WithError(code)
	if err != nil {
		return 0, err
	}

	if len(res) == 0 {
		return 0, err
	}

	return res[0], nil
}
