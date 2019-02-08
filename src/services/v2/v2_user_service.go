package v2

import (
	"../../objects"
	"../../repositories/v2"
	"github.com/jinzhu/copier"
)

type V2UserService struct {
	request        objects.V2UserObjectResponse
	userRepository v2.V2UserRepository
}

func V2UserServiceHandler() (V2UserService) {
	service := V2UserService{
		userRepository: v2.V2UserRepositoryHandler(),
	}
	return service
}

func (service *V2UserService) UpdateById(id int, requestObject objects.V2UserObjectRequest) (objects.V2UserObjectResponse, error) {

	user, err := service.userRepository.UpdateById(id, requestObject)
	if nil != err {
		return objects.V2UserObjectResponse{}, err
	}

	result := objects.V2UserObjectResponse{}
	copier.Copy(&result, &user)

	return result, nil

}
