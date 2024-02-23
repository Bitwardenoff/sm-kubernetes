//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2024 Bitwarden, Inc..

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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthToken) DeepCopyInto(out *AuthToken) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthToken.
func (in *AuthToken) DeepCopy() *AuthToken {
	if in == nil {
		return nil
	}
	out := new(AuthToken)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BitwardenSecret) DeepCopyInto(out *BitwardenSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BitwardenSecret.
func (in *BitwardenSecret) DeepCopy() *BitwardenSecret {
	if in == nil {
		return nil
	}
	out := new(BitwardenSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BitwardenSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BitwardenSecretList) DeepCopyInto(out *BitwardenSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BitwardenSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BitwardenSecretList.
func (in *BitwardenSecretList) DeepCopy() *BitwardenSecretList {
	if in == nil {
		return nil
	}
	out := new(BitwardenSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BitwardenSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BitwardenSecretSpec) DeepCopyInto(out *BitwardenSecretSpec) {
	*out = *in
	if in.SecretMap != nil {
		in, out := &in.SecretMap, &out.SecretMap
		*out = make([]SecretMap, len(*in))
		copy(*out, *in)
	}
	out.AuthToken = in.AuthToken
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BitwardenSecretSpec.
func (in *BitwardenSecretSpec) DeepCopy() *BitwardenSecretSpec {
	if in == nil {
		return nil
	}
	out := new(BitwardenSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BitwardenSecretStatus) DeepCopyInto(out *BitwardenSecretStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BitwardenSecretStatus.
func (in *BitwardenSecretStatus) DeepCopy() *BitwardenSecretStatus {
	if in == nil {
		return nil
	}
	out := new(BitwardenSecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretMap) DeepCopyInto(out *SecretMap) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretMap.
func (in *SecretMap) DeepCopy() *SecretMap {
	if in == nil {
		return nil
	}
	out := new(SecretMap)
	in.DeepCopyInto(out)
	return out
}
