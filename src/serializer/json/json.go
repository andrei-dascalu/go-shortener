package json

import (
	"github.com/andrei-dascalu/go-shortener/src/shortener"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

type Redirect struct{}

func (r *Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}

	return redirect, nil
}

func (r *Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	data, err := json.Marshal(input)

	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}

	return data, nil
}
