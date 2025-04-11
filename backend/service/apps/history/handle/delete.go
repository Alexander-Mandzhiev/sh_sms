package handle

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) Create(ctx context.Context, req *sentences.CreateSentenceRequest) (*sentences.SentenceResponse, error) {
	resp, err := s.service.Create(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
