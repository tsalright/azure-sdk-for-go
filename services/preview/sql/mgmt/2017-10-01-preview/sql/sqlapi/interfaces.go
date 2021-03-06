package sqlapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-10-01-preview/sql"
	"github.com/Azure/go-autorest/autorest"
	"github.com/satori/go.uuid"
)

// DatabaseOperationsClientAPI contains the set of methods on the DatabaseOperationsClient type.
type DatabaseOperationsClientAPI interface {
	Cancel(ctx context.Context, resourceGroupName string, serverName string, databaseName string, operationID uuid.UUID) (result autorest.Response, err error)
	ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.DatabaseOperationListResultPage, err error)
	ListByDatabaseComplete(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.DatabaseOperationListResultIterator, err error)
}

var _ DatabaseOperationsClientAPI = (*sql.DatabaseOperationsClient)(nil)

// ElasticPoolOperationsClientAPI contains the set of methods on the ElasticPoolOperationsClient type.
type ElasticPoolOperationsClientAPI interface {
	Cancel(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, operationID uuid.UUID) (result autorest.Response, err error)
	ListByElasticPool(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (result sql.ElasticPoolOperationListResultPage, err error)
	ListByElasticPoolComplete(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (result sql.ElasticPoolOperationListResultIterator, err error)
}

var _ ElasticPoolOperationsClientAPI = (*sql.ElasticPoolOperationsClient)(nil)

// DatabaseVulnerabilityAssessmentScansClientAPI contains the set of methods on the DatabaseVulnerabilityAssessmentScansClient type.
type DatabaseVulnerabilityAssessmentScansClientAPI interface {
	Export(ctx context.Context, resourceGroupName string, serverName string, databaseName string, scanID string) (result sql.DatabaseVulnerabilityAssessmentScansExport, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, scanID string) (result sql.VulnerabilityAssessmentScanRecord, err error)
	InitiateScan(ctx context.Context, resourceGroupName string, serverName string, databaseName string, scanID string) (result sql.DatabaseVulnerabilityAssessmentScansInitiateScanFuture, err error)
	ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.VulnerabilityAssessmentScanRecordListResultPage, err error)
	ListByDatabaseComplete(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.VulnerabilityAssessmentScanRecordListResultIterator, err error)
}

var _ DatabaseVulnerabilityAssessmentScansClientAPI = (*sql.DatabaseVulnerabilityAssessmentScansClient)(nil)

// ManagedDatabaseVulnerabilityAssessmentRuleBaselinesClientAPI contains the set of methods on the ManagedDatabaseVulnerabilityAssessmentRuleBaselinesClient type.
type ManagedDatabaseVulnerabilityAssessmentRuleBaselinesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, ruleID string, baselineName sql.VulnerabilityAssessmentPolicyBaselineName, parameters sql.DatabaseVulnerabilityAssessmentRuleBaseline) (result sql.DatabaseVulnerabilityAssessmentRuleBaseline, err error)
	Delete(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, ruleID string, baselineName sql.VulnerabilityAssessmentPolicyBaselineName) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, ruleID string, baselineName sql.VulnerabilityAssessmentPolicyBaselineName) (result sql.DatabaseVulnerabilityAssessmentRuleBaseline, err error)
}

var _ ManagedDatabaseVulnerabilityAssessmentRuleBaselinesClientAPI = (*sql.ManagedDatabaseVulnerabilityAssessmentRuleBaselinesClient)(nil)

// ManagedDatabaseVulnerabilityAssessmentScansClientAPI contains the set of methods on the ManagedDatabaseVulnerabilityAssessmentScansClient type.
type ManagedDatabaseVulnerabilityAssessmentScansClientAPI interface {
	Export(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, scanID string) (result sql.DatabaseVulnerabilityAssessmentScansExport, err error)
	Get(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, scanID string) (result sql.VulnerabilityAssessmentScanRecord, err error)
	InitiateScan(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, scanID string) (result sql.ManagedDatabaseVulnerabilityAssessmentScansInitiateScanFuture, err error)
	ListByDatabase(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string) (result sql.VulnerabilityAssessmentScanRecordListResultPage, err error)
	ListByDatabaseComplete(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string) (result sql.VulnerabilityAssessmentScanRecordListResultIterator, err error)
}

