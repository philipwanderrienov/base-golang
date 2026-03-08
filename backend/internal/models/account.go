package models

type Account struct {
	ID    string `json:"id" example:"1"`                    // Unique identifier for the account
	Name  string `json:"name" example:"Alice Smith"`        // Name of the account holder
	Email string `json:"email" example:"alice@example.com"` // Email address of the account holder
}

type CreateAccountRequest struct {
	Name  string `json:"name" binding:"required" example:"Alice Smith"`              // Name of the account holder
	Email string `json:"email" binding:"required,email" example:"alice@example.com"` // Email address of the account holder
}

type UpdateAccountRequest struct {
	Name  string `json:"name" example:"Alice Smith"`              // Name of the account holder
	Email string `json:"email" example:"alice@example.com"` // Email address of the account holder
}

type AccountResponse struct {
	Message string   `json:"message" example:"Account retrieved successfully"` // Message describing the result of the operation
	Data    *Account `json:"data"`                                             // Pointer to account, can be nil for empty responses
}

type AccountsListResponse struct {
	Message string    `json:"message" example:"Accounts retrieved successfully"` // Message describing the result of the operation
	Data    []Account `json:"data"`                                              // Slice of accounts
	Total   int       `json:"total" example:"2"`                                 // Total count of accounts
}

type AccountErrorResponse struct {
	Error string `json:"error" example:"Account not found"` // Error message describing what went wrong
	Code  int    `json:"code" example:"404"`                // HTTP status code representing the error
}
