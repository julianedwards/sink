// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/conversion_upload_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [ConversionUploadService.UploadClickConversions][google.ads.googleads.v2.services.ConversionUploadService.UploadClickConversions].
type UploadClickConversionsRequest struct {
	// The ID of the customer performing the upload.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The conversions that are being uploaded.
	Conversions []*ClickConversion `protobuf:"bytes,2,rep,name=conversions,proto3" json:"conversions,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// This should always be set to true.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadClickConversionsRequest) Reset()         { *m = UploadClickConversionsRequest{} }
func (m *UploadClickConversionsRequest) String() string { return proto.CompactTextString(m) }
func (*UploadClickConversionsRequest) ProtoMessage()    {}
func (*UploadClickConversionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_962dfdf4c09453fe, []int{0}
}

func (m *UploadClickConversionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadClickConversionsRequest.Unmarshal(m, b)
}
func (m *UploadClickConversionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadClickConversionsRequest.Marshal(b, m, deterministic)
}
func (m *UploadClickConversionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadClickConversionsRequest.Merge(m, src)
}
func (m *UploadClickConversionsRequest) XXX_Size() int {
	return xxx_messageInfo_UploadClickConversionsRequest.Size(m)
}
func (m *UploadClickConversionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadClickConversionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadClickConversionsRequest proto.InternalMessageInfo

func (m *UploadClickConversionsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *UploadClickConversionsRequest) GetConversions() []*ClickConversion {
	if m != nil {
		return m.Conversions
	}
	return nil
}

func (m *UploadClickConversionsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *UploadClickConversionsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// Response message for [ConversionUploadService.UploadClickConversions][google.ads.googleads.v2.services.ConversionUploadService.UploadClickConversions].
type UploadClickConversionsResponse struct {
	// Errors that pertain to conversion failures in the partial failure mode.
	// Returned when all errors occur inside the conversions. If any errors occur
	// outside the conversions (e.g. auth errors), we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,1,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// Returned for successfully processed conversions. Proto will be empty for
	// rows that received an error. Results are not returned when validate_only is
	// true.
	Results              []*ClickConversionResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *UploadClickConversionsResponse) Reset()         { *m = UploadClickConversionsResponse{} }
func (m *UploadClickConversionsResponse) String() string { return proto.CompactTextString(m) }
func (*UploadClickConversionsResponse) ProtoMessage()    {}
func (*UploadClickConversionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_962dfdf4c09453fe, []int{1}
}

func (m *UploadClickConversionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadClickConversionsResponse.Unmarshal(m, b)
}
func (m *UploadClickConversionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadClickConversionsResponse.Marshal(b, m, deterministic)
}
func (m *UploadClickConversionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadClickConversionsResponse.Merge(m, src)
}
func (m *UploadClickConversionsResponse) XXX_Size() int {
	return xxx_messageInfo_UploadClickConversionsResponse.Size(m)
}
func (m *UploadClickConversionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadClickConversionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadClickConversionsResponse proto.InternalMessageInfo

func (m *UploadClickConversionsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *UploadClickConversionsResponse) GetResults() []*ClickConversionResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// Request message for [ConversionUploadService.UploadCallConversions][google.ads.googleads.v2.services.ConversionUploadService.UploadCallConversions].
type UploadCallConversionsRequest struct {
	// The ID of the customer performing the upload.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The conversions that are being uploaded.
	Conversions []*CallConversion `protobuf:"bytes,2,rep,name=conversions,proto3" json:"conversions,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// This should always be set to true.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadCallConversionsRequest) Reset()         { *m = UploadCallConversionsRequest{} }
func (m *UploadCallConversionsRequest) String() string { return proto.CompactTextString(m) }
func (*UploadCallConversionsRequest) ProtoMessage()    {}
func (*UploadCallConversionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_962dfdf4c09453fe, []int{2}
}

func (m *UploadCallConversionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadCallConversionsRequest.Unmarshal(m, b)
}
func (m *UploadCallConversionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadCallConversionsRequest.Marshal(b, m, deterministic)
}
func (m *UploadCallConversionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadCallConversionsRequest.Merge(m, src)
}
func (m *UploadCallConversionsRequest) XXX_Size() int {
	return xxx_messageInfo_UploadCallConversionsRequest.Size(m)
}
func (m *UploadCallConversionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadCallConversionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadCallConversionsRequest proto.InternalMessageInfo

func (m *UploadCallConversionsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *UploadCallConversionsRequest) GetConversions() []*CallConversion {
	if m != nil {
		return m.Conversions
	}
	return nil
}

func (m *UploadCallConversionsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *UploadCallConversionsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// Response message for [ConversionUploadService.UploadCallConversions][google.ads.googleads.v2.services.ConversionUploadService.UploadCallConversions].
type UploadCallConversionsResponse struct {
	// Errors that pertain to conversion failures in the partial failure mode.
	// Returned when all errors occur inside the conversions. If any errors occur
	// outside the conversions (e.g. auth errors), we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,1,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// Returned for successfully processed conversions. Proto will be empty for
	// rows that received an error. Results are not returned when validate_only is
	// true.
	Results              []*CallConversionResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *UploadCallConversionsResponse) Reset()         { *m = UploadCallConversionsResponse{} }
func (m *UploadCallConversionsResponse) String() string { return proto.CompactTextString(m) }
func (*UploadCallConversionsResponse) ProtoMessage()    {}
func (*UploadCallConversionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_962dfdf4c09453fe, []int{3}
}

func (m *UploadCallConversionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadCallConversionsResponse.Unmarshal(m, b)
}
func (m *UploadCallConversionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadCallConversionsResponse.Marshal(b, m, deterministic)
}
func (m *UploadCallConversionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadCallConversionsResponse.Merge(m, src)
}
func (m *UploadCallConversionsResponse) XXX_Size() int {
	return xxx_messageInfo_UploadCallConversionsResponse.Size(m)
}
func (m *UploadCallConversionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadCallConversionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadCallConversionsResponse proto.InternalMessageInfo

func (m *UploadCallConversionsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *UploadCallConversionsResponse) GetResults() []*CallConversionResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// A click conversion.
type ClickConversion struct {
	// The Google click ID (gclid) associated with this conversion.
	Gclid *wrappers.StringValue `protobuf:"bytes,1,opt,name=gclid,proto3" json:"gclid,omitempty"`
	// Resource name of the conversion action associated with this conversion.
	// Note: Although this resource name consists of a customer id and a
	// conversion action id, validation will ignore the customer id and use the
	// conversion action id as the sole identifier of the conversion action.
	ConversionAction *wrappers.StringValue `protobuf:"bytes,2,opt,name=conversion_action,json=conversionAction,proto3" json:"conversion_action,omitempty"`
	// The date time at which the conversion occurred. Must be after
	// the click time. The timezone must be specified. The format is
	// "yyyy-mm-dd hh:mm:ss+|-hh:mm", e.g. “2019-01-01 12:32:45-08:00”.
	ConversionDateTime *wrappers.StringValue `protobuf:"bytes,3,opt,name=conversion_date_time,json=conversionDateTime,proto3" json:"conversion_date_time,omitempty"`
	// The value of the conversion for the advertiser.
	ConversionValue *wrappers.DoubleValue `protobuf:"bytes,4,opt,name=conversion_value,json=conversionValue,proto3" json:"conversion_value,omitempty"`
	// Currency associated with the conversion value. This is the ISO 4217
	// 3-character currency code. For example: USD, EUR.
	CurrencyCode *wrappers.StringValue `protobuf:"bytes,5,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// The order ID associated with the conversion. An order id can only be used
	// for one conversion per conversion action.
	OrderId *wrappers.StringValue `protobuf:"bytes,6,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	// Additional data about externally attributed conversions. This field
	// is required for conversions with an externally attributed conversion
	// action, but should not be set otherwise.
	ExternalAttributionData *ExternalAttributionData `protobuf:"bytes,7,opt,name=external_attribution_data,json=externalAttributionData,proto3" json:"external_attribution_data,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                 `json:"-"`
	XXX_unrecognized        []byte                   `json:"-"`
	XXX_sizecache           int32                    `json:"-"`
}

func (m *ClickConversion) Reset()         { *m = ClickConversion{} }
func (m *ClickConversion) String() string { return proto.CompactTextString(m) }
func (*ClickConversion) ProtoMessage()    {}
func (*ClickConversion) Descriptor() ([]byte, []int) {
	return fileDescriptor_962dfdf4c09453fe, []int{4}
}

func (m *ClickConversion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClickConversion.Unmarshal(m, b)
}
func (m *ClickConversion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClickConversion.Marshal(b, m, deterministic)
}
func (m *ClickConversion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClickConversion.Merge(m, src)
}
func (m *ClickConversion) XXX_Size() int {
	return xxx_messageInfo_ClickConversion.Size(m)
}
func (m *ClickConversion) XXX_DiscardUnknown() {
	xxx_messageInfo_ClickConversion.DiscardUnknown(m)
}

var xxx_messageInfo_ClickConversion proto.InternalMessageInfo

func (m *ClickConversion) GetGclid() *wrappers.StringValue {
	if m != nil {
		return m.Gclid
	}
	return nil
}

func (m *ClickConversion) GetConversionAction() *wrappers.StringValue {
	if m != nil {
		return m.ConversionAction
	}
	return nil
}

func (m *ClickConversion) GetConversionDateTime() *wrappers.StringValue {
	if m != nil {
		return m.ConversionDateTime
	}
	return nil
}

func (m *ClickConversion) GetConversionValue() *wrappers.DoubleValue {
	if m != nil {
		return m.ConversionValue
	}
	return nil
}

func (m *ClickConversion) GetCurrencyCode() *wrappers.StringValue {
	if m != nil {
		return m.CurrencyCode
	}
	return nil
}

func (m *ClickConversion) GetOrderId() *wrappers.StringValue {
	if m != nil {
		return m.OrderId
	}
	return nil
}

func (m *ClickConversion) GetExternalAttributionData() *ExternalAttributionData {
	if m != nil {
		return m.ExternalAttributionData
	}
	return nil
}

// A call conversion.
type CallConversion struct {
	// The caller id from which this call was placed. Caller id is expected to be
	// in E.164 format with preceding '+' sign. e.g. "+16502531234".
	CallerId *wrappers.StringValue `protobuf:"bytes,1,opt,name=caller_id,json=callerId,proto3" json:"caller_id,omitempty"`
	// The date time at which the call occurred. The timezone must be specified.
	// The format is "yyyy-mm-dd hh:mm:ss+|-hh:mm",
	// e.g. "2019-01-01 12:32:45-08:00".
	CallStartDateTime *wrappers.StringValue `protobuf:"bytes,2,opt,name=call_start_date_time,json=callStartDateTime,proto3" json:"call_start_date_time,omitempty"`
	// Resource name of the conversion action associated with this conversion.
	// Note: Although this resource name consists of a customer id and a
	// conversion action id, validation will ignore the customer id and use the
	// conversion action id as the sole identifier of the conversion action.
	ConversionAction *wrappers.StringValue `protobuf:"bytes,3,opt,name=conversion_action,json=conversionAction,proto3" json:"conversion_action,omitempty"`
	// The date time at which the conversion occurred. Must be after the call
	// time. The timezone must be specified. The format is
	// "yyyy-mm-dd hh:mm:ss+|-hh:mm", e.g. "2019-01-01 12:32:45-08:00".
	ConversionDateTime *wrappers.StringValue `protobuf:"bytes,4,opt,name=conversion_date_time,json=conversionDateTime,proto3" json:"conversion_date_time,omitempty"`
	// The value of the conversion for the advertiser.
	ConversionValue *wrappers.DoubleValue `protobuf:"bytes,5,opt,name=conversion_value,json=conversionValue,proto3" json:"conversion_value,omitempty"`
	// Currency associated with the conversion value. This is the ISO 4217
	// 3-character currency code. For example: USD, EUR.
	CurrencyCode         *wrappers.StringValue `protobuf:"bytes,6,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CallConversion) Reset()         { *m = CallConversion{} }
func (m *CallConversion) String() string { return proto.CompactTextString(m) }
func (*CallConversion) ProtoMessage()    {}
func (*CallConversion) Descriptor() ([]byte, []int) {
	return fileDescriptor_962dfdf4c09453fe, []int{5}
}

func (m *CallConversion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallConversion.Unmarshal(m, b)
}
func (m *CallConversion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallConversion.Marshal(b, m, deterministic)
}
func (m *CallConversion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallConversion.Merge(m, src)
}
func (m *CallConversion) XXX_Size() int {
	return xxx_messageInfo_CallConversion.Size(m)
}
func (m *CallConversion) XXX_DiscardUnknown() {
	xxx_messageInfo_CallConversion.DiscardUnknown(m)
}

var xxx_messageInfo_CallConversion proto.InternalMessageInfo

func (m *CallConversion) GetCallerId() *wrappers.StringValue {
	if m != nil {
		return m.CallerId
	}
	return nil
}

func (m *CallConversion) GetCallStartDateTime() *wrappers.StringValue {
	if m != nil {
		return m.CallStartDateTime
	}
	return nil
}

func (m *CallConversion) GetConversionAction() *wrappers.StringValue {
	if m != nil {
		return m.ConversionAction
	}
	return nil
}

func (m *CallConversion) GetConversionDateTime() *wrappers.StringValue {
	if m != nil {
		return m.ConversionDateTime
	}
	return nil
}

func (m *CallConversion) GetConversionValue() *wrappers.DoubleValue {
	if m != nil {
		return m.ConversionValue
	}
	return nil
}

func (m *CallConversion) GetCurrencyCode() *wrappers.StringValue {
	if m != nil {
		return m.CurrencyCode
	}
	return nil
}

// Contains additional information about externally attributed conversions.
type ExternalAttributionData struct {
	// Represents the fraction of the conversion that is attributed to the
	// Google Ads click.
	ExternalAttributionCredit *wrappers.DoubleValue `protobuf:"bytes,1,opt,name=external_attribution_credit,json=externalAttributionCredit,proto3" json:"external_attribution_credit,omitempty"`
	// Specifies the attribution model name.
	ExternalAttributionModel *wrappers.StringValue `protobuf:"bytes,2,opt,name=external_attribution_model,json=externalAttributionModel,proto3" json:"external_attribution_model,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}              `json:"-"`
	XXX_unrecognized         []byte                `json:"-"`
	XXX_sizecache            int32                 `json:"-"`
}

func (m *ExternalAttributionData) Reset()         { *m = ExternalAttributionData{} }
func (m *ExternalAttributionData) String() string { return proto.CompactTextString(m) }
func (*ExternalAttributionData) ProtoMessage()    {}
func (*ExternalAttributionData) Descriptor() ([]byte, []int) {
	return fileDescriptor_962dfdf4c09453fe, []int{6}
}

func (m *ExternalAttributionData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExternalAttributionData.Unmarshal(m, b)
}
func (m *ExternalAttributionData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExternalAttributionData.Marshal(b, m, deterministic)
}
func (m *ExternalAttributionData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalAttributionData.Merge(m, src)
}
func (m *ExternalAttributionData) XXX_Size() int {
	return xxx_messageInfo_ExternalAttributionData.Size(m)
}
func (m *ExternalAttributionData) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalAttributionData.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalAttributionData proto.InternalMessageInfo

func (m *ExternalAttributionData) GetExternalAttributionCredit() *wrappers.DoubleValue {
	if m != nil {
		return m.ExternalAttributionCredit
	}
	return nil
}

func (m *ExternalAttributionData) GetExternalAttributionModel() *wrappers.StringValue {
	if m != nil {
		return m.ExternalAttributionModel
	}
	return nil
}

// Identifying information for a successfully processed ClickConversion.
type ClickConversionResult struct {
	// The Google Click ID (gclid) associated with this conversion.
	Gclid *wrappers.StringValue `protobuf:"bytes,1,opt,name=gclid,proto3" json:"gclid,omitempty"`
	// Resource name of the conversion action associated with this conversion.
	ConversionAction *wrappers.StringValue `protobuf:"bytes,2,opt,name=conversion_action,json=conversionAction,proto3" json:"conversion_action,omitempty"`
	// The date time at which the conversion occurred. The format is
	// "yyyy-mm-dd hh:mm:ss+|-hh:mm", e.g. “2019-01-01 12:32:45-08:00”.
	ConversionDateTime   *wrappers.StringValue `protobuf:"bytes,3,opt,name=conversion_date_time,json=conversionDateTime,proto3" json:"conversion_date_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ClickConversionResult) Reset()         { *m = ClickConversionResult{} }
func (m *ClickConversionResult) String() string { return proto.CompactTextString(m) }
func (*ClickConversionResult) ProtoMessage()    {}
func (*ClickConversionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_962dfdf4c09453fe, []int{7}
}

func (m *ClickConversionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClickConversionResult.Unmarshal(m, b)
}
func (m *ClickConversionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClickConversionResult.Marshal(b, m, deterministic)
}
func (m *ClickConversionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClickConversionResult.Merge(m, src)
}
func (m *ClickConversionResult) XXX_Size() int {
	return xxx_messageInfo_ClickConversionResult.Size(m)
}
func (m *ClickConversionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ClickConversionResult.DiscardUnknown(m)
}

var xxx_messageInfo_ClickConversionResult proto.InternalMessageInfo

func (m *ClickConversionResult) GetGclid() *wrappers.StringValue {
	if m != nil {
		return m.Gclid
	}
	return nil
}

func (m *ClickConversionResult) GetConversionAction() *wrappers.StringValue {
	if m != nil {
		return m.ConversionAction
	}
	return nil
}

func (m *ClickConversionResult) GetConversionDateTime() *wrappers.StringValue {
	if m != nil {
		return m.ConversionDateTime
	}
	return nil
}

// Identifying information for a successfully processed CallConversionUpload.
type CallConversionResult struct {
	// The caller id from which this call was placed. Caller id is expected to be
	// in E.164 format with preceding '+' sign.
	CallerId *wrappers.StringValue `protobuf:"bytes,1,opt,name=caller_id,json=callerId,proto3" json:"caller_id,omitempty"`
	// The date time at which the call occurred. The format is
	// "yyyy-mm-dd hh:mm:ss+|-hh:mm", e.g. "2019-01-01 12:32:45-08:00".
	CallStartDateTime *wrappers.StringValue `protobuf:"bytes,2,opt,name=call_start_date_time,json=callStartDateTime,proto3" json:"call_start_date_time,omitempty"`
	// Resource name of the conversion action associated with this conversion.
	ConversionAction *wrappers.StringValue `protobuf:"bytes,3,opt,name=conversion_action,json=conversionAction,proto3" json:"conversion_action,omitempty"`
	// The date time at which the conversion occurred. The format is
	// "yyyy-mm-dd hh:mm:ss+|-hh:mm", e.g. "2019-01-01 12:32:45-08:00".
	ConversionDateTime   *wrappers.StringValue `protobuf:"bytes,4,opt,name=conversion_date_time,json=conversionDateTime,proto3" json:"conversion_date_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CallConversionResult) Reset()         { *m = CallConversionResult{} }
func (m *CallConversionResult) String() string { return proto.CompactTextString(m) }
func (*CallConversionResult) ProtoMessage()    {}
func (*CallConversionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_962dfdf4c09453fe, []int{8}
}

func (m *CallConversionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallConversionResult.Unmarshal(m, b)
}
func (m *CallConversionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallConversionResult.Marshal(b, m, deterministic)
}
func (m *CallConversionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallConversionResult.Merge(m, src)
}
func (m *CallConversionResult) XXX_Size() int {
	return xxx_messageInfo_CallConversionResult.Size(m)
}
func (m *CallConversionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CallConversionResult.DiscardUnknown(m)
}

var xxx_messageInfo_CallConversionResult proto.InternalMessageInfo

func (m *CallConversionResult) GetCallerId() *wrappers.StringValue {
	if m != nil {
		return m.CallerId
	}
	return nil
}

func (m *CallConversionResult) GetCallStartDateTime() *wrappers.StringValue {
	if m != nil {
		return m.CallStartDateTime
	}
	return nil
}

func (m *CallConversionResult) GetConversionAction() *wrappers.StringValue {
	if m != nil {
		return m.ConversionAction
	}
	return nil
}

func (m *CallConversionResult) GetConversionDateTime() *wrappers.StringValue {
	if m != nil {
		return m.ConversionDateTime
	}
	return nil
}

func init() {
	proto.RegisterType((*UploadClickConversionsRequest)(nil), "google.ads.googleads.v2.services.UploadClickConversionsRequest")
	proto.RegisterType((*UploadClickConversionsResponse)(nil), "google.ads.googleads.v2.services.UploadClickConversionsResponse")
	proto.RegisterType((*UploadCallConversionsRequest)(nil), "google.ads.googleads.v2.services.UploadCallConversionsRequest")
	proto.RegisterType((*UploadCallConversionsResponse)(nil), "google.ads.googleads.v2.services.UploadCallConversionsResponse")
	proto.RegisterType((*ClickConversion)(nil), "google.ads.googleads.v2.services.ClickConversion")
	proto.RegisterType((*CallConversion)(nil), "google.ads.googleads.v2.services.CallConversion")
	proto.RegisterType((*ExternalAttributionData)(nil), "google.ads.googleads.v2.services.ExternalAttributionData")
	proto.RegisterType((*ClickConversionResult)(nil), "google.ads.googleads.v2.services.ClickConversionResult")
	proto.RegisterType((*CallConversionResult)(nil), "google.ads.googleads.v2.services.CallConversionResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/conversion_upload_service.proto", fileDescriptor_962dfdf4c09453fe)
}

var fileDescriptor_962dfdf4c09453fe = []byte{
	// 945 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x57, 0x4d, 0x6f, 0xdc, 0x44,
	0x18, 0x96, 0x77, 0xf3, 0xd5, 0xd9, 0xb6, 0xa1, 0x43, 0x42, 0xdc, 0x6d, 0x28, 0x2b, 0x53, 0x89,
	0x28, 0x07, 0x1b, 0x5c, 0x44, 0x89, 0x11, 0x4d, 0xdd, 0xa4, 0x0d, 0x39, 0x14, 0x82, 0x17, 0x72,
	0xa8, 0x56, 0xb2, 0x26, 0xf6, 0x74, 0x65, 0x31, 0xeb, 0x31, 0x33, 0xe3, 0x85, 0x08, 0x71, 0x81,
	0x9f, 0xc0, 0x95, 0x13, 0x47, 0xfe, 0x01, 0xea, 0x89, 0x6b, 0xaf, 0xbd, 0x70, 0xe0, 0x82, 0x40,
	0x1c, 0x38, 0xf1, 0x13, 0x90, 0x3d, 0x9e, 0xf5, 0x66, 0xe5, 0x65, 0x9d, 0x68, 0xb9, 0xa0, 0xde,
	0x66, 0x67, 0xde, 0xf7, 0x79, 0x3f, 0x9e, 0x67, 0xc6, 0xef, 0x82, 0x7b, 0x7d, 0x4a, 0xfb, 0x04,
	0x5b, 0x28, 0xe4, 0x96, 0x5c, 0x66, 0xab, 0xa1, 0x6d, 0x71, 0xcc, 0x86, 0x51, 0x80, 0xb9, 0x15,
	0xd0, 0x78, 0x88, 0x19, 0x8f, 0x68, 0xec, 0xa7, 0x09, 0xa1, 0x28, 0xf4, 0x8b, 0x23, 0x33, 0x61,
	0x54, 0x50, 0xd8, 0x91, 0x6e, 0x26, 0x0a, 0xb9, 0x39, 0x42, 0x30, 0x87, 0xb6, 0xa9, 0x10, 0xda,
	0x9b, 0x2a, 0x46, 0x12, 0x59, 0x28, 0x8e, 0xa9, 0x40, 0x22, 0xa2, 0x31, 0x97, 0xfe, 0xed, 0x9b,
	0xc5, 0x69, 0xfe, 0xeb, 0x24, 0x7d, 0x62, 0x7d, 0xc1, 0x50, 0x92, 0x60, 0xa6, 0xce, 0x37, 0x8a,
	0x73, 0x96, 0x04, 0x16, 0x17, 0x48, 0xa4, 0x93, 0x07, 0x19, 0x6c, 0x40, 0x22, 0x1c, 0x0b, 0x79,
	0x60, 0xfc, 0xae, 0x81, 0x57, 0x3f, 0xcd, 0x53, 0xdd, 0x23, 0x51, 0xf0, 0xd9, 0xde, 0xa8, 0x00,
	0xee, 0xe1, 0xcf, 0x53, 0xcc, 0x05, 0x7c, 0x0d, 0xb4, 0x82, 0x94, 0x0b, 0x3a, 0xc0, 0xcc, 0x8f,
	0x42, 0x5d, 0xeb, 0x68, 0x5b, 0x97, 0x3c, 0xa0, 0xb6, 0x0e, 0x43, 0xd8, 0x05, 0xad, 0xb2, 0x6e,
	0xae, 0x37, 0x3a, 0xcd, 0xad, 0x96, 0xfd, 0x96, 0x39, 0xab, 0x54, 0x73, 0x22, 0xa0, 0x37, 0x8e,
	0x02, 0xdf, 0x00, 0xab, 0x09, 0x62, 0x22, 0x42, 0xc4, 0x7f, 0x82, 0x22, 0x92, 0x32, 0xac, 0x37,
	0x3b, 0xda, 0xd6, 0x8a, 0x77, 0xb5, 0xd8, 0x7e, 0x28, 0x77, 0xe1, 0xeb, 0xe0, 0xca, 0x10, 0x91,
	0x28, 0x44, 0x02, 0xfb, 0x34, 0x26, 0xa7, 0xfa, 0x42, 0x6e, 0x76, 0x59, 0x6d, 0x7e, 0x14, 0x93,
	0x53, 0xe3, 0xa9, 0x06, 0x6e, 0x4e, 0xab, 0x92, 0x27, 0x34, 0xe6, 0x18, 0x3e, 0x04, 0xeb, 0x13,
	0x01, 0x7d, 0xcc, 0x18, 0x65, 0x79, 0xc1, 0x2d, 0x1b, 0xaa, 0x7a, 0x58, 0x12, 0x98, 0xdd, 0xbc,
	0xb5, 0xde, 0xcb, 0x67, 0x53, 0x79, 0x90, 0x99, 0xc3, 0x8f, 0xc1, 0x32, 0xc3, 0x3c, 0x25, 0x42,
	0x75, 0xe2, 0xce, 0xf9, 0x3b, 0x91, 0xfb, 0x7b, 0x0a, 0xc7, 0xf8, 0x4d, 0x03, 0x9b, 0x45, 0xf6,
	0x88, 0x90, 0x8b, 0x50, 0xe4, 0x55, 0x51, 0xf4, 0x66, 0x8d, 0xc4, 0xce, 0xc4, 0xfb, 0x2f, 0x19,
	0xfa, 0xa9, 0xd4, 0xe1, 0x64, 0x8d, 0x73, 0x26, 0xe8, 0x68, 0x92, 0xa0, 0x77, 0xce, 0xdd, 0x87,
	0x09, 0x7e, 0xbe, 0x5f, 0x00, 0xab, 0x13, 0x14, 0x42, 0x1b, 0x2c, 0xf6, 0x03, 0x52, 0x90, 0xd1,
	0xb2, 0x37, 0x55, 0x0c, 0x75, 0x73, 0xcd, 0xae, 0x60, 0x51, 0xdc, 0x3f, 0x46, 0x24, 0xc5, 0x9e,
	0x34, 0x85, 0x87, 0xe0, 0xda, 0xd8, 0x03, 0x82, 0x82, 0xec, 0xe6, 0xeb, 0x8d, 0x1a, 0xfe, 0x2f,
	0x95, 0x6e, 0x6e, 0xee, 0x05, 0x3f, 0x04, 0x6b, 0x63, 0x50, 0x79, 0xeb, 0x45, 0x34, 0x90, 0x0c,
	0xcd, 0x42, 0x83, 0xa5, 0xe7, 0x3e, 0x12, 0xf8, 0x93, 0x68, 0x80, 0xe1, 0x01, 0x18, 0x8b, 0xe1,
	0x0f, 0x33, 0xbb, 0x9c, 0xc6, 0x2a, 0xac, 0x7d, 0x9a, 0x9e, 0x10, 0x2c, 0xb1, 0x56, 0x4b, 0xaf,
	0x7c, 0x03, 0xba, 0xe0, 0x4a, 0x90, 0x32, 0x86, 0xe3, 0xe0, 0xd4, 0x0f, 0x68, 0x88, 0xf5, 0xc5,
	0x1a, 0x19, 0x5d, 0x56, 0x2e, 0x7b, 0x34, 0xc4, 0xf0, 0x0e, 0x58, 0xa1, 0x2c, 0x94, 0x52, 0x5f,
	0xaa, 0xe1, 0xbd, 0x9c, 0x5b, 0x1f, 0x86, 0x30, 0x05, 0xd7, 0xf1, 0x97, 0x02, 0xb3, 0x18, 0x11,
	0x1f, 0x09, 0xc1, 0xa2, 0x93, 0x54, 0x14, 0xed, 0x41, 0xfa, 0x72, 0x8e, 0xb4, 0x33, 0x5b, 0x0b,
	0x0f, 0x0a, 0x08, 0xb7, 0x44, 0xd8, 0x47, 0x02, 0x79, 0x1b, 0xb8, 0xfa, 0xc0, 0xf8, 0xb9, 0x09,
	0xae, 0x9e, 0x15, 0x10, 0xdc, 0x01, 0x97, 0x02, 0x44, 0x48, 0x79, 0x5d, 0x67, 0xd5, 0xb0, 0x22,
	0xcd, 0x0f, 0x43, 0xf8, 0x08, 0xac, 0x65, 0x6b, 0x9f, 0x0b, 0xc4, 0xc4, 0x18, 0xb3, 0x75, 0x74,
	0x72, 0x2d, 0xf3, 0xec, 0x66, 0x8e, 0x23, 0x62, 0x2b, 0x35, 0xd7, 0x9c, 0xab, 0xe6, 0x16, 0xe6,
	0xa8, 0xb9, 0xc5, 0xb9, 0x68, 0x6e, 0xe9, 0xbc, 0x9a, 0x33, 0x9e, 0x6b, 0x60, 0x63, 0x0a, 0xf1,
	0xb0, 0x07, 0x6e, 0x54, 0xca, 0x2a, 0x60, 0x38, 0x8c, 0xc4, 0x54, 0x7a, 0xc7, 0x53, 0xbe, 0x5e,
	0xa1, 0x9d, 0xbd, 0xdc, 0x1d, 0x3e, 0x06, 0xed, 0x4a, 0xf4, 0x01, 0x0d, 0x31, 0xa9, 0xc5, 0xba,
	0x5e, 0x01, 0xfe, 0x28, 0xf3, 0x36, 0xfe, 0xd4, 0xc0, 0x7a, 0xe5, 0xb7, 0xe7, 0x7f, 0xf6, 0x7c,
	0x19, 0x4f, 0x1b, 0x60, 0xad, 0xea, 0x0d, 0x7f, 0x71, 0x11, 0x6b, 0x74, 0xcf, 0xfe, 0xbb, 0x09,
	0x36, 0xca, 0xce, 0xc9, 0xaf, 0x74, 0x57, 0x3e, 0x87, 0xf0, 0x57, 0x0d, 0xbc, 0x52, 0x3d, 0x59,
	0xc1, 0xdd, 0xd9, 0x6f, 0xe9, 0xbf, 0x4e, 0x9e, 0xed, 0x7b, 0x17, 0x07, 0x90, 0x33, 0x83, 0xb1,
	0xfb, 0xcd, 0xf3, 0x3f, 0xbe, 0x6b, 0xec, 0x18, 0x6f, 0x67, 0x53, 0xba, 0x9a, 0x87, 0xb8, 0xf5,
	0xd5, 0xd8, 0xb4, 0xf4, 0xfe, 0xf6, 0xd7, 0x4e, 0x5a, 0x89, 0xe2, 0x68, 0xdb, 0xf0, 0x17, 0x0d,
	0xac, 0x57, 0x8e, 0x25, 0xf0, 0x6e, 0xed, 0xe4, 0x2a, 0x67, 0xb6, 0xf6, 0xee, 0x85, 0xfd, 0x8b,
	0xda, 0xee, 0xe6, 0xb5, 0xbd, 0x6b, 0xdc, 0xae, 0x55, 0xdb, 0x59, 0x10, 0x47, 0xdb, 0x6e, 0xdf,
	0x78, 0xe6, 0xea, 0x65, 0xdc, 0x62, 0x95, 0x44, 0xdc, 0x0c, 0xe8, 0xe0, 0xfe, 0xb7, 0x0d, 0x70,
	0x2b, 0xa0, 0x83, 0x99, 0x39, 0xde, 0xdf, 0x9c, 0x22, 0x8c, 0xa3, 0x4c, 0x5c, 0x47, 0xda, 0xe3,
	0x0f, 0x0a, 0x84, 0x3e, 0x25, 0x28, 0xee, 0x9b, 0x94, 0xf5, 0xad, 0x3e, 0x8e, 0x73, 0xe9, 0x59,
	0x65, 0xcc, 0xe9, 0x7f, 0xa9, 0xde, 0x53, 0x8b, 0x1f, 0x1a, 0xcd, 0x03, 0xd7, 0xfd, 0xb1, 0xd1,
	0x39, 0x90, 0x80, 0x6e, 0xc8, 0x4d, 0xb9, 0xcc, 0x56, 0xc7, 0xb6, 0x59, 0x04, 0xe6, 0xcf, 0x94,
	0x49, 0xcf, 0x0d, 0x79, 0x6f, 0x64, 0xd2, 0x3b, 0xb6, 0x7b, 0xca, 0xe4, 0xaf, 0xc6, 0x2d, 0xb9,
	0xef, 0x38, 0x6e, 0xc8, 0x1d, 0x67, 0x64, 0xe4, 0x38, 0xc7, 0xb6, 0xe3, 0x28, 0xb3, 0x93, 0xa5,
	0x3c, 0xcf, 0xdb, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xd3, 0x6c, 0x0e, 0xf9, 0x0d, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConversionUploadServiceClient is the client API for ConversionUploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConversionUploadServiceClient interface {
	// Processes the given click conversions.
	UploadClickConversions(ctx context.Context, in *UploadClickConversionsRequest, opts ...grpc.CallOption) (*UploadClickConversionsResponse, error)
	// Processes the given call conversions.
	UploadCallConversions(ctx context.Context, in *UploadCallConversionsRequest, opts ...grpc.CallOption) (*UploadCallConversionsResponse, error)
}

type conversionUploadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConversionUploadServiceClient(cc grpc.ClientConnInterface) ConversionUploadServiceClient {
	return &conversionUploadServiceClient{cc}
}

func (c *conversionUploadServiceClient) UploadClickConversions(ctx context.Context, in *UploadClickConversionsRequest, opts ...grpc.CallOption) (*UploadClickConversionsResponse, error) {
	out := new(UploadClickConversionsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.ConversionUploadService/UploadClickConversions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversionUploadServiceClient) UploadCallConversions(ctx context.Context, in *UploadCallConversionsRequest, opts ...grpc.CallOption) (*UploadCallConversionsResponse, error) {
	out := new(UploadCallConversionsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.ConversionUploadService/UploadCallConversions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConversionUploadServiceServer is the server API for ConversionUploadService service.
type ConversionUploadServiceServer interface {
	// Processes the given click conversions.
	UploadClickConversions(context.Context, *UploadClickConversionsRequest) (*UploadClickConversionsResponse, error)
	// Processes the given call conversions.
	UploadCallConversions(context.Context, *UploadCallConversionsRequest) (*UploadCallConversionsResponse, error)
}

// UnimplementedConversionUploadServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConversionUploadServiceServer struct {
}

func (*UnimplementedConversionUploadServiceServer) UploadClickConversions(ctx context.Context, req *UploadClickConversionsRequest) (*UploadClickConversionsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method UploadClickConversions not implemented")
}
func (*UnimplementedConversionUploadServiceServer) UploadCallConversions(ctx context.Context, req *UploadCallConversionsRequest) (*UploadCallConversionsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method UploadCallConversions not implemented")
}

func RegisterConversionUploadServiceServer(s *grpc.Server, srv ConversionUploadServiceServer) {
	s.RegisterService(&_ConversionUploadService_serviceDesc, srv)
}

func _ConversionUploadService_UploadClickConversions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadClickConversionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversionUploadServiceServer).UploadClickConversions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.ConversionUploadService/UploadClickConversions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversionUploadServiceServer).UploadClickConversions(ctx, req.(*UploadClickConversionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConversionUploadService_UploadCallConversions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadCallConversionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversionUploadServiceServer).UploadCallConversions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.ConversionUploadService/UploadCallConversions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversionUploadServiceServer).UploadCallConversions(ctx, req.(*UploadCallConversionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConversionUploadService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.ConversionUploadService",
	HandlerType: (*ConversionUploadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadClickConversions",
			Handler:    _ConversionUploadService_UploadClickConversions_Handler,
		},
		{
			MethodName: "UploadCallConversions",
			Handler:    _ConversionUploadService_UploadCallConversions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/conversion_upload_service.proto",
}