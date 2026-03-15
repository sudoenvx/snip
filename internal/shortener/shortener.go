package shortener

import (
	"errors"
	"fmt"

	"github.com/sudoenvx/snip/internal/generator"
	"github.com/sudoenvx/snip/internal/validator"
)

type Result struct {
	Shorten string
	Code string
}

func ShortenUrl(url string) (*Result, error) {
	isValidUrl := validator.ValidateURL(url)
	if !isValidUrl {
		return nil, errors.New("Not a valid url")
	}

	code, err := generator.GenerateShortCode(16)
	if err != nil {
		return nil, err
	}

	shortenUrl := fmt.Sprintf("http://localhost:3000/e/%s", code)

	result := &Result{
		Shorten: shortenUrl,
		Code: code,
	}

	return result, nil
}