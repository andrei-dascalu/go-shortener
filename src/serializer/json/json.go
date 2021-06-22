package json

import (
	"github.com/andrei-dascalu/go-shortener/src/shortener"
	"github.com/pkg/errors"
)

type Redirect struct{}

func (r *Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}

	if err := redirect.UnmarshalJSON(input); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}

	return redirect, nil
}

func (r *Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	data, err := input.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}

	return data, nil
}
