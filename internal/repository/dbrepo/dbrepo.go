package dbrepo

import (
	"database/sql"

	"github.com/amitkrkun/bookings/internal/config"
	"github.com/amitkrkun/bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

type testDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

func NewTestingsRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &testDBRepo{
		App: a,
	}
}

// func (r *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
// 	// For testing purposes, assume that the room is available if the roomID is even
// 	// Replace this logic with your actual implementation for checking availability in the test environment
// 	return roomID%2 == 0, nil
// }

// func (r *testDBRepo) GetUserByID(id int) (model.user, error) {
// 	// Your implementation for retrieving a user by ID in the test environment
// 	// This is just a placeholder; replace it with your actual implementation
// 	return nil, nil
// }
