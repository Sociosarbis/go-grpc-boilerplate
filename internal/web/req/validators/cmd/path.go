package cmd

type PathValue struct {
	Path      *string `json:"path,omitempty"`
	Is_folder *bool   `json:"is_folder,omitempty"`
}
