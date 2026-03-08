package models

type Congregation struct {
	ID       string `json:"id" example:"1"`                    // Unique identifier for the congregation
	Name     string `json:"name" example:"First Congregation"` // Name of the congregation
	Location string `json:"location" example:"City A"`         // Location of the congregation
}

type CongregationsListResponse struct {
	Message string         `json:"message" example:"Congregations retrieved successfully"` // Message describing the result of the operation
	Data    []Congregation `json:"data"`                                                   // Slice of congregations
	Total   int            `json:"total" example:"2"`                                      // Total count of congregations
}

type CongregationResponse struct {
	Message string        `json:"message" example:"Congregation retrieved successfully"` // Message describing the result of the operation
	Data    *Congregation `json:"data"`                                                  // Pointer to congregation, can be nil for empty responses
}

type CongregationErrorResponse struct {
	Error string `json:"error" example:"Congregation not found"` // Error message describing what went wrong
	Code  int    `json:"code" example:"404"`                     // HTTP status code representing the error
}

type CreateCongregationRequest struct {
	Name     string `json:"name" binding:"required" example:"First Congregation"` // Name of the congregation
	Location string `json:"location" binding:"required" example:"City A"`         // Location of the congregation
}

type UpdateCongregationRequest struct {
	Name     string `json:"name" example:"First Congregation"` // Name of the congregation
	Location string `json:"location" example:"City A"`         // Location of the congregation
}
