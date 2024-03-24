package repository

import "gorm.io/gorm"

type journeyRepositoryDB struct {
	db *gorm.DB
}

func NewJourneyRepositoryDB(db *gorm.DB) journeyRepositoryDB {
	db.AutoMigrate(&journey{})
	return journeyRepositoryDB{db: db}
}

var tx *gorm.DB

func (r journeyRepositoryDB) BeginTransaction() error {
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

func (r journeyRepositoryDB) CloseTransaction(err error) error {
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

func (r journeyRepositoryDB) GetAllJourney() ([]journey, error) {
	journeys := []journey{}
	result := r.db.Find(&journeys)
	if result.Error != nil {
		return nil, result.Error
	}
	return journeys, nil
}

func (r journeyRepositoryDB) AddNewJourney(repoReq AddRepositoryRequest) error {
	journey := journey{StartMonthYear: repoReq.StartMonthYear, EndMonthYear: repoReq.EndMonthYear, Organization: repoReq.Organization, Detail: repoReq.Detail, PictureUrl: repoReq.PictureUrl, SortValue: repoReq.SortValue}
	if result := tx.Create(&journey); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r journeyRepositoryDB) UpdateJourney(repoReq UpdateRepositoryRequest) error {
	journey := journey{}
	if result := tx.First(&journey, repoReq.Id); result.Error != nil {
		return result.Error
	}
	journey.StartMonthYear = repoReq.StartMonthYear
	journey.EndMonthYear = repoReq.EndMonthYear
	journey.Organization = repoReq.Organization
	journey.Detail = repoReq.Detail
	journey.PictureUrl = repoReq.PictureUrl
	journey.SortValue = repoReq.SortValue
	if result := tx.Save(&journey); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r journeyRepositoryDB) DeleteJourney(id int) error {
	if result := tx.Delete(&journey{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}
