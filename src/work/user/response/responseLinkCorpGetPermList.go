package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseLinkCorpGetPermList struct {
	*response.ResponseWork

	UserIDs       []string `json:"userids"`
	DepartmentIDs []string `json:"department_ids"`
}
