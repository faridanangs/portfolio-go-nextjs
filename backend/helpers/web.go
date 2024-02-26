package helpers

import "portfolio/models/entities"

func AuthResponse(req *entities.AuthUser) *entities.AuthResponse {
	return &entities.AuthResponse{
		Id:        req.Id,
		Username:  req.Username,
		Email:     req.Email,
		CreatedAt: req.CreatedAt,
	}
}
