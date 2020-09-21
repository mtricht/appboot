package manifest

// File represents a file in the manifest
type File struct {
	File     string `json:"file"`
	Checksum string `json:"checksum"`
	URL      string `json:"url"`
	Size     int64  `json:"size"`
}
