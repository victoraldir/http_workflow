package file

import (
	"io/ioutil"

	"github.com/victoraldir/http-follower/internal/request/dto"
	"gopkg.in/yaml.v2"
)

func LoadWorkflowrequestFromYaml(filename string) (*dto.WorkflowRequest, error) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var workflow *dto.WorkflowRequest
	err = yaml.Unmarshal(yamlFile, &workflow)
	if err != nil {
		return nil, err
	}

	return workflow, nil
}
