package usecase

import (
	"context"

	"assessment1/domain"
)

type LoanUsecaseImpl struct {
	loanRepo domain.LoanRepository
}

func NewLoanUsecase(loanRepo domain.LoanRepository) domain.LoanUsecase {
	return &LoanUsecaseImpl{loanRepo}
}

func (u *LoanUsecaseImpl) ApplyForLoan(ctx context.Context, loan domain.Loan) (domain.Loan, error) {
	return u.loanRepo.ApplyForLoan(ctx, loan)
}

func (u *LoanUsecaseImpl) GetLoanStatus(ctx context.Context, loanID string) (domain.Loan, error) {
	return u.loanRepo.GetLoanStatus(ctx, loanID)
}

func (u *LoanUsecaseImpl) GetAllLoans(ctx context.Context, status string, order string) ([]domain.Loan, error) {
	return u.loanRepo.GetAllLoans(ctx, status, order)
}

func (u *LoanUsecaseImpl) UpdateLoanStatus(ctx context.Context, loanID string, status string) (domain.Loan, error) {
	return u.loanRepo.UpdateLoanStatus(ctx, loanID, status)
}

func (u *LoanUsecaseImpl) DeleteLoan(ctx context.Context, loanID string) error {
	return u.loanRepo.DeleteLoan(ctx, loanID)
}
