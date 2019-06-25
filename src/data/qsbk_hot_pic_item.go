package data

type QsbkHotPicItem struct {
	Id                    string
	AuthorNickName        string
	AuthorGender          string
	AuthorAge             string
	AuthorImgUrl          string
	Content               string
	ThumbImgUrl           string
	StatsVoteContent      string
	StatsCommentContent   string
	StatsCommentDetailUrl string
	Md5                   string
}
type QsbkHotPicItemList struct {
	ItemList []QsbkHotPicItem
}
