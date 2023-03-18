package cursor

type CommentRequest struct {
	ToComment      string `json:"toComment"`
	Filename       string `json:"filename"`
	CachedComments struct {
	} `json:"cachedComments"`
}
