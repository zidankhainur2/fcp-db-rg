package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
	FetchWithClass() (*[]model.StudentClass, error)
}

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	var students []model.Student
    err := s.db.Find(&students).Error
    if err != nil {
        return nil, err
    }
    return students, nil
}


func (s *studentRepoImpl) Store(student *model.Student) error {
	return s.db.Create(student).Error
}


func (s *studentRepoImpl) Update(id int, student *model.Student) error {
    result := s.db.Model(&model.Student{}).Where("id = ?", id).Updates(student)
    if result.RowsAffected == 0 {
        return gorm.ErrRecordNotFound
    }
    return result.Error
}

func (s *studentRepoImpl) Delete(id int) error {
	return s.db.Where("id = ?", id).Delete(&model.Student{}).Error
}


func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	var student model.Student
	if err := s.db.Where("id = ?", id).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (s *studentRepoImpl) FetchWithClass() (*[]model.StudentClass, error) {
	var studentClasses []model.StudentClass
	query := `
		SELECT students.name, students.address, classes.name AS class_name, classes.professor, classes.room_number
		FROM students
		JOIN classes ON students.class_id = classes.id
	`
	if err := s.db.Raw(query).Scan(&studentClasses).Error; err != nil {
		return nil, err
	}
	if studentClasses == nil {
		studentClasses = []model.StudentClass{}
	}
	return &studentClasses, nil
}
	

