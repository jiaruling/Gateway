package controller

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jiaruling/Gateway/dao"
	"github.com/jiaruling/Gateway/dto"
	"github.com/jiaruling/Gateway/middleware"
	"github.com/jiaruling/Gateway/public"
	"github.com/jiaruling/golang_utils/lib"
	// "time"
	// "github.com/gin-gonic/gin"
	// "github.com/jiaruling/Gateway/dao"
	// "github.com/jiaruling/Gateway/dto"
	// "github.com/jiaruling/Gateway/middleware"
	// "github.com/jiaruling/Gateway/public"
	// "github.com/jiaruling/golang_utils/lib"
	// "github.com/pkg/errors"
)

type APPController struct {
}

// APPControllerRegister admin路由注册
func APPRegister(router *gin.RouterGroup) {
	admin := APPController{}
	router.GET("/app_list", admin.APPList)
	router.GET("/app_detail", admin.APPDetail)
	router.GET("/app_stat", admin.AppStatistics)
	router.DELETE("/app_delete", admin.APPDelete)
	router.POST("/app_add", admin.AppAdd)
	router.PUT("/app_update", admin.AppUpdate)
}

// wait for test: 租户列表
// APPList godoc
// @Summary 租户列表
// @Description 租户列表
// @Tags 租户管理
// @ID /app/app_list
// @Accept  json
// @Produce  json
// @Param info query string false "关键词"
// @Param page_size query string true "每页多少条"
// @Param page_no query string true "页码"
// @Success 200 {object} middleware.Response{data=dto.APPListOutput} "success"
// @Router /app/app_list [get]
func (admin *APPController) APPList(c *gin.Context) {
	params := &dto.APPListInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	info := &dao.App{}
	list, total, err := info.APPList(lib.GetMysqlGorm(), params)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	outputList := []dto.APPListItemOutput{}
	for _, item := range list {
		// appCounter, err := public.FlowCounterHandler.GetCounter(public.FlowAppPrefix + item.AppID)
		// if err != nil {
		// 	middleware.ResponseError(c, 2003, err)
		// 	c.Abort()
		// 	return
		// }
		Qpd, Qps := 0, 0
		outputList = append(outputList, dto.APPListItemOutput{
			ID:       item.ID,
			AppID:    item.AppID,
			Name:     item.Name,
			Secret:   item.Secret,
			WhiteIPS: item.WhiteIPS,
			Qpd:      item.Qpd,
			Qps:      item.Qps,
			RealQpd:  int64(Qpd),
			RealQps:  int64(Qps),
		})
	}
	output := dto.APPListOutput{
		List:  outputList,
		Total: total,
	}
	middleware.ResponseSuccess(c, output)
	return
}

// wait for test: 租户详情
// APPDetail godoc
// @Summary 租户详情
// @Description 租户详情
// @Tags 租户管理
// @ID /app/app_detail
// @Accept  json
// @Produce  json
// @Param id query string true "租户ID"
// @Success 200 {object} middleware.Response{data=dao.App} "success"
// @Router /app/app_detail [get]
func (admin *APPController) APPDetail(c *gin.Context) {
	params := &dto.APPDetailInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	search := &dao.App{
		ID: params.ID,
	}
	detail, err := search.Find(lib.GetMysqlGorm(), search)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	middleware.ResponseSuccess(c, detail)
	return
}

// wait for test: 租户删除
// APPDelete godoc
// @Summary 租户删除
// @Description 租户删除
// @Tags 租户管理
// @ID /app/app_delete
// @Accept  json
// @Produce  json
// @Param id query string true "租户ID"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /app/app_delete [delete]
func (admin *APPController) APPDelete(c *gin.Context) {
	params := &dto.APPDetailInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	search := &dao.App{
		ID: params.ID,
	}
	info, err := search.Find(lib.GetMysqlGorm(), search)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	info.IsDelete = 1
	if err := info.Save(lib.GetMysqlGorm()); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}

