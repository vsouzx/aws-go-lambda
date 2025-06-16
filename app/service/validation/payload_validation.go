package validation

import (
	"encoding/json"

	"github.com/go-playground/validator"
)

type PayloadValidation struct{}

func (pv *PayloadValidation) ValidatePayload(body string, dto any) error {
	if err := json.Unmarshal([]byte(body), &dto); err != nil {
		return err
	}

	if err := validator.New().Struct(dto); err != nil {
		return err
	}

	return nil
}
