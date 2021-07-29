package strexec

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type StrexecSuite struct {
	suite.Suite
}

func (suite *StrexecSuite) TestEchoOneArg() {
	cmd := "echo arg"
	execCmd, err := Command(cmd)
	suite.Assert().NoError(err)
	out, err := execCmd.CombinedOutput()
	suite.Assert().NoError(err)
	expected := "arg\n"
	actual := string(out)
	suite.Assert().Equal(expected, actual)
}

func (suite *StrexecSuite) TestEchoTwoArgs() {
	cmd := "echo two args"
	execCmd, err := Command(cmd)
	suite.Assert().NoError(err)
	out, err := execCmd.CombinedOutput()
	suite.Assert().NoError(err)
	expected := "two args\n"
	actual := string(out)
	suite.Assert().Equal(expected, actual)
}

func (suite *StrexecSuite) TestEchoWithFlag() {
	cmd := "echo -n two args"
	execCmd, err := Command(cmd)
	suite.Assert().NoError(err)
	out, err := execCmd.CombinedOutput()
	suite.Assert().NoError(err)
	expected := "two args"
	actual := string(out)
	suite.Assert().Equal(expected, actual)
}

func (suite *StrexecSuite) TestHistory() {
	cmd := "history"
	execCmd, err := Command(cmd)
	suite.Assert().NoError(err)
	out, err := execCmd.CombinedOutput()
	suite.Assert().NoError(err)
	expected := ""
	actual := string(out)
	suite.Assert().Equal(expected, actual)
}

func TestStrexecSuite(t *testing.T) {
	suite.Run(t, new(StrexecSuite))
}
