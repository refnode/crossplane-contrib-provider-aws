/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package dbinstance

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/rds"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/rds/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateDescribeDBInstancesInput returns input for read
// operation.
func GenerateDescribeDBInstancesInput(cr *svcapitypes.DBInstance) *svcsdk.DescribeDBInstancesInput {
	res := &svcsdk.DescribeDBInstancesInput{}

	if cr.Status.AtProvider.DBInstanceIdentifier != nil {
		res.SetDBInstanceIdentifier(*cr.Status.AtProvider.DBInstanceIdentifier)
	}

	return res
}

// GenerateDBInstance returns the current state in the form of *svcapitypes.DBInstance.
func GenerateDBInstance(resp *svcsdk.DescribeDBInstancesOutput) *svcapitypes.DBInstance {
	cr := &svcapitypes.DBInstance{}

	found := false
	for _, elem := range resp.DBInstances {
		if elem.ActivityStreamEngineNativeAuditFieldsIncluded != nil {
			cr.Status.AtProvider.ActivityStreamEngineNativeAuditFieldsIncluded = elem.ActivityStreamEngineNativeAuditFieldsIncluded
		} else {
			cr.Status.AtProvider.ActivityStreamEngineNativeAuditFieldsIncluded = nil
		}
		if elem.ActivityStreamKinesisStreamName != nil {
			cr.Status.AtProvider.ActivityStreamKinesisStreamName = elem.ActivityStreamKinesisStreamName
		} else {
			cr.Status.AtProvider.ActivityStreamKinesisStreamName = nil
		}
		if elem.ActivityStreamKmsKeyId != nil {
			cr.Status.AtProvider.ActivityStreamKMSKeyID = elem.ActivityStreamKmsKeyId
		} else {
			cr.Status.AtProvider.ActivityStreamKMSKeyID = nil
		}
		if elem.ActivityStreamMode != nil {
			cr.Status.AtProvider.ActivityStreamMode = elem.ActivityStreamMode
		} else {
			cr.Status.AtProvider.ActivityStreamMode = nil
		}
		if elem.ActivityStreamPolicyStatus != nil {
			cr.Status.AtProvider.ActivityStreamPolicyStatus = elem.ActivityStreamPolicyStatus
		} else {
			cr.Status.AtProvider.ActivityStreamPolicyStatus = nil
		}
		if elem.ActivityStreamStatus != nil {
			cr.Status.AtProvider.ActivityStreamStatus = elem.ActivityStreamStatus
		} else {
			cr.Status.AtProvider.ActivityStreamStatus = nil
		}
		if elem.AllocatedStorage != nil {
			cr.Spec.ForProvider.AllocatedStorage = elem.AllocatedStorage
		} else {
			cr.Spec.ForProvider.AllocatedStorage = nil
		}
		if elem.AssociatedRoles != nil {
			f7 := []*svcapitypes.DBInstanceRole{}
			for _, f7iter := range elem.AssociatedRoles {
				f7elem := &svcapitypes.DBInstanceRole{}
				if f7iter.FeatureName != nil {
					f7elem.FeatureName = f7iter.FeatureName
				}
				if f7iter.RoleArn != nil {
					f7elem.RoleARN = f7iter.RoleArn
				}
				if f7iter.Status != nil {
					f7elem.Status = f7iter.Status
				}
				f7 = append(f7, f7elem)
			}
			cr.Status.AtProvider.AssociatedRoles = f7
		} else {
			cr.Status.AtProvider.AssociatedRoles = nil
		}
		if elem.AutoMinorVersionUpgrade != nil {
			cr.Spec.ForProvider.AutoMinorVersionUpgrade = elem.AutoMinorVersionUpgrade
		} else {
			cr.Spec.ForProvider.AutoMinorVersionUpgrade = nil
		}
		if elem.AutomaticRestartTime != nil {
			cr.Status.AtProvider.AutomaticRestartTime = &metav1.Time{*elem.AutomaticRestartTime}
		} else {
			cr.Status.AtProvider.AutomaticRestartTime = nil
		}
		if elem.AutomationMode != nil {
			cr.Status.AtProvider.AutomationMode = elem.AutomationMode
		} else {
			cr.Status.AtProvider.AutomationMode = nil
		}
		if elem.AvailabilityZone != nil {
			cr.Spec.ForProvider.AvailabilityZone = elem.AvailabilityZone
		} else {
			cr.Spec.ForProvider.AvailabilityZone = nil
		}
		if elem.AwsBackupRecoveryPointArn != nil {
			cr.Status.AtProvider.AWSBackupRecoveryPointARN = elem.AwsBackupRecoveryPointArn
		} else {
			cr.Status.AtProvider.AWSBackupRecoveryPointARN = nil
		}
		if elem.BackupRetentionPeriod != nil {
			cr.Spec.ForProvider.BackupRetentionPeriod = elem.BackupRetentionPeriod
		} else {
			cr.Spec.ForProvider.BackupRetentionPeriod = nil
		}
		if elem.BackupTarget != nil {
			cr.Spec.ForProvider.BackupTarget = elem.BackupTarget
		} else {
			cr.Spec.ForProvider.BackupTarget = nil
		}
		if elem.CACertificateIdentifier != nil {
			cr.Spec.ForProvider.CACertificateIdentifier = elem.CACertificateIdentifier
		} else {
			cr.Spec.ForProvider.CACertificateIdentifier = nil
		}
		if elem.CertificateDetails != nil {
			f16 := &svcapitypes.CertificateDetails{}
			if elem.CertificateDetails.CAIdentifier != nil {
				f16.CAIdentifier = elem.CertificateDetails.CAIdentifier
			}
			if elem.CertificateDetails.ValidTill != nil {
				f16.ValidTill = &metav1.Time{*elem.CertificateDetails.ValidTill}
			}
			cr.Status.AtProvider.CertificateDetails = f16
		} else {
			cr.Status.AtProvider.CertificateDetails = nil
		}
		if elem.CharacterSetName != nil {
			cr.Spec.ForProvider.CharacterSetName = elem.CharacterSetName
		} else {
			cr.Spec.ForProvider.CharacterSetName = nil
		}
		if elem.CopyTagsToSnapshot != nil {
			cr.Spec.ForProvider.CopyTagsToSnapshot = elem.CopyTagsToSnapshot
		} else {
			cr.Spec.ForProvider.CopyTagsToSnapshot = nil
		}
		if elem.CustomIamInstanceProfile != nil {
			cr.Spec.ForProvider.CustomIAMInstanceProfile = elem.CustomIamInstanceProfile
		} else {
			cr.Spec.ForProvider.CustomIAMInstanceProfile = nil
		}
		if elem.CustomerOwnedIpEnabled != nil {
			cr.Status.AtProvider.CustomerOwnedIPEnabled = elem.CustomerOwnedIpEnabled
		} else {
			cr.Status.AtProvider.CustomerOwnedIPEnabled = nil
		}
		if elem.DBClusterIdentifier != nil {
			cr.Spec.ForProvider.DBClusterIdentifier = elem.DBClusterIdentifier
		} else {
			cr.Spec.ForProvider.DBClusterIdentifier = nil
		}
		if elem.DBInstanceArn != nil {
			cr.Status.AtProvider.DBInstanceARN = elem.DBInstanceArn
		} else {
			cr.Status.AtProvider.DBInstanceARN = nil
		}
		if elem.DBInstanceAutomatedBackupsReplications != nil {
			f23 := []*svcapitypes.DBInstanceAutomatedBackupsReplication{}
			for _, f23iter := range elem.DBInstanceAutomatedBackupsReplications {
				f23elem := &svcapitypes.DBInstanceAutomatedBackupsReplication{}
				if f23iter.DBInstanceAutomatedBackupsArn != nil {
					f23elem.DBInstanceAutomatedBackupsARN = f23iter.DBInstanceAutomatedBackupsArn
				}
				f23 = append(f23, f23elem)
			}
			cr.Status.AtProvider.DBInstanceAutomatedBackupsReplications = f23
		} else {
			cr.Status.AtProvider.DBInstanceAutomatedBackupsReplications = nil
		}
		if elem.DBInstanceClass != nil {
			cr.Spec.ForProvider.DBInstanceClass = elem.DBInstanceClass
		} else {
			cr.Spec.ForProvider.DBInstanceClass = nil
		}
		if elem.DBInstanceIdentifier != nil {
			cr.Status.AtProvider.DBInstanceIdentifier = elem.DBInstanceIdentifier
		} else {
			cr.Status.AtProvider.DBInstanceIdentifier = nil
		}
		if elem.DBInstanceStatus != nil {
			cr.Status.AtProvider.DBInstanceStatus = elem.DBInstanceStatus
		} else {
			cr.Status.AtProvider.DBInstanceStatus = nil
		}
		if elem.DBName != nil {
			cr.Spec.ForProvider.DBName = elem.DBName
		} else {
			cr.Spec.ForProvider.DBName = nil
		}
		if elem.DBParameterGroups != nil {
			f28 := []*svcapitypes.DBParameterGroupStatus_SDK{}
			for _, f28iter := range elem.DBParameterGroups {
				f28elem := &svcapitypes.DBParameterGroupStatus_SDK{}
				if f28iter.DBParameterGroupName != nil {
					f28elem.DBParameterGroupName = f28iter.DBParameterGroupName
				}
				if f28iter.ParameterApplyStatus != nil {
					f28elem.ParameterApplyStatus = f28iter.ParameterApplyStatus
				}
				f28 = append(f28, f28elem)
			}
			cr.Status.AtProvider.DBParameterGroups = f28
		} else {
			cr.Status.AtProvider.DBParameterGroups = nil
		}
		if elem.DBSecurityGroups != nil {
			f29 := []*svcapitypes.DBSecurityGroupMembership{}
			for _, f29iter := range elem.DBSecurityGroups {
				f29elem := &svcapitypes.DBSecurityGroupMembership{}
				if f29iter.DBSecurityGroupName != nil {
					f29elem.DBSecurityGroupName = f29iter.DBSecurityGroupName
				}
				if f29iter.Status != nil {
					f29elem.Status = f29iter.Status
				}
				f29 = append(f29, f29elem)
			}
			cr.Status.AtProvider.DBSecurityGroups = f29
		} else {
			cr.Status.AtProvider.DBSecurityGroups = nil
		}
		if elem.DBSubnetGroup != nil {
			f30 := &svcapitypes.DBSubnetGroup{}
			if elem.DBSubnetGroup.DBSubnetGroupArn != nil {
				f30.DBSubnetGroupARN = elem.DBSubnetGroup.DBSubnetGroupArn
			}
			if elem.DBSubnetGroup.DBSubnetGroupDescription != nil {
				f30.DBSubnetGroupDescription = elem.DBSubnetGroup.DBSubnetGroupDescription
			}
			if elem.DBSubnetGroup.DBSubnetGroupName != nil {
				f30.DBSubnetGroupName = elem.DBSubnetGroup.DBSubnetGroupName
			}
			if elem.DBSubnetGroup.SubnetGroupStatus != nil {
				f30.SubnetGroupStatus = elem.DBSubnetGroup.SubnetGroupStatus
			}
			if elem.DBSubnetGroup.Subnets != nil {
				f30f4 := []*svcapitypes.Subnet{}
				for _, f30f4iter := range elem.DBSubnetGroup.Subnets {
					f30f4elem := &svcapitypes.Subnet{}
					if f30f4iter.SubnetAvailabilityZone != nil {
						f30f4elemf0 := &svcapitypes.AvailabilityZone{}
						if f30f4iter.SubnetAvailabilityZone.Name != nil {
							f30f4elemf0.Name = f30f4iter.SubnetAvailabilityZone.Name
						}
						f30f4elem.SubnetAvailabilityZone = f30f4elemf0
					}
					if f30f4iter.SubnetIdentifier != nil {
						f30f4elem.SubnetIdentifier = f30f4iter.SubnetIdentifier
					}
					if f30f4iter.SubnetOutpost != nil {
						f30f4elemf2 := &svcapitypes.Outpost{}
						if f30f4iter.SubnetOutpost.Arn != nil {
							f30f4elemf2.ARN = f30f4iter.SubnetOutpost.Arn
						}
						f30f4elem.SubnetOutpost = f30f4elemf2
					}
					if f30f4iter.SubnetStatus != nil {
						f30f4elem.SubnetStatus = f30f4iter.SubnetStatus
					}
					f30f4 = append(f30f4, f30f4elem)
				}
				f30.Subnets = f30f4
			}
			if elem.DBSubnetGroup.SupportedNetworkTypes != nil {
				f30f5 := []*string{}
				for _, f30f5iter := range elem.DBSubnetGroup.SupportedNetworkTypes {
					var f30f5elem string
					f30f5elem = *f30f5iter
					f30f5 = append(f30f5, &f30f5elem)
				}
				f30.SupportedNetworkTypes = f30f5
			}
			if elem.DBSubnetGroup.VpcId != nil {
				f30.VPCID = elem.DBSubnetGroup.VpcId
			}
			cr.Status.AtProvider.DBSubnetGroup = f30
		} else {
			cr.Status.AtProvider.DBSubnetGroup = nil
		}
		if elem.DBSystemId != nil {
			cr.Status.AtProvider.DBSystemID = elem.DBSystemId
		} else {
			cr.Status.AtProvider.DBSystemID = nil
		}
		if elem.DbInstancePort != nil {
			cr.Status.AtProvider.DBInstancePort = elem.DbInstancePort
		} else {
			cr.Status.AtProvider.DBInstancePort = nil
		}
		if elem.DbiResourceId != nil {
			cr.Status.AtProvider.DBIResourceID = elem.DbiResourceId
		} else {
			cr.Status.AtProvider.DBIResourceID = nil
		}
		if elem.DeletionProtection != nil {
			cr.Spec.ForProvider.DeletionProtection = elem.DeletionProtection
		} else {
			cr.Spec.ForProvider.DeletionProtection = nil
		}
		if elem.DomainMemberships != nil {
			f35 := []*svcapitypes.DomainMembership{}
			for _, f35iter := range elem.DomainMemberships {
				f35elem := &svcapitypes.DomainMembership{}
				if f35iter.Domain != nil {
					f35elem.Domain = f35iter.Domain
				}
				if f35iter.FQDN != nil {
					f35elem.FQDN = f35iter.FQDN
				}
				if f35iter.IAMRoleName != nil {
					f35elem.IAMRoleName = f35iter.IAMRoleName
				}
				if f35iter.Status != nil {
					f35elem.Status = f35iter.Status
				}
				f35 = append(f35, f35elem)
			}
			cr.Status.AtProvider.DomainMemberships = f35
		} else {
			cr.Status.AtProvider.DomainMemberships = nil
		}
		if elem.EnabledCloudwatchLogsExports != nil {
			f36 := []*string{}
			for _, f36iter := range elem.EnabledCloudwatchLogsExports {
				var f36elem string
				f36elem = *f36iter
				f36 = append(f36, &f36elem)
			}
			cr.Status.AtProvider.EnabledCloudwatchLogsExports = f36
		} else {
			cr.Status.AtProvider.EnabledCloudwatchLogsExports = nil
		}
		if elem.Endpoint != nil {
			f37 := &svcapitypes.Endpoint{}
			if elem.Endpoint.Address != nil {
				f37.Address = elem.Endpoint.Address
			}
			if elem.Endpoint.HostedZoneId != nil {
				f37.HostedZoneID = elem.Endpoint.HostedZoneId
			}
			if elem.Endpoint.Port != nil {
				f37.Port = elem.Endpoint.Port
			}
			cr.Status.AtProvider.Endpoint = f37
		} else {
			cr.Status.AtProvider.Endpoint = nil
		}
		if elem.Engine != nil {
			cr.Spec.ForProvider.Engine = elem.Engine
		} else {
			cr.Spec.ForProvider.Engine = nil
		}
		if elem.EngineVersion != nil {
			cr.Spec.ForProvider.EngineVersion = elem.EngineVersion
		} else {
			cr.Spec.ForProvider.EngineVersion = nil
		}
		if elem.EnhancedMonitoringResourceArn != nil {
			cr.Status.AtProvider.EnhancedMonitoringResourceARN = elem.EnhancedMonitoringResourceArn
		} else {
			cr.Status.AtProvider.EnhancedMonitoringResourceARN = nil
		}
		if elem.IAMDatabaseAuthenticationEnabled != nil {
			cr.Status.AtProvider.IAMDatabaseAuthenticationEnabled = elem.IAMDatabaseAuthenticationEnabled
		} else {
			cr.Status.AtProvider.IAMDatabaseAuthenticationEnabled = nil
		}
		if elem.InstanceCreateTime != nil {
			cr.Status.AtProvider.InstanceCreateTime = &metav1.Time{*elem.InstanceCreateTime}
		} else {
			cr.Status.AtProvider.InstanceCreateTime = nil
		}
		if elem.Iops != nil {
			cr.Spec.ForProvider.IOPS = elem.Iops
		} else {
			cr.Spec.ForProvider.IOPS = nil
		}
		if elem.KmsKeyId != nil {
			cr.Spec.ForProvider.KMSKeyID = elem.KmsKeyId
		} else {
			cr.Spec.ForProvider.KMSKeyID = nil
		}
		if elem.LatestRestorableTime != nil {
			cr.Status.AtProvider.LatestRestorableTime = &metav1.Time{*elem.LatestRestorableTime}
		} else {
			cr.Status.AtProvider.LatestRestorableTime = nil
		}
		if elem.LicenseModel != nil {
			cr.Spec.ForProvider.LicenseModel = elem.LicenseModel
		} else {
			cr.Spec.ForProvider.LicenseModel = nil
		}
		if elem.ListenerEndpoint != nil {
			f47 := &svcapitypes.Endpoint{}
			if elem.ListenerEndpoint.Address != nil {
				f47.Address = elem.ListenerEndpoint.Address
			}
			if elem.ListenerEndpoint.HostedZoneId != nil {
				f47.HostedZoneID = elem.ListenerEndpoint.HostedZoneId
			}
			if elem.ListenerEndpoint.Port != nil {
				f47.Port = elem.ListenerEndpoint.Port
			}
			cr.Status.AtProvider.ListenerEndpoint = f47
		} else {
			cr.Status.AtProvider.ListenerEndpoint = nil
		}
		if elem.MasterUserSecret != nil {
			f48 := &svcapitypes.MasterUserSecret{}
			if elem.MasterUserSecret.KmsKeyId != nil {
				f48.KMSKeyID = elem.MasterUserSecret.KmsKeyId
			}
			if elem.MasterUserSecret.SecretArn != nil {
				f48.SecretARN = elem.MasterUserSecret.SecretArn
			}
			if elem.MasterUserSecret.SecretStatus != nil {
				f48.SecretStatus = elem.MasterUserSecret.SecretStatus
			}
			cr.Status.AtProvider.MasterUserSecret = f48
		} else {
			cr.Status.AtProvider.MasterUserSecret = nil
		}
		if elem.MasterUsername != nil {
			cr.Spec.ForProvider.MasterUsername = elem.MasterUsername
		} else {
			cr.Spec.ForProvider.MasterUsername = nil
		}
		if elem.MaxAllocatedStorage != nil {
			cr.Spec.ForProvider.MaxAllocatedStorage = elem.MaxAllocatedStorage
		} else {
			cr.Spec.ForProvider.MaxAllocatedStorage = nil
		}
		if elem.MonitoringInterval != nil {
			cr.Spec.ForProvider.MonitoringInterval = elem.MonitoringInterval
		} else {
			cr.Spec.ForProvider.MonitoringInterval = nil
		}
		if elem.MonitoringRoleArn != nil {
			cr.Spec.ForProvider.MonitoringRoleARN = elem.MonitoringRoleArn
		} else {
			cr.Spec.ForProvider.MonitoringRoleARN = nil
		}
		if elem.MultiAZ != nil {
			cr.Spec.ForProvider.MultiAZ = elem.MultiAZ
		} else {
			cr.Spec.ForProvider.MultiAZ = nil
		}
		if elem.NcharCharacterSetName != nil {
			cr.Spec.ForProvider.NcharCharacterSetName = elem.NcharCharacterSetName
		} else {
			cr.Spec.ForProvider.NcharCharacterSetName = nil
		}
		if elem.NetworkType != nil {
			cr.Spec.ForProvider.NetworkType = elem.NetworkType
		} else {
			cr.Spec.ForProvider.NetworkType = nil
		}
		if elem.OptionGroupMemberships != nil {
			f56 := []*svcapitypes.OptionGroupMembership{}
			for _, f56iter := range elem.OptionGroupMemberships {
				f56elem := &svcapitypes.OptionGroupMembership{}
				if f56iter.OptionGroupName != nil {
					f56elem.OptionGroupName = f56iter.OptionGroupName
				}
				if f56iter.Status != nil {
					f56elem.Status = f56iter.Status
				}
				f56 = append(f56, f56elem)
			}
			cr.Status.AtProvider.OptionGroupMemberships = f56
		} else {
			cr.Status.AtProvider.OptionGroupMemberships = nil
		}
		if elem.PendingModifiedValues != nil {
			f57 := &svcapitypes.PendingModifiedValues{}
			if elem.PendingModifiedValues.AllocatedStorage != nil {
				f57.AllocatedStorage = elem.PendingModifiedValues.AllocatedStorage
			}
			if elem.PendingModifiedValues.AutomationMode != nil {
				f57.AutomationMode = elem.PendingModifiedValues.AutomationMode
			}
			if elem.PendingModifiedValues.BackupRetentionPeriod != nil {
				f57.BackupRetentionPeriod = elem.PendingModifiedValues.BackupRetentionPeriod
			}
			if elem.PendingModifiedValues.CACertificateIdentifier != nil {
				f57.CACertificateIdentifier = elem.PendingModifiedValues.CACertificateIdentifier
			}
			if elem.PendingModifiedValues.DBInstanceClass != nil {
				f57.DBInstanceClass = elem.PendingModifiedValues.DBInstanceClass
			}
			if elem.PendingModifiedValues.DBInstanceIdentifier != nil {
				f57.DBInstanceIdentifier = elem.PendingModifiedValues.DBInstanceIdentifier
			}
			if elem.PendingModifiedValues.DBSubnetGroupName != nil {
				f57.DBSubnetGroupName = elem.PendingModifiedValues.DBSubnetGroupName
			}
			if elem.PendingModifiedValues.EngineVersion != nil {
				f57.EngineVersion = elem.PendingModifiedValues.EngineVersion
			}
			if elem.PendingModifiedValues.IAMDatabaseAuthenticationEnabled != nil {
				f57.IAMDatabaseAuthenticationEnabled = elem.PendingModifiedValues.IAMDatabaseAuthenticationEnabled
			}
			if elem.PendingModifiedValues.Iops != nil {
				f57.IOPS = elem.PendingModifiedValues.Iops
			}
			if elem.PendingModifiedValues.LicenseModel != nil {
				f57.LicenseModel = elem.PendingModifiedValues.LicenseModel
			}
			if elem.PendingModifiedValues.MasterUserPassword != nil {
				f57.MasterUserPassword = elem.PendingModifiedValues.MasterUserPassword
			}
			if elem.PendingModifiedValues.MultiAZ != nil {
				f57.MultiAZ = elem.PendingModifiedValues.MultiAZ
			}
			if elem.PendingModifiedValues.PendingCloudwatchLogsExports != nil {
				f57f13 := &svcapitypes.PendingCloudwatchLogsExports{}
				if elem.PendingModifiedValues.PendingCloudwatchLogsExports.LogTypesToDisable != nil {
					f57f13f0 := []*string{}
					for _, f57f13f0iter := range elem.PendingModifiedValues.PendingCloudwatchLogsExports.LogTypesToDisable {
						var f57f13f0elem string
						f57f13f0elem = *f57f13f0iter
						f57f13f0 = append(f57f13f0, &f57f13f0elem)
					}
					f57f13.LogTypesToDisable = f57f13f0
				}
				if elem.PendingModifiedValues.PendingCloudwatchLogsExports.LogTypesToEnable != nil {
					f57f13f1 := []*string{}
					for _, f57f13f1iter := range elem.PendingModifiedValues.PendingCloudwatchLogsExports.LogTypesToEnable {
						var f57f13f1elem string
						f57f13f1elem = *f57f13f1iter
						f57f13f1 = append(f57f13f1, &f57f13f1elem)
					}
					f57f13.LogTypesToEnable = f57f13f1
				}
				f57.PendingCloudwatchLogsExports = f57f13
			}
			if elem.PendingModifiedValues.Port != nil {
				f57.Port = elem.PendingModifiedValues.Port
			}
			if elem.PendingModifiedValues.ProcessorFeatures != nil {
				f57f15 := []*svcapitypes.ProcessorFeature{}
				for _, f57f15iter := range elem.PendingModifiedValues.ProcessorFeatures {
					f57f15elem := &svcapitypes.ProcessorFeature{}
					if f57f15iter.Name != nil {
						f57f15elem.Name = f57f15iter.Name
					}
					if f57f15iter.Value != nil {
						f57f15elem.Value = f57f15iter.Value
					}
					f57f15 = append(f57f15, f57f15elem)
				}
				f57.ProcessorFeatures = f57f15
			}
			if elem.PendingModifiedValues.ResumeFullAutomationModeTime != nil {
				f57.ResumeFullAutomationModeTime = &metav1.Time{*elem.PendingModifiedValues.ResumeFullAutomationModeTime}
			}
			if elem.PendingModifiedValues.StorageThroughput != nil {
				f57.StorageThroughput = elem.PendingModifiedValues.StorageThroughput
			}
			if elem.PendingModifiedValues.StorageType != nil {
				f57.StorageType = elem.PendingModifiedValues.StorageType
			}
			cr.Status.AtProvider.PendingModifiedValues = f57
		} else {
			cr.Status.AtProvider.PendingModifiedValues = nil
		}
		if elem.PerformanceInsightsEnabled != nil {
			cr.Status.AtProvider.PerformanceInsightsEnabled = elem.PerformanceInsightsEnabled
		} else {
			cr.Status.AtProvider.PerformanceInsightsEnabled = nil
		}
		if elem.PerformanceInsightsKMSKeyId != nil {
			cr.Spec.ForProvider.PerformanceInsightsKMSKeyID = elem.PerformanceInsightsKMSKeyId
		} else {
			cr.Spec.ForProvider.PerformanceInsightsKMSKeyID = nil
		}
		if elem.PerformanceInsightsRetentionPeriod != nil {
			cr.Spec.ForProvider.PerformanceInsightsRetentionPeriod = elem.PerformanceInsightsRetentionPeriod
		} else {
			cr.Spec.ForProvider.PerformanceInsightsRetentionPeriod = nil
		}
		if elem.PreferredBackupWindow != nil {
			cr.Spec.ForProvider.PreferredBackupWindow = elem.PreferredBackupWindow
		} else {
			cr.Spec.ForProvider.PreferredBackupWindow = nil
		}
		if elem.PreferredMaintenanceWindow != nil {
			cr.Spec.ForProvider.PreferredMaintenanceWindow = elem.PreferredMaintenanceWindow
		} else {
			cr.Spec.ForProvider.PreferredMaintenanceWindow = nil
		}
		if elem.ProcessorFeatures != nil {
			f63 := []*svcapitypes.ProcessorFeature{}
			for _, f63iter := range elem.ProcessorFeatures {
				f63elem := &svcapitypes.ProcessorFeature{}
				if f63iter.Name != nil {
					f63elem.Name = f63iter.Name
				}
				if f63iter.Value != nil {
					f63elem.Value = f63iter.Value
				}
				f63 = append(f63, f63elem)
			}
			cr.Spec.ForProvider.ProcessorFeatures = f63
		} else {
			cr.Spec.ForProvider.ProcessorFeatures = nil
		}
		if elem.PromotionTier != nil {
			cr.Spec.ForProvider.PromotionTier = elem.PromotionTier
		} else {
			cr.Spec.ForProvider.PromotionTier = nil
		}
		if elem.PubliclyAccessible != nil {
			cr.Spec.ForProvider.PubliclyAccessible = elem.PubliclyAccessible
		} else {
			cr.Spec.ForProvider.PubliclyAccessible = nil
		}
		if elem.ReadReplicaDBClusterIdentifiers != nil {
			f66 := []*string{}
			for _, f66iter := range elem.ReadReplicaDBClusterIdentifiers {
				var f66elem string
				f66elem = *f66iter
				f66 = append(f66, &f66elem)
			}
			cr.Status.AtProvider.ReadReplicaDBClusterIdentifiers = f66
		} else {
			cr.Status.AtProvider.ReadReplicaDBClusterIdentifiers = nil
		}
		if elem.ReadReplicaDBInstanceIdentifiers != nil {
			f67 := []*string{}
			for _, f67iter := range elem.ReadReplicaDBInstanceIdentifiers {
				var f67elem string
				f67elem = *f67iter
				f67 = append(f67, &f67elem)
			}
			cr.Status.AtProvider.ReadReplicaDBInstanceIdentifiers = f67
		} else {
			cr.Status.AtProvider.ReadReplicaDBInstanceIdentifiers = nil
		}
		if elem.ReadReplicaSourceDBInstanceIdentifier != nil {
			cr.Status.AtProvider.ReadReplicaSourceDBInstanceIdentifier = elem.ReadReplicaSourceDBInstanceIdentifier
		} else {
			cr.Status.AtProvider.ReadReplicaSourceDBInstanceIdentifier = nil
		}
		if elem.ReplicaMode != nil {
			cr.Status.AtProvider.ReplicaMode = elem.ReplicaMode
		} else {
			cr.Status.AtProvider.ReplicaMode = nil
		}
		if elem.ResumeFullAutomationModeTime != nil {
			cr.Status.AtProvider.ResumeFullAutomationModeTime = &metav1.Time{*elem.ResumeFullAutomationModeTime}
		} else {
			cr.Status.AtProvider.ResumeFullAutomationModeTime = nil
		}
		if elem.SecondaryAvailabilityZone != nil {
			cr.Status.AtProvider.SecondaryAvailabilityZone = elem.SecondaryAvailabilityZone
		} else {
			cr.Status.AtProvider.SecondaryAvailabilityZone = nil
		}
		if elem.StatusInfos != nil {
			f72 := []*svcapitypes.DBInstanceStatusInfo{}
			for _, f72iter := range elem.StatusInfos {
				f72elem := &svcapitypes.DBInstanceStatusInfo{}
				if f72iter.Message != nil {
					f72elem.Message = f72iter.Message
				}
				if f72iter.Normal != nil {
					f72elem.Normal = f72iter.Normal
				}
				if f72iter.Status != nil {
					f72elem.Status = f72iter.Status
				}
				if f72iter.StatusType != nil {
					f72elem.StatusType = f72iter.StatusType
				}
				f72 = append(f72, f72elem)
			}
			cr.Status.AtProvider.StatusInfos = f72
		} else {
			cr.Status.AtProvider.StatusInfos = nil
		}
		if elem.StorageEncrypted != nil {
			cr.Spec.ForProvider.StorageEncrypted = elem.StorageEncrypted
		} else {
			cr.Spec.ForProvider.StorageEncrypted = nil
		}
		if elem.StorageThroughput != nil {
			cr.Spec.ForProvider.StorageThroughput = elem.StorageThroughput
		} else {
			cr.Spec.ForProvider.StorageThroughput = nil
		}
		if elem.StorageType != nil {
			cr.Spec.ForProvider.StorageType = elem.StorageType
		} else {
			cr.Spec.ForProvider.StorageType = nil
		}
		if elem.TagList != nil {
			f76 := []*svcapitypes.Tag{}
			for _, f76iter := range elem.TagList {
				f76elem := &svcapitypes.Tag{}
				if f76iter.Key != nil {
					f76elem.Key = f76iter.Key
				}
				if f76iter.Value != nil {
					f76elem.Value = f76iter.Value
				}
				f76 = append(f76, f76elem)
			}
			cr.Status.AtProvider.TagList = f76
		} else {
			cr.Status.AtProvider.TagList = nil
		}
		if elem.TdeCredentialArn != nil {
			cr.Spec.ForProvider.TDECredentialARN = elem.TdeCredentialArn
		} else {
			cr.Spec.ForProvider.TDECredentialARN = nil
		}
		if elem.Timezone != nil {
			cr.Spec.ForProvider.Timezone = elem.Timezone
		} else {
			cr.Spec.ForProvider.Timezone = nil
		}
		if elem.VpcSecurityGroups != nil {
			f79 := []*svcapitypes.VPCSecurityGroupMembership{}
			for _, f79iter := range elem.VpcSecurityGroups {
				f79elem := &svcapitypes.VPCSecurityGroupMembership{}
				if f79iter.Status != nil {
					f79elem.Status = f79iter.Status
				}
				if f79iter.VpcSecurityGroupId != nil {
					f79elem.VPCSecurityGroupID = f79iter.VpcSecurityGroupId
				}
				f79 = append(f79, f79elem)
			}
			cr.Status.AtProvider.VPCSecurityGroups = f79
		} else {
			cr.Status.AtProvider.VPCSecurityGroups = nil
		}
		found = true
		break
	}
	if !found {
		return cr
	}

	return cr
}