var _ ManagedDatabaseVulnerabilityAssessmentScansClientAPI = (*sql.ManagedDatabaseVulnerabilityAssessmentScansClient)(nil)

// ManagedDatabaseVulnerabilityAssessmentsClientAPI contains the set of methods on the ManagedDatabaseVulnerabilityAssessmentsClient type.
type ManagedDatabaseVulnerabilityAssessmentsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters sql.DatabaseVulnerabilityAssessment) (result sql.DatabaseVulnerabilityAssessment, err error)
	Delete(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string) (result sql.DatabaseVulnerabilityAssessment, err error)
	ListByDatabase(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string) (result sql.DatabaseVulnerabilityAssessmentListResultPage, err error)
	ListByDatabaseComplete(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string) (result sql.DatabaseVulnerabilityAssessmentListResultIterator, err error)
}

var _ ManagedDatabaseVulnerabilityAssessmentsClientAPI = (*sql.ManagedDatabaseVulnerabilityAssessmentsClient)(nil)

// CapabilitiesClientAPI contains the set of methods on the CapabilitiesClient type.
type CapabilitiesClientAPI interface {
	ListByLocation(ctx context.Context, locationName string, include sql.CapabilityGroup) (result sql.LocationCapabilities, err error)
}

var _ CapabilitiesClientAPI = (*sql.CapabilitiesClient)(nil)

// DatabasesClientAPI contains the set of methods on the DatabasesClient type.
type DatabasesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters sql.Database) (result sql.DatabasesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.DatabasesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.Database, err error)
	ListByElasticPool(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (result sql.DatabaseListResultPage, err error)
	ListByElasticPoolComplete(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (result sql.DatabaseListResultIterator, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.DatabaseListResultPage, err error)
	ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string) (result sql.DatabaseListResultIterator, err error)
	Pause(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.DatabasesPauseFuture, err error)
	Rename(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters sql.ResourceMoveDefinition) (result autorest.Response, err error)
	Resume(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.DatabasesResumeFuture, err error)
	Update(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters sql.DatabaseUpdate) (result sql.DatabasesUpdateFuture, err error)
	UpgradeDataWarehouse(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.DatabasesUpgradeDataWarehouseFuture, err error)
}

var _ DatabasesClientAPI = (*sql.DatabasesClient)(nil)

// ElasticPoolsClientAPI contains the set of methods on the ElasticPoolsClient type.
type ElasticPoolsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, parameters sql.ElasticPool) (result sql.ElasticPoolsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (result sql.ElasticPoolsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (result sql.ElasticPool, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string, skip *int32) (result sql.ElasticPoolListResultPage, err error)
	ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string, skip *int32) (result sql.ElasticPoolListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, parameters sql.ElasticPoolUpdate) (result sql.ElasticPoolsUpdateFuture, err error)
}

var _ ElasticPoolsClientAPI = (*sql.ElasticPoolsClient)(nil)

// InstanceFailoverGroupsClientAPI contains the set of methods on the InstanceFailoverGroupsClient type.
type InstanceFailoverGroupsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, parameters sql.InstanceFailoverGroup) (result sql.InstanceFailoverGroupsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string) (result sql.InstanceFailoverGroupsDeleteFuture, err error)
	Failover(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string) (result sql.InstanceFailoverGroupsFailoverFuture, err error)
	ForceFailoverAllowDataLoss(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string) (result sql.InstanceFailoverGroupsForceFailoverAllowDataLossFuture, err error)
	Get(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string) (result sql.InstanceFailoverGroup, err error)
	ListByLocation(ctx context.Context, resourceGroupName string, locationName string) (result sql.InstanceFailoverGroupListResultPage, err error)
	ListByLocationComplete(ctx context.Context, resourceGroupName string, locationName string) (result sql.InstanceFailoverGroupListResultIterator, err error)
}

var _ InstanceFailoverGroupsClientAPI = (*sql.InstanceFailoverGroupsClient)(nil)

