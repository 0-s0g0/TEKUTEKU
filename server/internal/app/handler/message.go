package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/0-s0g0/TEKUTEKU/server/internal/app/handler/schema"
	"github.com/0-s0g0/TEKUTEKU/server/internal/domain/entity"
	"github.com/0-s0g0/TEKUTEKU/server/pkg/null"

	"github.com/0-s0g0/TEKUTEKU/server/internal/app/service"
)

type IMessageHandler interface {
	GET() func(http.ResponseWriter, *http.Request) error
	POST() func(http.ResponseWriter, *http.Request) error
	ReplyPOST() func(http.ResponseWriter, *http.Request) error
}

type MessageHandler struct {
	ms service.IMessageService
}

func NewMessageHandler(ms service.IMessageService) IMessageHandler {
	return &MessageHandler{
		ms: ms,
	}
}

// GET implements IMessageHandler.
func (m *MessageHandler) GET() func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, s *http.Request) error {
		message, err := m.ms.GetAll(s.Context())
		if err != nil {
			return err
		}
		m := make([]schema.Message, 0, len(message))
		for _, v := range message {
			if v.Reply != nil {
				reply := make([]schema.Message, 0, len(v.Reply))
				for _, r := range v.Reply {
					reply = append(reply, schema.Message{
						ID:        r.ID,
						School:    r.School,
						Message:   r.Message,
						Likes:     r.Likes,
						X:         r.X,
						Y:         r.Y,
						FloatTime: r.FloatTime,
						CreatedAt: r.CreatedAt.String(),
					})
				}
				m = append(m, schema.Message{
					ID:        v.ID,
					School:    v.School,
					Message:   v.Message,
					Likes:     v.Likes,
					X:         v.X,
					Y:         v.Y,
					FloatTime: v.FloatTime,
					CreatedAt: v.CreatedAt.String(),
					Reply:     reply,
				})
			} else {
				m = append(m, schema.Message{
					ID:        v.ID,
					School:    v.School,
					Message:   v.Message,
					Likes:     v.Likes,
					X:         v.X,
					Y:         v.Y,
					FloatTime: v.FloatTime,
					CreatedAt: v.CreatedAt.String(),
				})
			}
		}
		res := schema.MessageGETResponse{
			Messages: m,
		}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			return err
		}
		w.WriteHeader(http.StatusCreated)
		return nil

	}
}

// POST implements IMessageHandler.
func (m *MessageHandler) POST() func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req schema.MessagePOSTRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return err
		}
		message := entity.Message{
			School:  req.School,
			Message: req.Message,
		}
		createdMessage, err := m.ms.Create(r.Context(), message)
		if err != nil {
			return err
		}
		log.Printf("createdMessage: %+v", createdMessage)

		w.WriteHeader(http.StatusCreated)
		return nil
	}
}

// Reply implements IMessageHandler.
func (m *MessageHandler) ReplyPOST() func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req schema.ReplyPOSTRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return err
		}
		message := entity.Message{
			School:   req.School,
			Message:  req.Message,
			ParentID: null.New(req.To),
		}
		createdMessage, err := m.ms.Create(r.Context(), message)
		if err != nil {
			return err
		}
		log.Printf("createdMessage: %+v", createdMessage)
		w.WriteHeader(http.StatusCreated)

		return nil
	}
}
