package types

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
)

// NetworkCreateRequest is the request message sent to the server for network create call.
//
// Deprecated: use [network.CreateRequest].
type NetworkCreateRequest = network.CreateRequest

// NetworkCreate is the expected body of the "create network" http request message
//
// Deprecated: use [network.CreateOptions].
type NetworkCreate = network.CreateOptions

// NetworkListOptions holds parameters to filter the list of networks with.
//
// Deprecated: use [network.ListOptions].
type NetworkListOptions = network.ListOptions

// NetworkCreateResponse is the response message sent by the server for network create call.
//
// Deprecated: use [network.CreateResponse].
type NetworkCreateResponse = network.CreateResponse

// NetworkInspectOptions holds parameters to inspect network.
//
// Deprecated: use [network.InspectOptions].
type NetworkInspectOptions = network.InspectOptions

// NetworkConnect represents the data to be used to connect a container to the network
//
// Deprecated: use [network.ConnectOptions].
type NetworkConnect = network.ConnectOptions

// NetworkDisconnect represents the data to be used to disconnect a container from the network
//
// Deprecated: use [network.DisconnectOptions].
type NetworkDisconnect = network.DisconnectOptions

// EndpointResource contains network resources allocated and used for a container in a network.
//
// Deprecated: use [network.EndpointResource].
type EndpointResource = network.EndpointResource

// NetworkResource is the body of the "get network" http response message/
//
// Deprecated: use [network.Inspect] or [network.Summary] (for list operations).
type NetworkResource = network.Inspect

// NetworksPruneReport contains the response for Engine API:
// POST "/networks/prune"
//
// Deprecated: use [network.PruneReport].
type NetworksPruneReport = network.PruneReport

// ExecConfig is a small subset of the Config struct that holds the configuration
// for the exec feature of docker.
//
// Deprecated: use [container.ExecOptions].
type ExecConfig = container.ExecOptions

// ExecStartCheck is a temp struct used by execStart
// Config fields is part of ExecConfig in runconfig package
//
// Deprecated: use [container.ExecStartOptions] or [container.ExecAttachOptions].
type ExecStartCheck = container.ExecStartOptions

// ContainerExecInspect holds information returned by exec inspect.
//
// Deprecated: use [container.ExecInspect].
type ContainerExecInspect = container.ExecInspect

// ContainersPruneReport contains the response for Engine API:
// POST "/containers/prune"
//
// Deprecated: use [container.PruneReport].
type ContainersPruneReport = container.PruneReport

// ContainerPathStat is used to encode the header from
// GET "/containers/{name:.*}/archive"
// "Name" is the file or directory name.
//
// Deprecated: use [container.PathStat].
type ContainerPathStat = container.PathStat
