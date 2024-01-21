package api

import (
	"context"
	"encoding/json"
	kafka "github.com/segmentio/kafka-go"
)

type Reader[T any] struct {
	r       *kafka.Reader
	onError func(item T)
}

type Writer[T any] struct {
	w *kafka.Writer
}

func NewReader[T any](addr, topic, group string, onError func(T)) (Reader[T], func() error) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{addr},
		GroupID: group,
		Topic:   topic,
	})
	return Reader[T]{r: r, onError: onError}, r.Close
}

func (r Reader[T]) Read(handler func(items T) error) error {
	for {
		message, err := r.r.FetchMessage(context.TODO())
		if err != nil {
			return err
		}
		var t T
		json.Unmarshal(message.Value, &t)

		err = handler(t)
		if err != nil {
			r.onError(t)
		}

		r.r.CommitMessages(context.TODO(), message)
	}
}

func NewWriter[T any](addr, topic string) (Writer[T], func() error) {
	w := &kafka.Writer{
		Addr:     kafka.TCP(addr),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	return Writer[T]{w: w}, w.Close
}

func (w *Writer[T]) WriteBatch(ctx context.Context, items ...T) error {
	messages := make([]kafka.Message, len(items))
	for i, item := range items {
		b, _ := json.Marshal(item) // using a naive approach for serialization
		messages[i] = kafka.Message{
			Value: b,
		}
	}
	return w.w.WriteMessages(ctx, messages...)
}
