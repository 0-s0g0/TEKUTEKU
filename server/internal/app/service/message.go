package service

import (
	"context"
	"math/rand"
	"time"

	"github.com/0-s0g0/TEKUTEKU/server/internal/domain/entity"
	"github.com/0-s0g0/TEKUTEKU/server/internal/domain/repository"
	"github.com/0-s0g0/TEKUTEKU/server/pkg/uuid"
)

type IMessageService interface {
	GetAll(ctx context.Context) ([]entity.Message, error)
	GetByID(ctx context.Context, id string) (*entity.Message, error)
	GetByTimeRange(ctx context.Context, from, to time.Time) ([]entity.Message, error)
	Create(ctx context.Context, message entity.Message) (*entity.Message, error)
	GiveLike(ctx context.Context, id string) error
}

type MessageService struct {
	mr repository.IMessageRepository
}

func NewMessageService(mr repository.IMessageRepository) IMessageService {
	return &MessageService{
		mr: mr,
	}
}

// Create implements IMessageService.
func (m *MessageService) Create(ctx context.Context, message entity.Message) (*entity.Message, error) {
	mess := entity.Message{
		ID:        uuid.New(),
		School:    message.School,
		Message:   message.Message,
		X:         rand.Intn(15)*25 + 100,
		Y:         rand.Intn(15)*25 + 100,
		FloatTime: float32(rand.Intn(10))*0.2 + 5.0,
		CreatedAt: time.Now(),
		ParentID:  message.ParentID,
	}
	created, err := m.mr.Create(ctx, mess)
	if err != nil {
		return nil, err
	}
	return created, nil
}

// GetAll implements IMessageService.
func (m *MessageService) GetAll(ctx context.Context) ([]entity.Message, error) {
	messages, err := m.mr.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	collection := make(map[string][]entity.Message)
	for _, m := range messages {
		if !m.ParentID.Valid {
			continue
		}
		if _, ok := collection[m.ParentID.Value]; !ok {
			collection[m.ParentID.Value] = make([]entity.Message, 0, 10)
		}
		collection[m.ParentID.Value] = append(collection[m.ParentID.Value], m)
	}
	mess := make([]entity.Message, 0, len(messages))
	for _, m := range messages {
		n := entity.Message{
			ID:        m.ID,
			School:    m.School,
			Message:   m.Message,
			Likes:     m.Likes,
			X:         m.X,
			Y:         m.Y,
			FloatTime: m.FloatTime,
			CreatedAt: m.CreatedAt,
			ParentID:  m.ParentID,
		}
		for k, v := range collection {
			if m.ID == k {
				n.Reply = v
			}
		}
		mess = append(mess, n)
	}
	return mess, nil
}

// GetByID implements IMessageService.
func (m *MessageService) GetByID(ctx context.Context, id string) (*entity.Message, error) {
	return nil, nil
}

// GetByTimeRange implements IMessageService.
func (m *MessageService) GetByTimeRange(ctx context.Context, from time.Time, to time.Time) ([]entity.Message, error) {
	return nil, nil
}

// GiveLike implements IMessageService.
func (m *MessageService) GiveLike(ctx context.Context, id string) error {
	if err := m.mr.GiveLike(ctx, id); err != nil {
		return err
	}
	return nil
}
