package breweryService

import (
	"beer-production-api/entities/brewery"
	"beer-production-api/entities/user"
	"fmt"
)

type InviteUserToBrewery struct {
	BreweryRepository brewery.BreweryRepository
	UserRepository user.UserRepository
}



func NewInviteUserToBrewery(breweryRepositry brewery.BreweryRepository, userRepositry user.UserRepository) *InviteUserToBrewery {
	return &InviteUserToBrewery{BreweryRepository: breweryRepositry, UserRepository: userRepositry}
}

func (iU *InviteUserToBrewery) Execute(invite brewery.InviteUserInputDTO) (*brewery.Invite, error) {
	invitingUserRole, err := iU.BreweryRepository.GetUserRole(invite.InvitingUserId, invite.BreweryId)
	if err != nil {
		return nil, err
	}

	if (invite.InvitingUserId == invite.InvitedUserId) {
		return nil, fmt.Errorf("You can't invite yourself")
	}

	if (invitingUserRole != "creator" && invitingUserRole != "admin") {
		return nil, fmt.Errorf("Inviting user dont have permission")
	}

	if (invite.Role == "creator") {
		return nil, fmt.Errorf("Can't invite to be creator")
	}
	
	createdInvite, err := iU.BreweryRepository.InviteUserToBrewery(invite)
	if err != nil {
		return nil, err
	}

	return createdInvite, nil
}