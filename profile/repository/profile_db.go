package repository

import (
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type profileRepositoryDB struct {
	db *gorm.DB
}

func NewProfileRepositoryDB(db *gorm.DB) profileRepositoryDB {
	db.AutoMigrate(&profile{})
	return profileRepositoryDB{db: db}
}

var tx *gorm.DB

func (r profileRepositoryDB) BeginTransaction() error {
	tx = r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r profileRepositoryDB) CloseTransaction(err error) error {
	if err != nil {
		err = tx.Rollback().Error
	} else {
		err = tx.Commit().Error
	}
	tx = nil
	if err != nil {
		return err
	}
	return nil
}

func (r profileRepositoryDB) GetAllMetric() ([]profile, error) {
	profiles := []profile{}
	result := r.db.Find(&profiles)
	if result.Error != nil {
		return nil, result.Error
	}
	return profiles, nil
}

func (r profileRepositoryDB) AddNewMetric(repoReq AddRepositoryRequest) error {
	profile := profile{MetricName: repoReq.MetricName, MetricValue: repoReq.MetricValue, SortValue: repoReq.SortValue}
	if result := tx.Create(&profile); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r profileRepositoryDB) UpdateMetric(repoReq UpdateRepositoryRequest) error {
	profile := profile{}
	if result := tx.First(&profile, repoReq.Id); result.Error != nil {
		return result.Error
	}
	profile.MetricValue = repoReq.MetricValue
	profile.SortValue = repoReq.SortValue
	if result := tx.Save(&profile); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r profileRepositoryDB) DeleteMetric(id int) error {
	if result := tx.Delete(&profile{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}
