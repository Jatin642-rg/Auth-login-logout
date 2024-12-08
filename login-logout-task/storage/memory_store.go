package storage

import "errors"

var users = map[string]string{}
var tokens = map[string]bool{}

func InitializeStore() {
	users = make(map[string]string)
	tokens = make(map[string]bool)
}

func AddUser(email, password string) error {
	if _, exists := users[email]; exists {
		return errors.New("user already exists")
	}
	users[email] = password
	return nil
}

func GetUser(email string) (struct{ Password string }, error) {
	password, exists := users[email]
	if !exists {
		return struct{ Password string }{}, errors.New("user not found")
	}
	return struct{ Password string }{Password: password}, nil
}

func RevokeToken(token string) {
	tokens[token] = true
}

func IsTokenRevoked(token string) bool {
	return tokens[token]
}