// GenerateCreateDBInstanceInput returns a create input.
func GenerateCreateDBInstanceInput(cr *svcapitypes.DBInstance) *svcsdk.CreateDBInstanceInput {
	res := &svcsdk.CreateDBInstanceInput{}

	if cr.Spec.ForProvider.AllocatedStorage != nil {
		res.SetAllocatedStorage(*cr.Spec.ForProvider.AllocatedStorage)
	}
	if cr.Spec.ForProvider.AutoMinorVersionUpgrade != nil {
		res.SetAutoMinorVersionUpgrade(*cr.Spec.ForProvider.AutoMinorVersionUpgrade)
	}
	if cr.Spec.ForProvider.AvailabilityZone != nil {
		res.SetAvailabilityZone(*cr.Spec.ForProvider.AvailabilityZone)
	}
	if cr.Spec.ForProvider.BackupRetentionPeriod != nil {
		res.SetBackupRetentionPeriod(*cr.Spec.ForProvider.BackupRetentionPeriod)
	}
	if cr.Spec.ForProvider.BackupTarget != nil {
		res.SetBackupTarget(*cr.Spec.ForProvider.BackupTarget)
	}
	if cr.Spec.ForProvider.CACertificateIdentifier != nil {
		res.SetCACertificateIdentifier(*cr.Spec.ForProvider.CACertificateIdentifier)
	}
	if cr.Spec.ForProvider.CharacterSetName != nil {
		res.SetCharacterSetName(*cr.Spec.ForProvider.CharacterSetName)
	}
	if cr.Spec.ForProvider.CopyTagsToSnapshot != nil {
		res.SetCopyTagsToSnapshot(*cr.Spec.ForProvider.CopyTagsToSnapshot)
	}
	if cr.Spec.ForProvider.CustomIAMInstanceProfile != nil {
		res.SetCustomIamInstanceProfile(*cr.Spec.ForProvider.CustomIAMInstanceProfile)
	}
	if cr.Spec.ForProvider.DBClusterIdentifier != nil {
		res.SetDBClusterIdentifier(*cr.Spec.ForProvider.DBClusterIdentifier)
	}
	if cr.Spec.ForProvider.DBInstanceClass != nil {
		res.SetDBInstanceClass(*cr.Spec.ForProvider.DBInstanceClass)
	}
	if cr.Spec.ForProvider.DBName != nil {
		res.SetDBName(*cr.Spec.ForProvider.DBName)
	}
	if cr.Spec.ForProvider.DBParameterGroupName != nil {
		res.SetDBParameterGroupName(*cr.Spec.ForProvider.DBParameterGroupName)
	}
	if cr.Spec.ForProvider.DBSubnetGroupName != nil {
		res.SetDBSubnetGroupName(*cr.Spec.ForProvider.DBSubnetGroupName)
	}
	if cr.Spec.ForProvider.DeletionProtection != nil {
		res.SetDeletionProtection(*cr.Spec.ForProvider.DeletionProtection)
	}
	if cr.Spec.ForProvider.Domain != nil {
		res.SetDomain(*cr.Spec.ForProvider.Domain)
	}
	if cr.Spec.ForProvider.DomainIAMRoleName != nil {
		res.SetDomainIAMRoleName(*cr.Spec.ForProvider.DomainIAMRoleName)
	}
	if cr.Spec.ForProvider.EnableCloudwatchLogsExports != nil {
		f17 := []*string{}
		for _, f17iter := range cr.Spec.ForProvider.EnableCloudwatchLogsExports {
			var f17elem string
			f17elem = *f17iter
			f17 = append(f17, &f17elem)
		}
		res.SetEnableCloudwatchLogsExports(f17)
	}
	if cr.Spec.ForProvider.EnableCustomerOwnedIP != nil {
		res.SetEnableCustomerOwnedIp(*cr.Spec.ForProvider.EnableCustomerOwnedIP)
	}
	if cr.Spec.ForProvider.EnableIAMDatabaseAuthentication != nil {
		res.SetEnableIAMDatabaseAuthentication(*cr.Spec.ForProvider.EnableIAMDatabaseAuthentication)
	}
	if cr.Spec.ForProvider.EnablePerformanceInsights != nil {
		res.SetEnablePerformanceInsights(*cr.Spec.ForProvider.EnablePerformanceInsights)
	}
	if cr.Spec.ForProvider.Engine != nil {
		res.SetEngine(*cr.Spec.ForProvider.Engine)
	}
	if cr.Spec.ForProvider.EngineVersion != nil {
		res.SetEngineVersion(*cr.Spec.ForProvider.EngineVersion)
	}
	if cr.Spec.ForProvider.IOPS != nil {
		res.SetIops(*cr.Spec.ForProvider.IOPS)
	}
	if cr.Spec.ForProvider.KMSKeyID != nil {
		res.SetKmsKeyId(*cr.Spec.ForProvider.KMSKeyID)
	}
	if cr.Spec.ForProvider.LicenseModel != nil {
		res.SetLicenseModel(*cr.Spec.ForProvider.LicenseModel)
	}
	if cr.Spec.ForProvider.ManageMasterUserPassword != nil {
		res.SetManageMasterUserPassword(*cr.Spec.ForProvider.ManageMasterUserPassword)
	}
	if cr.Spec.ForProvider.MasterUserSecretKMSKeyID != nil {
		res.SetMasterUserSecretKmsKeyId(*cr.Spec.ForProvider.MasterUserSecretKMSKeyID)
	}
	if cr.Spec.ForProvider.MasterUsername != nil {
		res.SetMasterUsername(*cr.Spec.ForProvider.MasterUsername)
	}
	if cr.Spec.ForProvider.MaxAllocatedStorage != nil {
		res.SetMaxAllocatedStorage(*cr.Spec.ForProvider.MaxAllocatedStorage)
	}
	if cr.Spec.ForProvider.MonitoringInterval != nil {
		res.SetMonitoringInterval(*cr.Spec.ForProvider.MonitoringInterval)
	}
	if cr.Spec.ForProvider.MonitoringRoleARN != nil {
		res.SetMonitoringRoleArn(*cr.Spec.ForProvider.MonitoringRoleARN)
	}
	if cr.Spec.ForProvider.MultiAZ != nil {
		res.SetMultiAZ(*cr.Spec.ForProvider.MultiAZ)
	}
	if cr.Spec.ForProvider.NcharCharacterSetName != nil {
		res.SetNcharCharacterSetName(*cr.Spec.ForProvider.NcharCharacterSetName)
	}
	if cr.Spec.ForProvider.NetworkType != nil {
		res.SetNetworkType(*cr.Spec.ForProvider.NetworkType)
	}
	if cr.Spec.ForProvider.OptionGroupName != nil {
		res.SetOptionGroupName(*cr.Spec.ForProvider.OptionGroupName)
	}
	if cr.Spec.ForProvider.PerformanceInsightsKMSKeyID != nil {
		res.SetPerformanceInsightsKMSKeyId(*cr.Spec.ForProvider.PerformanceInsightsKMSKeyID)
	}
	if cr.Spec.ForProvider.PerformanceInsightsRetentionPeriod != nil {
		res.SetPerformanceInsightsRetentionPeriod(*cr.Spec.ForProvider.PerformanceInsightsRetentionPeriod)
	}
	if cr.Spec.ForProvider.Port != nil {
		res.SetPort(*cr.Spec.ForProvider.Port)
	}
	if cr.Spec.ForProvider.PreferredBackupWindow != nil {
		res.SetPreferredBackupWindow(*cr.Spec.ForProvider.PreferredBackupWindow)
	}
	if cr.Spec.ForProvider.PreferredMaintenanceWindow != nil {
		res.SetPreferredMaintenanceWindow(*cr.Spec.ForProvider.PreferredMaintenanceWindow)
	}
	if cr.Spec.ForProvider.ProcessorFeatures != nil {
		f41 := []*svcsdk.ProcessorFeature{}
		for _, f41iter := range cr.Spec.ForProvider.ProcessorFeatures {
			f41elem := &svcsdk.ProcessorFeature{}
			if f41iter.Name != nil {
				f41elem.SetName(*f41iter.Name)
			}
			if f41iter.Value != nil {
				f41elem.SetValue(*f41iter.Value)
			}
			f41 = append(f41, f41elem)
		}
		res.SetProcessorFeatures(f41)
	}
	if cr.Spec.ForProvider.PromotionTier != nil {
		res.SetPromotionTier(*cr.Spec.ForProvider.PromotionTier)
	}
	if cr.Spec.ForProvider.PubliclyAccessible != nil {
		res.SetPubliclyAccessible(*cr.Spec.ForProvider.PubliclyAccessible)
	}
	if cr.Spec.ForProvider.StorageEncrypted != nil {
		res.SetStorageEncrypted(*cr.Spec.ForProvider.StorageEncrypted)
	}
	if cr.Spec.ForProvider.StorageThroughput != nil {
		res.SetStorageThroughput(*cr.Spec.ForProvider.StorageThroughput)
	}
	if cr.Spec.ForProvider.StorageType != nil {
		res.SetStorageType(*cr.Spec.ForProvider.StorageType)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f47 := []*svcsdk.Tag{}
		for _, f47iter := range cr.Spec.ForProvider.Tags {
			f47elem := &svcsdk.Tag{}
			if f47iter.Key != nil {
				f47elem.SetKey(*f47iter.Key)
			}
			if f47iter.Value != nil {
				f47elem.SetValue(*f47iter.Value)
			}
			f47 = append(f47, f47elem)
		}
		res.SetTags(f47)
	}
	if cr.Spec.ForProvider.TDECredentialARN != nil {
		res.SetTdeCredentialArn(*cr.Spec.ForProvider.TDECredentialARN)
	}
	if cr.Spec.ForProvider.TDECredentialPassword != nil {
		res.SetTdeCredentialPassword(*cr.Spec.ForProvider.TDECredentialPassword)
	}
	if cr.Spec.ForProvider.Timezone != nil {
		res.SetTimezone(*cr.Spec.ForProvider.Timezone)
	}

	return res
}

