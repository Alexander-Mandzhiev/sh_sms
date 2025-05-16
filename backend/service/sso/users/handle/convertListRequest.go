package handle

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/sso/users"
	"backend/service/sso/models"
)

func (s *serverAPI) convertListRequest(req *users.ListRequest) (*models.ListRequest, error) {
	clientID, err := utils.ValidateAndReturnUUID(req.GetClientId())
	if err != nil {
		return nil, s.convertError(err)
	}

	listReq := &models.ListRequest{
		Page:     int(req.Page),
		Count:    int(req.Count),
		ClientID: &clientID,
	}

	if req.GetEmailFilter() != "" {
		emailFilter := req.GetEmailFilter()
		listReq.EmailFilter = &emailFilter
	}

	if req.GetPhoneFilter() != "" {
		phoneFilter := req.GetPhoneFilter()
		listReq.PhoneFilter = &phoneFilter
	}

	if req.ActiveOnly != nil {
		listReq.ActiveOnly = req.ActiveOnly
	}

	return listReq, nil
}
