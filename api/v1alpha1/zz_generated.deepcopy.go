//go:build !ignore_autogenerated

/*
Copyright 2024.

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
	"k8s.io/api/autoscaling/v2"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthConfig) DeepCopyInto(out *AuthConfig) {
	*out = *in
	if in.OAuthConfig != nil {
		in, out := &in.OAuthConfig, &out.OAuthConfig
		*out = new(OAuthConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthConfig.
func (in *AuthConfig) DeepCopy() *AuthConfig {
	if in == nil {
		return nil
	}
	out := new(AuthConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelDeployment) DeepCopyInto(out *ModelDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelDeployment.
func (in *ModelDeployment) DeepCopy() *ModelDeployment {
	if in == nil {
		return nil
	}
	out := new(ModelDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModelDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelDeploymentList) DeepCopyInto(out *ModelDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ModelDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelDeploymentList.
func (in *ModelDeploymentList) DeepCopy() *ModelDeploymentList {
	if in == nil {
		return nil
	}
	out := new(ModelDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModelDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelDeploymentSpec) DeepCopyInto(out *ModelDeploymentSpec) {
	*out = *in
	out.ModelSourceRef = in.ModelSourceRef
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ServerlessConfig != nil {
		in, out := &in.ServerlessConfig, &out.ServerlessConfig
		*out = new(ServerlessConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.OffloadingConfig != nil {
		in, out := &in.OffloadingConfig, &out.OffloadingConfig
		*out = new(OffloadingConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelDeploymentSpec.
func (in *ModelDeploymentSpec) DeepCopy() *ModelDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(ModelDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelDeploymentStatus) DeepCopyInto(out *ModelDeploymentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.CreatedAtOnBeamlit.DeepCopyInto(&out.CreatedAtOnBeamlit)
	in.UpdatedAtOnBeamlit.DeepCopyInto(&out.UpdatedAtOnBeamlit)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelDeploymentStatus.
func (in *ModelDeploymentStatus) DeepCopy() *ModelDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(ModelDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OAuthConfig) DeepCopyInto(out *OAuthConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OAuthConfig.
func (in *OAuthConfig) DeepCopy() *OAuthConfig {
	if in == nil {
		return nil
	}
	out := new(OAuthConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OffloadingBehavior) DeepCopyInto(out *OffloadingBehavior) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OffloadingBehavior.
func (in *OffloadingBehavior) DeepCopy() *OffloadingBehavior {
	if in == nil {
		return nil
	}
	out := new(OffloadingBehavior)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OffloadingConfig) DeepCopyInto(out *OffloadingConfig) {
	*out = *in
	if in.ServiceRef != nil {
		in, out := &in.ServiceRef, &out.ServiceRef
		*out = new(ServiceReference)
		**out = **in
	}
	if in.RemoteBackend != nil {
		in, out := &in.RemoteBackend, &out.RemoteBackend
		*out = new(RemoteBackend)
		(*in).DeepCopyInto(*out)
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]v2.MetricSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Behavior != nil {
		in, out := &in.Behavior, &out.Behavior
		*out = new(OffloadingBehavior)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OffloadingConfig.
func (in *OffloadingConfig) DeepCopy() *OffloadingConfig {
	if in == nil {
		return nil
	}
	out := new(OffloadingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteBackend) DeepCopyInto(out *RemoteBackend) {
	*out = *in
	if in.AuthConfig != nil {
		in, out := &in.AuthConfig, &out.AuthConfig
		*out = new(AuthConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.HeadersToAdd != nil {
		in, out := &in.HeadersToAdd, &out.HeadersToAdd
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteBackend.
func (in *RemoteBackend) DeepCopy() *RemoteBackend {
	if in == nil {
		return nil
	}
	out := new(RemoteBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerlessConfig) DeepCopyInto(out *ServerlessConfig) {
	*out = *in
	if in.Metric != nil {
		in, out := &in.Metric, &out.Metric
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
	if in.ScaleUpMinimum != nil {
		in, out := &in.ScaleUpMinimum, &out.ScaleUpMinimum
		*out = new(int32)
		**out = **in
	}
	if in.ScaleDownDelay != nil {
		in, out := &in.ScaleDownDelay, &out.ScaleDownDelay
		*out = new(string)
		**out = **in
	}
	if in.StableWindow != nil {
		in, out := &in.StableWindow, &out.StableWindow
		*out = new(string)
		**out = **in
	}
	if in.LastPodRetentionPeriod != nil {
		in, out := &in.LastPodRetentionPeriod, &out.LastPodRetentionPeriod
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerlessConfig.
func (in *ServerlessConfig) DeepCopy() *ServerlessConfig {
	if in == nil {
		return nil
	}
	out := new(ServerlessConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceReference) DeepCopyInto(out *ServiceReference) {
	*out = *in
	out.ObjectReference = in.ObjectReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceReference.
func (in *ServiceReference) DeepCopy() *ServiceReference {
	if in == nil {
		return nil
	}
	out := new(ServiceReference)
	in.DeepCopyInto(out)
	return out
}