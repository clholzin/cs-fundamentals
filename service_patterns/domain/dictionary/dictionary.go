package dictionary

// Word
type Word struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Words
type Words []Word

// Result
type Result Words

// Repository
type Repository interface {
	Add(key, value string) error
	Search(key string) string
}
