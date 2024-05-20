package scraper

import (
	"strconv"
	"strings"
)

func parseStringField(field string) string {
	return sanitizeField(field)
}

func parseFloatField(field string) (float64, error) {
	if field == "-" {
		return 0, nil
	}

	field = sanitizeField(field)

	value, err := strconv.ParseFloat(field, 64)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func parseIntField(field string) (int, error) {
	if field == "-" {
		return 0, nil
	}

	field = sanitizeField(field)

	value, err := strconv.Atoi(field)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func sanitizeField(field string) (sanitized string) {
	sanitized = strings.TrimSpace(field)
	sanitized = strings.ReplaceAll(sanitized, "\n", "")
	sanitized = strings.ReplaceAll(sanitized, "\t", "")
	sanitized = strings.ReplaceAll(sanitized, "\r", "")
	sanitized = strings.ReplaceAll(sanitized, ".", "")
	sanitized = strings.ReplaceAll(sanitized, ",", ".")
	sanitized = strings.ReplaceAll(sanitized, "%", "")

	return
}
