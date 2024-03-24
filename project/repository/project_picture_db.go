package repository

func (r projectRepositoryDB) GetAllProjectPicture() ([]projectPicture, error) {
	projectPictures := []projectPicture{}
	result := r.db.Find(&projectPictures)
	if result.Error != nil {
		return nil, result.Error
	}
	return projectPictures, nil
}

func (r projectRepositoryDB) AddNewProjectPicture(repoReq AddProjectPictureRepositoryRequest) error {
	projectPicture := projectPicture{ProjectId: repoReq.ProjectId, PictureUrl: repoReq.PictureUrl, SortValue: repoReq.SortValue}
	if result := tx.Create(&projectPicture); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r projectRepositoryDB) UpdateProjectPicture(repoReq UpdateProjectPictureRepositoryRequest) error {
	projectPicture := projectPicture{}
	if result := tx.First(&projectPicture, repoReq.Id); result.Error != nil {
		return result.Error
	}
	projectPicture.SortValue = repoReq.SortValue
	if result := tx.Save(&projectPicture); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r projectRepositoryDB) DeleteProjectPictureById(id int) error {
	if result := tx.Delete(&projectPicture{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r projectRepositoryDB) DeleteProjectPictureByProjectId(id int) error {
	if result := tx.Delete(&projectPicture{}, "project_id = ?", id); result.Error != nil {
		return result.Error
	}
	return nil
}
