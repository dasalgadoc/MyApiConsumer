package model

import "time"

const DateFormat = time.RFC3339

type EngineResponse struct {
	TransactionCompliant bool       `json:"is_transaction_compliant"`
	TransactionDetail    string     `json:"transaction_detail"`
	Payer                EngineUser `json:"payer"`
	Collector            EngineUser `json:"collector"`
}

type EngineUser struct {
	UserId       int64         `json:"user_id"`
	SiteId       string        `json:"site_id"`
	Regulations  []Regulation  `json:"regulations"`
	Restrictions []Restriction `json:"restrictions"`
	Compliant    bool          `json:"is_compliant"`
	Detail       string        `json:"detail"`
	UserAmount   int           `json:"associated_users_amount"`
}

type Regulation struct {
	Name             string    `json:"name"`
	Status           string    `json:"status"`
	Level            string    `json:"level"`
	EvaluationResult string    `json:"evaluation_result"`
	AssignationDate  time.Time `json:"assignation_date"`
	LastUpdated      time.Time `json:"last_updated"`
}

type Restriction struct {
	Name            string    `json:"name"`
	Status          string    `json:"status"`
	AssignationDate time.Time `json:"assignation_date"`
	LastUpdated     time.Time `json:"last_updated"`
}
