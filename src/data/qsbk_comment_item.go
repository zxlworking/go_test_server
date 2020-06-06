package data

type QsbkCommentItem struct {
	Id                  string
	JokeId              string
	ArticleId           string
	CommentUserId       string
	CommentUserImg      string
	CommentUserNickName string
	CommentUserGender   string
	CommentUserAge      string
	CommentUserContent  string
	CommentType         string
}
type QsbkCommentList struct {
	ItemList []QsbkCommentItem
}
