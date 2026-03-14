package generator

import "testing"

func TestGenerateShortCode(t *testing.T) {
	_, err := GenerateShortCode(10)
	if err != nil {
		t.Error(err)
	}
}