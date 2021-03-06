package install_istio

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/rotisserie/eris"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common"
	common_config "github.com/solo-io/service-mesh-hub/cli/pkg/common/config"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/kube"
	"github.com/solo-io/service-mesh-hub/cli/pkg/options"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/mesh/install/istio/operator"
	"github.com/solo-io/service-mesh-hub/pkg/common/docker"
)

var (
	ConflictingControlPlaneSettings   = eris.New("Cannot both use a pre-configured Mesh profile and provide an IstioOperator Custom Resource")
	FailedToParseControlPlaneSettings = func(err error) error {
		return eris.Wrap(err, "Failed to parse the provided IstioOperator resource")
	}
	FailedToParseControlPlaneWithProfile = func(err error, profile string) error {
		return eris.Wrapf(err, "Failed to parse the pre-configured IstioOperator with profile %s", profile)
	}
	FailedToWriteControlPlane = func(err error) error {
		return eris.Wrap(err, "Failed to write the provided IstioOperator resource")
	}
	TooManyControlPlaneResources = func(numResources int) error {
		return eris.Errorf("Expected the IstioOperator manifest to have only a single resource, found %d", numResources)
	}
	UnknownControlPlaneKind = func(kind string) error {
		return eris.Errorf("Expected the manifest to contain an IstioOperator, but found %s", kind)
	}
	ContextNotFound = func(contextName string) error {
		return eris.Errorf("Context '%s' not found in kubeconfig file", contextName)
	}

	istioControlPlaneKind = "IstioOperator"
)

type IstioInstaller interface {
	Install() error
}

func NewIstioInstaller(
	out io.Writer,
	in io.Reader,
	clientfactory common.ClientsFactory,
	opts *options.Options,
	kubeConfigPath,
	kubeContext string,
	kubeLoader common_config.KubeLoader,
	imageNameParser docker.ImageNameParser,
	fileReader common.FileReader,
) (IstioInstaller, error) {

	clients, err := clientfactory(opts)
	if err != nil {
		return nil, err
	}
	restClientGetter := kubeLoader.RESTClientGetter(kubeConfigPath, kubeContext)
	unstructuredKubeClient := clients.UnstructuredKubeClientFactory(restClientGetter)

	rawCfg, err := kubeLoader.GetRawConfigForContext(kubeConfigPath, kubeContext)
	if err != nil {
		return nil, err
	}

	contextName := kubeContext
	if contextName == "" {
		contextName = rawCfg.CurrentContext
	}
	context, ok := rawCfg.Contexts[contextName]
	if !ok {
		return nil, ContextNotFound(contextName)
	}
	clusterName := context.Cluster

	operatorManager := clients.IstioClients.OperatorManagerFactory(
		unstructuredKubeClient,
		clients.IstioClients.OperatorManifestBuilder,
		clients.DeploymentClient,
		imageNameParser,
		&opts.Mesh.Install.InstallationConfig,
	)

	return &istioInstaller{
		unstructuredKubeClient: unstructuredKubeClient,
		manifestBuilder:        clients.IstioClients.OperatorManifestBuilder,
		istioInstallOptions:    &opts.Mesh.Install,
		out:                    out,
		in:                     in,
		clusterName:            clusterName,
		operatorManager:        operatorManager,
		fileReader:             fileReader,
	}, nil
}

type istioInstaller struct {
	unstructuredKubeClient kube.UnstructuredKubeClient
	manifestBuilder        operator.InstallerManifestBuilder
	istioInstallOptions    *options.MeshInstall
	out                    io.Writer
	in                     io.Reader
	clusterName            string
	operatorManager        operator.OperatorManager
	fileReader             common.FileReader
}

func (i *istioInstaller) Install() error {
	namespace := i.istioInstallOptions.InstallationConfig.InstallNamespace

	istioControlPlane, err := i.loadIstioOperator()
	if err != nil {
		return err
	}

	if i.istioInstallOptions.DryRun {
		manifest, err := i.manifestBuilder.Build(&i.istioInstallOptions.InstallationConfig)
		if err != nil {
			return err
		}

		fmt.Fprintln(i.out, manifest)

		if istioControlPlane != "" {
			fmt.Fprintf(i.out, "\n---\n%s\n", istioControlPlane)
		}
		return nil
	}

	err = i.installOperator(namespace)
	if err != nil {
		return err
	}

	err = i.writeControlPlaneResource(namespace, istioControlPlane)
	if err != nil {
		return err
	}

	return nil
}

