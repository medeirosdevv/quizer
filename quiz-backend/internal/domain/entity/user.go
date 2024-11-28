package entity

import (
    "golang.org/x/crypto/bcrypt"
    "time"
)

type User struct {
    ID        uint      `json:"id" gorm:"primarykey"`
    Name      string    `json:"name"`
    Email     string    `json:"email" gorm:"unique"`
    Username  string    `json:"username" gorm:"unique"`
    Password  string    `json:"-"` 
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(name, email, username, password string) (*User, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    return &User{
        Name:     name,
        Email:    email,
        Username: username,
        Password: string(hashedPassword),
    }, nil
}

func (u *User) CheckPassword(password string) error {
    return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}