// BackupShortTermRetentionPoliciesClientAPI contains the set of methods on the BackupShortTermRetentionPoliciesClient type.
type BackupShortTermRetentionPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters sql.BackupShortTermRetentionPolicy) (result sql.BackupShortTermRetentionPoliciesCreateOrUpdateFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.BackupShortTermRetentionPolicy, err error)
	ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.BackupShortTermRetentionPolicyListResultPage, err error)
	ListByDatabaseComplete(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.BackupShortTermRetentionPolicyListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters sql.BackupShortTermRetentionPolicy) (result sql.BackupShortTermRetentionPoliciesUpdateFuture, err error)
}

var _ BackupShortTermRetentionPoliciesClientAPI = (*sql.BackupShortTermRetentionPoliciesClient)(nil)

// TdeCertificatesClientAPI contains the set of methods on the TdeCertificatesClient type.
type TdeCertificatesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, serverName string, parameters sql.TdeCertificate) (result sql.TdeCertificatesCreateFuture, err error)
}

var _ TdeCertificatesClientAPI = (*sql.TdeCertificatesClient)(nil)

// ManagedInstanceTdeCertificatesClientAPI contains the set of methods on the ManagedInstanceTdeCertificatesClient type.
type ManagedInstanceTdeCertificatesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters sql.TdeCertificate) (result sql.ManagedInstanceTdeCertificatesCreateFuture, err error)
}

var _ ManagedInstanceTdeCertificatesClientAPI = (*sql.ManagedInstanceTdeCertificatesClient)(nil)

// ManagedInstanceKeysClientAPI contains the set of methods on the ManagedInstanceKeysClient type.
type ManagedInstanceKeysClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, keyName string, parameters sql.ManagedInstanceKey) (result sql.ManagedInstanceKeysCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, managedInstanceName string, keyName string) (result sql.ManagedInstanceKeysDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, managedInstanceName string, keyName string) (result sql.ManagedInstanceKey, err error)
	ListByInstance(ctx context.Context, resourceGroupName string, managedInstanceName string, filter string) (result sql.ManagedInstanceKeyListResultPage, err error)
	ListByInstanceComplete(ctx context.Context, resourceGroupName string, managedInstanceName string, filter string) (result sql.ManagedInstanceKeyListResultIterator, err error)
}

var _ ManagedInstanceKeysClientAPI = (*sql.ManagedInstanceKeysClient)(nil)

// ManagedInstanceEncryptionProtectorsClientAPI contains the set of methods on the ManagedInstanceEncryptionProtectorsClient type.
type ManagedInstanceEncryptionProtectorsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters sql.ManagedInstanceEncryptionProtector) (result sql.ManagedInstanceEncryptionProtectorsCreateOrUpdateFuture, err error)
	Get(ctx context.Context, resourceGroupName string, managedInstanceName string) (result sql.ManagedInstanceEncryptionProtector, err error)
	ListByInstance(ctx context.Context, resourceGroupName string, managedInstanceName string) (result sql.ManagedInstanceEncryptionProtectorListResultPage, err error)
	ListByInstanceComplete(ctx context.Context, resourceGroupName string, managedInstanceName string) (result sql.ManagedInstanceEncryptionProtectorListResultIterator, err error)
	Revalidate(ctx context.Context, resourceGroupName string, managedInstanceName string) (result sql.ManagedInstanceEncryptionProtectorsRevalidateFuture, err error)
}

var _ ManagedInstanceEncryptionProtectorsClientAPI = (*sql.ManagedInstanceEncryptionProtectorsClient)(nil)

// RecoverableManagedDatabasesClientAPI contains the set of methods on the RecoverableManagedDatabasesClient type.
type RecoverableManagedDatabasesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, managedInstanceName string, recoverableDatabaseName string) (result sql.RecoverableManagedDatabase, err error)
	ListByInstance(ctx context.Context, resourceGroupName string, managedInstanceName string) (result sql.RecoverableManagedDatabaseListResultPage, err error)
	ListByInstanceComplete(ctx context.Context, resourceGroupName string, managedInstanceName string) (result sql.RecoverableManagedDatabaseListResultIterator, err error)
}

var _ RecoverableManagedDatabasesClientAPI = (*sql.RecoverableManagedDatabasesClient)(nil)