// GenerateModifyDBInstanceInput returns an update input.
func GenerateModifyDBInstanceInput(cr *svcapitypes.DBInstance) *svcsdk.ModifyDBInstanceInput {
	res := &svcsdk.ModifyDBInstanceInput{}

	if cr.Spec.ForProvider.AllocatedStorage != nil {
		res.SetAllocatedStorage(*cr.Spec.ForProvider.AllocatedStorage)
	}
	if cr.Spec.ForProvider.AllowMajorVersionUpgrade != nil {
		res.SetAllowMajorVersionUpgrade(*cr.Spec.ForProvider.AllowMajorVersionUpgrade)
	}
	if cr.Spec.ForProvider.AutoMinorVersionUpgrade != nil {
		res.SetAutoMinorVersionUpgrade(*cr.Spec.ForProvider.AutoMinorVersionUpgrade)
	}
	if cr.Status.AtProvider.AutomationMode != nil {
		res.SetAutomationMode(*cr.Status.AtProvider.AutomationMode)
	}
	if cr.Status.AtProvider.AWSBackupRecoveryPointARN != nil {
		res.SetAwsBackupRecoveryPointArn(*cr.Status.AtProvider.AWSBackupRecoveryPointARN)
	}
	if cr.Spec.ForProvider.BackupRetentionPeriod != nil {
		res.SetBackupRetentionPeriod(*cr.Spec.ForProvider.BackupRetentionPeriod)
	}
	if cr.Spec.ForProvider.CACertificateIdentifier != nil {
		res.SetCACertificateIdentifier(*cr.Spec.ForProvider.CACertificateIdentifier)
	}
	if cr.Spec.ForProvider.CopyTagsToSnapshot != nil {
		res.SetCopyTagsToSnapshot(*cr.Spec.ForProvider.CopyTagsToSnapshot)
	}
	if cr.Spec.ForProvider.DBInstanceClass != nil {
		res.SetDBInstanceClass(*cr.Spec.ForProvider.DBInstanceClass)
	}
	if cr.Spec.ForProvider.DBParameterGroupName != nil {
		res.SetDBParameterGroupName(*cr.Spec.ForProvider.DBParameterGroupName)
	}
	if cr.Spec.ForProvider.DeletionProtection != nil {
		res.SetDeletionProtection(*cr.Spec.ForProvider.DeletionProtection)
	}
	if cr.Spec.ForProvider.Domain != nil {
		res.SetDomain(*cr.Spec.ForProvider.Domain)
	}
	if cr.Spec.ForProvider.DomainIAMRoleName != nil {
		res.SetDomainIAMRoleName(*cr.Spec.ForProvider.DomainIAMRoleName)
	}
	if cr.Spec.ForProvider.EnableCustomerOwnedIP != nil {
		res.SetEnableCustomerOwnedIp(*cr.Spec.ForProvider.EnableCustomerOwnedIP)
	}
	if cr.Spec.ForProvider.EnableIAMDatabaseAuthentication != nil {
		res.SetEnableIAMDatabaseAuthentication(*cr.Spec.ForProvider.EnableIAMDatabaseAuthentication)
	}
	if cr.Spec.ForProvider.EnablePerformanceInsights != nil {
		res.SetEnablePerformanceInsights(*cr.Spec.ForProvider.EnablePerformanceInsights)
	}
	if cr.Spec.ForProvider.EngineVersion != nil {
		res.SetEngineVersion(*cr.Spec.ForProvider.EngineVersion)
	}
	if cr.Spec.ForProvider.IOPS != nil {
		res.SetIops(*cr.Spec.ForProvider.IOPS)
	}
	if cr.Spec.ForProvider.LicenseModel != nil {
		res.SetLicenseModel(*cr.Spec.ForProvider.LicenseModel)
	}
	if cr.Spec.ForProvider.ManageMasterUserPassword != nil {
		res.SetManageMasterUserPassword(*cr.Spec.ForProvider.ManageMasterUserPassword)
	}
	if cr.Spec.ForProvider.MasterUserSecretKMSKeyID != nil {
		res.SetMasterUserSecretKmsKeyId(*cr.Spec.ForProvider.MasterUserSecretKMSKeyID)
	}
	if cr.Spec.ForProvider.MaxAllocatedStorage != nil {
		res.SetMaxAllocatedStorage(*cr.Spec.ForProvider.MaxAllocatedStorage)
	}
	if cr.Spec.ForProvider.MonitoringInterval != nil {
		res.SetMonitoringInterval(*cr.Spec.ForProvider.MonitoringInterval)
	}
	if cr.Spec.ForProvider.MonitoringRoleARN != nil {
		res.SetMonitoringRoleArn(*cr.Spec.ForProvider.MonitoringRoleARN)
	}
	if cr.Spec.ForProvider.MultiAZ != nil {
		res.SetMultiAZ(*cr.Spec.ForProvider.MultiAZ)
	}
	if cr.Spec.ForProvider.NetworkType != nil {
		res.SetNetworkType(*cr.Spec.ForProvider.NetworkType)
	}
	if cr.Spec.ForProvider.OptionGroupName != nil {
		res.SetOptionGroupName(*cr.Spec.ForProvider.OptionGroupName)
	}
	if cr.Spec.ForProvider.PerformanceInsightsKMSKeyID != nil {
		res.SetPerformanceInsightsKMSKeyId(*cr.Spec.ForProvider.PerformanceInsightsKMSKeyID)
	}
	if cr.Spec.ForProvider.PerformanceInsightsRetentionPeriod != nil {
		res.SetPerformanceInsightsRetentionPeriod(*cr.Spec.ForProvider.PerformanceInsightsRetentionPeriod)
	}
	if cr.Spec.ForProvider.PreferredBackupWindow != nil {
		res.SetPreferredBackupWindow(*cr.Spec.ForProvider.PreferredBackupWindow)
	}
	if cr.Spec.ForProvider.PreferredMaintenanceWindow != nil {
		res.SetPreferredMaintenanceWindow(*cr.Spec.ForProvider.PreferredMaintenanceWindow)
	}
	if cr.Spec.ForProvider.ProcessorFeatures != nil {
		f36 := []*svcsdk.ProcessorFeature{}
		for _, f36iter := range cr.Spec.ForProvider.ProcessorFeatures {
			f36elem := &svcsdk.ProcessorFeature{}
			if f36iter.Name != nil {
				f36elem.SetName(*f36iter.Name)
			}
			if f36iter.Value != nil {
				f36elem.SetValue(*f36iter.Value)
			}
			f36 = append(f36, f36elem)
		}
		res.SetProcessorFeatures(f36)
	}
	if cr.Spec.ForProvider.PromotionTier != nil {
		res.SetPromotionTier(*cr.Spec.ForProvider.PromotionTier)
	}
	if cr.Spec.ForProvider.PubliclyAccessible != nil {
		res.SetPubliclyAccessible(*cr.Spec.ForProvider.PubliclyAccessible)
	}
	if cr.Status.AtProvider.ReplicaMode != nil {
		res.SetReplicaMode(*cr.Status.AtProvider.ReplicaMode)
	}
	if cr.Spec.ForProvider.StorageThroughput != nil {
		res.SetStorageThroughput(*cr.Spec.ForProvider.StorageThroughput)
	}
	if cr.Spec.ForProvider.StorageType != nil {
		res.SetStorageType(*cr.Spec.ForProvider.StorageType)
	}
	if cr.Spec.ForProvider.TDECredentialARN != nil {
		res.SetTdeCredentialArn(*cr.Spec.ForProvider.TDECredentialARN)
	}
	if cr.Spec.ForProvider.TDECredentialPassword != nil {
		res.SetTdeCredentialPassword(*cr.Spec.ForProvider.TDECredentialPassword)
	}

	return res
}

// GenerateDeleteDBInstanceInput returns a deletion input.
func GenerateDeleteDBInstanceInput(cr *svcapitypes.DBInstance) *svcsdk.DeleteDBInstanceInput {
	res := &svcsdk.DeleteDBInstanceInput{}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "DBInstanceNotFound"
}
