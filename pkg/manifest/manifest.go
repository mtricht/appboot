package manifest

// Entry represents a file in the manifest
type Entry struct {
	File     string `json:"file"`
	Checksum string `json:"checksum"`
	URL      string `json:"url"`
	Size     int64  `json:"size"`
}
