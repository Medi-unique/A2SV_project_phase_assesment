package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"assessment1/domain"
)

type LoanController struct {
	loanUsecase domain.LoanUsecase
}

func NewLoanController(loanUsecase domain.LoanUsecase) *LoanController {
	return &LoanController{loanUsecase}
}

func (c *LoanController) ApplyForLoan(ctx *gin.Context) {
	var loan domain.Loan
	if err := ctx.ShouldBindJSON(&loan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	loan, err := c.loanUsecase.ApplyForLoan(ctx, loan)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, loan)
}

func (c *LoanController) GetLoanStatus(ctx *gin.Context) {
	loanID := ctx.Param("id")
	loan, err := c.loanUsecase.GetLoanStatus(ctx, loanID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, loan)
}

func (c *LoanController) GetAllLoans(ctx *gin.Context) {
	status := ctx.DefaultQuery("status", "all")
	order := ctx.DefaultQuery("order", "asc")
	loans, err := c.loanUsecase.GetAllLoans(ctx, status, order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, loans)
}

func (c *LoanController) UpdateLoanStatus(ctx *gin.Context) {
	loanID := ctx.Param("id")
	var status struct {
		Status string `json:"status"`
	}
	if err := ctx.ShouldBindJSON(&status); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	loan, err := c.loanUsecase.UpdateLoanStatus(ctx, loanID, status.Status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, loan)
}

func (c *LoanController) DeleteLoan(ctx *gin.Context) {
	loanID := ctx.Param("id")
	err := c.loanUsecase.DeleteLoan(ctx, loanID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Loan deleted successfully"})
}
