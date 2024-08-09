package response

import "admin-gin/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
