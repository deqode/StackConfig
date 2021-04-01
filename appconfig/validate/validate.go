package validate

import (
	"errors"
	"github.com/deqodelabs/IaaC/appconfig/pb"
	"github.com/deqodelabs/IaaC/appconfig/validate/forwardedPortCheck"
	"github.com/deqodelabs/IaaC/appconfig/validate/resourceCheck"
	"github.com/deqodelabs/IaaC/appconfig/validate/runtimeCheck"
)

func CustomValidate(app *pb.AppConfig) error {
	if app == nil {
		return errors.New("app is nil")
	}
	for _ , v := range app.Services {
		err := forwardedPortCheck.CustomForwardedPortValidation(v.Network.ForwardedPorts)
		if err != nil{
			return errors.New("invalid forwarded ports")
		}
		err = resourceCheck.CustomResourceValidation(v.Resources)
		if err != nil{
			return errors.New("invalid resource")
		}
		err = runtimeCheck.CustomValidateRuntime(v)
		if err != nil{
			return errors.New("runtime config unavailable")
		}
	}
	return nil
}