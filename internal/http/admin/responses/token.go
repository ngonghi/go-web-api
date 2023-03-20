package responses

import "github.com/ngonghi/vian-backend/internal/models"

// Token ...
type Token struct {
	Token string `json:"token"`
	ID    int64  `json:"id"`
	Name  string `json:"name"`
}

// NewToken ... Create new User response
func NewToken(token string, adminUser models.AdminUser) *Token {
	response := &Token{
		Token: token,
		ID:    adminUser.ID,
		Name:  adminUser.Name,
	}
	return response
}
