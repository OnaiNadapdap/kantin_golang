package helper 

func Contains(slice []string, item string) bool {
    for _, s := range slice {
        if s == item {
            return true
        }
    }
    return false
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}
type ResponsePaginate struct {
	Meta      Meta        `json:"meta"`
	Data      interface{} `json:"data"`
	PerPage   int         `json:"per_page"`
	Page      int         `json:"page"`
	TotalData int64       `json:"total_data"`
	TotalPage int         `json:"total_page"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}