package database

import (
    "quiz-backend/internal/domain/entity"
    "github.com/glebarez/sqlite" 
    "gorm.io/gorm"
)

func NewSQLiteDB(dbPath string) (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    err = db.AutoMigrate(&entity.User{})
    if err != nil {
        return nil, err
    }

    return db, nil
}