// wait for test: 租户添加
// AppAdd godoc
// @Summary 租户添加
// @Description 租户添加
// @Tags 租户管理
// @ID /app/app_add
// @Accept  json
// @Produce  json
// @Param body body dto.APPAddHttpInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /app/app_add [post]
func (admin *APPController) AppAdd(c *gin.Context) {
	params := &dto.APPAddHttpInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	//验证app_id是否被占用
	search := &dao.App{
		AppID: params.AppID,
	}
	if _, err := search.Find(lib.GetMysqlGorm(), search); err == nil {
		middleware.ResponseError(c, 2002, errors.New("租户ID被占用，请重新输入"))
		return
	}
	if params.Secret == "" {
		params.Secret = public.MD5(params.AppID)
	}
	tx := lib.GetMysqlGorm()
	info := &dao.App{
		AppID:    params.AppID,
		Name:     params.Name,
		Secret:   params.Secret,
		WhiteIPS: params.WhiteIPS,
		Qps:      params.Qps,
		Qpd:      params.Qpd,
	}
	if err := info.Save(tx); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}

// wait for test: 租户更新
// AppUpdate godoc
// @Summary 租户更新
// @Description 租户更新
// @Tags 租户管理
// @ID /app/app_update
// @Accept  json
// @Produce  json
// @Param body body dto.APPUpdateHttpInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /app/app_update [put]
func (admin *APPController) AppUpdate(c *gin.Context) {
	params := &dto.APPUpdateHttpInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	search := &dao.App{
		ID: params.ID,
	}
	info, err := search.Find(lib.GetMysqlGorm(), search)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	if params.Secret == "" {
		params.Secret = public.MD5(params.AppID)
	}
	info.Name = params.Name
	info.Secret = params.Secret
	info.WhiteIPS = params.WhiteIPS
	info.Qps = params.Qps
	info.Qpd = params.Qpd
	if err := info.Save(lib.GetMysqlGorm()); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}

// wait for test: 租户统计
// AppStatistics godoc
// @Summary 租户统计
// @Description 租户统计
// @Tags 租户管理
// @ID /app/app_stat
// @Accept  json
// @Produce  json
// @Param id query string true "租户ID"
// @Success 200 {object} middleware.Response{data=dto.StatisticsOutput} "success"
// @Router /app/app_stat [get]
func (admin *APPController) AppStatistics(c *gin.Context) {
	params := &dto.APPDetailInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	// search := &dao.App{
	// 	ID: params.ID,
	// }
	// detail, err := search.Find(lib.GetMysqlGorm(), search)
	// if err != nil {
	// 	middleware.ResponseError(c, 2002, err)
	// 	return
	// }

	//今日流量全天小时级访问统计
	todayStat := []int64{}
	// counter, err := public.FlowCounterHandler.GetCounter(public.FlowAppPrefix + detail.AppID)
	// if err != nil {
	// 	middleware.ResponseError(c, 2002, err)
	// 	c.Abort()
	// 	return
	// }
	// currentTime := time.Now()
	for i := 0; i <= time.Now().In(lib.TimeLocation).Hour(); i++ {
		// dateTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), i, 0, 0, 0, lib.TimeLocation)
		// hourData, _ := counter.GetHourData(dateTime)
		// todayStat = append(todayStat, hourData)
		todayStat = append(todayStat, 0)
	}

	//昨日流量全天小时级访问统计
	yesterdayStat := []int64{}
	// yesterTime := currentTime.Add(-1 * time.Duration(time.Hour*24))
	for i := 0; i <= 23; i++ {
		// dateTime := time.Date(yesterTime.Year(), yesterTime.Month(), yesterTime.Day(), i, 0, 0, 0, lib.TimeLocation)
		// hourData, _ := counter.GetHourData(dateTime)
		// yesterdayStat = append(yesterdayStat, hourData)
		yesterdayStat = append(yesterdayStat, 0)
	}
	stat := dto.StatisticsOutput{
		Today:     todayStat,
		Yesterday: yesterdayStat,
	}
	middleware.ResponseSuccess(c, stat)
	return
}
