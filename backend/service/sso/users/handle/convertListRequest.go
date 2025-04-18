package handle

import (
	"backend/protos/gen/go/sso/users"
	"backend/service/constants"
	"backend/service/sso/models"
	"fmt"
	"github.com/google/uuid"
)

func (s *serverAPI) convertListRequest(req *users.ListRequest) (*models.ListRequest, error) {
	var clientID *uuid.UUID
	if req.ClientId != "" {
		parsed, err := uuid.Parse(req.ClientId)
		if err != nil {
			return nil, fmt.Errorf("%w: invalid client_id", constants.ErrInvalidArgument)
		}
		clientID = &parsed
	}

	listReq := &models.ListRequest{
		Page:     int(req.Page),
		Count:    int(req.Count),
		ClientID: clientID,
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
