package post

type Post struct {
	Id      int64
	Comment string `orm:"type(text)"`
}
