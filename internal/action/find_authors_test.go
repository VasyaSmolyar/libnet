package action_test

import (
	"context"
	"libnet/internal/action"
	"libnet/internal/action/dimock"
	"libnet/internal/model"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFindAuthors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	title := "It"
	res := []*model.Author{
		{
			LastName: "King",
		},
	}

	m := dimock.NewMockAuthorRepository(ctrl)
	m.
		EXPECT().
		FindAuthor(ctx, title).
		Return(res, nil)

	act := action.NewFindAuthors(m)
	books, err := act.Do(ctx, title)
	assert.NoError(t, err)
	assert.Equal(t, res, books)
}
