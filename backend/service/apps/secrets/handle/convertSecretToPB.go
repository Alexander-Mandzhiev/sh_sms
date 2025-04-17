package handle

import (
	pb "backend/protos/gen/go/apps/secrets"
	"backend/service/apps/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *serverAPI) convertSecretToPB(secret *models.Secret) *pb.Secret {
	pbSecret := &pb.Secret{
		ClientId:      secret.ClientID,
		AppId:         int32(secret.AppID),
		SecretType:    secret.SecretType,
		CurrentSecret: secret.CurrentSecret,
		Algorithm:     secret.Algorithm,
		SecretVersion: int32(secret.SecretVersion),
		GeneratedAt:   timestamppb.New(secret.GeneratedAt),
	}

	if secret.RevokedAt != nil {
		pbSecret.RevokedAt = timestamppb.New(*secret.RevokedAt)
	}

	return pbSecret
}

func (s *serverAPI) convertSecretsToPB(secrets []*models.Secret) []*pb.Secret {
	pbSecrets := make([]*pb.Secret, 0, len(secrets))
	for _, secr := range secrets {
		pbSecrets = append(pbSecrets, s.convertSecretToPB(secr))
	}
	return pbSecrets
}
