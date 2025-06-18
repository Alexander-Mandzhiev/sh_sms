package groups_models

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/private_school"
	"strings"

	"github.com/google/uuid"
)

type CreateGroup struct {
	ClientID  uuid.UUID
	Name      string
	CuratorID *uuid.UUID
}

func CreateGroupFromProto(req *private_school_v1.CreateGroupRequest) (*CreateGroup, error) {
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

	group := &CreateGroup{
		ClientID: clientID,
		Name:     name,
	}

	if req.CuratorId != nil {
		if *req.CuratorId == "" {
			group.CuratorID = nil
		} else {
			curatorID, err := utils.ValidateStringAndReturnUUID(*req.CuratorId)
			if err != nil {
				return nil, ErrInvalidCuratorID
			}
			group.CuratorID = &curatorID
		}
	}

	return group, nil
}
func (cg *CreateGroup) ToProto() *private_school_v1.CreateGroupRequest {
	var curatorID *string
	if cg.CuratorID != nil {
		curatorStr := cg.CuratorID.String()
		curatorID = &curatorStr
	} else {
		curatorID = nil
	}

	return &private_school_v1.CreateGroupRequest{
		ClientId:  cg.ClientID.String(),
		Name:      cg.Name,
		CuratorId: curatorID,
	}
}
