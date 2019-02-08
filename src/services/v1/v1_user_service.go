package v1

import (
	"../../objects"
	"../../repositories/v1"
	"github.com/jinzhu/copier"
)

type V1UserService struct {
	request        objects.V1UserObjectResponse
	userRepository v1.V1UserRepository
}

func V1UserServiceHandler() (V1UserService) {
	service := V1UserService{
		userRepository: v1.V1UserRepositoryHandler(),
	}
	return service
}

func (service *V1UserService) GetById(id int) (objects.V1UserObjectResponse, error) {
	user, err := service.userRepository.GetById(id)
	if nil != err {
		return objects.V1UserObjectResponse{}, err
	}
	result := objects.V1UserObjectResponse{}
	copier.Copy(&result, &user)
	return result, nil
}

func (service *V1UserService) UpdateById(id int, requestObject objects.V1UserObjectRequest) (objects.V1UserObjectResponse, error) {

	user, err := service.userRepository.UpdateById(id, requestObject)
	if nil != err {
		return objects.V1UserObjectResponse{}, err
	}

	result := objects.V1UserObjectResponse{}
	copier.Copy(&result, &user)

	return result, nil

}
