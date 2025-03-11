package temporal

import (
	// "errors"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func Test_SuccessfulWorkflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	env.OnActivity(FetchIssuersActivity, mock.Anything).Return(FetchIssuersActivityResponse{}, nil)
	env.ExecuteWorkflow(Scheduler)

	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

}

func Test_UnsuccessfulWorkflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	env.OnActivity(FetchIssuersActivity, mock.Anything).Return(FetchIssuersActivityResponse{}, errors.New("error"))
	env.ExecuteWorkflow(Scheduler)

	require.True(t, env.IsWorkflowCompleted())
	require.Error(t, env.GetWorkflowError())

}

