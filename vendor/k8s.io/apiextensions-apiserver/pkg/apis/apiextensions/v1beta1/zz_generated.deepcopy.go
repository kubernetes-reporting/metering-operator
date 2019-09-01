// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this ***REMOVED***le except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the speci***REMOVED***c language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConversionRequest) DeepCopyInto(out *ConversionRequest) {
	*out = *in
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = make([]runtime.RawExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConversionRequest.
func (in *ConversionRequest) DeepCopy() *ConversionRequest {
	if in == nil {
		return nil
	}
	out := new(ConversionRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConversionResponse) DeepCopyInto(out *ConversionResponse) {
	*out = *in
	if in.ConvertedObjects != nil {
		in, out := &in.ConvertedObjects, &out.ConvertedObjects
		*out = make([]runtime.RawExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Result.DeepCopyInto(&out.Result)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConversionResponse.
func (in *ConversionResponse) DeepCopy() *ConversionResponse {
	if in == nil {
		return nil
	}
	out := new(ConversionResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConversionReview) DeepCopyInto(out *ConversionReview) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(ConversionRequest)
		(*in).DeepCopyInto(*out)
	}
	if in.Response != nil {
		in, out := &in.Response, &out.Response
		*out = new(ConversionResponse)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConversionReview.
func (in *ConversionReview) DeepCopy() *ConversionReview {
	if in == nil {
		return nil
	}
	out := new(ConversionReview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConversionReview) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceColumnDe***REMOVED***nition) DeepCopyInto(out *CustomResourceColumnDe***REMOVED***nition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceColumnDe***REMOVED***nition.
func (in *CustomResourceColumnDe***REMOVED***nition) DeepCopy() *CustomResourceColumnDe***REMOVED***nition {
	if in == nil {
		return nil
	}
	out := new(CustomResourceColumnDe***REMOVED***nition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceConversion) DeepCopyInto(out *CustomResourceConversion) {
	*out = *in
	if in.WebhookClientCon***REMOVED***g != nil {
		in, out := &in.WebhookClientCon***REMOVED***g, &out.WebhookClientCon***REMOVED***g
		*out = new(WebhookClientCon***REMOVED***g)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceConversion.
func (in *CustomResourceConversion) DeepCopy() *CustomResourceConversion {
	if in == nil {
		return nil
	}
	out := new(CustomResourceConversion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDe***REMOVED***nition) DeepCopyInto(out *CustomResourceDe***REMOVED***nition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDe***REMOVED***nition.
func (in *CustomResourceDe***REMOVED***nition) DeepCopy() *CustomResourceDe***REMOVED***nition {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDe***REMOVED***nition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomResourceDe***REMOVED***nition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDe***REMOVED***nitionCondition) DeepCopyInto(out *CustomResourceDe***REMOVED***nitionCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDe***REMOVED***nitionCondition.
func (in *CustomResourceDe***REMOVED***nitionCondition) DeepCopy() *CustomResourceDe***REMOVED***nitionCondition {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDe***REMOVED***nitionCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDe***REMOVED***nitionList) DeepCopyInto(out *CustomResourceDe***REMOVED***nitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CustomResourceDe***REMOVED***nition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDe***REMOVED***nitionList.
func (in *CustomResourceDe***REMOVED***nitionList) DeepCopy() *CustomResourceDe***REMOVED***nitionList {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDe***REMOVED***nitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomResourceDe***REMOVED***nitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDe***REMOVED***nitionNames) DeepCopyInto(out *CustomResourceDe***REMOVED***nitionNames) {
	*out = *in
	if in.ShortNames != nil {
		in, out := &in.ShortNames, &out.ShortNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Categories != nil {
		in, out := &in.Categories, &out.Categories
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDe***REMOVED***nitionNames.
func (in *CustomResourceDe***REMOVED***nitionNames) DeepCopy() *CustomResourceDe***REMOVED***nitionNames {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDe***REMOVED***nitionNames)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDe***REMOVED***nitionSpec) DeepCopyInto(out *CustomResourceDe***REMOVED***nitionSpec) {
	*out = *in
	in.Names.DeepCopyInto(&out.Names)
	if in.Validation != nil {
		in, out := &in.Validation, &out.Validation
		*out = new(CustomResourceValidation)
		(*in).DeepCopyInto(*out)
	}
	if in.Subresources != nil {
		in, out := &in.Subresources, &out.Subresources
		*out = new(CustomResourceSubresources)
		(*in).DeepCopyInto(*out)
	}
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]CustomResourceDe***REMOVED***nitionVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AdditionalPrinterColumns != nil {
		in, out := &in.AdditionalPrinterColumns, &out.AdditionalPrinterColumns
		*out = make([]CustomResourceColumnDe***REMOVED***nition, len(*in))
		copy(*out, *in)
	}
	if in.Conversion != nil {
		in, out := &in.Conversion, &out.Conversion
		*out = new(CustomResourceConversion)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDe***REMOVED***nitionSpec.
func (in *CustomResourceDe***REMOVED***nitionSpec) DeepCopy() *CustomResourceDe***REMOVED***nitionSpec {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDe***REMOVED***nitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDe***REMOVED***nitionStatus) DeepCopyInto(out *CustomResourceDe***REMOVED***nitionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CustomResourceDe***REMOVED***nitionCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.AcceptedNames.DeepCopyInto(&out.AcceptedNames)
	if in.StoredVersions != nil {
		in, out := &in.StoredVersions, &out.StoredVersions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDe***REMOVED***nitionStatus.
func (in *CustomResourceDe***REMOVED***nitionStatus) DeepCopy() *CustomResourceDe***REMOVED***nitionStatus {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDe***REMOVED***nitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDe***REMOVED***nitionVersion) DeepCopyInto(out *CustomResourceDe***REMOVED***nitionVersion) {
	*out = *in
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(CustomResourceValidation)
		(*in).DeepCopyInto(*out)
	}
	if in.Subresources != nil {
		in, out := &in.Subresources, &out.Subresources
		*out = new(CustomResourceSubresources)
		(*in).DeepCopyInto(*out)
	}
	if in.AdditionalPrinterColumns != nil {
		in, out := &in.AdditionalPrinterColumns, &out.AdditionalPrinterColumns
		*out = make([]CustomResourceColumnDe***REMOVED***nition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDe***REMOVED***nitionVersion.
func (in *CustomResourceDe***REMOVED***nitionVersion) DeepCopy() *CustomResourceDe***REMOVED***nitionVersion {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDe***REMOVED***nitionVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceSubresourceScale) DeepCopyInto(out *CustomResourceSubresourceScale) {
	*out = *in
	if in.LabelSelectorPath != nil {
		in, out := &in.LabelSelectorPath, &out.LabelSelectorPath
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceSubresourceScale.
func (in *CustomResourceSubresourceScale) DeepCopy() *CustomResourceSubresourceScale {
	if in == nil {
		return nil
	}
	out := new(CustomResourceSubresourceScale)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceSubresourceStatus) DeepCopyInto(out *CustomResourceSubresourceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceSubresourceStatus.
func (in *CustomResourceSubresourceStatus) DeepCopy() *CustomResourceSubresourceStatus {
	if in == nil {
		return nil
	}
	out := new(CustomResourceSubresourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceSubresources) DeepCopyInto(out *CustomResourceSubresources) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(CustomResourceSubresourceStatus)
		**out = **in
	}
	if in.Scale != nil {
		in, out := &in.Scale, &out.Scale
		*out = new(CustomResourceSubresourceScale)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceSubresources.
func (in *CustomResourceSubresources) DeepCopy() *CustomResourceSubresources {
	if in == nil {
		return nil
	}
	out := new(CustomResourceSubresources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceValidation) DeepCopyInto(out *CustomResourceValidation) {
	*out = *in
	if in.OpenAPIV3Schema != nil {
		in, out := &in.OpenAPIV3Schema, &out.OpenAPIV3Schema
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceValidation.
func (in *CustomResourceValidation) DeepCopy() *CustomResourceValidation {
	if in == nil {
		return nil
	}
	out := new(CustomResourceValidation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDocumentation) DeepCopyInto(out *ExternalDocumentation) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDocumentation.
func (in *ExternalDocumentation) DeepCopy() *ExternalDocumentation {
	if in == nil {
		return nil
	}
	out := new(ExternalDocumentation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSON) DeepCopyInto(out *JSON) {
	*out = *in
	if in.Raw != nil {
		in, out := &in.Raw, &out.Raw
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSON.
func (in *JSON) DeepCopy() *JSON {
	if in == nil {
		return nil
	}
	out := new(JSON)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in JSONSchemaDe***REMOVED***nitions) DeepCopyInto(out *JSONSchemaDe***REMOVED***nitions) {
	{
		in := &in
		*out = make(JSONSchemaDe***REMOVED***nitions, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONSchemaDe***REMOVED***nitions.
func (in JSONSchemaDe***REMOVED***nitions) DeepCopy() JSONSchemaDe***REMOVED***nitions {
	if in == nil {
		return nil
	}
	out := new(JSONSchemaDe***REMOVED***nitions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in JSONSchemaDependencies) DeepCopyInto(out *JSONSchemaDependencies) {
	{
		in := &in
		*out = make(JSONSchemaDependencies, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONSchemaDependencies.
func (in JSONSchemaDependencies) DeepCopy() JSONSchemaDependencies {
	if in == nil {
		return nil
	}
	out := new(JSONSchemaDependencies)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSONSchemaProps) DeepCopyInto(out *JSONSchemaProps) {
	clone := in.DeepCopy()
	*out = *clone
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSONSchemaPropsOrArray) DeepCopyInto(out *JSONSchemaPropsOrArray) {
	*out = *in
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = (*in).DeepCopy()
	}
	if in.JSONSchemas != nil {
		in, out := &in.JSONSchemas, &out.JSONSchemas
		*out = make([]JSONSchemaProps, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONSchemaPropsOrArray.
func (in *JSONSchemaPropsOrArray) DeepCopy() *JSONSchemaPropsOrArray {
	if in == nil {
		return nil
	}
	out := new(JSONSchemaPropsOrArray)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSONSchemaPropsOrBool) DeepCopyInto(out *JSONSchemaPropsOrBool) {
	*out = *in
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONSchemaPropsOrBool.
func (in *JSONSchemaPropsOrBool) DeepCopy() *JSONSchemaPropsOrBool {
	if in == nil {
		return nil
	}
	out := new(JSONSchemaPropsOrBool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSONSchemaPropsOrStringArray) DeepCopyInto(out *JSONSchemaPropsOrStringArray) {
	*out = *in
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = (*in).DeepCopy()
	}
	if in.Property != nil {
		in, out := &in.Property, &out.Property
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONSchemaPropsOrStringArray.
func (in *JSONSchemaPropsOrStringArray) DeepCopy() *JSONSchemaPropsOrStringArray {
	if in == nil {
		return nil
	}
	out := new(JSONSchemaPropsOrStringArray)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceReference) DeepCopyInto(out *ServiceReference) {
	*out = *in
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	return
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookClientCon***REMOVED***g) DeepCopyInto(out *WebhookClientCon***REMOVED***g) {
	*out = *in
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(ServiceReference)
		(*in).DeepCopyInto(*out)
	}
	if in.CABundle != nil {
		in, out := &in.CABundle, &out.CABundle
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookClientCon***REMOVED***g.
func (in *WebhookClientCon***REMOVED***g) DeepCopy() *WebhookClientCon***REMOVED***g {
	if in == nil {
		return nil
	}
	out := new(WebhookClientCon***REMOVED***g)
	in.DeepCopyInto(out)
	return out
}
