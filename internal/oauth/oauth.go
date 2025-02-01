package auth

import (
    "time"
    "errors"
    "database/sql"
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
)

// User represents a user in the database
type User struct {
    ID           int    `json:"id"`
    Username     string `json:"username"`
    PasswordHash string `json:"-"`
    Email        string `json:"email"`
}

// Claims represents the JWT claims
type Claims struct {
    UserID   int    `json:"user_id"`
    Username string `json:"username"`
    jwt.RegisteredClaims
}

// GenerateToken generates a JWT token for a user
func GenerateToken(user *User, secretKey string, expiration time.Duration) (string, error) {
    claims := &Claims{
        UserID:   user.ID,
        Username: user.Username,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            Issuer:    "edi-system",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secretKey))
}

// ValidateCredentials validates a user's credentials against the database
func ValidateCredentials(db *sql.DB, username, password string) (*User, error) {
    var user User
    query := `SELECT id, username, password_hash, email FROM users WHERE username = $1`
    row := db.QueryRow(query, username)
    err := row.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Email)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("user not found")
        }
        return nil, err
    }

    // Compare the provided password with the hashed password
    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
        return nil, errors.New("invalid password")
    }

    return &user, nil
}
