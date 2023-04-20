package controller

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jiaruling/Gateway/dao"
	"github.com/jiaruling/Gateway/dto"
	"github.com/jiaruling/Gateway/global"
	"github.com/jiaruling/Gateway/middleware"
	"github.com/jiaruling/golang_utils/lib"
	// "time"
	// "github.com/e421083458/go_gateway/public"
	// "github.com/e421083458/golang_common/lib"
	// "github.com/gin-gonic/gin"
	// "github.com/jiaruling/Gateway/dao"
	// "github.com/jiaruling/Gateway/dto"
	// "github.com/jiaruling/Gateway/middleware"
	// "github.com/pkg/errors"
)

type DashboardController struct{}

func DashboardRegister(group *gin.RouterGroup) {
	service := &DashboardController{}
	group.GET("/panel_group_data", service.PanelGroupData)
	group.GET("/flow_stat", service.FlowStat)
	group.GET("/service_stat", service.ServiceStat)
}

// wait for test: 指标统计
// PanelGroupData godoc
// @Summary 指标统计
// @Description 指标统计
// @Tags 首页大盘
// @ID /dashboard/panel_group_data
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.PanelGroupDataOutput} "success"
// @Router /dashboard/panel_group_data [get]
func (service *DashboardController) PanelGroupData(c *gin.Context) {
	tx := lib.GetMysqlGorm()
	serviceInfo := &dao.ServiceInfo{}
	_, serviceNum, err := serviceInfo.PageList(tx, &dto.ServiceListInput{PageSize: 1, PageNo: 1})
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	app := &dao.App{}
	_, appNum, err := app.APPList(tx, &dto.APPListInput{PageNo: 1, PageSize: 1})
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	// counter, err := public.FlowCounterHandler.GetCounter(public.FlowTotal)
	// if err != nil {
	// 	middleware.ResponseError(c, 2003, err)
	// 	return
	// }
	Qpd, Qps := 0, 0
	out := &dto.PanelGroupDataOutput{
		ServiceNum:      serviceNum,
		AppNum:          appNum,
		TodayRequestNum: int64(Qpd),
		CurrentQPS:      int64(Qps),
	}
	middleware.ResponseSuccess(c, out)
}

// wait for test: 服务类型占比
// ServiceStat godoc
// @Summary 服务统计
// @Description 服务统计
// @Tags 首页大盘
// @ID /dashboard/service_stat
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.DashServiceStatOutput} "success"
// @Router /dashboard/service_stat [get]
func (service *DashboardController) ServiceStat(c *gin.Context) {
	tx := lib.GetMysqlGorm()
	serviceInfo := &dao.ServiceInfo{}
	list, err := serviceInfo.GroupByLoadType(tx)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	legend := []string{}
	for index, item := range list {
		name, ok := global.LoadTypeMap[item.LoadType]
		if !ok {
			middleware.ResponseError(c, 2003, errors.New("load_type not found"))
			return
		}
		list[index].Name = name
		legend = append(legend, name)
	}
	out := &dto.DashServiceStatOutput{
		Legend: legend,
		Data:   list,
	}
	middleware.ResponseSuccess(c, out)
}

// wait for test: 流量统计
// FlowStat godoc
// @Summary 服务统计
// @Description 服务统计
// @Tags 首页大盘
// @ID /dashboard/flow_stat
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.ServiceStatOutput} "success"
// @Router /dashboard/flow_stat [get]
func (service *DashboardController) FlowStat(c *gin.Context) {
	// counter, err := public.FlowCounterHandler.GetCounter(public.FlowTotal)
	// if err != nil {
	// 	middleware.ResponseError(c, 2001, err)
	// 	return
	// }
	todayList := []int64{}
	currentTime := time.Now()
	for i := 0; i <= currentTime.Hour(); i++ {
		// dateTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), i, 0, 0, 0, lib.TimeLocation)
		// hourData, _ := counter.GetHourData(dateTime)
		// todayList = append(todayList, hourData)
		todayList = append(todayList, 0)
	}

	yesterdayList := []int64{}
	// yesterTime := currentTime.Add(-1 * time.Duration(time.Hour*24))
	for i := 0; i <= 23; i++ {
		// dateTime := time.Date(yesterTime.Year(), yesterTime.Month(), yesterTime.Day(), i, 0, 0, 0, lib.TimeLocation)
		// hourData, _ := counter.GetHourData(dateTime)
		// yesterdayList = append(yesterdayList, hourData)
		yesterdayList = append(yesterdayList, 0)
	}
	middleware.ResponseSuccess(c, &dto.ServiceStatOutput{
		Today:     todayList,
		Yesterday: yesterdayList,
	})
}
