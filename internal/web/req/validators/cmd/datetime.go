package cmd

type DateTimeConfigItem struct {
	Enable *bool   `json:"enable,omitempty"`
	Format *string `json:"format,omitempty"`
}

type DateTimeConfig struct {
	Format *string         `json:"format,omitempty" validate:"omitempty,min=1"`
	Date   *DateTimeConfig `json:"date,omitempty"`
	Time   *DateTimeConfig `json:"time,omitempty"`
}

type DateTimeValue struct {
	Date *int64 `json:"date,omitempty"`
	Time *int64 `json:"time:omitempty"`
}
