package register

import (
	"context"
	"io"

	"github.com/google/wire"
	"github.com/solo-io/service-mesh-hub/cli/pkg/cliconstants"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common"
	common_config "github.com/solo-io/service-mesh-hub/cli/pkg/common/config"
	"github.com/solo-io/service-mesh-hub/cli/pkg/options"
	"github.com/spf13/cobra"
)

type RegistrationCmd *cobra.Command

var RegistrationSet = wire.NewSet(
	ClusterRegistrationCmd,
)

func ClusterRegistrationCmd(
	ctx context.Context,
	kubeClientsFactory common.KubeClientsFactory,
	clientsFactory common.ClientsFactory,
	opts *options.Options,
	out io.Writer,
	kubeLoader common_config.KubeLoader,
) RegistrationCmd {

	register := &cobra.Command{
		Use:   cliconstants.ClusterRegisterCommand.Use,
		Short: cliconstants.ClusterRegisterCommand.Short,
		Long:  cliconstants.ClusterRegisterCommand.Long,
		RunE: func(cmd *cobra.Command, args []string) error {
			binaryName := common.GetBinaryName(cmd)
			return RegisterCluster(ctx, binaryName, cmd.Flags(), clientsFactory, kubeClientsFactory, opts, out, kubeLoader)
		},
	}

	options.AddClusterRegisterFlags(register, opts)
	return register
}
