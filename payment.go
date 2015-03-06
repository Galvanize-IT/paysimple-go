package paysimple

import (
	"time"
)

type Status int64

const (
	Pending Status = iota
	Posted
	Settled
	Failed
	_
	Voided
	Reversed
	_
	_
	ReversePosted
	Chargeback
	_
	Authorized
	_
	Returned
	_
	ReverseNSF
	_
	RefundSettled
)

type Failure struct {
	Code               string
	Description        string
	MerchantActionText string
	IsDecline          bool
}

type Payment struct {
	ID                  int64 `json:"Id"`
	CustomerID          int64 `json:"CustomerId"`
	CustomerFirstName   string
	CustomerLastName    string
	CustomerCompany     string
	ReferenceID         int64 `json:"ReferenceId"`
	Status              string
	RecurringScheduleID int64 `json:"RecurringScheduleId"`
	PaymentType         string
	PaymentSubType      string
	ProviderAuthCode    string
	TraceNumber         string
	PaymentDate         time.Time
	ReturnDate          *time.Time
	EstimatedSettleDate time.Time
	ActualSettledDate   *time.Time
	CanVoidUntil        time.Time
	Amount              float64
	FailureData         Failure
	AccountID           int64 `json:"AccountId"`
	InvoiceID           int64 `json:"InvoiceId"`
	IsDebit             bool
	InvoiceNumber       string // TODO or int64?
	PurchaseOrderNumber string // TODO or int64?
	OrderID             int64  `json:"OrderId"`
	Description         string
	Latitude            *float64
	Longitude           *float64
	LastModified        time.Time
	CreateOn            time.Time
	// SuccessReceiptOptions
	// FailureReceiptOptions
}

type PaymentsResponse struct {
	Meta
	Response []Payment
}
