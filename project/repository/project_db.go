package repository

import "gorm.io/gorm"

type projectRepositoryDB struct {
	db *gorm.DB
}

func NewProjectRepositoryDB(db *gorm.DB) projectRepositoryDB {
	db.AutoMigrate(&projectDescript{})
	db.AutoMigrate(&projectTag{})
	db.AutoMigrate(&projectPicture{})
	db.AutoMigrate(&projectTopic{})
	db.AutoMigrate(&group{})
	db.AutoMigrate(&groupProject{})
	return projectRepositoryDB{db: db}
}

var tx *gorm.DB

func (r projectRepositoryDB) BeginTransaction() error {
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

func (r projectRepositoryDB) CloseTransaction(err error) error {
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
