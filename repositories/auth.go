// Create package repositories here ...
package repositories

// import the required packages here ...
import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

// AuthRepository interface here ...
type AuthRepository interface {
	Register(user models.User) (models.User, error)
}

// RepositoryAuth function here ...
func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

// Register method here ...
func (r *repository) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}
