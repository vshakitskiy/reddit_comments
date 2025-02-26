package pubsub

import (
	"fmt"

	"github.com/vshakitskiy/reddit_comments/internal/model"
)

type CommentPubSub struct {
	ps *PubSub[*model.Comment]
}

func NewCommentPubSub() *CommentPubSub {
	return &CommentPubSub{
		ps: NewPubSub[*model.Comment](),
	}
}

func (p *CommentPubSub) SubscribeToComments(postID string) Subscriber[*model.Comment] {
	return p.ps.Subscribe(fmt.Sprintf("post:%s", postID))
}

func (p *CommentPubSub) PublishComment(postID string, comment *model.Comment) {
	p.ps.Publish(fmt.Sprintf("post:%s", postID), comment)
}
