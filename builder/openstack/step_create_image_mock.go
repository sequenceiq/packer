package openstack

import (
	"fmt"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
)

type stepCreateImageMock struct{}

func (s *stepCreateImageMock) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(Config)
	ui := state.Get("ui").(packer.Ui)

	// Create the image
	ui.Say(fmt.Sprintf("[MOCK] Creating the image: %s", config.ImageName))
	imageId := config.ImageName

	// Set the Image ID in the state
	ui.Message(fmt.Sprintf("Image: %s", imageId))
	state.Put("image", imageId)

	return multistep.ActionContinue
}

func (s *stepCreateImageMock) Cleanup(multistep.StateBag) {
	// No cleanup...
}
