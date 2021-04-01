package resourceCheck

import(
	"errors"
	"github.com/deqodelabs/IaaC/appconfig/pb"
)

func CustomResourceValidation(resource *pb.Resource) error{
	if resource.Cpu != 1 && (resource.Cpu % 2) != 0 {
		return errors.New("runtime config not valid")
	}
	return nil
}