package Request

type CreateCommentJSON struct {
	IdeaID  int    `json:"IdeaID"`
	Content string `json:"Content"`
	UserID  int    `json:"UserID"`
}

type DeleteCommentJSON struct {
	CommentID int `json:"CommentID"`
}

type GetCommentJSON struct {
	IdeaID int `json:"IdeaID"`
}
