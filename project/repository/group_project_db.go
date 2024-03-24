package repository

func (r projectRepositoryDB) GetAllGroupProject() ([]groupProject, error) {
	groupProjects := []groupProject{}
	result := r.db.Find(&groupProjects)
	if result.Error != nil {
		return nil, result.Error
	}
	return groupProjects, nil
}

func (r projectRepositoryDB) AddGroupProject(repoReq AddGroupProjectRepositoryRequest) error {
	groupProject := groupProject{GroupId: repoReq.GroupId, ProjectId: repoReq.ProjectId, SortValue: repoReq.SortValue}
	if result := tx.Create(&groupProject); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r projectRepositoryDB) UpdateGroupProject(repoReq UpdateGroupProjectRepositoryRequest) error {
	groupProject := groupProject{}
	if result := tx.First(&groupProject, repoReq.Id); result.Error != nil {
		return result.Error
	}
	groupProject.SortValue = repoReq.SortValue
	if result := tx.Save(&groupProject); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r projectRepositoryDB) DeleteGroupProjectById(id int) error {
	if result := tx.Delete(&groupProject{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r projectRepositoryDB) DeleteGroupProjectByGroupId(id int) error {
	if result := tx.Delete(&groupProject{}, "group_id = ?", id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r projectRepositoryDB) DeleteGroupProjectByProjectId(id int) error {
	if result := tx.Delete(&groupProject{}, "project_id = ?", id); result.Error != nil {
		return result.Error
	}
	return nil
}
