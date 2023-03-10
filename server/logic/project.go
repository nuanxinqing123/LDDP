package logic

import (
	"LDDP/server/dao"
	"LDDP/server/model"
	res "LDDP/utils/response"
	"go.uber.org/zap"
	"strconv"
)

// GetDivisionProjectData 分页查询
func GetDivisionProjectData(page string) (res.ResCode, model.ProjectPage) {
	var data []model.Project
	var pPage model.ProjectPage
	if page == "" {
		// 空值，默认获取前20条数据(第一页)
		data = dao.GetDivisionProjectData(1)
	} else {
		// String转Int
		intPage, err := strconv.Atoi(page)
		if err != nil {
			// 类型转换失败，查询默认获取前20条数据(第一页)
			zap.L().Error(err.Error())
			data = dao.GetDivisionProjectData(1)
		} else {
			// 查询指定页数的数据
			data = dao.GetDivisionProjectData(intPage)
		}
	}

	// 查询总页数
	count := dao.GetProjectDataPage()
	// 计算页数
	z := count / 20
	var y int64
	y = count % 20

	if y != 0 {
		pPage.Page = z + 1
	} else {
		pPage.Page = z
	}
	pPage.PageData = data

	return res.CodeSuccess, pPage
}

// GetDivisionProjectDataAll 获取全部项目数据
func GetDivisionProjectDataAll() (res.ResCode, []model.Project) {
	return res.CodeSuccess, dao.GetDivisionProjectDataAll()
}

// ProjectAdd 创建项目
func ProjectAdd(p *model.ProjectAdd) (res.ResCode, string) {
	// 检查项目类名是否重复
	if dao.ProjectSearchOne(p.ProjectType).ProjectType != "" {
		return res.CodeProjectError, "项目类名存在重复"
	}

	pr := &model.Project{
		ProjectType:       p.ProjectType,
		ProjectPrice:      p.ProjectPrice,
		ProjectAgentPrice: p.ProjectAgentPrice,
		ProjectAPI:        -1,
		ProjectRemarks:    p.ProjectRemarks,
		ProjectTips:       "",
		ProjectAgentState: false,
		ProjectState:      false,
	}
	dao.CreateProject(pr)
	return res.CodeSuccess, "创建成功"
}

// ProjectUpdate 修改项目
func ProjectUpdate(p *model.ProjectUpdate) res.ResCode {
	data := dao.ProjectSearchByID(p.ID)
	data.ProjectType = p.ProjectType
	data.ProjectPrice = p.ProjectPrice
	data.ProjectAgentPrice = p.ProjectAgentPrice
	data.ProjectAPI = p.ProjectAPI
	data.ProjectRemarks = p.ProjectRemarks
	data.ProjectState = p.ProjectState
	data.ProjectTips = p.ProjectTips
	data.ProjectAgentState = p.ProjectAgentState

	dao.UpdateProject(data)
	return res.CodeSuccess
}

// UserTaskList 可选任务列表
func UserTaskList() (res.ResCode, []model.ProjectList) {
	// 获取面板所有已启动的项目
	data := dao.GetTrueProject()

	// 列表删除
	var pl []model.ProjectList
	for _, d := range data {
		var p model.ProjectList
		p.ProjectType = d.ProjectType
		p.ProjectPrice = d.ProjectPrice
		p.ProjectRemarks = d.ProjectRemarks
		p.ProjectTips = d.ProjectTips
		pl = append(pl, p)
	}

	return res.CodeSuccess, pl
}
