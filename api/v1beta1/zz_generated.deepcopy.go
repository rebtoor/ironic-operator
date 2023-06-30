//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Red Hat Inc.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/openstack-k8s-operators/lib-common/modules/common/condition"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DHCPRange) DeepCopyInto(out *DHCPRange) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DHCPRange.
func (in *DHCPRange) DeepCopy() *DHCPRange {
	if in == nil {
		return nil
	}
	out := new(DHCPRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ironic) DeepCopyInto(out *Ironic) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ironic.
func (in *Ironic) DeepCopy() *Ironic {
	if in == nil {
		return nil
	}
	out := new(Ironic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ironic) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicAPI) DeepCopyInto(out *IronicAPI) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicAPI.
func (in *IronicAPI) DeepCopy() *IronicAPI {
	if in == nil {
		return nil
	}
	out := new(IronicAPI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IronicAPI) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicAPIList) DeepCopyInto(out *IronicAPIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IronicAPI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicAPIList.
func (in *IronicAPIList) DeepCopy() *IronicAPIList {
	if in == nil {
		return nil
	}
	out := new(IronicAPIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IronicAPIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicAPISpec) DeepCopyInto(out *IronicAPISpec) {
	*out = *in
	in.IronicAPITemplate.DeepCopyInto(&out.IronicAPITemplate)
	out.PasswordSelectors = in.PasswordSelectors
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicAPISpec.
func (in *IronicAPISpec) DeepCopy() *IronicAPISpec {
	if in == nil {
		return nil
	}
	out := new(IronicAPISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicAPIStatus) DeepCopyInto(out *IronicAPIStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.APIEndpoints != nil {
		in, out := &in.APIEndpoints, &out.APIEndpoints
		*out = make(map[string]map[string]string, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicAPIStatus.
func (in *IronicAPIStatus) DeepCopy() *IronicAPIStatus {
	if in == nil {
		return nil
	}
	out := new(IronicAPIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicAPITemplate) DeepCopyInto(out *IronicAPITemplate) {
	*out = *in
	in.IronicServiceTemplate.DeepCopyInto(&out.IronicServiceTemplate)
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExternalEndpoints != nil {
		in, out := &in.ExternalEndpoints, &out.ExternalEndpoints
		*out = make([]MetalLBConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicAPITemplate.
func (in *IronicAPITemplate) DeepCopy() *IronicAPITemplate {
	if in == nil {
		return nil
	}
	out := new(IronicAPITemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicConductor) DeepCopyInto(out *IronicConductor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicConductor.
func (in *IronicConductor) DeepCopy() *IronicConductor {
	if in == nil {
		return nil
	}
	out := new(IronicConductor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IronicConductor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicConductorList) DeepCopyInto(out *IronicConductorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IronicConductor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicConductorList.
func (in *IronicConductorList) DeepCopy() *IronicConductorList {
	if in == nil {
		return nil
	}
	out := new(IronicConductorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IronicConductorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicConductorSpec) DeepCopyInto(out *IronicConductorSpec) {
	*out = *in
	in.IronicConductorTemplate.DeepCopyInto(&out.IronicConductorTemplate)
	out.PasswordSelectors = in.PasswordSelectors
	if in.KeystoneVars != nil {
		in, out := &in.KeystoneVars, &out.KeystoneVars
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicConductorSpec.
func (in *IronicConductorSpec) DeepCopy() *IronicConductorSpec {
	if in == nil {
		return nil
	}
	out := new(IronicConductorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicConductorStatus) DeepCopyInto(out *IronicConductorStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicConductorStatus.
func (in *IronicConductorStatus) DeepCopy() *IronicConductorStatus {
	if in == nil {
		return nil
	}
	out := new(IronicConductorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicConductorTemplate) DeepCopyInto(out *IronicConductorTemplate) {
	*out = *in
	in.IronicServiceTemplate.DeepCopyInto(&out.IronicServiceTemplate)
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DHCPRanges != nil {
		in, out := &in.DHCPRanges, &out.DHCPRanges
		*out = make([]DHCPRange, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicConductorTemplate.
func (in *IronicConductorTemplate) DeepCopy() *IronicConductorTemplate {
	if in == nil {
		return nil
	}
	out := new(IronicConductorTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicDBSyncDebug) DeepCopyInto(out *IronicDBSyncDebug) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicDBSyncDebug.
func (in *IronicDBSyncDebug) DeepCopy() *IronicDBSyncDebug {
	if in == nil {
		return nil
	}
	out := new(IronicDBSyncDebug)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicDebug) DeepCopyInto(out *IronicDebug) {
	*out = *in
	out.IronicDBSyncDebug = in.IronicDBSyncDebug
	out.IronicServiceDebug = in.IronicServiceDebug
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicDebug.
func (in *IronicDebug) DeepCopy() *IronicDebug {
	if in == nil {
		return nil
	}
	out := new(IronicDebug)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicDefaults) DeepCopyInto(out *IronicDefaults) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicDefaults.
func (in *IronicDefaults) DeepCopy() *IronicDefaults {
	if in == nil {
		return nil
	}
	out := new(IronicDefaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicImages) DeepCopyInto(out *IronicImages) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicImages.
func (in *IronicImages) DeepCopy() *IronicImages {
	if in == nil {
		return nil
	}
	out := new(IronicImages)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicInspector) DeepCopyInto(out *IronicInspector) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicInspector.
func (in *IronicInspector) DeepCopy() *IronicInspector {
	if in == nil {
		return nil
	}
	out := new(IronicInspector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IronicInspector) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicInspectorList) DeepCopyInto(out *IronicInspectorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IronicInspector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicInspectorList.
func (in *IronicInspectorList) DeepCopy() *IronicInspectorList {
	if in == nil {
		return nil
	}
	out := new(IronicInspectorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IronicInspectorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicInspectorPasswordSelector) DeepCopyInto(out *IronicInspectorPasswordSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicInspectorPasswordSelector.
func (in *IronicInspectorPasswordSelector) DeepCopy() *IronicInspectorPasswordSelector {
	if in == nil {
		return nil
	}
	out := new(IronicInspectorPasswordSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicInspectorSpec) DeepCopyInto(out *IronicInspectorSpec) {
	*out = *in
	in.IronicInspectorTemplate.DeepCopyInto(&out.IronicInspectorTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicInspectorSpec.
func (in *IronicInspectorSpec) DeepCopy() *IronicInspectorSpec {
	if in == nil {
		return nil
	}
	out := new(IronicInspectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicInspectorStatus) DeepCopyInto(out *IronicInspectorStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.APIEndpoints != nil {
		in, out := &in.APIEndpoints, &out.APIEndpoints
		*out = make(map[string]map[string]string, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicInspectorStatus.
func (in *IronicInspectorStatus) DeepCopy() *IronicInspectorStatus {
	if in == nil {
		return nil
	}
	out := new(IronicInspectorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicInspectorTemplate) DeepCopyInto(out *IronicInspectorTemplate) {
	*out = *in
	out.PasswordSelectors = in.PasswordSelectors
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Debug = in.Debug
	if in.DefaultConfigOverwrite != nil {
		in, out := &in.DefaultConfigOverwrite, &out.DefaultConfigOverwrite
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExternalEndpoints != nil {
		in, out := &in.ExternalEndpoints, &out.ExternalEndpoints
		*out = make([]MetalLBConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DHCPRanges != nil {
		in, out := &in.DHCPRanges, &out.DHCPRanges
		*out = make([]DHCPRange, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicInspectorTemplate.
func (in *IronicInspectorTemplate) DeepCopy() *IronicInspectorTemplate {
	if in == nil {
		return nil
	}
	out := new(IronicInspectorTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicList) DeepCopyInto(out *IronicList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Ironic, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicList.
func (in *IronicList) DeepCopy() *IronicList {
	if in == nil {
		return nil
	}
	out := new(IronicList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IronicList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicNeutronAgent) DeepCopyInto(out *IronicNeutronAgent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicNeutronAgent.
func (in *IronicNeutronAgent) DeepCopy() *IronicNeutronAgent {
	if in == nil {
		return nil
	}
	out := new(IronicNeutronAgent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IronicNeutronAgent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicNeutronAgentList) DeepCopyInto(out *IronicNeutronAgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IronicNeutronAgent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicNeutronAgentList.
func (in *IronicNeutronAgentList) DeepCopy() *IronicNeutronAgentList {
	if in == nil {
		return nil
	}
	out := new(IronicNeutronAgentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IronicNeutronAgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicNeutronAgentSpec) DeepCopyInto(out *IronicNeutronAgentSpec) {
	*out = *in
	in.IronicNeutronAgentTemplate.DeepCopyInto(&out.IronicNeutronAgentTemplate)
	out.PasswordSelectors = in.PasswordSelectors
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicNeutronAgentSpec.
func (in *IronicNeutronAgentSpec) DeepCopy() *IronicNeutronAgentSpec {
	if in == nil {
		return nil
	}
	out := new(IronicNeutronAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicNeutronAgentStatus) DeepCopyInto(out *IronicNeutronAgentStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicNeutronAgentStatus.
func (in *IronicNeutronAgentStatus) DeepCopy() *IronicNeutronAgentStatus {
	if in == nil {
		return nil
	}
	out := new(IronicNeutronAgentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicNeutronAgentTemplate) DeepCopyInto(out *IronicNeutronAgentTemplate) {
	*out = *in
	in.IronicServiceTemplate.DeepCopyInto(&out.IronicServiceTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicNeutronAgentTemplate.
func (in *IronicNeutronAgentTemplate) DeepCopy() *IronicNeutronAgentTemplate {
	if in == nil {
		return nil
	}
	out := new(IronicNeutronAgentTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicServiceDebug) DeepCopyInto(out *IronicServiceDebug) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicServiceDebug.
func (in *IronicServiceDebug) DeepCopy() *IronicServiceDebug {
	if in == nil {
		return nil
	}
	out := new(IronicServiceDebug)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicServiceTemplate) DeepCopyInto(out *IronicServiceTemplate) {
	*out = *in
	out.Debug = in.Debug
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.DefaultConfigOverwrite != nil {
		in, out := &in.DefaultConfigOverwrite, &out.DefaultConfigOverwrite
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicServiceTemplate.
func (in *IronicServiceTemplate) DeepCopy() *IronicServiceTemplate {
	if in == nil {
		return nil
	}
	out := new(IronicServiceTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicSpec) DeepCopyInto(out *IronicSpec) {
	*out = *in
	out.Images = in.Images
	out.PasswordSelectors = in.PasswordSelectors
	out.Debug = in.Debug
	if in.DefaultConfigOverwrite != nil {
		in, out := &in.DefaultConfigOverwrite, &out.DefaultConfigOverwrite
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.IronicAPI.DeepCopyInto(&out.IronicAPI)
	if in.IronicConductors != nil {
		in, out := &in.IronicConductors, &out.IronicConductors
		*out = make([]IronicConductorTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.IronicInspector.DeepCopyInto(&out.IronicInspector)
	in.IronicNeutronAgent.DeepCopyInto(&out.IronicNeutronAgent)
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicSpec.
func (in *IronicSpec) DeepCopy() *IronicSpec {
	if in == nil {
		return nil
	}
	out := new(IronicSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicStatus) DeepCopyInto(out *IronicStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.APIEndpoints != nil {
		in, out := &in.APIEndpoints, &out.APIEndpoints
		*out = make(map[string]map[string]string, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.IronicConductorReadyCount != nil {
		in, out := &in.IronicConductorReadyCount, &out.IronicConductorReadyCount
		*out = make(map[string]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicStatus.
func (in *IronicStatus) DeepCopy() *IronicStatus {
	if in == nil {
		return nil
	}
	out := new(IronicStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetalLBConfig) DeepCopyInto(out *MetalLBConfig) {
	*out = *in
	if in.LoadBalancerIPs != nil {
		in, out := &in.LoadBalancerIPs, &out.LoadBalancerIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetalLBConfig.
func (in *MetalLBConfig) DeepCopy() *MetalLBConfig {
	if in == nil {
		return nil
	}
	out := new(MetalLBConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PasswordSelector) DeepCopyInto(out *PasswordSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PasswordSelector.
func (in *PasswordSelector) DeepCopy() *PasswordSelector {
	if in == nil {
		return nil
	}
	out := new(PasswordSelector)
	in.DeepCopyInto(out)
	return out
}
