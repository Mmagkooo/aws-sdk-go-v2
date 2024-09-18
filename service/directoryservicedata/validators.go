// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservicedata

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAddGroupMember struct {
}

func (*validateOpAddGroupMember) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAddGroupMember) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AddGroupMemberInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAddGroupMemberInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateGroup struct {
}

func (*validateOpCreateGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateUser struct {
}

func (*validateOpCreateUser) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateUser) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateUserInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateUserInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteGroup struct {
}

func (*validateOpDeleteGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteUser struct {
}

func (*validateOpDeleteUser) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteUser) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteUserInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteUserInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeGroup struct {
}

func (*validateOpDescribeGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeUser struct {
}

func (*validateOpDescribeUser) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeUser) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeUserInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeUserInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisableUser struct {
}

func (*validateOpDisableUser) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisableUser) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisableUserInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisableUserInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListGroupMembers struct {
}

func (*validateOpListGroupMembers) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListGroupMembers) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListGroupMembersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListGroupMembersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListGroupsForMember struct {
}

func (*validateOpListGroupsForMember) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListGroupsForMember) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListGroupsForMemberInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListGroupsForMemberInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListGroups struct {
}

func (*validateOpListGroups) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListGroups) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListGroupsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListGroupsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListUsers struct {
}

func (*validateOpListUsers) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListUsers) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListUsersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListUsersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRemoveGroupMember struct {
}

func (*validateOpRemoveGroupMember) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRemoveGroupMember) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RemoveGroupMemberInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRemoveGroupMemberInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSearchGroups struct {
}

func (*validateOpSearchGroups) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSearchGroups) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SearchGroupsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSearchGroupsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSearchUsers struct {
}

func (*validateOpSearchUsers) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSearchUsers) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SearchUsersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSearchUsersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateGroup struct {
}

func (*validateOpUpdateGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateUser struct {
}

func (*validateOpUpdateUser) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateUser) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateUserInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateUserInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAddGroupMemberValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAddGroupMember{}, middleware.After)
}

func addOpCreateGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateGroup{}, middleware.After)
}

func addOpCreateUserValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateUser{}, middleware.After)
}

func addOpDeleteGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteGroup{}, middleware.After)
}

func addOpDeleteUserValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteUser{}, middleware.After)
}

func addOpDescribeGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeGroup{}, middleware.After)
}

func addOpDescribeUserValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeUser{}, middleware.After)
}

func addOpDisableUserValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisableUser{}, middleware.After)
}

func addOpListGroupMembersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListGroupMembers{}, middleware.After)
}

func addOpListGroupsForMemberValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListGroupsForMember{}, middleware.After)
}

func addOpListGroupsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListGroups{}, middleware.After)
}

func addOpListUsersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListUsers{}, middleware.After)
}

func addOpRemoveGroupMemberValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRemoveGroupMember{}, middleware.After)
}

func addOpSearchGroupsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSearchGroups{}, middleware.After)
}

func addOpSearchUsersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSearchUsers{}, middleware.After)
}

func addOpUpdateGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateGroup{}, middleware.After)
}

func addOpUpdateUserValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateUser{}, middleware.After)
}

func validateOpAddGroupMemberInput(v *AddGroupMemberInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AddGroupMemberInput"}
	if v.DirectoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DirectoryId"))
	}
	if v.GroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GroupName"))
	}
	if v.MemberName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemberName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateGroupInput(v *CreateGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateGroupInput"}
	if v.DirectoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DirectoryId"))
	}
	if v.SAMAccountName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SAMAccountName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateUserInput(v *CreateUserInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateUserInput"}
	if v.DirectoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DirectoryId"))
	}
	if v.SAMAccountName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SAMAccountName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteGroupInput(v *DeleteGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteGroupInput"}
	if v.DirectoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DirectoryId"))
	}
	if v.SAMAccountName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SAMAccountName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteUserInput(v *DeleteUserInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteUserInput"}
	if v.DirectoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DirectoryId"))
	}
	if v.SAMAccountName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SAMAccountName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeGroupInput(v *DescribeGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeGroupInput"}
	if v.DirectoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DirectoryId"))
	}
	if v.SAMAccountName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SAMAccountName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeUserInput(v *DescribeUserInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeUserInput"}
	if v.DirectoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DirectoryId"))
	}
	if v.SAMAccountName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SAMAccountName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisableUserInput(v *DisableUserInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisableUserInput"}
	if v.DirectoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DirectoryId"))
	}
	if v.SAMAccountName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SAMAccountName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListGroupMembersInput(v *ListGroupMembersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListGroupMembersInput"}
	if v.DirectoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DirectoryId"))
	}
	if v.SAMAccountName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SAMAccountName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListGroupsForMemberInput(v *ListGroupsForMemberInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListGroupsForMemberInput"}
	if v.DirectoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DirectoryId"))
	}
	if v.SAMAccountName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SAMAccountName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListGroupsInput(v *ListGroupsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListGroupsInput"}
	if v.DirectoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DirectoryId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListUsersInput(v *ListUsersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListUsersInput"}
	if v.DirectoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DirectoryId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRemoveGroupMemberInput(v *RemoveGroupMemberInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RemoveGroupMemberInput"}
	if v.DirectoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DirectoryId"))
	}
	if v.GroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GroupName"))
	}
	if v.MemberName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemberName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSearchGroupsInput(v *SearchGroupsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SearchGroupsInput"}
	if v.DirectoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DirectoryId"))
	}
	if v.SearchString == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SearchString"))
	}
	if v.SearchAttributes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SearchAttributes"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSearchUsersInput(v *SearchUsersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SearchUsersInput"}
	if v.DirectoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DirectoryId"))
	}
	if v.SearchString == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SearchString"))
	}
	if v.SearchAttributes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SearchAttributes"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateGroupInput(v *UpdateGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateGroupInput"}
	if v.DirectoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DirectoryId"))
	}
	if v.SAMAccountName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SAMAccountName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateUserInput(v *UpdateUserInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateUserInput"}
	if v.DirectoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DirectoryId"))
	}
	if v.SAMAccountName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SAMAccountName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
