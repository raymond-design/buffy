package users

import (
	"context"
	"sync"

	"github.com/gofrs/uuid"
	userspb "github.com/raymond-design/buffy/proto/users/v1"
)

type Server struct {
	users []*userspb.User
	mutex sync.Mutex

	*userspb.UnimplementedUserServiceServer
}

func (s *Server) AddUser(ctx context.Context, req *userspb.AddUserRequest) (*userspb.AddUserResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	user := &userspb.User{
		UserName: req.UserName,
		Id:       uuid.Must(uuid.NewV4()).String(),
	}

	s.users = append(s.users, user)

	return &userspb.AddUserResponse{
		User: user,
	}, nil
}
