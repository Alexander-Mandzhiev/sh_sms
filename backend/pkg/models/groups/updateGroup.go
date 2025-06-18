package groups_models

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/private_school"
	"strings"

	"github.com/google/uuid"
)

type UpdateGroup struct {
	PublicID  uuid.UUID
	ClientID  uuid.UUID
	Name      string
	CuratorID *uuid.UUID
}

func UpdateGroupFromProto(req *private_school_v1.UpdateGroupRequest) (*UpdateGroup, error) {
	publicID, err := utils.ValidateStringAndReturnUUID(req.GetId())
	if err != nil {
		return nil, ErrInvalidGroupID
	}

	clientID, err := utils.ValidateStringAndReturnUUID(req.GetClientId())
	if err != nil {
		return nil, ErrInvalidClientID
	}

	name := strings.TrimSpace(req.GetName())
	if name == "" {
		return nil, ErrGroupNameRequired
	}
	if len(name) > MaxGroupName {
		return nil, ErrGroupNameTooLong
	}

	update := &UpdateGroup{
		PublicID: publicID,
		ClientID: clientID,
		Name:     name,
	}

	if req.CuratorId == "" {
		update.CuratorID = nil
	} else {
		curatorID, err := utils.ValidateStringAndReturnUUID(req.CuratorId)
		if err != nil {
			return nil, ErrInvalidCuratorID
		}
		update.CuratorID = &curatorID
	}

	return update, nil
}

func (ug *UpdateGroup) ToProto() *private_school_v1.UpdateGroupRequest {
	var curatorIDStr string
	if ug.CuratorID != nil {
		curatorIDStr = ug.CuratorID.String()
	} else {
		curatorIDStr = ""
	}

	return &private_school_v1.UpdateGroupRequest{
		Id:        ug.PublicID.String(),
		ClientId:  ug.ClientID.String(),
		Name:      ug.Name,
		CuratorId: curatorIDStr,
	}
}
