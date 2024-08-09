package response

type ProGrabRecord struct {
	CreatedTime string `json:"created_time"`
	UserName    string `json:"user_name" `
	CategoryNum int    `json:"category_num"`
	GoNum       int    `json:"go_num"`
	PageNum     int    `json:"page_num"`
	RunTime     string `json:"run_time"`
	GrabNum     int64  `json:"grab_num"`
	Uuid        string `json:"uuid"`
}
