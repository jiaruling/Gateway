package dao

import (
	"sync"
	"time"

	"github.com/jiaruling/Gateway/dto"
	"github.com/jiaruling/golang_utils/lib"
	"gorm.io/gorm"
)

type App struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	AppID     string    `json:"app_id" gorm:"column:app_id" description:"租户id"`
	Name      string    `json:"name" gorm:"column:name" description:"租户名称"`
	Secret    string    `json:"secret" gorm:"column:secret" description:"密钥"`
	WhiteIPS  string    `json:"white_ips" gorm:"column:white_ips" description:"ip白名单，支持前缀匹配"`
	Qpd       int64     `json:"qpd" gorm:"column:qpd" description:"日请求量限制"`
	Qps       int64     `json:"qps" gorm:"column:qps" description:"每秒请求量限制"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"添加时间	"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	IsDelete  int8      `json:"is_delete" gorm:"column:is_delete" description:"是否已删除；0：否；1：是"`
}

func (t *App) TableName() string {
	return "gateway_app"
}

func (t *App) Find(tx *gorm.DB, search *App) (*App, error) {
	model := &App{}
	err := tx.Where(search).First(model).Error
	return model, err
}

func (t *App) Save(tx *gorm.DB) error {
	if err := tx.Save(t).Error; err != nil {
		return err
	}
	return nil
}

func (t *App) APPList(tx *gorm.DB, params *dto.APPListInput) ([]App, int64, error) {
	var list []App
	var count int64
	pageNo := params.PageNo
	pageSize := params.PageSize

	//limit offset,pagesize
	offset := (pageNo - 1) * pageSize
	query := tx.Table(t.TableName()).Select("*")
	query = query.Where("is_delete=?", 0)
	errCount := query.Count(&count).Error
	if errCount != nil {
		return nil, 0, errCount
	}
	if params.Info != "" {
		query = query.Where(" (name like ? or app_id like ?)", "%"+params.Info+"%", "%"+params.Info+"%")
	}
	err := query.Limit(pageSize).Offset(offset).Order("id desc").Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	return list, count, nil
}

var AppManagerHandler *AppManager

func init() {
	AppManagerHandler = NewAppManager()
}

type AppManager struct {
	AppMap   map[string]*App
	AppSlice []*App
	Locker   sync.RWMutex
	init     sync.Once
	err      error
}

func NewAppManager() *AppManager {
	return &AppManager{
		AppMap:   map[string]*App{},
		AppSlice: []*App{},
		Locker:   sync.RWMutex{},
		init:     sync.Once{},
	}
}

func (s *AppManager) GetAppList() []*App {
	return s.AppSlice
}

func (s *AppManager) LoadOnce() error {
	s.init.Do(func() {
		appInfo := &App{}
		tx := lib.GetMysqlGorm()
		params := &dto.APPListInput{PageNo: 1, PageSize: 99999}
		list, _, err := appInfo.APPList(tx, params)
		if err != nil {
			s.err = err
			return
		}
		s.Locker.Lock()
		defer s.Locker.Unlock()
		for _, listItem := range list {
			tmpItem := listItem
			s.AppMap[listItem.AppID] = &tmpItem
			s.AppSlice = append(s.AppSlice, &tmpItem)
		}
	})
	return s.err
}

func (s *AppManager) Add(item *App) {
	s.Locker.Lock()
	defer s.Locker.Unlock()
	s.AppMap[item.AppID] = item
	s.AppSlice = append(s.AppSlice, item)
	return
}

func (s *AppManager) Update(item *App) {
	s.Locker.Lock()
	defer s.Locker.Unlock()
	s.AppMap[item.AppID] = item
	for i, appItem := range s.AppSlice {
		if appItem.AppID == item.AppID {
			s.AppSlice[i] = item
			break
		}
	}
	return
}

func (s *AppManager) Delete(item *App) {
	s.Locker.Lock()
	defer s.Locker.Unlock()
	delete(s.AppMap, item.AppID)
	len := len(s.AppSlice)
	for i, app := range s.AppSlice {
		if app.AppID == item.AppID {
			s.deleteSlice(i, len)
			break
		}
	}
	return
}

func (s *AppManager) deleteSlice(i, len int) {
	if len == 1 {
		s.AppSlice = []*App{}
		return
	}
	if i == 0 {
		s.AppSlice = s.AppSlice[1:]
		return
	}
	if i == len-1 {
		s.AppSlice = s.AppSlice[:len-1]
		return
	}
	s.AppSlice = append(s.AppSlice[:i], s.AppSlice[i+1:]...)
	return
}
