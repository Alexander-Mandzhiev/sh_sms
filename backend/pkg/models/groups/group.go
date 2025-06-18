package groups_models

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/private_school"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"time"
)

type Group struct {
	InternalID int64
	PublicID   uuid.UUID
	ClientID   uuid.UUID
	Name       string
	CuratorID  *uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (g *Group) ToProto() *private_school_v1.GroupResponse {
	var curatorID *string
	if g.CuratorID != nil {
		curatorStr := g.CuratorID.String()
		curatorID = &curatorStr
	}

	return &private_school_v1.GroupResponse{
		Id:        g.PublicID.String(),
		ClientId:  g.ClientID.String(),
		Name:      g.Name,
		CuratorId: curatorID,
		CreatedAt: timestamppb.New(g.CreatedAt),
		UpdatedAt: timestamppb.New(g.UpdatedAt),
	}
}

func GroupFromProto(req *private_school_v1.GroupResponse) (*Group, error) {
	clientID, err := utils.ValidateStringAndReturnUUID(req.GetClientId())
	if err != nil {
		return nil, ErrInvalidClientID
	}
	publicID, err := utils.ValidateStringAndReturnUUID(req.GetId())
	if err != nil {
		return nil, ErrInvalidGroupID
	}

	name := strings.TrimSpace(req.GetName())
	if name == "" {
		return nil, ErrGroupNameRequired
	}
	if len(name) > MaxGroupName {
		return nil, ErrGroupNameTooLong
	}

	group := &Group{
		PublicID:  publicID,
		ClientID:  clientID,
		Name:      name,
		CreatedAt: req.CreatedAt.AsTime(),
		UpdatedAt: req.UpdatedAt.AsTime(),
	}

	if req.CuratorId != nil && *req.CuratorId != "" {
		curatorID, err := utils.ValidateStringAndReturnUUID(*req.CuratorId)
		if err != nil {
			return nil, ErrInvalidCuratorID
		}
		group.CuratorID = &curatorID
	} else {
		group.CuratorID = nil
	}

	return group, nil
}
