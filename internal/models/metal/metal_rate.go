package metal

import (
	"fmt"
	"kcloudb1/internal/config"
	"time"
)

type MetalRate struct {
	ID              int64     `json:"ID"`
	MetalID         int64     `json:"metal_id"`
	Rate            float32   `json:"rate"`
	Date            time.Time `json:"date"`
	ChangePercent1D float32   `json:"change_percent_1d"`
	CreatedAt       time.Time `json:"created_at"`
}

func (MetalRate) TableName() string {
	return "one_metal_rate"
}

func (m *MetalRate) Create() error {
	return config.DB.Create(m).Error
}

func (m *MetalRate) Update() error {
	return config.DB.Updates(m).Error
}

func (m *MetalRate) Delete() error {
	return config.DB.Delete(m).Error
}

func (m *MetalRate) Get() error {
	return config.DB.First(m, m.ID).Error
}

func (m *MetalRate) LastByMetalID(metalID any) error {

	return config.DB.Order("created_at desc").First(&m, "metal_id = ?", metalID).Error
}

func (m *MetalRate) GetMetalRateByStartToEnd(metalID, startDate, endDate, order any) ([]MetalRate, error) {
	var metalRates []MetalRate
	err := config.DB.Order(fmt.Sprintf("date %v", order)).Where("metal_id = ? AND date BETWEEN to_date(?, 'YYYY-MM-DD') AND to_date(?, 'YYYY-MM-DD')", metalID, startDate, endDate).Find(&metalRates).Error
	return metalRates, err
}

func (m *MetalRate) GetMetalRateByKey(metalID, key, order any) ([]MetalRate, error) {
	var metalRates []MetalRate
	qry := config.DB.Order(fmt.Sprintf("date %v", order)).Find(&metalRates)
	if key == "week" {
		err := qry.Where("metal_id = ? AND date BETWEEN to_date(?, 'YYYY-MM-DD') AND to_date(?, 'YYYY-MM-DD')", metalID, time.Now().AddDate(0, 0, -7).Local().Format("2006-01-02"), time.Now().Format("2006-01-02")).Find(&metalRates).Error
		return metalRates, err
	} else if key == "month" {
		err := qry.Where("metal_id = ? AND date BETWEEN to_date(?, 'YYYY-MM-DD') AND to_date(?, 'YYYY-MM-DD')", metalID, time.Now().AddDate(0, -1, 0).Format("2006-01-02"), time.Now().Format("2006-01-02")).Find(&metalRates).Error
		return metalRates, err
	} else if key == "year" {
		err := qry.Where("metal_id = ? AND date BETWEEN to_date(?, 'YYYY-MM-DD') AND to_date(?, 'YYYY-MM-DD')", metalID, time.Now().AddDate(-1, 0, 0).Format("2006-01-02"), time.Now().Format("2006-01-02")).Find(&metalRates).Error
		return metalRates, err
	} else if key == "today" {
		err := qry.Where("metal_id = ? AND date = to_date(?, 'YYYY-MM-DD') ", metalID, time.Now().Format("2006-01-02")).Find(&metalRates).Error
		if err != nil {
			return metalRates, err
		} else {
			if len(metalRates) == 0 {
				return m.GetMetalRateByKey(metalID, "last", order)
			} else {
				return metalRates, nil
			}
		}

	} else if key == "last" {
		err := config.DB.Order(fmt.Sprintf("date %v", "desc")).First(&metalRates, "metal_id = ?", metalID).Error
		return metalRates, err
	}
	return metalRates, nil
}

func (m *MetalRate) GetAll() ([]MetalRate, error) {
	var metalRates []MetalRate
	err := config.DB.Find(&metalRates).Error
	return metalRates, err
}
