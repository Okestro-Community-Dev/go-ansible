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

func TestDefaultStdoutCallbackExecute(t *testing.T) {
	t.Parallel()
	t.Run("Testing Default stdout callback execution", func(t *testing.T) {
		exec := execute.NewMockExecute()

		exec.On("WithOutput", mock.Anything).Return(exec)
		exec.On("AddEnvVar", configuration.AnsibleStdoutCallback, DefaultStdoutCallback)
		exec.On("Execute", mock.Anything).Return(nil)

		e := NewDefaultStdoutCallbackExecute(nil).WithExecutor(exec)
		err := e.Execute(context.TODO())

		assert.Nil(t, err)
		exec.AssertExpectations(t)
	})

	t.Run("Testing error on Default stdout callback when execute function returns an error", func(t *testing.T) {
		exec := execute.NewMockExecute()

		exec.On("WithOutput", mock.Anything).Return(exec)
		exec.On("AddEnvVar", configuration.AnsibleStdoutCallback, DefaultStdoutCallback)
		exec.On("Execute", mock.Anything).Return(errors.New("some error"))

		e := NewDefaultStdoutCallbackExecute(exec)
		err := e.Execute(context.TODO())

		assert.ErrorContains(t, err, "some error")
	})

	t.Run("Testing error on Default stdout callback when executor is not provided", func(t *testing.T) {
		e := NewDefaultStdoutCallbackExecute(nil)
		err := e.Execute(context.TODO())

		assert.ErrorContains(t, err, "DefaultStdoutCallbackExecute executor requires an executor")
	})
}
