package configuration

import "github.com/Okestro-Community-Dev/go-ansible/v2/pkg/execute"

type ExecutorEnvVarSetter interface {
	execute.Executor
	AddEnvVar(key, value string)
}
