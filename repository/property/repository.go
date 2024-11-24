package property

import (
	"geoproperty_be/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	DB *gorm.DB
}

// Generate implements domain.PropertyRepository.
func (r *Repository) Generate(query string) ([]int, error) {
	var id []int

	if err := r.DB.Raw(query).Scan(&id).Error; err != nil {
		return nil, err
	}

	return id, nil
}

// Delete implements domain.PropertyRepository.
func (r *Repository) Delete(uid string) error {
	if err := r.DB.Where("uuid = ?", uid).Delete(&domain.Property[string, string]{}).Error; err != nil {
		return err
	}

	return nil
}

// Find implements domain.PropertyRepository.
func (r *Repository) Find(param map[string]any) (*[]domain.Property[string, string], error) {
	var properties []domain.Property[string, string]

	if err := r.DB.Preload(clause.Associations).Where(param).Find(&properties).Error; err != nil {
		return nil, err
	}

	return &properties, nil
}

// Insert implements domain.PropertyRepository.
func (r *Repository) Insert(property domain.Property[string, string]) (*domain.Property[string, string], error) {
	if err := r.DB.Preload(clause.Associations).Create(&property).Find(&property).Error; err != nil {
		return nil, err
	}

	return &property, nil
}

// Update implements domain.PropertyRepository.
func (*Repository) Update(property domain.Property[string, string]) (*domain.Property[string, string], error) {
	panic("unimplemented")
}

// FinByPolygon implements domain.PropertyRepository.
func (r *Repository) FinByPolygon(polygon string) (*[]domain.Property[string, string], error) {
	var properties []domain.Property[string, string]

	if err := r.DB.Preload(clause.Associations).Where("ST_Intersects(geometry, ?)", polygon).Find(&properties).Error; err != nil {
		return nil, err
	}

	return &properties, nil
}

func NewRepository(db *gorm.DB) domain.PropertyRepository {
	return &Repository{
		DB: db,
	}
}
