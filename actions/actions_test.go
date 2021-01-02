package actions

import (
	"testing"

	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/suite"
)

type ActionSuite struct {
	*suite.Action
}

func (as *ActionSuite) T() *testing.T {
	panic("implement me")
}

func (as *ActionSuite) SetT(t *testing.T) {
	panic("implement me")
}

func Test_ActionSuite(t *testing.T) {
	action, err := suite.NewActionWithFixtures(App(), packr.New("Test_ActionSuite", "../fixtures"))
	if err != nil {
		t.Fatal(err)
	}

	as := &ActionSuite{
		Action: action,
	}
	suite.Run(t, as)
}
