package temporal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.temporal.io/sdk/testsuite"
)

func TestFetchIssuersActivity(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()
	env.RegisterActivity(FetchIssuersActivity)

	val, err := env.ExecuteActivity(FetchIssuersActivity)
	var result FetchIssuersActivityResponse
	assert.NoError(t, val.Get(&result))
	assert.NoError(t, err)
}
