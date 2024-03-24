package repository

func (r projectRepositoryDB) GetAllProjectTag() ([]projectTag, error) {
	projectTags := []projectTag{}
	result := r.db.Find(&projectTags)
	if result.Error != nil {
		return nil, result.Error
	}
	return projectTags, nil
}

func (r projectRepositoryDB) AddNewProjectTag(repoReq AddProjectTagRepositoryRequest) error {
	projectTag := projectTag{ProjectId: repoReq.ProjectId, Main: repoReq.Main, Sub: repoReq.Sub, SortValue: repoReq.SortValue}
	if result := tx.Create(&projectTag); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r projectRepositoryDB) UpdateProjectTag(repoReq UpdateProjectTagRepositoryRequest) error {
	projectTag := projectTag{}
	if result := tx.First(&projectTag, repoReq.Id); result.Error != nil {
		return result.Error
	}
	projectTag.SortValue = repoReq.SortValue
	if result := tx.Save(&projectTag); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r projectRepositoryDB) DeleteProjectTagById(id int) error {
	if result := tx.Delete(&projectTag{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r projectRepositoryDB) DeleteProjectTagByProjectId(id int) error {
	if result := tx.Delete(&projectTag{}, "project_id = ?", id); result.Error != nil {
		return result.Error
	}
	return nil
}
