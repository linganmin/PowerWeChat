package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseBroadcastGetFollowers struct {
	*response.ResponseMiniProgram

	Followers []*power.HashMap `json:"followers"`
	PageBreak string           `json:"page_break,omitempty"`
}