func (i *istioInstaller) installOperator(namespace string) error {
	installNeeded, err := i.operatorManager.ValidateOperatorNamespace(i.clusterName)
	if err != nil {
		return eris.Wrapf(err, "Mesh operator namespace validation failed for cluster '%s' in namespace '%s'", i.clusterName, namespace)
	}

	// install the operator if it didn't exist already
	if installNeeded {
		fmt.Fprintf(i.out, "Installing the Mesh operator to cluster '%s' in namespace '%s'\n", i.clusterName, namespace)

		err := i.operatorManager.Install()
		if err != nil {
			return err
		}
	} else {
		fmt.Fprintf(i.out, "The Mesh operator is already installed to cluster '%s' in namespace '%s' and is suitable for use. Continuing with the Mesh installation.\n", i.clusterName, namespace)
	}

	return nil
}

func (i *istioInstaller) loadIstioOperator() (string, error) {
	userPath := i.istioInstallOptions.ManifestPath
	profile := i.istioInstallOptions.Profile

	if userPath != "" && profile != "" {
		return "", ConflictingControlPlaneSettings
	}

	if userPath != "" {
		userSpecifiedControlPlane, err := i.loadControlPlaneFromUserFlagConfig()
		if err != nil {
			return "", err
		}
		return userSpecifiedControlPlane, nil
	} else if profile != "" {
		preConfiguredProfile, err := i.manifestBuilder.GetOperatorSpecWithProfile(profile, i.istioInstallOptions.InstallationConfig.InstallNamespace)
		if err != nil {
			return "", FailedToParseControlPlaneWithProfile(err, profile)
		}
		return preConfiguredProfile, nil
	}

	return "", nil
}

// returns "", nil if the user did not provide an IstioOperator
func (i *istioInstaller) loadControlPlaneFromUserFlagConfig() (string, error) {
	path := i.istioInstallOptions.ManifestPath

	var contents []byte
	if path == "-" {
		var err error
		contents, err = ioutil.ReadAll(i.in)
		if err != nil {
			return "", err
		}
	} else {
		fileExists, err := i.fileReader.Exists(path)
		if err != nil {
			return "", eris.Wrapf(err, "Unexpected error while reading IstioControlPlane spec")
		} else if !fileExists {
			return "", eris.Errorf("Path to IstioOperator spec does not exist: %s", i.istioInstallOptions.ManifestPath)
		}

		contents, err = i.fileReader.Read(path)
		if err != nil {
			return "", err
		}
	}

	stringContents := string(contents)
	return stringContents, nil
}

// the userProvidedControlPlane may be nil if they didn't provide one
func (i *istioInstaller) writeControlPlaneResource(namespace, istioControlPlaneToWrite string) error {
	if istioControlPlaneToWrite == "" {
		fmt.Fprintf(i.out,
			"\nThe Mesh operator has been installed to cluster '%s' in namespace '%s'. No IstioOperator custom resource was provided to meshctl, so Mesh is currently not fully installed yet. Write a IstioOperator CR to cluster '%s' to complete your installation\n",
			i.clusterName,
			namespace,
			i.clusterName,
		)
		return nil
	}

	// write the control plane
	resources, err := i.unstructuredKubeClient.BuildResources(namespace, istioControlPlaneToWrite)
	if err != nil {
		return FailedToParseControlPlaneSettings(err)
	}

	if len(resources) != 1 {
		return TooManyControlPlaneResources(len(resources))
	}

	istioControlPlane := resources[0]
	resourceKind := istioControlPlane.Object.GetObjectKind().GroupVersionKind().Kind
	if resourceKind != istioControlPlaneKind {
		return UnknownControlPlaneKind(resourceKind)
	}

	_, err = i.unstructuredKubeClient.Create(namespace, resources)
	if err != nil {
		return FailedToWriteControlPlane(err)
	}

	fmt.Fprintf(i.out, "\nThe IstioOperator has been written to cluster '%s' in namespace '%s'. The Mesh operator should process it momentarily and install Mesh.\n", i.clusterName, namespace)
	return nil
}
