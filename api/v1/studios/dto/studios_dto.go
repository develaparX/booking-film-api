package dto

type CreateUStudioRequest struct {
	Name     string `json:"name" validate:"required"`
	Capacity string `json:"capacity" validate:"required,capacity"`
}

type UpdateStudioRequest struct {
	Id       string `json:"id"`
	Capacity string `json:"capacity" validate:"required"`
}

type StudioResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"email"`
}
