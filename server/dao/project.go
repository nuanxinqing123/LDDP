package dao

import "LDDP/server/model"

// GetProjectDataPage 获取Project表总数据
func GetProjectDataPage() int64 {
	var c []model.Project
	result := DB.Find(&c)
	return result.RowsAffected
}

// GetDivisionProjectData 分页查询
func GetDivisionProjectData(page int) []model.Project {
	var p []model.Project
	if page == 1 {
		// 获取前20条数据
		DB.Order("id desc").Limit(20).Offset(0).Find(&p)
	} else {
		DB.Order("id desc").Limit(20).Offset((page - 1) * 20).Find(&p)
	}
	return p
}

// GetDivisionProjectDataAll 获取全部项目数据
func GetDivisionProjectDataAll() []model.Project {
	var p []model.Project
	DB.Find(&p)
	return p
}

// ProjectSearch 模糊搜索
func ProjectSearch(tp, s string) []model.Project {
	var v []model.Project
	if tp == "项目类名" {
		DB.Where("project_type LIKE ?", "%"+s+"%").Find(&v)
	} else {
		DB.Where("project_remarks LIKE ?", "%"+s+"%").Find(&v)
	}

	return v
}

// ProjectSearchByID ID搜索项目
func ProjectSearchByID(id int) model.Project {
	var v model.Project
	DB.Where("id = ?", id).First(&v)
	return v
}

// CreateProject 创建项目
func CreateProject(pr *model.Project) {
	DB.Create(&pr)
}

// UpdateProject 修改项目
func UpdateProject(p model.Project) {
	DB.Save(&p)
}

// DelProject 删除项目
func DelProject(p *model.ProjectDelete) {
	var pr model.Project
	DB.Where("id = ? ", p.ID).First(&pr)
	DB.Delete(&pr)
}

// ProjectSearchOne 精确搜索
func ProjectSearchOne(s string) model.Project {
	var v model.Project
	DB.Where("project_type = ?", s).First(&v)
	return v
}

// ProjectSearchOneTrue 精确搜索True
func ProjectSearchOneTrue(s string) model.Project {
	var v model.Project
	DB.Where("project_type = ? AND project_state = ?", s, true).First(&v)
	return v
}

// GetTrueProject 查询已启动项目
func GetTrueProject() []model.Project {
	var p []model.Project
	DB.Where("project_state = ?", true).Find(&p)
	return p
}

// ProjectSearchByForwardID 查询项目绑定转发
func ProjectSearchByForwardID(id int) bool {
	var p model.Project
	DB.Where("project_api = ?", id).First(&p)
	if p.ID == 0 {
		return false
	} else {
		return true
	}
}
