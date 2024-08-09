package request

type ProGrabLog struct {
	UUID string `json:"uuid" form:"uuid" validate:"required"`
}
