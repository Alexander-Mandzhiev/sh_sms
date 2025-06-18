package groups_models

import "backend/protos/gen/go/private_school"

type GroupListResponse struct {
	Groups     []*Group
	NextCursor int64 // internal_id последней записи (0 если конец)
}

func (gl *GroupListResponse) ToProto() *private_school_v1.ListGroupsResponse {
	groupsProto := make([]*private_school_v1.GroupResponse, 0, len(gl.Groups))
	for _, group := range gl.Groups {
		groupsProto = append(groupsProto, group.ToProto())
	}

	return &private_school_v1.ListGroupsResponse{
		Groups:     groupsProto,
		NextCursor: gl.NextCursor,
	}
}

func GroupListFromProto(resp *private_school_v1.ListGroupsResponse) (*GroupListResponse, error) {
	groupList := &GroupListResponse{
		NextCursor: resp.GetNextCursor(),
	}
	for _, groupProto := range resp.GetGroups() {
		group, err := GroupFromProto(groupProto)
		if err != nil {
			return nil, err
		}
		groupList.Groups = append(groupList.Groups, group)
	}

	return groupList, nil
}
