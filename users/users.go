package users

import (
	"sync"
	"github.com/gofrc/uuid"
	userspb "github.com/raymond-design/buffy/proto/users/v1;users"
)
type Server struct {
	users []*userspb.User
	mutex sync.Mutex

	*userspb.UnimplementedUserSerivceServer
}

func (s *Server) AddUser(ctx context.Context, req *userspb.AddUserRequest) (*userspb.AddUserResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	user := &userspb.User{
		Username: req.user_name,
		Id: uuid.Must(uuid.NewV4().String()),
	}

	s.users = append(s.users, user)

	returnv &userspb.AddUserResponse{
		User: user,
	}, nil
}