package repositories

import (
	"shepays/db"
	"shepays/models"
)

type KycUserDocRepository struct {
	Db *db.Database
}

func (s *KycUserDocRepository) CreateKycUserDoc(w *models.KYCPAN) error {
	return s.Db.CreateKycUserDoc_(w)
}

func (s *KycUserDocRepository) ReadKycUserDoc(userId string) (*models.KYCPAN, error) {
	return s.Db.ReadKycUserDoc_(userId)
}

func (s *KycUserDocRepository) UpdateKycUserDoc(w *models.KYCPAN) error {
	return s.Db.UpdateKycUserDoc_(w)
}
