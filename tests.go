package cursor

type TestsRequest struct {
	CachedTests      []any  `json:"cachedTests"`
	FileName         string `json:"fileName"`
	FileContents     string `json:"fileContents"`
	TestFileName     string `json:"testFileName"`
	TestFileContents string `json:"testFileContents"`
}
