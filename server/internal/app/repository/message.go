package repository

import (
	"context"
	"time"

	"github.com/0-s0g0/TEKUTEKU/server/pkg/errors"
	"github.com/0-s0g0/TEKUTEKU/server/pkg/null"
	timeformat "github.com/0-s0g0/TEKUTEKU/server/pkg/time_format"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/0-s0g0/TEKUTEKU/server/db/sql/query"
	"github.com/0-s0g0/TEKUTEKU/server/internal/domain/entity"
	"github.com/0-s0g0/TEKUTEKU/server/internal/domain/repository"
)

type MessageRepository struct {
	queries *query.Queries
}

func NewMessageRepository(queries *query.Queries) repository.IMessageRepository {
	return &MessageRepository{queries: queries}
}

// Create implements repository.IMessageRepository.
func (m *MessageRepository) Create(ctx context.Context, message entity.Message) (*entity.Message, error) {
	_, err := m.queries.CreateMessage(ctx, query.CreateMessageParams{
		MessageID: message.ID,
		School:    int32(message.School),
		Message:   message.Message,
		X:         int32(message.X),
		Y:         int32(message.Y),
		FloatTime: message.FloatTime,
		CreatedAt: timeformat.Format(message.CreatedAt),
		ParentID:  pgtype.Text{String: message.ParentID.Value, Valid: message.ParentID.Valid},
	})
	if err != nil {
		return nil, errors.HandleDBError(err)
	}
	// createdMessage := &entity.Message{
	// 	ID:        massage.MessageID,
	// 	School:    int(massage.School),
	// 	Message:   massage.Message,
	// 	Likes:     int(massage.Likes),
	// 	X:         int(massage.X),
	// 	Y:         int(massage.Y),
	// 	FloatTime: massage.FloatTime,
	// 	CreatedAt: massage.CreatedAt.Time,
	// }
	return nil, nil
}

// GetAll implements repository.IMessageRepository.
func (m *MessageRepository) GetAll(ctx context.Context) ([]entity.Message, error) {
	massage, err := m.queries.GetAllMessage(ctx)
	if err != nil {
		return nil, errors.HandleDBError(err)
	}
	a := make([]entity.Message, 0, len(massage))
	for _, m := range massage {
		a = append(a, entity.Message{
			ID:        m.MessageID,
			School:    int(m.School),
			Message:   m.Message,
			Likes:     int(m.Likes),
			X:         int(m.X),
			Y:         int(m.Y),
			FloatTime: m.FloatTime,
			CreatedAt: timeformat.Parse(m.CreatedAt),
			ParentID:  null.Null[string]{Value: m.ParentID.String, Valid: m.ParentID.Valid},
		})
	}
	return a, nil
}

// GetByID implements repository.IMessageRepository.
func (m *MessageRepository) GetByID(ctx context.Context, id string) (*entity.Message, error) {
	return nil, nil
}

// GetByTimeRange implements repository.IMessageRepository.
func (m *MessageRepository) GetByTimeRange(ctx context.Context, from time.Time, to time.Time) ([]entity.Message, error) {
	return nil, nil
}

// GiveLike implements repository.IMessageRepository.
func (m *MessageRepository) GiveLike(ctx context.Context, id string) error {
	if err := m.queries.IncrementLikes(ctx, id); err != nil {
		return err
	}
	return nil
}
