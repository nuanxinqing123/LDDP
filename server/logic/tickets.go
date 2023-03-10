package logic

import (
	"LDDP/server/dao"
	"LDDP/server/model"
	res "LDDP/utils/response"
	"os"
	"strconv"
	"strings"

	"github.com/segmentio/ksuid"
	"go.uber.org/zap"
)

// TicketsDivisionData Tickets分页查询
func TicketsDivisionData(page string) (res.ResCode, model.TicketsPageData) {
	var data []model.Tickets
	var pageData model.TicketsPageData

	if page == "" {
		// 空值，默认获取前20条数据(第一页)
		data = dao.TicketsDivisionTicketsData(1)
	} else {
		// String转Int
		intPage, err := strconv.Atoi(page)
		if err != nil {
			// 类型转换失败，查询默认获取前20条数据(第一页)
			zap.L().Error(err.Error())
			data = dao.TicketsDivisionTicketsData(1)
		} else {
			// 查询指定页数的数据
			data = dao.TicketsDivisionTicketsData(intPage)
		}
	}

	// 查询总页数
	count := dao.GetTicketsDataPage()
	// 计算页数
	z := count / 20
	var y int64
	y = count % 20

	if y != 0 {
		pageData.Page = z + 1
	} else {
		pageData.Page = z
	}
	pageData.PageData = data

	// 数据模糊处理
	for i := 0; i < len(pageData.PageData); i++ {
		str := pageData.PageData[i].TicketsKey
		s2 := strings.SplitAfter(str, "")
		s2[4] = "*"
		s2[5] = "*"
		s2[6] = "*"
		s2[7] = "*"
		s2[8] = "*"
		s2[9] = "*"
		s2[10] = "*"
		s2[11] = "*"
		pageData.PageData[i].TicketsKey = strings.Join(s2, "")
	}

	return res.CodeSuccess, pageData
}

// TicketsSearch Tickets数据查询
func TicketsSearch(tp, state, s string) (res.ResCode, []model.Tickets) {
	var t []model.Tickets

	if tp == "卡密值" {
		// 卡密值模糊搜索
		t = dao.TicketsValueSearch(state, s)
	} else {
		// 标识模糊搜索
		t = dao.TicketsRemarksSearch(state, s)
	}

	return res.CodeSuccess, t
}

// TicketsAdd 批量生成卡密
func TicketsAdd(p *model.CreateTickets) res.ResCode {
	// 判断本地是否还有遗留文件
	_, err := os.Stat("Tickets.txt")
	if err == nil {
		// 删除旧文件
		err = os.Remove("Tickets.txt")
		if err != nil {
			zap.L().Error(err.Error())
			return res.CodeServerBusy
		}
	}

	// 创建记录数组
	var li []string

	// 创建对象
	tickets := new(model.Tickets)

	tickets.TicketsKeyState = true
	tickets.TicketsKeyPoints = p.TicketsKeyPoints
	tickets.TicketsKeyRemarks = p.TicketsKeyRemarks
	// 获取生成数量
	for i := 0; i < p.TicketsKeyCount; i++ {
		// 生成用户UID
		uid := ksuid.New()
		tickets.TicketsKey = p.TicketsKeyPrefix + uid.String()

		// 加入数组
		li = append(li, tickets.TicketsKey)

		// 写入数据库
		dao.TicketsAdd(tickets)
	}

	// 创建Tickets.txt并写入数据
	filepath := "Tickets.txt"
	_, err = os.Create(filepath)
	if err != nil {
		// 记录错误
		zap.L().Error("[生成卡密]序列化数据失败:" + err.Error())
		return res.CodeTicketsError
	}
	// 打开JSON文件
	f, err := os.Open(filepath)
	if err != nil {
		// 记录错误
		zap.L().Error("[生成卡密]序列化数据失败:" + err.Error())
		return res.CodeTicketsError
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			// 记录错误
			zap.L().Error("[生成卡密]序列化数据失败:" + err.Error())
		}
	}(f)

	// 写入Tickets数据
	for i := 0; i < len(li); i++ {
		_ = os.WriteFile(filepath, []byte(li[i]+"\n"), 0777)
	}

	return res.CodeSuccess
}

// TicketsDataDelete Tickets删除本地数据
func TicketsDataDelete() {
	//time.Sleep(time.Second * 3)
	err := os.Remove("Tickets.txt")
	if err != nil {
		zap.L().Error(err.Error())
	}
}

// TicketsDelete 删除Tickets
func TicketsDelete(p *model.DelTickets) res.ResCode {
	// 删除Tickets数据
	if err := dao.TicketsDelete(p); err != nil {
		zap.L().Error("Error update database, err:", zap.Error(err))
		return res.CodeTicketsError
	}
	return res.CodeSuccess
}
