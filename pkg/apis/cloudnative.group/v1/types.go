package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type VirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineSpec   `json:"spec,omitempty"`
	Status VirtualMachineStatus `json:"status,omitempty"`
}

type VirtualMachineSpec struct {
	Name   string `json:"name,omitempty"`
	Image  string `json:"image,omitempty"`
	Cpu    int    `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
}

type VirtualMachinePhase string

const (
	VirtualMachineNone                       VirtualMachinePhase = ""
	VirtualMachinePending                    VirtualMachinePhase = "Pending"
	VirtualMachineRunning                    VirtualMachinePhase = "Running"
	VirtualMachineSucceeded                  VirtualMachinePhase = "Succeeded"
	VirtualMachineFailed                     VirtualMachinePhase = "Failed"
	VirtualMachineUnknown                    VirtualMachinePhase = "Unknown"
	VirtualMachineUnschedulable              VirtualMachinePhase = "Unschedulable"
	VirtualMachineStopped                    VirtualMachinePhase = "Stopped"
	VirtualMachineTerminating                VirtualMachinePhase = "Terminating"
	VirtualMachineDeleting                   VirtualMachinePhase = "Deleting"
	VirtualMachineDeleted                    VirtualMachinePhase = "Deleted"
	VirtualMachinePaused                     VirtualMachinePhase = "Paused"
	VirtualMachineSuspended                  VirtualMachinePhase = "Suspended"
	VirtualMachineEvicted                    VirtualMachinePhase = "Evicted"
	VirtualMachineDraining                   VirtualMachinePhase = "Draining"
	VirtualMachineCompleted                  VirtualMachinePhase = "Completed"
	VirtualMachineAborted                    VirtualMachinePhase = "Aborted"
	VirtualMachineUnknownVirtualMachinePhase VirtualMachinePhase = "UnknownVirtualMachinePhase"
	VirtualMachineUnset                      VirtualMachinePhase = "Unset"
	VirtualMachineScheduled                  VirtualMachinePhase = "Scheduled"
	VirtualMachineScheduling                 VirtualMachinePhase = "Scheduling"
	VirtualMachineScheduledToRunning         VirtualMachinePhase = "ScheduledToRunning"
	VirtualMachineScheduledToStopped         VirtualMachinePhase = "ScheduledToStopped"
)

type ResourceUsage struct {
	CPU    float64 `json:"cpu"`
	Memory float64 `json:"memory"`
}

type ServerStatus struct {
	ID    string        `json:"id"`
	State string        `json:"state"`
	Usage ResourceUsage `json:"usage"`
}

type VirtualMachineStatus struct {
	Phase          VirtualMachinePhase `json:"phase"`
	Reason         string              `json:"reason,omitempty"`
	Server         ServerStatus        `json:"server,omitempty"`
	LastUpdateTime metav1.Time         `json:"lastUpdateTime"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type VirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []VirtualMachine `json:"items"`
}
