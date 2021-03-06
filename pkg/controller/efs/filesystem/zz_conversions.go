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

package filesystem

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/efs"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/crossplane/provider-aws/apis/efs/v1alpha1"
	awsclients "github.com/crossplane/provider-aws/pkg/clients"
)

// GenerateDescribeFileSystemsInput returns input for read
// operation.
func GenerateDescribeFileSystemsInput(cr *svcapitypes.FileSystem) *svcsdk.DescribeFileSystemsInput {
	res := &svcsdk.DescribeFileSystemsInput{}

	if cr.Status.AtProvider.CreationToken != nil {
		res.SetCreationToken(*cr.Status.AtProvider.CreationToken)
	}
	if cr.Status.AtProvider.FileSystemID != nil {
		res.SetFileSystemId(*cr.Status.AtProvider.FileSystemID)
	}

	return res
}

// GenerateFileSystem returns the current state in the form of *svcapitypes.FileSystem.
func GenerateFileSystem(resp *svcsdk.DescribeFileSystemsOutput) *svcapitypes.FileSystem {
	cr := &svcapitypes.FileSystem{}

	found := false
	for _, elem := range resp.FileSystems {
		if elem.CreationTime != nil {
			cr.Status.AtProvider.CreationTime = &metav1.Time{*elem.CreationTime}
		}
		if elem.CreationToken != nil {
			cr.Status.AtProvider.CreationToken = elem.CreationToken
		}
		if elem.Encrypted != nil {
			cr.Spec.ForProvider.Encrypted = elem.Encrypted
		}
		if elem.FileSystemArn != nil {
			cr.Status.AtProvider.FileSystemARN = elem.FileSystemArn
		}
		if elem.FileSystemId != nil {
			cr.Status.AtProvider.FileSystemID = elem.FileSystemId
		}
		if elem.KmsKeyId != nil {
			cr.Spec.ForProvider.KMSKeyID = elem.KmsKeyId
		}
		if elem.LifeCycleState != nil {
			cr.Status.AtProvider.LifeCycleState = elem.LifeCycleState
		}
		if elem.Name != nil {
			cr.Status.AtProvider.Name = elem.Name
		}
		if elem.NumberOfMountTargets != nil {
			cr.Status.AtProvider.NumberOfMountTargets = elem.NumberOfMountTargets
		}
		if elem.OwnerId != nil {
			cr.Status.AtProvider.OwnerID = elem.OwnerId
		}
		if elem.PerformanceMode != nil {
			cr.Spec.ForProvider.PerformanceMode = elem.PerformanceMode
		}
		if elem.SizeInBytes != nil {
			f11 := &svcapitypes.FileSystemSize{}
			if elem.SizeInBytes.Timestamp != nil {
				f11.Timestamp = &metav1.Time{*elem.SizeInBytes.Timestamp}
			}
			if elem.SizeInBytes.Value != nil {
				f11.Value = elem.SizeInBytes.Value
			}
			if elem.SizeInBytes.ValueInIA != nil {
				f11.ValueInIA = elem.SizeInBytes.ValueInIA
			}
			if elem.SizeInBytes.ValueInStandard != nil {
				f11.ValueInStandard = elem.SizeInBytes.ValueInStandard
			}
			cr.Status.AtProvider.SizeInBytes = f11
		}
		if elem.Tags != nil {
			f12 := []*svcapitypes.Tag{}
			for _, f12iter := range elem.Tags {
				f12elem := &svcapitypes.Tag{}
				if f12iter.Key != nil {
					f12elem.Key = f12iter.Key
				}
				if f12iter.Value != nil {
					f12elem.Value = f12iter.Value
				}
				f12 = append(f12, f12elem)
			}
			cr.Spec.ForProvider.Tags = f12
		}
		if elem.ThroughputMode != nil {
			cr.Spec.ForProvider.ThroughputMode = elem.ThroughputMode
		}
		found = true
		break
	}
	if !found {
		return cr
	}

	return cr
}

