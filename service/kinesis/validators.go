// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAddTagsToStream struct {
}

func (*validateOpAddTagsToStream) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAddTagsToStream) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AddTagsToStreamInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAddTagsToStreamInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateStream struct {
}

func (*validateOpCreateStream) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateStream) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateStreamInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateStreamInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDecreaseStreamRetentionPeriod struct {
}

func (*validateOpDecreaseStreamRetentionPeriod) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDecreaseStreamRetentionPeriod) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DecreaseStreamRetentionPeriodInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDecreaseStreamRetentionPeriodInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteResourcePolicy struct {
}

func (*validateOpDeleteResourcePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteResourcePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteResourcePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteResourcePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisableEnhancedMonitoring struct {
}

func (*validateOpDisableEnhancedMonitoring) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisableEnhancedMonitoring) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisableEnhancedMonitoringInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisableEnhancedMonitoringInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpEnableEnhancedMonitoring struct {
}

func (*validateOpEnableEnhancedMonitoring) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpEnableEnhancedMonitoring) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*EnableEnhancedMonitoringInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpEnableEnhancedMonitoringInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetRecords struct {
}

func (*validateOpGetRecords) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetRecords) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetRecordsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetRecordsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetResourcePolicy struct {
}

func (*validateOpGetResourcePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetResourcePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetResourcePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetResourcePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetShardIterator struct {
}

func (*validateOpGetShardIterator) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetShardIterator) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetShardIteratorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetShardIteratorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpIncreaseStreamRetentionPeriod struct {
}

func (*validateOpIncreaseStreamRetentionPeriod) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpIncreaseStreamRetentionPeriod) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*IncreaseStreamRetentionPeriodInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpIncreaseStreamRetentionPeriodInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListShards struct {
}

func (*validateOpListShards) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListShards) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListShardsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListShardsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListStreamConsumers struct {
}

func (*validateOpListStreamConsumers) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListStreamConsumers) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListStreamConsumersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListStreamConsumersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpMergeShards struct {
}

func (*validateOpMergeShards) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpMergeShards) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*MergeShardsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpMergeShardsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutRecord struct {
}

func (*validateOpPutRecord) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutRecord) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutRecordInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutRecordInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutRecords struct {
}

func (*validateOpPutRecords) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutRecords) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutRecordsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutRecordsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutResourcePolicy struct {
}

func (*validateOpPutResourcePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutResourcePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutResourcePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutResourcePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRegisterStreamConsumer struct {
}

func (*validateOpRegisterStreamConsumer) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRegisterStreamConsumer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RegisterStreamConsumerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRegisterStreamConsumerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRemoveTagsFromStream struct {
}

func (*validateOpRemoveTagsFromStream) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRemoveTagsFromStream) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RemoveTagsFromStreamInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRemoveTagsFromStreamInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSplitShard struct {
}

func (*validateOpSplitShard) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSplitShard) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SplitShardInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSplitShardInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartStreamEncryption struct {
}

func (*validateOpStartStreamEncryption) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartStreamEncryption) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartStreamEncryptionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartStreamEncryptionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopStreamEncryption struct {
}

func (*validateOpStopStreamEncryption) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopStreamEncryption) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopStreamEncryptionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopStreamEncryptionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSubscribeToShard struct {
}

func (*validateOpSubscribeToShard) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSubscribeToShard) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SubscribeToShardInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSubscribeToShardInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateShardCount struct {
}

func (*validateOpUpdateShardCount) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateShardCount) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateShardCountInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateShardCountInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateStreamMode struct {
}

func (*validateOpUpdateStreamMode) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateStreamMode) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateStreamModeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateStreamModeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAddTagsToStreamValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAddTagsToStream{}, middleware.After)
}

func addOpCreateStreamValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateStream{}, middleware.After)
}

func addOpDecreaseStreamRetentionPeriodValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDecreaseStreamRetentionPeriod{}, middleware.After)
}

func addOpDeleteResourcePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteResourcePolicy{}, middleware.After)
}

func addOpDisableEnhancedMonitoringValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisableEnhancedMonitoring{}, middleware.After)
}

func addOpEnableEnhancedMonitoringValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpEnableEnhancedMonitoring{}, middleware.After)
}

func addOpGetRecordsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetRecords{}, middleware.After)
}

func addOpGetResourcePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetResourcePolicy{}, middleware.After)
}

func addOpGetShardIteratorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetShardIterator{}, middleware.After)
}

func addOpIncreaseStreamRetentionPeriodValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpIncreaseStreamRetentionPeriod{}, middleware.After)
}

func addOpListShardsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListShards{}, middleware.After)
}

func addOpListStreamConsumersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListStreamConsumers{}, middleware.After)
}

func addOpMergeShardsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpMergeShards{}, middleware.After)
}

func addOpPutRecordValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutRecord{}, middleware.After)
}

func addOpPutRecordsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutRecords{}, middleware.After)
}

func addOpPutResourcePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutResourcePolicy{}, middleware.After)
}

func addOpRegisterStreamConsumerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRegisterStreamConsumer{}, middleware.After)
}

func addOpRemoveTagsFromStreamValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRemoveTagsFromStream{}, middleware.After)
}

func addOpSplitShardValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSplitShard{}, middleware.After)
}

func addOpStartStreamEncryptionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartStreamEncryption{}, middleware.After)
}

func addOpStopStreamEncryptionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopStreamEncryption{}, middleware.After)
}

func addOpSubscribeToShardValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSubscribeToShard{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateShardCountValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateShardCount{}, middleware.After)
}

func addOpUpdateStreamModeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateStreamMode{}, middleware.After)
}

func validatePutRecordsRequestEntry(v *types.PutRecordsRequestEntry) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutRecordsRequestEntry"}
	if v.Data == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Data"))
	}
	if v.PartitionKey == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PartitionKey"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePutRecordsRequestEntryList(v []types.PutRecordsRequestEntry) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutRecordsRequestEntryList"}
	for i := range v {
		if err := validatePutRecordsRequestEntry(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateShardFilter(v *types.ShardFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ShardFilter"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateStartingPosition(v *types.StartingPosition) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartingPosition"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateStreamModeDetails(v *types.StreamModeDetails) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StreamModeDetails"}
	if len(v.StreamMode) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("StreamMode"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAddTagsToStreamInput(v *AddTagsToStreamInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AddTagsToStreamInput"}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateStreamInput(v *CreateStreamInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateStreamInput"}
	if v.StreamName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StreamName"))
	}
	if v.StreamModeDetails != nil {
		if err := validateStreamModeDetails(v.StreamModeDetails); err != nil {
			invalidParams.AddNested("StreamModeDetails", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDecreaseStreamRetentionPeriodInput(v *DecreaseStreamRetentionPeriodInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DecreaseStreamRetentionPeriodInput"}
	if v.RetentionPeriodHours == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RetentionPeriodHours"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteResourcePolicyInput(v *DeleteResourcePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteResourcePolicyInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisableEnhancedMonitoringInput(v *DisableEnhancedMonitoringInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisableEnhancedMonitoringInput"}
	if v.ShardLevelMetrics == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ShardLevelMetrics"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpEnableEnhancedMonitoringInput(v *EnableEnhancedMonitoringInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EnableEnhancedMonitoringInput"}
	if v.ShardLevelMetrics == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ShardLevelMetrics"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetRecordsInput(v *GetRecordsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetRecordsInput"}
	if v.ShardIterator == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ShardIterator"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetResourcePolicyInput(v *GetResourcePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetResourcePolicyInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetShardIteratorInput(v *GetShardIteratorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetShardIteratorInput"}
	if v.ShardId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ShardId"))
	}
	if len(v.ShardIteratorType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ShardIteratorType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpIncreaseStreamRetentionPeriodInput(v *IncreaseStreamRetentionPeriodInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "IncreaseStreamRetentionPeriodInput"}
	if v.RetentionPeriodHours == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RetentionPeriodHours"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListShardsInput(v *ListShardsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListShardsInput"}
	if v.ShardFilter != nil {
		if err := validateShardFilter(v.ShardFilter); err != nil {
			invalidParams.AddNested("ShardFilter", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListStreamConsumersInput(v *ListStreamConsumersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListStreamConsumersInput"}
	if v.StreamARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StreamARN"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpMergeShardsInput(v *MergeShardsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MergeShardsInput"}
	if v.ShardToMerge == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ShardToMerge"))
	}
	if v.AdjacentShardToMerge == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AdjacentShardToMerge"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutRecordInput(v *PutRecordInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutRecordInput"}
	if v.Data == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Data"))
	}
	if v.PartitionKey == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PartitionKey"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutRecordsInput(v *PutRecordsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutRecordsInput"}
	if v.Records == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Records"))
	} else if v.Records != nil {
		if err := validatePutRecordsRequestEntryList(v.Records); err != nil {
			invalidParams.AddNested("Records", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutResourcePolicyInput(v *PutResourcePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutResourcePolicyInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if v.Policy == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Policy"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRegisterStreamConsumerInput(v *RegisterStreamConsumerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RegisterStreamConsumerInput"}
	if v.StreamARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StreamARN"))
	}
	if v.ConsumerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConsumerName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRemoveTagsFromStreamInput(v *RemoveTagsFromStreamInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RemoveTagsFromStreamInput"}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSplitShardInput(v *SplitShardInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SplitShardInput"}
	if v.ShardToSplit == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ShardToSplit"))
	}
	if v.NewStartingHashKey == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NewStartingHashKey"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartStreamEncryptionInput(v *StartStreamEncryptionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartStreamEncryptionInput"}
	if len(v.EncryptionType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("EncryptionType"))
	}
	if v.KeyId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KeyId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopStreamEncryptionInput(v *StopStreamEncryptionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopStreamEncryptionInput"}
	if len(v.EncryptionType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("EncryptionType"))
	}
	if v.KeyId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KeyId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSubscribeToShardInput(v *SubscribeToShardInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SubscribeToShardInput"}
	if v.ConsumerARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConsumerARN"))
	}
	if v.ShardId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ShardId"))
	}
	if v.StartingPosition == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StartingPosition"))
	} else if v.StartingPosition != nil {
		if err := validateStartingPosition(v.StartingPosition); err != nil {
			invalidParams.AddNested("StartingPosition", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateShardCountInput(v *UpdateShardCountInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateShardCountInput"}
	if v.TargetShardCount == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TargetShardCount"))
	}
	if len(v.ScalingType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ScalingType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateStreamModeInput(v *UpdateStreamModeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateStreamModeInput"}
	if v.StreamARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StreamARN"))
	}
	if v.StreamModeDetails == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StreamModeDetails"))
	} else if v.StreamModeDetails != nil {
		if err := validateStreamModeDetails(v.StreamModeDetails); err != nil {
			invalidParams.AddNested("StreamModeDetails", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
