package forwardedPortCheck

import (
	"errors"
	"github.com/deqodelabs/IaaC/stackconfig/pb"
)

func CustomForwardedPortValidation(forwardedPorts []*pb.ForwardedPort) error{
	if forwardedPorts == nil{
		return nil
	}
	for _, forwardedPort := range forwardedPorts{
		if forwardedPort.HostPort >= 10400 && forwardedPort.HostPort <= 10500 {
			return errors.New("host port should not be in range [10400, 10500]")
		}
		if forwardedPort.ContainerPort >= 10400 && forwardedPort.ContainerPort <= 10500 {
			return errors.New("container port should not be in range [10400, 10500]")
		}
	}
	return nil
}
