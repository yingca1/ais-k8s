// +build !ignore_autogenerated

/*
 * Copyright (c) 2021, NVIDIA CORPORATION. All rights reserved.
 */

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AIStore) DeepCopyInto(out *AIStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIStore.
func (in *AIStore) DeepCopy() *AIStore {
	if in == nil {
		return nil
	}
	out := new(AIStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AIStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AIStoreList) DeepCopyInto(out *AIStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AIStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIStoreList.
func (in *AIStoreList) DeepCopy() *AIStoreList {
	if in == nil {
		return nil
	}
	out := new(AIStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AIStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AIStoreSpec) DeepCopyInto(out *AIStoreSpec) {
	*out = *in
	if in.ConfigToUpdate != nil {
		in, out := &in.ConfigToUpdate, &out.ConfigToUpdate
		*out = new(ConfigToUpdate)
		(*in).DeepCopyInto(*out)
	}
	in.ProxySpec.DeepCopyInto(&out.ProxySpec)
	in.TargetSpec.DeepCopyInto(&out.TargetSpec)
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.DisablePodAntiAffinity != nil {
		in, out := &in.DisablePodAntiAffinity, &out.DisablePodAntiAffinity
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIStoreSpec.
func (in *AIStoreSpec) DeepCopy() *AIStoreSpec {
	if in == nil {
		return nil
	}
	out := new(AIStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AIStoreStatus) DeepCopyInto(out *AIStoreStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIStoreStatus.
func (in *AIStoreStatus) DeepCopy() *AIStoreStatus {
	if in == nil {
		return nil
	}
	out := new(AIStoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthConfToUpdate) DeepCopyInto(out *AuthConfToUpdate) {
	*out = *in
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthConfToUpdate.
func (in *AuthConfToUpdate) DeepCopy() *AuthConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(AuthConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CksumConfToUpdate) DeepCopyInto(out *CksumConfToUpdate) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.ValidateColdGet != nil {
		in, out := &in.ValidateColdGet, &out.ValidateColdGet
		*out = new(bool)
		**out = **in
	}
	if in.ValidateWarmGet != nil {
		in, out := &in.ValidateWarmGet, &out.ValidateWarmGet
		*out = new(bool)
		**out = **in
	}
	if in.ValidateObjMove != nil {
		in, out := &in.ValidateObjMove, &out.ValidateObjMove
		*out = new(bool)
		**out = **in
	}
	if in.EnableReadRange != nil {
		in, out := &in.EnableReadRange, &out.EnableReadRange
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CksumConfToUpdate.
func (in *CksumConfToUpdate) DeepCopy() *CksumConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(CksumConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientConfToUpdate) DeepCopyInto(out *ClientConfToUpdate) {
	*out = *in
	if in.TimeoutStr != nil {
		in, out := &in.TimeoutStr, &out.TimeoutStr
		*out = new(string)
		**out = **in
	}
	if in.TimeoutLongStr != nil {
		in, out := &in.TimeoutLongStr, &out.TimeoutLongStr
		*out = new(string)
		**out = **in
	}
	if in.ListObjectsStr != nil {
		in, out := &in.ListObjectsStr, &out.ListObjectsStr
		*out = new(string)
		**out = **in
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientConfToUpdate.
func (in *ClientConfToUpdate) DeepCopy() *ClientConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(ClientConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompressionConfToUpdate) DeepCopyInto(out *CompressionConfToUpdate) {
	*out = *in
	if in.BlockMaxSize != nil {
		in, out := &in.BlockMaxSize, &out.BlockMaxSize
		*out = new(int)
		**out = **in
	}
	if in.Checksum != nil {
		in, out := &in.Checksum, &out.Checksum
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompressionConfToUpdate.
func (in *CompressionConfToUpdate) DeepCopy() *CompressionConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(CompressionConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigToUpdate) DeepCopyInto(out *ConfigToUpdate) {
	*out = *in
	if in.Confdir != nil {
		in, out := &in.Confdir, &out.Confdir
		*out = new(string)
		**out = **in
	}
	if in.Mirror != nil {
		in, out := &in.Mirror, &out.Mirror
		*out = new(MirrorConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.EC != nil {
		in, out := &in.EC, &out.EC
		*out = new(ECConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Log != nil {
		in, out := &in.Log, &out.Log
		*out = new(LogConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Periodic != nil {
		in, out := &in.Periodic, &out.Periodic
		*out = new(PeriodConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(TimeoutConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Client != nil {
		in, out := &in.Client, &out.Client
		*out = new(ClientConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.LRU != nil {
		in, out := &in.LRU, &out.LRU
		*out = new(LRUConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Disk != nil {
		in, out := &in.Disk, &out.Disk
		*out = new(DiskConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Rebalance != nil {
		in, out := &in.Rebalance, &out.Rebalance
		*out = new(RebalanceConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Replication != nil {
		in, out := &in.Replication, &out.Replication
		*out = new(ReplicationConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Cksum != nil {
		in, out := &in.Cksum, &out.Cksum
		*out = new(CksumConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Versioning != nil {
		in, out := &in.Versioning, &out.Versioning
		*out = new(VersionConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Net != nil {
		in, out := &in.Net, &out.Net
		*out = new(NetConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.FSHC != nil {
		in, out := &in.FSHC, &out.FSHC
		*out = new(FSHCConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = new(AuthConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Keepalive != nil {
		in, out := &in.Keepalive, &out.Keepalive
		*out = new(KeepaliveConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Downloader != nil {
		in, out := &in.Downloader, &out.Downloader
		*out = new(DownloaderConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.DSort != nil {
		in, out := &in.DSort, &out.DSort
		*out = new(DSortConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Compression != nil {
		in, out := &in.Compression, &out.Compression
		*out = new(CompressionConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.MDWrite != nil {
		in, out := &in.MDWrite, &out.MDWrite
		*out = new(string)
		**out = **in
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	if in.Vmodule != nil {
		in, out := &in.Vmodule, &out.Vmodule
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigToUpdate.
func (in *ConfigToUpdate) DeepCopy() *ConfigToUpdate {
	if in == nil {
		return nil
	}
	out := new(ConfigToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DSortConfToUpdate) DeepCopyInto(out *DSortConfToUpdate) {
	*out = *in
	if in.DuplicatedRecords != nil {
		in, out := &in.DuplicatedRecords, &out.DuplicatedRecords
		*out = new(string)
		**out = **in
	}
	if in.MissingShards != nil {
		in, out := &in.MissingShards, &out.MissingShards
		*out = new(string)
		**out = **in
	}
	if in.EKMMalformedLine != nil {
		in, out := &in.EKMMalformedLine, &out.EKMMalformedLine
		*out = new(string)
		**out = **in
	}
	if in.EKMMissingKey != nil {
		in, out := &in.EKMMissingKey, &out.EKMMissingKey
		*out = new(string)
		**out = **in
	}
	if in.DefaultMaxMemUsage != nil {
		in, out := &in.DefaultMaxMemUsage, &out.DefaultMaxMemUsage
		*out = new(string)
		**out = **in
	}
	if in.CallTimeoutStr != nil {
		in, out := &in.CallTimeoutStr, &out.CallTimeoutStr
		*out = new(string)
		**out = **in
	}
	if in.Compression != nil {
		in, out := &in.Compression, &out.Compression
		*out = new(string)
		**out = **in
	}
	if in.DSorterMemThreshold != nil {
		in, out := &in.DSorterMemThreshold, &out.DSorterMemThreshold
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DSortConfToUpdate.
func (in *DSortConfToUpdate) DeepCopy() *DSortConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(DSortConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaemonSpec) DeepCopyInto(out *DaemonSpec) {
	*out = *in
	out.ServiceSpec = in.ServiceSpec
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerSecurity != nil {
		in, out := &in.ContainerSecurity, &out.ContainerSecurity
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaemonSpec.
func (in *DaemonSpec) DeepCopy() *DaemonSpec {
	if in == nil {
		return nil
	}
	out := new(DaemonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskConfToUpdate) DeepCopyInto(out *DiskConfToUpdate) {
	*out = *in
	if in.DiskUtilLowWM != nil {
		in, out := &in.DiskUtilLowWM, &out.DiskUtilLowWM
		*out = new(int64)
		**out = **in
	}
	if in.DiskUtilHighWM != nil {
		in, out := &in.DiskUtilHighWM, &out.DiskUtilHighWM
		*out = new(int64)
		**out = **in
	}
	if in.DiskUtilMaxWM != nil {
		in, out := &in.DiskUtilMaxWM, &out.DiskUtilMaxWM
		*out = new(int64)
		**out = **in
	}
	if in.IostatTimeLongStr != nil {
		in, out := &in.IostatTimeLongStr, &out.IostatTimeLongStr
		*out = new(string)
		**out = **in
	}
	if in.IostatTimeShortStr != nil {
		in, out := &in.IostatTimeShortStr, &out.IostatTimeShortStr
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskConfToUpdate.
func (in *DiskConfToUpdate) DeepCopy() *DiskConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(DiskConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DownloaderConfToUpdate) DeepCopyInto(out *DownloaderConfToUpdate) {
	*out = *in
	if in.TimeoutStr != nil {
		in, out := &in.TimeoutStr, &out.TimeoutStr
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DownloaderConfToUpdate.
func (in *DownloaderConfToUpdate) DeepCopy() *DownloaderConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(DownloaderConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ECConfToUpdate) DeepCopyInto(out *ECConfToUpdate) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ObjSizeLimit != nil {
		in, out := &in.ObjSizeLimit, &out.ObjSizeLimit
		*out = new(int64)
		**out = **in
	}
	if in.DataSlices != nil {
		in, out := &in.DataSlices, &out.DataSlices
		*out = new(int)
		**out = **in
	}
	if in.ParitySlices != nil {
		in, out := &in.ParitySlices, &out.ParitySlices
		*out = new(int)
		**out = **in
	}
	if in.Compression != nil {
		in, out := &in.Compression, &out.Compression
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ECConfToUpdate.
func (in *ECConfToUpdate) DeepCopy() *ECConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(ECConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FSHCConfToUpdate) DeepCopyInto(out *FSHCConfToUpdate) {
	*out = *in
	if in.TestFileCount != nil {
		in, out := &in.TestFileCount, &out.TestFileCount
		*out = new(int)
		**out = **in
	}
	if in.ErrorLimit != nil {
		in, out := &in.ErrorLimit, &out.ErrorLimit
		*out = new(int)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FSHCConfToUpdate.
func (in *FSHCConfToUpdate) DeepCopy() *FSHCConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(FSHCConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPConfToUpdate) DeepCopyInto(out *HTTPConfToUpdate) {
	*out = *in
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.WriteBufferSize != nil {
		in, out := &in.WriteBufferSize, &out.WriteBufferSize
		*out = new(int)
		**out = **in
	}
	if in.ReadBufferSize != nil {
		in, out := &in.ReadBufferSize, &out.ReadBufferSize
		*out = new(int)
		**out = **in
	}
	if in.UseHTTPS != nil {
		in, out := &in.UseHTTPS, &out.UseHTTPS
		*out = new(bool)
		**out = **in
	}
	if in.SkipVerify != nil {
		in, out := &in.SkipVerify, &out.SkipVerify
		*out = new(bool)
		**out = **in
	}
	if in.Chunked != nil {
		in, out := &in.Chunked, &out.Chunked
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPConfToUpdate.
func (in *HTTPConfToUpdate) DeepCopy() *HTTPConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(HTTPConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeepaliveConfToUpdate) DeepCopyInto(out *KeepaliveConfToUpdate) {
	*out = *in
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(KeepaliveTrackerConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(KeepaliveTrackerConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.RetryFactor != nil {
		in, out := &in.RetryFactor, &out.RetryFactor
		*out = new(uint8)
		**out = **in
	}
	if in.TimeoutFactor != nil {
		in, out := &in.TimeoutFactor, &out.TimeoutFactor
		*out = new(uint8)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeepaliveConfToUpdate.
func (in *KeepaliveConfToUpdate) DeepCopy() *KeepaliveConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(KeepaliveConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeepaliveTrackerConfToUpdate) DeepCopyInto(out *KeepaliveTrackerConfToUpdate) {
	*out = *in
	if in.IntervalStr != nil {
		in, out := &in.IntervalStr, &out.IntervalStr
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Factor != nil {
		in, out := &in.Factor, &out.Factor
		*out = new(uint8)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeepaliveTrackerConfToUpdate.
func (in *KeepaliveTrackerConfToUpdate) DeepCopy() *KeepaliveTrackerConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(KeepaliveTrackerConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4ConfToUpdate) DeepCopyInto(out *L4ConfToUpdate) {
	*out = *in
	if in.Proto != nil {
		in, out := &in.Proto, &out.Proto
		*out = new(string)
		**out = **in
	}
	if in.PortStr != nil {
		in, out := &in.PortStr, &out.PortStr
		*out = new(string)
		**out = **in
	}
	if in.PortIntraControlStr != nil {
		in, out := &in.PortIntraControlStr, &out.PortIntraControlStr
		*out = new(string)
		**out = **in
	}
	if in.PortIntraDataStr != nil {
		in, out := &in.PortIntraDataStr, &out.PortIntraDataStr
		*out = new(string)
		**out = **in
	}
	if in.SndRcvBufSize != nil {
		in, out := &in.SndRcvBufSize, &out.SndRcvBufSize
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4ConfToUpdate.
func (in *L4ConfToUpdate) DeepCopy() *L4ConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(L4ConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LRUConfToUpdate) DeepCopyInto(out *LRUConfToUpdate) {
	*out = *in
	if in.LowWM != nil {
		in, out := &in.LowWM, &out.LowWM
		*out = new(int64)
		**out = **in
	}
	if in.HighWM != nil {
		in, out := &in.HighWM, &out.HighWM
		*out = new(int64)
		**out = **in
	}
	if in.OOS != nil {
		in, out := &in.OOS, &out.OOS
		*out = new(int64)
		**out = **in
	}
	if in.DontEvictTimeStr != nil {
		in, out := &in.DontEvictTimeStr, &out.DontEvictTimeStr
		*out = new(string)
		**out = **in
	}
	if in.CapacityUpdTimeStr != nil {
		in, out := &in.CapacityUpdTimeStr, &out.CapacityUpdTimeStr
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LRUConfToUpdate.
func (in *LRUConfToUpdate) DeepCopy() *LRUConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(LRUConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogConfToUpdate) DeepCopyInto(out *LogConfToUpdate) {
	*out = *in
	if in.Dir != nil {
		in, out := &in.Dir, &out.Dir
		*out = new(string)
		**out = **in
	}
	if in.Level != nil {
		in, out := &in.Level, &out.Level
		*out = new(string)
		**out = **in
	}
	if in.MaxSize != nil {
		in, out := &in.MaxSize, &out.MaxSize
		*out = new(uint64)
		**out = **in
	}
	if in.MaxTotal != nil {
		in, out := &in.MaxTotal, &out.MaxTotal
		*out = new(uint64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogConfToUpdate.
func (in *LogConfToUpdate) DeepCopy() *LogConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(LogConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MirrorConfToUpdate) DeepCopyInto(out *MirrorConfToUpdate) {
	*out = *in
	if in.Copies != nil {
		in, out := &in.Copies, &out.Copies
		*out = new(int64)
		**out = **in
	}
	if in.Burst != nil {
		in, out := &in.Burst, &out.Burst
		*out = new(int)
		**out = **in
	}
	if in.UtilThresh != nil {
		in, out := &in.UtilThresh, &out.UtilThresh
		*out = new(int64)
		**out = **in
	}
	if in.OptimizePUT != nil {
		in, out := &in.OptimizePUT, &out.OptimizePUT
		*out = new(bool)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MirrorConfToUpdate.
func (in *MirrorConfToUpdate) DeepCopy() *MirrorConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(MirrorConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Mount) DeepCopyInto(out *Mount) {
	*out = *in
	out.Size = in.Size.DeepCopy()
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Mount.
func (in *Mount) DeepCopy() *Mount {
	if in == nil {
		return nil
	}
	out := new(Mount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetConfToUpdate) DeepCopyInto(out *NetConfToUpdate) {
	*out = *in
	if in.L4 != nil {
		in, out := &in.L4, &out.L4
		*out = new(L4ConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = new(HTTPConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetConfToUpdate.
func (in *NetConfToUpdate) DeepCopy() *NetConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(NetConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoDiskIO) DeepCopyInto(out *NoDiskIO) {
	*out = *in
	out.DryObjSize = in.DryObjSize.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoDiskIO.
func (in *NoDiskIO) DeepCopy() *NoDiskIO {
	if in == nil {
		return nil
	}
	out := new(NoDiskIO)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeriodConfToUpdate) DeepCopyInto(out *PeriodConfToUpdate) {
	*out = *in
	if in.StatsTimeStr != nil {
		in, out := &in.StatsTimeStr, &out.StatsTimeStr
		*out = new(string)
		**out = **in
	}
	if in.RetrySyncTimeStr != nil {
		in, out := &in.RetrySyncTimeStr, &out.RetrySyncTimeStr
		*out = new(string)
		**out = **in
	}
	if in.NotifTimeStr != nil {
		in, out := &in.NotifTimeStr, &out.NotifTimeStr
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeriodConfToUpdate.
func (in *PeriodConfToUpdate) DeepCopy() *PeriodConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(PeriodConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RebalanceConfToUpdate) DeepCopyInto(out *RebalanceConfToUpdate) {
	*out = *in
	if in.DestRetryTimeStr != nil {
		in, out := &in.DestRetryTimeStr, &out.DestRetryTimeStr
		*out = new(string)
		**out = **in
	}
	if in.QuiesceStr != nil {
		in, out := &in.QuiesceStr, &out.QuiesceStr
		*out = new(string)
		**out = **in
	}
	if in.Compression != nil {
		in, out := &in.Compression, &out.Compression
		*out = new(string)
		**out = **in
	}
	if in.Multiplier != nil {
		in, out := &in.Multiplier, &out.Multiplier
		*out = new(uint8)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RebalanceConfToUpdate.
func (in *RebalanceConfToUpdate) DeepCopy() *RebalanceConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(RebalanceConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationConfToUpdate) DeepCopyInto(out *ReplicationConfToUpdate) {
	*out = *in
	if in.OnColdGet != nil {
		in, out := &in.OnColdGet, &out.OnColdGet
		*out = new(bool)
		**out = **in
	}
	if in.OnPut != nil {
		in, out := &in.OnPut, &out.OnPut
		*out = new(bool)
		**out = **in
	}
	if in.OnLRUEviction != nil {
		in, out := &in.OnLRUEviction, &out.OnLRUEviction
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationConfToUpdate.
func (in *ReplicationConfToUpdate) DeepCopy() *ReplicationConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(ReplicationConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpec) DeepCopyInto(out *ServiceSpec) {
	*out = *in
	out.ServicePort = in.ServicePort
	out.PublicPort = in.PublicPort
	out.IntraControlPort = in.IntraControlPort
	out.IntraDataPort = in.IntraDataPort
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpec.
func (in *ServiceSpec) DeepCopy() *ServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetSpec) DeepCopyInto(out *TargetSpec) {
	*out = *in
	in.DaemonSpec.DeepCopyInto(&out.DaemonSpec)
	in.NoDiskIO.DeepCopyInto(&out.NoDiskIO)
	if in.Mounts != nil {
		in, out := &in.Mounts, &out.Mounts
		*out = make([]Mount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetSpec.
func (in *TargetSpec) DeepCopy() *TargetSpec {
	if in == nil {
		return nil
	}
	out := new(TargetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeoutConfToUpdate) DeepCopyInto(out *TimeoutConfToUpdate) {
	*out = *in
	if in.CplaneOperationStr != nil {
		in, out := &in.CplaneOperationStr, &out.CplaneOperationStr
		*out = new(string)
		**out = **in
	}
	if in.MaxKeepaliveStr != nil {
		in, out := &in.MaxKeepaliveStr, &out.MaxKeepaliveStr
		*out = new(string)
		**out = **in
	}
	if in.MaxHostBusyStr != nil {
		in, out := &in.MaxHostBusyStr, &out.MaxHostBusyStr
		*out = new(string)
		**out = **in
	}
	if in.StartupStr != nil {
		in, out := &in.StartupStr, &out.StartupStr
		*out = new(string)
		**out = **in
	}
	if in.SendFileStr != nil {
		in, out := &in.SendFileStr, &out.SendFileStr
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeoutConfToUpdate.
func (in *TimeoutConfToUpdate) DeepCopy() *TimeoutConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(TimeoutConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionConfToUpdate) DeepCopyInto(out *VersionConfToUpdate) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ValidateWarmGet != nil {
		in, out := &in.ValidateWarmGet, &out.ValidateWarmGet
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionConfToUpdate.
func (in *VersionConfToUpdate) DeepCopy() *VersionConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(VersionConfToUpdate)
	in.DeepCopyInto(out)
	return out
}
