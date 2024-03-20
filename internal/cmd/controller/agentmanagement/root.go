package agentmanagement

import (
	"fmt"

	command "github.com/rancher/fleet/internal/cmd"
	"github.com/rancher/fleet/pkg/version"
	"github.com/spf13/cobra"
)

type AgentManagement struct {
	Kubeconfig       string `usage:"kubeconfig file"`
	Namespace        string `usage:"namespace to watch" env:"NAMESPACE"`
	DisableBootstrap bool   `usage:"disable local cluster components" name:"disable-bootstrap"`
}

func (a *AgentManagement) Run(cmd *cobra.Command, args []string) error {
	if a.Namespace == "" {
		return fmt.Errorf("--namespace or env NAMESPACE is required to be set")
	}
	return start(cmd.Context(), a.Kubeconfig, a.Namespace, a.DisableBootstrap)
}

func App() *cobra.Command {
	return command.Command(&AgentManagement{}, cobra.Command{
		Version: version.FriendlyVersion(),
		Use:     "agentmanagement",
	})
}
