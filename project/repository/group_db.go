package repository

func (r projectRepositoryDB) GetAllGroup() ([]group, error) {
	groups := []group{}
	result := r.db.Find(&groups)
	if result.Error != nil {
		return nil, result.Error
	}
	return groups, nil
}

func (r projectRepositoryDB) AddNewGroup(repoReq AddGroupRepositoryRequest) error {
	group := group{GroupName: repoReq.GroupName, Detail: repoReq.Detail, SortValue: repoReq.SortValue}
	if result := tx.Create(&group); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r projectRepositoryDB) UpdateGroup(repoReq UpdateGroupRepositoryRequest) error {
	group := group{}
	if result := tx.First(&group, repoReq.Id); result.Error != nil {
		return result.Error
	}
	group.GroupName = repoReq.GroupName
	group.Detail = repoReq.Detail
	group.SortValue = repoReq.SortValue
	if result := tx.Save(&group); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r projectRepositoryDB) DeleteGroup(id int) error {
	if result := tx.Delete(&group{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}