func lateInitialize(cr *svcapitypes.FileSystem, resp *svcsdk.DescribeFileSystemsOutput) error {
	for _, resource := range resp.FileSystems {
		cr.Spec.ForProvider.Encrypted = awsclients.LateInitializeBoolPtr(cr.Spec.ForProvider.Encrypted, resource.Encrypted)
		cr.Spec.ForProvider.KMSKeyID = awsclients.LateInitializeStringPtr(cr.Spec.ForProvider.KMSKeyID, resource.KmsKeyId)
		cr.Spec.ForProvider.PerformanceMode = awsclients.LateInitializeStringPtr(cr.Spec.ForProvider.PerformanceMode, resource.PerformanceMode)
		if len(resource.Tags) != 0 && len(cr.Spec.ForProvider.Tags) == 0 {
			cr.Spec.ForProvider.Tags = make([]*svcapitypes.Tag, len(resource.Tags))
			for i0 := range resource.Tags {
				if resource.Tags[i0] != nil {
					if cr.Spec.ForProvider.Tags[i0] == nil {
						cr.Spec.ForProvider.Tags[i0] = &svcapitypes.Tag{}
					}
					cr.Spec.ForProvider.Tags[i0].Key = awsclients.LateInitializeStringPtr(cr.Spec.ForProvider.Tags[i0].Key, resource.Tags[i0].Key)
					cr.Spec.ForProvider.Tags[i0].Value = awsclients.LateInitializeStringPtr(cr.Spec.ForProvider.Tags[i0].Value, resource.Tags[i0].Value)
				}
			}
		}
		cr.Spec.ForProvider.ThroughputMode = awsclients.LateInitializeStringPtr(cr.Spec.ForProvider.ThroughputMode, resource.ThroughputMode)
	}
	return nil
}

func basicUpToDateCheck(cr *svcapitypes.FileSystem, resp *svcsdk.DescribeFileSystemsOutput) bool {
	for _, resource := range resp.FileSystems {
		if awsclients.StringValue(cr.Spec.ForProvider.ThroughputMode) != awsclients.StringValue(resource.ThroughputMode) {
			return false
		}
	}
	return true
}

// GenerateCreateFileSystemInput returns a create input.
func GenerateCreateFileSystemInput(cr *svcapitypes.FileSystem) *svcsdk.CreateFileSystemInput {
	res := &svcsdk.CreateFileSystemInput{}

	if cr.Spec.ForProvider.Encrypted != nil {
		res.SetEncrypted(*cr.Spec.ForProvider.Encrypted)
	}
	if cr.Spec.ForProvider.KMSKeyID != nil {
		res.SetKmsKeyId(*cr.Spec.ForProvider.KMSKeyID)
	}
	if cr.Spec.ForProvider.PerformanceMode != nil {
		res.SetPerformanceMode(*cr.Spec.ForProvider.PerformanceMode)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f3 := []*svcsdk.Tag{}
		for _, f3iter := range cr.Spec.ForProvider.Tags {
			f3elem := &svcsdk.Tag{}
			if f3iter.Key != nil {
				f3elem.SetKey(*f3iter.Key)
			}
			if f3iter.Value != nil {
				f3elem.SetValue(*f3iter.Value)
			}
			f3 = append(f3, f3elem)
		}
		res.SetTags(f3)
	}
	if cr.Spec.ForProvider.ThroughputMode != nil {
		res.SetThroughputMode(*cr.Spec.ForProvider.ThroughputMode)
	}

	return res
}

// GenerateUpdateFileSystemInput returns an update input.
func GenerateUpdateFileSystemInput(cr *svcapitypes.FileSystem) *svcsdk.UpdateFileSystemInput {
	res := &svcsdk.UpdateFileSystemInput{}

	if cr.Status.AtProvider.FileSystemID != nil {
		res.SetFileSystemId(*cr.Status.AtProvider.FileSystemID)
	}
	if cr.Spec.ForProvider.ThroughputMode != nil {
		res.SetThroughputMode(*cr.Spec.ForProvider.ThroughputMode)
	}

	return res
}

// GenerateDeleteFileSystemInput returns a deletion input.
func GenerateDeleteFileSystemInput(cr *svcapitypes.FileSystem) *svcsdk.DeleteFileSystemInput {
	res := &svcsdk.DeleteFileSystemInput{}

	if cr.Status.AtProvider.FileSystemID != nil {
		res.SetFileSystemId(*cr.Status.AtProvider.FileSystemID)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "FileSystemNotFound"
}
