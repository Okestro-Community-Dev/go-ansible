package stdoutcallback

import (
	"context"
	"errors"
	"testing"

	"github.com/Okestro-Community-Dev/go-ansible/v2/pkg/execute"
	"github.com/Okestro-Community-Dev/go-ansible/v2/pkg/execute/configuration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDenseStdoutCallbackExecute(t *testing.T) {
	t.Parallel()
	t.Run("Testing Dense stdout callback execution", func(t *testing.T) {
		exec := execute.NewMockExecute()

		exec.On("WithOutput", mock.Anything).Return(exec)
		exec.On("AddEnvVar", configuration.AnsibleStdoutCallback, DenseStdoutCallback)
		exec.On("Execute", mock.Anything).Return(nil)

		e := NewDenseStdoutCallbackExecute(nil).WithExecutor(exec)
		err := e.Execute(context.TODO())

		assert.Nil(t, err)
		exec.AssertExpectations(t)
	})

	t.Run("Testing error on Dense stdout callback when execute function returns an error", func(t *testing.T) {
		exec := execute.NewMockExecute()

		exec.On("WithOutput", mock.Anything).Return(exec)
		exec.On("AddEnvVar", configuration.AnsibleStdoutCallback, DenseStdoutCallback)
		exec.On("Execute", mock.Anything).Return(errors.New("some error"))

		e := NewDenseStdoutCallbackExecute(exec)
		err := e.Execute(context.TODO())

		assert.ErrorContains(t, err, "some error")
	})

	t.Run("Testing error on Dense stdout callback when executor is not provided", func(t *testing.T) {
		e := NewDenseStdoutCallbackExecute(nil)
		err := e.Execute(context.TODO())

		assert.ErrorContains(t, err, "DenseStdoutCallbackExecute executor requires an executor")
	})
}
