package wire

import (
	"github.com/google/wire"
	zephyr_discovery "github.com/solo-io/service-mesh-hub/pkg/clients/zephyr/discovery"
	zephyr_networking "github.com/solo-io/service-mesh-hub/pkg/clients/zephyr/networking"
	"github.com/solo-io/service-mesh-hub/services/apiserver/pkg/server"
	"github.com/solo-io/service-mesh-hub/services/common/multicluster"
)

// just used to package everything up for wire
type ApiServerContext struct {
	MultiClusterDeps       multicluster.MultiClusterDependencies
	ManagementPlaneClients ManagementPlaneClients
	Server                 server.GrpcServer
}

var managementPlaneClientsSet = wire.NewSet(
	zephyr_discovery.NewMeshServiceClient,
	zephyr_discovery.NewMeshWorkloadClient,
	zephyr_discovery.NewMeshClient,
	zephyr_discovery.NewControllerRuntimeKubernetesClusterClient,
	zephyr_networking.NewVirtualMeshClient,
)

type ManagementPlaneClients struct {
	MeshServiceClient       zephyr_discovery.MeshServiceClient
	MeshWorkloadClient      zephyr_discovery.MeshWorkloadClient
	MeshClient              zephyr_discovery.MeshClient
	KubernetesClusterClient zephyr_discovery.KubernetesClusterClient
	VirtualMeshClient       zephyr_networking.VirtualMeshClient
}

func ApiServerContextProvider(
	meshServiceClient zephyr_discovery.MeshServiceClient,
	meshWorkloadClient zephyr_discovery.MeshWorkloadClient,
	meshClient zephyr_discovery.MeshClient,
	kubernetesClusterClient zephyr_discovery.KubernetesClusterClient,
	virtualMeshClient zephyr_networking.VirtualMeshClient,
	grpcServer server.GrpcServer,
	multiClusterDeps multicluster.MultiClusterDependencies,
) ApiServerContext {

	return ApiServerContext{
		ManagementPlaneClients: ManagementPlaneClients{
			MeshServiceClient:       meshServiceClient,
			MeshWorkloadClient:      meshWorkloadClient,
			MeshClient:              meshClient,
			KubernetesClusterClient: kubernetesClusterClient,
			VirtualMeshClient:       virtualMeshClient,
		},
		Server:           grpcServer,
		MultiClusterDeps: multiClusterDeps,
	}
}