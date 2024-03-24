package service

type projectDescriptResponse struct {
	Id             int    `json:"id" example:"1"`
	ProjectName    string `json:"project_name" example:"Smart Nutrition Calculator"`
	About          string `json:"about" example:"just a calculator"`
	WebsiteUrl     string `json:"website_url"`
	GithubUrl      string `json:"github_url"`
	DockerImageUrl string `json:"docker_image_url"`
}

type AddProjectDescriptServiceRequest struct {
	ProjectName    string `json:"project_name" example:"Smart Nutrition Calculator"`
	About          string `json:"about" example:"just a calculator"`
	WebsiteUrl     string `json:"website_url"`
	GithubUrl      string `json:"github_url"`
	DockerImageUrl string `json:"docker_image_url"`
}

type UpdateProjectDescriptServiceRequest struct {
	Id             int    `json:"id" example:"1"`
	ProjectName    string `json:"project_name" example:"Smart Nutrition Calculator"`
	About          string `json:"about" example:"just a calculator"`
	WebsiteUrl     string `json:"website_url"`
	GithubUrl      string `json:"github_url"`
	DockerImageUrl string `json:"docker_image_url"`
}

type projectTagResponse struct {
	Id        int    `json:"id" example:"1"`
	ProjectId int    `json:"project_id" example:"1"`
	Main      string `json:"main" example:"Back-End"`
	Sub       string `json:"sub" example:"Go"`
	SortValue int    `json:"sort_value" example:"1"`
}

type AddProjectTagServiceRequest struct {
	ProjectId int    `json:"project_id" example:"1"`
	Main      string `json:"main" example:"Back-End"`
	Sub       string `json:"sub" example:"Go"`
	SortValue int    `json:"sort_value" example:"1"`
}

type UpdateProjectTagServiceRequest struct {
	Id        int `json:"id" example:"1"`
	SortValue int `json:"sort_value" example:"1"`
}

type projectPictureResponse struct {
	Id         int    `json:"id" example:"1"`
	ProjectId  int    `json:"project_id" example:"1"`
	PictureUrl string `json:"picture_url"`
	SortValue  int    `json:"sort_value" example:"1"`
}

type AddProjectPictureServiceRequest struct {
	ProjectId  int    `json:"project_id" example:"1"`
	PictureUrl string `json:"picture_url"`
	SortValue  int    `json:"sort_value" example:"1"`
}

type UpdateProjectPictureServiceRequest struct {
	Id        int `json:"id" example:"1"`
	SortValue int `json:"sort_value" example:"1"`
}

type projectTopicResponse struct {
	Id        int    `json:"id" example:"1"`
	ProjectId int    `json:"project_id" example:"1"`
	TopicName string `json:"topic_name" example:"Background"`
	Detail    string `json:"detail" example:"I get the idea from my own problem"`
	SortValue int    `json:"sort_value" example:"1"`
}

type AddProjectTopicServiceRequest struct {
	ProjectId int    `json:"project_id" example:"1"`
	TopicName string `json:"topic_name" example:"Background"`
	Detail    string `json:"detail" example:"I get the idea from my own problem"`
	SortValue int    `json:"sort_value" example:"1"`
}

type UpdateProjectTopicServiceRequest struct {
	Id        int    `json:"id" example:"1"`
	TopicName string `json:"topic_name" example:"Background"`
	Detail    string `json:"detail" example:"I get the idea from my own problem"`
	SortValue int    `json:"sort_value" example:"1"`
}

type groupResponse struct {
	Id        int    `json:"id" example:"1"`
	GroupName string `json:"group_name" example:"Complete"`
	Detail    string `json:"detail" example:"All projects that already done and can run successfully"`
	SortValue int    `json:"sort_value" example:"1"`
}

type AddGroupServiceRequest struct {
	GroupName string `json:"group_name" example:"Complete"`
	Detail    string `json:"detail" example:"All projects that already done and can run successfully"`
	SortValue int    `json:"sort_value" example:"1"`
}

type UpdateGroupServiceRequest struct {
	Id        int    `json:"id" example:"1"`
	GroupName string `json:"group_name" example:"Complete"`
	Detail    string `json:"detail" example:"All projects that already done and can run successfully"`
	SortValue int    `json:"sort_value" example:"1"`
}

type groupProjectResponse struct {
	Id        int `json:"id" example:"1"`
	GroupId   int `json:"group_id" example:"1"`
	ProjectId int `json:"project_id" example:"1"`
	SortValue int `json:"sort_value" example:"1"`
}

type AddGroupProjectServiceRequest struct {
	GroupId   int `json:"group_id" example:"1"`
	ProjectId int `json:"project_id" example:"1"`
	SortValue int `json:"sort_value" example:"1"`
}

type UpdateGroupProjectServiceRequest struct {
	Id        int `json:"id" example:"1"`
	SortValue int `json:"sort_value" example:"1"`
}

type DeleteServiceRequest struct {
	Id int `json:"id" example:"1"`
}

type ProjectService interface {
	GetAllProjectDescript() ([]projectDescriptResponse, error)
	AddNewProjectDescript(AddProjectDescriptServiceRequest) error
	UpdateProjectDescript(UpdateProjectDescriptServiceRequest) error
	DeleteProjectDescript(DeleteServiceRequest) error
	GetAllProjectTag() ([]projectTagResponse, error)
	AddNewProjectTag([]AddProjectTagServiceRequest) error
	UpdateProjectTag([]UpdateProjectTagServiceRequest) error
	DeleteProjectTag([]DeleteServiceRequest) error
	GetAllProjectPicture() ([]projectPictureResponse, error)
	AddNewProjectPicture([]AddProjectPictureServiceRequest) error
	UpdateProjectPicture([]UpdateProjectPictureServiceRequest) error
	DeleteProjectPicture([]DeleteServiceRequest) error
	GetProjectTopicByProjectId(int) ([]projectTopicResponse, error)
	AddNewProjectTopic(AddProjectTopicServiceRequest) error
	UpdateProjectTopic([]UpdateProjectTopicServiceRequest) error
	DeleteProjectTopic(DeleteServiceRequest) error
	GetAllGroup() ([]groupResponse, error)
	AddNewGroup(AddGroupServiceRequest) error
	UpdateGroup([]UpdateGroupServiceRequest) error
	DeleteGroup(DeleteServiceRequest) error
	GetAllGroupProject() ([]groupProjectResponse, error)
	AddGroupProject([]AddGroupProjectServiceRequest) error
	UpdateGroupProject([]UpdateGroupProjectServiceRequest) error
	DeleteGroupProject([]DeleteServiceRequest) error
}
