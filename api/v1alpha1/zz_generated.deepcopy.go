// +build !ignore_autogenerated

/*
Copyright 2021.

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineDeletionRemediation) DeepCopyInto(out *MachineDeletionRemediation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineDeletionRemediation.
func (in *MachineDeletionRemediation) DeepCopy() *MachineDeletionRemediation {
	if in == nil {
		return nil
	}
	out := new(MachineDeletionRemediation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineDeletionRemediation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineDeletionRemediationList) DeepCopyInto(out *MachineDeletionRemediationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineDeletionRemediation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineDeletionRemediationList.
func (in *MachineDeletionRemediationList) DeepCopy() *MachineDeletionRemediationList {
	if in == nil {
		return nil
	}
	out := new(MachineDeletionRemediationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineDeletionRemediationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineDeletionRemediationSpec) DeepCopyInto(out *MachineDeletionRemediationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineDeletionRemediationSpec.
func (in *MachineDeletionRemediationSpec) DeepCopy() *MachineDeletionRemediationSpec {
	if in == nil {
		return nil
	}
	out := new(MachineDeletionRemediationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineDeletionRemediationStatus) DeepCopyInto(out *MachineDeletionRemediationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineDeletionRemediationStatus.
func (in *MachineDeletionRemediationStatus) DeepCopy() *MachineDeletionRemediationStatus {
	if in == nil {
		return nil
	}
	out := new(MachineDeletionRemediationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineDeletionRemediationTemplate) DeepCopyInto(out *MachineDeletionRemediationTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineDeletionRemediationTemplate.
func (in *MachineDeletionRemediationTemplate) DeepCopy() *MachineDeletionRemediationTemplate {
	if in == nil {
		return nil
	}
	out := new(MachineDeletionRemediationTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineDeletionRemediationTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineDeletionRemediationTemplateList) DeepCopyInto(out *MachineDeletionRemediationTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineDeletionRemediationTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineDeletionRemediationTemplateList.
func (in *MachineDeletionRemediationTemplateList) DeepCopy() *MachineDeletionRemediationTemplateList {
	if in == nil {
		return nil
	}
	out := new(MachineDeletionRemediationTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineDeletionRemediationTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineDeletionRemediationTemplateSpec) DeepCopyInto(out *MachineDeletionRemediationTemplateSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineDeletionRemediationTemplateSpec.
func (in *MachineDeletionRemediationTemplateSpec) DeepCopy() *MachineDeletionRemediationTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(MachineDeletionRemediationTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineDeletionRemediationTemplateStatus) DeepCopyInto(out *MachineDeletionRemediationTemplateStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineDeletionRemediationTemplateStatus.
func (in *MachineDeletionRemediationTemplateStatus) DeepCopy() *MachineDeletionRemediationTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(MachineDeletionRemediationTemplateStatus)
	in.DeepCopyInto(out)
	return out
}
