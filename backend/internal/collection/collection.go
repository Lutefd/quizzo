package collection

import (
	"github.com/Lutefd/quizzo/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuizzesCollection interface {
	InsertQuiz(quiz model.Quiz) error
	GetQuizzes() ([]model.Quiz, error)
	GetQuizById(id primitive.ObjectID) (*model.Quiz, error)
	UpdateQuiz(quiz model.Quiz) error
	DeleteQuiz(id primitive.ObjectID) error
}
