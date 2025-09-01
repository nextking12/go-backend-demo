package repositories

import (
	"database/sql"

	"go-backend-demo/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	query := "SELECT id, name, email, created_at FROM users ORDER BY id"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) GetByID(id int) (*models.User, error) {
	query := "SELECT id, name, email, created_at FROM users WHERE id = $1"
	var user models.User
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, models.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(req *models.CreateUserRequest) (*models.User, error) {
	query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, created_at"
	var user models.User
	user.Name = req.Name
	user.Email = req.Email

	err := r.db.QueryRow(query, req.Name, req.Email).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(id int, req *models.UpdateUserRequest) (*models.User, error) {
	query := "UPDATE users SET name = $1, email = $2 WHERE id = $3"
	result, err := r.db.Exec(query, req.Name, req.Email, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, models.ErrUserNotFound
	}

	return r.GetByID(id)
}

func (r *UserRepository) Delete(id int) error {
	query := "DELETE FROM users WHERE id = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return models.ErrUserNotFound
	}
	return nil
}
