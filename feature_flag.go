package enalog

type FeatureFlag struct {
	Name   string `json:"name"`
	UserId string `json:"user_id"`
}

type FeatureFlagRes struct {
	Variant  string `json:"variant"`
	FlagType string `json:"flag_type"`
}
