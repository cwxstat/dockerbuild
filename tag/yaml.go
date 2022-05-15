package tag

type yaml interface {
	CommentsWithTags(string, string) (string, error)
	UnMarshal(string) error
	ImageVersion(string, string)
	NextMinor() error
	NextMajor() error
}
