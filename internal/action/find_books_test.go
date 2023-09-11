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

func TestFindBooks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	lastName := "King"
	res := []*model.Book{
		{
			Title: "It",
		},
	}

	m := dimock.NewMockBookRepository(ctrl)
	m.
		EXPECT().
		FindBooks(ctx, lastName).
		Return(res, nil)

	act := action.NewFindBooks(m)
	authors, err := act.Do(ctx, lastName)
	assert.NoError(t, err)
	assert.Equal(t, res, authors)
}
