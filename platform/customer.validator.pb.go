// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: platform/customer.proto

package platform

import (
	fmt "fmt"
	math "math"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/dinhtp/lets-run-pbtype/gateway"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CreateUpdateCustomerRequest) Validate() error {
	if this.Platform != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Platform); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Platform", err)
		}
	}
	if this.Customer != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Customer); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Customer", err)
		}
	}
	return nil
}
func (this *OneCustomerRequest) Validate() error {
	if this.Platform != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Platform); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Platform", err)
		}
	}
	return nil
}
func (this *ListCustomerRequest) Validate() error {
	if this.Platform != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Platform); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Platform", err)
		}
	}
	return nil
}
func (this *ListCustomerResponse) Validate() error {
	for _, item := range this.Items {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Items", err)
			}
		}
	}
	return nil
}
