package v1

import (
	"context"

	"github.com/hatajoe/8am/app/interface/rpc/v1.0/protocol"
	"github.com/hatajoe/8am/app/usecase"
)

type userService struct {
	userUsecase usecase.UserUsecase
}

// Verify interface compliance
var _ protocol.UserServiceServer = (*userService)(nil)

func NewUserService(userUsecase usecase.UserUsecase) *userService {
	return &userService{
		userUsecase: userUsecase,
	}
}

// Handler here should only:
// - do validation so that usecase can receive the correct input
// - do post request processing (sanitize, etc..) and return response to caller

// Should deletegate actual work to ITicketUseCase.CreateTicket(ctx, &usecase.Ticket{})
// ticketUseCase will do business logics by calling:
//   - domain.TicketService concrete struct
//   - domain.TicketRepository interface
//
// TicketRepository is the interface to persistent layer implementations, concrete implementation
// will be injected later by DI container
// TicketService include simple business logic related to domain models like: not allow duplicated domain ticket id, not allowing start time > end time, ..

func (s *userService) ListUser(ctx context.Context, in *protocol.ListUserRequestType) (*protocol.ListUserResponseType, error) {
	users, err := s.userUsecase.ListUser()
	if err != nil {
		return nil, err
	}

	res := &protocol.ListUserResponseType{
		Users: toUser(users),
	}

	return res, nil
}

func (s *userService) RegisterUser(ctx context.Context, in *protocol.RegisterUserRequestType) (*protocol.RegisterUserResponseType, error) {
	if err := s.userUsecase.RegisterUser(in.GetEmail()); err != nil {
		return &protocol.RegisterUserResponseType{}, err
	}
	return &protocol.RegisterUserResponseType{}, nil
}

func toUser(users []*usecase.User) []*protocol.User {
	res := make([]*protocol.User, len(users))
	for i, user := range users {
		res[i] = &protocol.User{
			Id:    user.ID,
			Email: user.Email,
		}
	}
	return res
}
