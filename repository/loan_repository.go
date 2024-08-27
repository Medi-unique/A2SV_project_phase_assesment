package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"assessment1/domain"
)

type LoanRepositoryImpl struct {
	collection *mongo.Collection
}

func NewLoanRepositoryImpl(collection *mongo.Collection) domain.LoanRepository {
	return &LoanRepositoryImpl{collection}
}

func (r *LoanRepositoryImpl) ApplyForLoan(ctx context.Context, loan domain.Loan) (domain.Loan, error) {
	_, err := r.collection.InsertOne(ctx, loan)
	return loan, err
}

func (r *LoanRepositoryImpl) GetLoanStatus(ctx context.Context, loanID string) (domain.Loan, error) {
	var loan domain.Loan
	err := r.collection.FindOne(ctx, bson.M{"_id": loanID}).Decode(&loan)
	return loan, err
}

func (r *LoanRepositoryImpl) GetAllLoans(ctx context.Context, status string, order string) ([]domain.Loan, error) {
	var loans []domain.Loan
	filter := bson.M{}
	if status != "all" {
		filter["status"] = status
	}
	opts := options.Find().SetSort(bson.D{{"createdAt", 1}})
	if order == "desc" {
		opts.SetSort(bson.D{{"createdAt", -1}})
	}
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &loans)
	return loans, err
}

func (r *LoanRepositoryImpl) UpdateLoanStatus(ctx context.Context, loanID string, status string) (domain.Loan, error) {
	var loan domain.Loan
	update := bson.M{"$set": bson.M{"status": status, "updatedAt": time.Now()}}
	err := r.collection.FindOneAndUpdate(ctx, bson.M{"_id": loanID}, update).Decode(&loan)
	return loan, err
}

func (r *LoanRepositoryImpl) DeleteLoan(ctx context.Context, loanID string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": loanID})
	return err
}
