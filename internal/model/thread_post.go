package model

type ThreadPost struct {
	Thread *Thread `json:"thread"`
	Posts  []*Post `json:"posts"`
	Tags   []*Tag  `json:"tags"`
}
