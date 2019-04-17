/*
Copyright 2018 The OpenEBS Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=backupcstor

// BackupCStor describes a cstor volume resource created as custom resource
type BackupCStor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupCStorSpec   `json:"spec"`
	Status            BackupCStorStatus `json:"status"`
}

// BackupCStorSpec is the spec for a BackupCstor resource
type BackupCStorSpec struct {
	//	BackupID     string `json:"backupId"`
	BackupName   string `json:"backupName"`
	VolumeName   string `json:"volumeName"`
	SnapName     string `json:"snapName"`
	PrevSnapName string `json:"prevSnapName"`
	BackupDest   string `json:"backupDest"`
}

// BackupCStorStatus is to hold result of action.
type BackupCStorStatus string

// Status written onto BackupCstor objects.
const (
	// BKPCStorStatusEmpty ensures the create operation is to be done, if import fails.
	BKPCStorStatusEmpty BackupCStorStatus = ""

	// BKPCStorStatusDone , backup is completed.
	BKPCStorStatusDone BackupCStorStatus = "Done"
	// BKPCStorStatusFailed , backup is failed.
	BKPCStorStatusFailed BackupCStorStatus = "Failed"
	// BKPCStorStatusInit , backup is initialized.
	BKPCStorStatusInit BackupCStorStatus = "Init"
	// BKPCStorStatusPending , backup is pending.
	BKPCStorStatusPending BackupCStorStatus = "Pending"
	// BKPCStorStatusInProgress , backup is in progress.
	BKPCStorStatusInProgress BackupCStorStatus = "Pending"
	// BKPCStorStatusInvalid , backup operation is invalid.
	BKPCStorStatusInvalid BackupCStorStatus = "Invalid"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=backupcstor

// BackupCStorList is a list of BackupCstor resources
type BackupCStorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []BackupCStor `json:"items"`
}