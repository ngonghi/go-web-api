package responses

import "github.com/ngonghi/vian-backend/internal/models"

// AdminUser ...
type AdminUser struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"created_at"`
}

// AdminUsers ...
type AdminUsers []AdminUser

// NewAdminUser ... Create new AdminUser response
func NewAdminUser(model models.AdminUser) *AdminUser {
	response := &AdminUser{
		ID:        model.ID,
		Name:      model.Name,
		Email:     model.Email,
		Password:  "",
		CreatedAt: model.CreatedAt.Unix(),
	}
	return response
}

// NewAdminUsers ... Create new AdminUsers response
func NewAdminUsers(models *[]models.AdminUser) AdminUsers {
	response := AdminUsers{}
	for _, model := range *models {
		response = append(response, *NewAdminUser(model))
	}

	return response
}
