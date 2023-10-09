package entity

type Paginate struct {
	List      interface{} `json:"list"`
	TotalPage float64     `json:"total_page"`
	TotalData int64       `json:"total_data"`
	Page      int64       `json:"page"`
	Limit     int64       `json:"limit"`
}
