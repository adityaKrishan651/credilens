package services

import (
	"errors"
	"strings"

	"credilens-backend/internal/constants"
)

func ValidateAnalyseInput(inputType string, content string) error {
	if strings.TrimSpace(content) == "" {
		return errors.New("content cannot be empty")
	}

	switch inputType {
	case constants.InputText, constants.InputURL, constants.InputDOM:
		return nil

	case constants.InputImage:
		return nil

	default:
		return errors.New("unsupported input type")
	}
}
