package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type SessionsRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailName(name string) error
	SessionAvailToken(token string) (model.Session, error)
}

type sessionsRepoImpl struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) *sessionsRepoImpl {
	return &sessionsRepoImpl{db}
}

func (s *sessionsRepoImpl) AddSessions(session model.Session) error {
	return s.db.Create(&session).Error
}

func (s *sessionsRepoImpl) DeleteSession(token string) error {
	return s.db.Where("token = ?", token).Delete(&model.Session{}).Error
}


func (s *sessionsRepoImpl) UpdateSessions(session model.Session) error {
    result := s.db.Model(&model.Session{}).Where("username = ?", session.Username).Updates(session)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func (s *sessionsRepoImpl) SessionAvailName(name string) error {
    return s.db.Where("username = ?", name).First(&model.Session{}).Error
}


func (s *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
	var session model.Session
	err := s.db.Where("token = ?", token).First(&session).Error
    if err != nil {
        return model.Session{}, err
    }
    return session, nil
}


