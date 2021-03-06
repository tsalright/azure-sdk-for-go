// +build go1.9

// Copyright 2020 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package netapp

import original "github.com/Azure/azure-sdk-for-go/services/netapp/mgmt/2019-11-01/netapp"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CheckNameResourceTypes = original.CheckNameResourceTypes

const (
	MicrosoftNetAppnetAppAccounts                              CheckNameResourceTypes = original.MicrosoftNetAppnetAppAccounts
	MicrosoftNetAppnetAppAccountscapacityPools                 CheckNameResourceTypes = original.MicrosoftNetAppnetAppAccountscapacityPools
	MicrosoftNetAppnetAppAccountscapacityPoolsvolumes          CheckNameResourceTypes = original.MicrosoftNetAppnetAppAccountscapacityPoolsvolumes
	MicrosoftNetAppnetAppAccountscapacityPoolsvolumessnapshots CheckNameResourceTypes = original.MicrosoftNetAppnetAppAccountscapacityPoolsvolumessnapshots
)

type EndpointType = original.EndpointType

const (
	Dst EndpointType = original.Dst
	Src EndpointType = original.Src
)

type InAvailabilityReasonType = original.InAvailabilityReasonType

const (
	AlreadyExists InAvailabilityReasonType = original.AlreadyExists
	Invalid       InAvailabilityReasonType = original.Invalid
)

type MirrorState = original.MirrorState

const (
	Broken        MirrorState = original.Broken
	Mirrored      MirrorState = original.Mirrored
	Uninitialized MirrorState = original.Uninitialized
)

type RelationshipStatus = original.RelationshipStatus

const (
	Idle         RelationshipStatus = original.Idle
	Transferring RelationshipStatus = original.Transferring
)

type ReplicationSchedule = original.ReplicationSchedule

const (
	OneZerominutely ReplicationSchedule = original.OneZerominutely
	Daily           ReplicationSchedule = original.Daily
	Hourly          ReplicationSchedule = original.Hourly
	Monthly         ReplicationSchedule = original.Monthly
	Weekly          ReplicationSchedule = original.Weekly
)

type ServiceLevel = original.ServiceLevel

const (
	Premium  ServiceLevel = original.Premium
	Standard ServiceLevel = original.Standard
	Ultra    ServiceLevel = original.Ultra
)

type Account = original.Account
type AccountList = original.AccountList
type AccountPatch = original.AccountPatch
type AccountProperties = original.AccountProperties
type AccountsClient = original.AccountsClient
type AccountsCreateOrUpdateFuture = original.AccountsCreateOrUpdateFuture
type AccountsDeleteFuture = original.AccountsDeleteFuture
type ActiveDirectory = original.ActiveDirectory
type AuthorizeRequest = original.AuthorizeRequest
type BaseClient = original.BaseClient
type CapacityPool = original.CapacityPool
type CapacityPoolList = original.CapacityPoolList
type CapacityPoolPatch = original.CapacityPoolPatch
type Dimension = original.Dimension
type ExportPolicyRule = original.ExportPolicyRule
type MetricSpecification = original.MetricSpecification
type MountTarget = original.MountTarget
type MountTargetProperties = original.MountTargetProperties
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationProperties = original.OperationProperties
type OperationsClient = original.OperationsClient
type PoolPatchProperties = original.PoolPatchProperties
type PoolProperties = original.PoolProperties
type PoolsClient = original.PoolsClient
type PoolsCreateOrUpdateFuture = original.PoolsCreateOrUpdateFuture
type PoolsDeleteFuture = original.PoolsDeleteFuture
type PoolsUpdateFuture = original.PoolsUpdateFuture
type ReplicationObject = original.ReplicationObject
type ReplicationStatus = original.ReplicationStatus
type ResourceClient = original.ResourceClient
type ResourceNameAvailability = original.ResourceNameAvailability
type ResourceNameAvailabilityRequest = original.ResourceNameAvailabilityRequest
type ServiceSpecification = original.ServiceSpecification
type Snapshot = original.Snapshot
type SnapshotProperties = original.SnapshotProperties
type SnapshotsClient = original.SnapshotsClient
type SnapshotsCreateFuture = original.SnapshotsCreateFuture
type SnapshotsDeleteFuture = original.SnapshotsDeleteFuture
type SnapshotsList = original.SnapshotsList
type SnapshotsUpdateFuture = original.SnapshotsUpdateFuture
type Volume = original.Volume
type VolumeList = original.VolumeList
type VolumePatch = original.VolumePatch
type VolumePatchProperties = original.VolumePatchProperties
type VolumePatchPropertiesExportPolicy = original.VolumePatchPropertiesExportPolicy
type VolumeProperties = original.VolumeProperties
type VolumePropertiesDataProtection = original.VolumePropertiesDataProtection
type VolumePropertiesExportPolicy = original.VolumePropertiesExportPolicy
type VolumeRevert = original.VolumeRevert
type VolumesAuthorizeReplicationFuture = original.VolumesAuthorizeReplicationFuture
type VolumesBreakReplicationFuture = original.VolumesBreakReplicationFuture
type VolumesClient = original.VolumesClient
type VolumesCreateOrUpdateFuture = original.VolumesCreateOrUpdateFuture
type VolumesDeleteFuture = original.VolumesDeleteFuture
type VolumesDeleteReplicationFuture = original.VolumesDeleteReplicationFuture
type VolumesResyncReplicationFuture = original.VolumesResyncReplicationFuture
type VolumesRevertFuture = original.VolumesRevertFuture
type VolumesUpdateFuture = original.VolumesUpdateFuture

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPoolsClient(subscriptionID string) PoolsClient {
	return original.NewPoolsClient(subscriptionID)
}
func NewPoolsClientWithBaseURI(baseURI string, subscriptionID string) PoolsClient {
	return original.NewPoolsClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceClient(subscriptionID string) ResourceClient {
	return original.NewResourceClient(subscriptionID)
}
func NewResourceClientWithBaseURI(baseURI string, subscriptionID string) ResourceClient {
	return original.NewResourceClientWithBaseURI(baseURI, subscriptionID)
}
func NewSnapshotsClient(subscriptionID string) SnapshotsClient {
	return original.NewSnapshotsClient(subscriptionID)
}
func NewSnapshotsClientWithBaseURI(baseURI string, subscriptionID string) SnapshotsClient {
	return original.NewSnapshotsClientWithBaseURI(baseURI, subscriptionID)
}
func NewVolumesClient(subscriptionID string) VolumesClient {
	return original.NewVolumesClient(subscriptionID)
}
func NewVolumesClientWithBaseURI(baseURI string, subscriptionID string) VolumesClient {
	return original.NewVolumesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCheckNameResourceTypesValues() []CheckNameResourceTypes {
	return original.PossibleCheckNameResourceTypesValues()
}
func PossibleEndpointTypeValues() []EndpointType {
	return original.PossibleEndpointTypeValues()
}
func PossibleInAvailabilityReasonTypeValues() []InAvailabilityReasonType {
	return original.PossibleInAvailabilityReasonTypeValues()
}
func PossibleMirrorStateValues() []MirrorState {
	return original.PossibleMirrorStateValues()
}
func PossibleRelationshipStatusValues() []RelationshipStatus {
	return original.PossibleRelationshipStatusValues()
}
func PossibleReplicationScheduleValues() []ReplicationSchedule {
	return original.PossibleReplicationScheduleValues()
}
func PossibleServiceLevelValues() []ServiceLevel {
	return original.PossibleServiceLevelValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
