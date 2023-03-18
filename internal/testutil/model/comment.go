package model

import (
	"github.com/opcotech/elemo/internal/model"
	"github.com/opcotech/elemo/internal/testutil"
)

// NewComment creates a new comment with the given user ID and text. The
// comment is not created in the database.
func NewComment(createdBy model.ID) *model.Comment {
	comment, err := model.NewComment(testutil.GenerateRandomString(10), createdBy)
	if err != nil {
		panic(err)
	}
	return comment
}
