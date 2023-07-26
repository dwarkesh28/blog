package validation

func ErrorMessages() map[string]string {
	return map[string]string{
		"required": "the field is required",
		"email":    "the dield must have a valid address",
		"min":      "should have more than the limit",
		"max":      "should have less than the limit",
	}
}