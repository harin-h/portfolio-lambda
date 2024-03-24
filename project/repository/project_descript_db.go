package repository

func (r projectRepositoryDB) GetAllProjectDescript() ([]projectDescript, error) {
	projectDescripts := []projectDescript{}
	result := r.db.Find(&projectDescripts)
	if result.Error != nil {
		return nil, result.Error
	}
	return projectDescripts, nil
}

func (r projectRepositoryDB) AddNewProjectDescript(repoReq AddProjectDescriptRepositoryRequest) error {
	projectDescript := projectDescript{ProjectName: repoReq.ProjectName, About: repoReq.About, WebsiteUrl: repoReq.WebsiteUrl, GithubUrl: repoReq.GithubUrl, DockerImageUrl: repoReq.DockerImageUrl}
	if result := tx.Create(&projectDescript); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r projectRepositoryDB) UpdateProjectDescript(repoReq UpdateProjectDescriptRepositoryRequest) error {
	projectDescript := projectDescript{}
	if result := tx.First(&projectDescript, repoReq.Id); result.Error != nil {
		return result.Error
	}
	projectDescript.ProjectName = repoReq.ProjectName
	projectDescript.About = repoReq.About
	projectDescript.WebsiteUrl = repoReq.WebsiteUrl
	projectDescript.GithubUrl = repoReq.GithubUrl
	projectDescript.DockerImageUrl = repoReq.DockerImageUrl
	if result := tx.Save(&projectDescript); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r projectRepositoryDB) DeleteProjectDescript(id int) error {
	if result := tx.Delete(&projectDescript{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}
