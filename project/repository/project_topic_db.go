package repository

func (r projectRepositoryDB) GetProjectTopicByProjectId(id int) ([]projectTopic, error) {
	projectTopics := []projectTopic{}
	result := r.db.Find(&projectTopics, "project_id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return projectTopics, nil
}

func (r projectRepositoryDB) AddNewProjectTopic(repoReq AddProjectTopicRepositoryRequest) error {
	projectTopic := projectTopic{ProjectId: repoReq.ProjectId, TopicName: repoReq.TopicName, Detail: repoReq.Detail, SortValue: repoReq.SortValue}
	if result := tx.Create(&projectTopic); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r projectRepositoryDB) UpdateProjectTopic(repoReq UpdateProjectTopicRepositoryRequest) error {
	projectTopic := projectTopic{}
	if result := tx.First(&projectTopic, repoReq.Id); result.Error != nil {
		return result.Error
	}
	projectTopic.TopicName = repoReq.TopicName
	projectTopic.Detail = repoReq.Detail
	projectTopic.SortValue = repoReq.SortValue
	if result := tx.Save(&projectTopic); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r projectRepositoryDB) DeleteProjectTopicById(id int) error {
	if result := tx.Delete(&projectTopic{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r projectRepositoryDB) DeleteProjectTopicByProjectId(id int) error {
	if result := tx.Delete(&projectTopic{}, "project_id = ?", id); result.Error != nil {
		return result.Error
	}
	return nil
}
