// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/redis/v1alpha1/redisdestination.proto

// RedisDestination defines policies that apply to redis traffic intended for a redis service

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "istio.io/api/networking/v1alpha3"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using RedisDestination within kubernetes types, where deepcopy-gen is used.
func (in *RedisDestination) DeepCopyInto(out *RedisDestination) {
	p := proto.Clone(in).(*RedisDestination)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisDestination. Required by controller-gen.
func (in *RedisDestination) DeepCopy() *RedisDestination {
	if in == nil {
		return nil
	}
	out := new(RedisDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RedisDestination. Required by controller-gen.
func (in *RedisDestination) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ConnectionPoolSettings within kubernetes types, where deepcopy-gen is used.
func (in *ConnectionPoolSettings) DeepCopyInto(out *ConnectionPoolSettings) {
	p := proto.Clone(in).(*ConnectionPoolSettings)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPoolSettings. Required by controller-gen.
func (in *ConnectionPoolSettings) DeepCopy() *ConnectionPoolSettings {
	if in == nil {
		return nil
	}
	out := new(ConnectionPoolSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPoolSettings. Required by controller-gen.
func (in *ConnectionPoolSettings) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Auth within kubernetes types, where deepcopy-gen is used.
func (in *Auth) DeepCopyInto(out *Auth) {
	p := proto.Clone(in).(*Auth)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Auth. Required by controller-gen.
func (in *Auth) DeepCopy() *Auth {
	if in == nil {
		return nil
	}
	out := new(Auth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Auth. Required by controller-gen.
func (in *Auth) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using SecretReference within kubernetes types, where deepcopy-gen is used.
func (in *SecretReference) DeepCopyInto(out *SecretReference) {
	p := proto.Clone(in).(*SecretReference)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretReference. Required by controller-gen.
func (in *SecretReference) DeepCopy() *SecretReference {
	if in == nil {
		return nil
	}
	out := new(SecretReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new SecretReference. Required by controller-gen.
func (in *SecretReference) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using PlainAuth within kubernetes types, where deepcopy-gen is used.
func (in *PlainAuth) DeepCopyInto(out *PlainAuth) {
	p := proto.Clone(in).(*PlainAuth)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlainAuth. Required by controller-gen.
func (in *PlainAuth) DeepCopy() *PlainAuth {
	if in == nil {
		return nil
	}
	out := new(PlainAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new PlainAuth. Required by controller-gen.
func (in *PlainAuth) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RedisSettings within kubernetes types, where deepcopy-gen is used.
func (in *RedisSettings) DeepCopyInto(out *RedisSettings) {
	p := proto.Clone(in).(*RedisSettings)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisSettings. Required by controller-gen.
func (in *RedisSettings) DeepCopy() *RedisSettings {
	if in == nil {
		return nil
	}
	out := new(RedisSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RedisSettings. Required by controller-gen.
func (in *RedisSettings) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TrafficPolicy within kubernetes types, where deepcopy-gen is used.
func (in *TrafficPolicy) DeepCopyInto(out *TrafficPolicy) {
	p := proto.Clone(in).(*TrafficPolicy)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficPolicy. Required by controller-gen.
func (in *TrafficPolicy) DeepCopy() *TrafficPolicy {
	if in == nil {
		return nil
	}
	out := new(TrafficPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TrafficPolicy. Required by controller-gen.
func (in *TrafficPolicy) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}