package validation

func ErrorMessages() map[string]string {
	return map[string]string{
		"required": "This field is required",
		"email":    "This field must be valid email",
		"min":      "Should be more than limit of characters",
		"max":      "Should be less than limit of characters",
	}
}
