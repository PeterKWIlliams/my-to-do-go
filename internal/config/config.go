package config

import (
	"github.com/PeterKWIlliams/my-to-do-go/internal/database"
)

type Config struct {
	DB *database.Queries
}
