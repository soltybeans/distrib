package inserter

import (
	"database/sql"
	"log/slog"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func InsertData() {
	db, err := initDB()
	if err != nil {
		slog.Error("Error setting up DB", "error", err.Error())
	}
	defer db.Close()

	insertPlayerQuery := `INSERT INTO players (id, player_name)
					VALUES($1, $2)
					`
	randomPlayerNumOne := strconv.Itoa(rand.Intn(100000))
	playerOneUuid := uuid.New()
	playerTwoUuid := uuid.New()
	randomPlayerNumTwo := strconv.Itoa(rand.Intn(100000))
	_, err = db.Exec(insertPlayerQuery, playerOneUuid, "randomPlayer"+randomPlayerNumOne)
	if err != nil {
		slog.Error("Error inserting player", "error", err.Error())
	}
	_, err = db.Exec(insertPlayerQuery, playerTwoUuid, "randomPlayer"+randomPlayerNumTwo)
	if err != nil {
		slog.Error("Error inserting player", "error", err.Error())
	}

	createGameCmd := `INSERT INTO games (id, move_number, created_date, player_one, player_two)
						VALUES($1, $2, $3, $4, $5)
						`
	_, err = db.Exec(createGameCmd, uuid.New(), 0, time.Now().UTC(), playerOneUuid, playerTwoUuid)
	if err != nil {
		slog.Error("Error creating game", "error", err.Error())
	}
}

func initDB() (*sql.DB, error) {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "cdc")

	connStr := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser +
		" password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"

	return sql.Open("pgx", connStr)

}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
