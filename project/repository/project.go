package repository

import (
	"gorm.io/gorm"
)

type projectDescript struct {
	gorm.Model
	ProjectName    string
	About          string
	WebsiteUrl     string
	GithubUrl      string
	DockerImageUrl string
}

type AddProjectDescriptRepositoryRequest struct {
	ProjectName    string
	About          string
	WebsiteUrl     string
	GithubUrl      string
	DockerImageUrl string
}

type UpdateProjectDescriptRepositoryRequest struct {
	Id             int
	ProjectName    string
	About          string
	WebsiteUrl     string
	GithubUrl      string
	DockerImageUrl string
}

type projectTag struct {
	gorm.Model
	ProjectId int
	Main      string
	Sub       string
	SortValue int
}

type AddProjectTagRepositoryRequest struct {
	ProjectId int
	Main      string
	Sub       string
	SortValue int
}

type UpdateProjectTagRepositoryRequest struct {
	Id        int
	SortValue int
}

type projectPicture struct {
	gorm.Model
	ProjectId  int
	PictureUrl string
	SortValue  int
}

type AddProjectPictureRepositoryRequest struct {
	ProjectId  int
	PictureUrl string
	SortValue  int
}

type UpdateProjectPictureRepositoryRequest struct {
	Id        int
	SortValue int
}

type projectTopic struct {
	gorm.Model
	ProjectId int
	TopicName string
	Detail    string
	SortValue int
}

type AddProjectTopicRepositoryRequest struct {
	ProjectId int
	TopicName string
	Detail    string
	SortValue int
}

type UpdateProjectTopicRepositoryRequest struct {
	Id        int
	TopicName string
	Detail    string
	SortValue int
}

type group struct {
	gorm.Model
	GroupName string
	Detail    string
	SortValue int
}

type AddGroupRepositoryRequest struct {
	GroupName string
	Detail    string
	SortValue int
}

type UpdateGroupRepositoryRequest struct {
	Id        int
	GroupName string
	Detail    string
	SortValue int
}

type groupProject struct {
	gorm.Model
	GroupId   int
	ProjectId int
	SortValue int
}

type AddGroupProjectRepositoryRequest struct {
	GroupId   int
	ProjectId int
	SortValue int
}

type UpdateGroupProjectRepositoryRequest struct {
	Id        int
	SortValue int
}

type ProfileRepository interface {
	BeginTransaction() error
	CloseTransaction(error) error
	GetAllProjectDescript() ([]projectDescript, error)
	AddNewProjectDescript(AddProjectDescriptRepositoryRequest) error
	UpdateProjectDescript(UpdateProjectDescriptRepositoryRequest) error
	DeleteProjectDescript(int) error
	GetAllProjectTag() ([]projectTag, error)
	AddNewProjectTag(AddProjectTagRepositoryRequest) error
	UpdateProjectTag(UpdateProjectTagRepositoryRequest) error
	DeleteProjectTagById(int) error
	DeleteProjectTagByProjectId(int) error
	GetAllProjectPicture() ([]projectPicture, error)
	AddNewProjectPicture(AddProjectPictureRepositoryRequest) error
	UpdateProjectPicture(UpdateProjectPictureRepositoryRequest) error
	DeleteProjectPictureById(int) error
	DeleteProjectPictureByProjectId(int) error
	GetProjectTopicByProjectId(int) ([]projectTopic, error)
	AddNewProjectTopic(AddProjectTopicRepositoryRequest) error
	UpdateProjectTopic(UpdateProjectTopicRepositoryRequest) error
	DeleteProjectTopicById(int) error
	DeleteProjectTopicByProjectId(int) error
	GetAllGroup() ([]group, error)
	AddNewGroup(AddGroupRepositoryRequest) error
	UpdateGroup(UpdateGroupRepositoryRequest) error
	DeleteGroup(int) error
	GetAllGroupProject() ([]groupProject, error)
	AddGroupProject(AddGroupProjectRepositoryRequest) error
	UpdateGroupProject(UpdateGroupProjectRepositoryRequest) error
	DeleteGroupProjectById(int) error
	DeleteGroupProjectByGroupId(int) error
	DeleteGroupProjectByProjectId(int) error
}
