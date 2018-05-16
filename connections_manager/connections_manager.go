package connections_manager

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math/rand"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/tochk/cockroachdb-admin-backend/configuration"
	"github.com/tochk/cockroachdb-admin-backend/user"
)

var connections map[string]*sqlx.DB

var InvalidTokenError = errors.New("invalid token")

func Init() {
	connections = make(map[string]*sqlx.DB)
}

func generateKey() []byte {
	const letterBytes = "1234567890abcdefghijklmnopqrstuvwxyz.,*_-=+"
	salt := make([]byte, 128)
	for i := range salt {
		salt[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return salt
}

func generateToken(login, password string) (string) {
	hasher := sha256.New()
	hasher.Write([]byte(login))
	token := hex.EncodeToString(hasher.Sum(nil))
	token += password
	hasher = sha256.New()
	hasher.Write([]byte(token))
	token = hex.EncodeToString(hasher.Sum(nil))
	hasher = sha256.New()
	hasher.Write(generateKey())
	token = hex.EncodeToString(hasher.Sum(nil))
	return token
}

func Connect(usr user.User) (token string, err error) {
	token = generateToken(usr.Login, usr.Password)
	connections[token], err = sqlx.Connect("postgres", "host="+configuration.Database.Host+" port="+strconv.Itoa(configuration.Database.Port)+" user="+usr.Login+" password="+usr.Password)
	return token, err
}

func GetConnection(token string) (*sqlx.DB, error) {
	if connections[token] == nil {
		return nil, InvalidTokenError
	}
	return connections[token], nil
}

//todo kill connections while inactive
