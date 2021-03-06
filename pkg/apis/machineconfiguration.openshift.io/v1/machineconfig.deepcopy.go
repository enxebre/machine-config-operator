package v1

import (
	ignv2_2 "github.com/coreos/ignition/config/v2_2"
	ignv2_2types "github.com/coreos/ignition/config/v2_2/types"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto copying the receiver, writing into out. in must be non-nil.
func (in *MachineConfig) DeepCopyInto(out *MachineConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy copying the receiver, creating a new MachineConfig.
func (in *MachineConfig) DeepCopy() *MachineConfig {
	if in == nil {
		return nil
	}
	out := new(MachineConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject copying the receiver, creating a new runtime.Object.
func (in *MachineConfig) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto copying the receiver, writing into out. in must be non-nil.
func (in *MachineConfigList) DeepCopyInto(out *MachineConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy copying the receiver, creating a new MachineConfigList.
func (in *MachineConfigList) DeepCopy() *MachineConfigList {
	if in == nil {
		return nil
	}
	out := new(MachineConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject copying the receiver, creating a new runtime.Object.
func (in *MachineConfigList) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto copying the receiver, writing into out. in must be non-nil.
func (in *MachineConfigSpec) DeepCopyInto(out *MachineConfigSpec) {
	*out = *in
	out.Config = deepCopyIgnConfig(in.Config)
	return
}

func deepCopyIgnConfig(in ignv2_2types.Config) ignv2_2types.Config {
	var out ignv2_2types.Config

	// https://github.com/coreos/ignition/blob/d19b2021cf397de7c31774c13805bbc3aa655646/config/v2_2/append.go#L41
	out.Ignition.Version = in.Ignition.Version

	return ignv2_2.Append(out, in)
}

// DeepCopy copying the receiver, creating a new MachineConfigSpec.
func (in *MachineConfigSpec) DeepCopy() *MachineConfigSpec {
	if in == nil {
		return nil
	}
	out := new(MachineConfigSpec)
	in.DeepCopyInto(out)
	return out
}
