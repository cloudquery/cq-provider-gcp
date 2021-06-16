package client

import "time"

func ParseISODate(d string) (*time.Time, error) {
	if d == "" {
		return nil, nil
	}
	location, err := time.LoadLocation("UTC")
	if err != nil {
		return nil, err
	}
	date, err := time.ParseInLocation(time.RFC3339, d, location)
	if err != nil {
		return nil, err
	}
	return &date, err
}
