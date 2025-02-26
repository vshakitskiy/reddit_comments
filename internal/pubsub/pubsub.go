package pubsub

import "sync"

type Subscriber[T any] chan T

type PubSub[T any] struct {
	mu          sync.RWMutex
	subscribers map[string][]Subscriber[T]
}

func NewPubSub[T any]() *PubSub[T] {
	return &PubSub[T]{
		subscribers: make(map[string][]Subscriber[T]),
	}
}

func (p *PubSub[T]) Subscribe(topic string) Subscriber[T] {
	ch := make(Subscriber[T], 1)

	p.mu.Lock()
	defer p.mu.Unlock()

	p.subscribers[topic] = append(p.subscribers[topic], ch)

	return ch
}

func (p *PubSub[T]) Publish(topic string, data T) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	for _, subscriber := range p.subscribers[topic] {
		subscriber <- data
	}
}

func (p *PubSub[T]) Unsubscribe(topic string, subscriber Subscriber[T]) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, s := range p.subscribers[topic] {
		if s == subscriber {
			p.subscribers[topic] = append(
				p.subscribers[topic][:i],
				p.subscribers[topic][i+1:]...,
			)
			break
		}
	}
}
