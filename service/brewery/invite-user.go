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
	dbBrewery, err := iU.BreweryRepository.GetByName(invite.BreweryName)
	if err != nil {
		return nil, fmt.Errorf("Can't find brewery")
	}

	invitingUserRole, err := iU.BreweryRepository.GetUserRole(invite.InvitingUserId, dbBrewery.ID)
	if err != nil {
		return nil, err
	}

	invitedUser, err := iU.UserRepository.GetByEmail(invite.InvitedUserEmail)
	if err != nil {
		return nil, fmt.Errorf("Can't find invited user")
	}

	if (invite.InvitingUserId == invitedUser.ID) {
		return nil, fmt.Errorf("You can't invite yourself")
	}

	if (invitingUserRole != "creator" && invitingUserRole != "admin") {
		return nil, fmt.Errorf("Inviting user dont have permission")
	}

	if (invite.Role == "creator") {
		return nil, fmt.Errorf("Can't invite to be creator")
	}

	repoInvite := brewery.InviteUserRepoInputDto{
		InvitingUserId: invite.InvitingUserId,
		InvitedUserId: invitedUser.ID,
		Role: invite.Role,
		BreweryId: dbBrewery.ID,
	}
	
	createdInvite, err := iU.BreweryRepository.InviteUserToBrewery(repoInvite)
	if err != nil {
		return nil, err
	}

	return createdInvite, nil
}