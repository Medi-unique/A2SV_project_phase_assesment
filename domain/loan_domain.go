package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Loan struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID    primitive.ObjectID `bson:"userId" json:"userId"`
	Amount    float64            `bson:"amount" json:"amount"`
	Status    string             `bson:"status" json:"status"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type LoanUsecase interface {
	ApplyForLoan(ctx context.Context, loan Loan) (Loan, error)
	GetLoanStatus(ctx context.Context, loanID string) (Loan, error)
	GetAllLoans(ctx context.Context, status string, order string) ([]Loan, error)
	UpdateLoanStatus(ctx context.Context, loanID string, status string) (Loan, error)
	DeleteLoan(ctx context.Context, loanID string) error
}

type LoanRepository interface {
	ApplyForLoan(ctx context.Context, loan Loan) (Loan, error)
	GetLoanStatus(ctx context.Context, loanID string) (Loan, error)
	GetAllLoans(ctx context.Context, status string, order string) ([]Loan, error)
	UpdateLoanStatus(ctx context.Context, loanID string, status string) (Loan, error)
	DeleteLoan(ctx context.Context, loanID string) error
}
