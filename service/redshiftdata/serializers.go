// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftdata

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/redshiftdata/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"path"
)

type awsAwsjson11_serializeOpBatchExecuteStatement struct {
}

func (*awsAwsjson11_serializeOpBatchExecuteStatement) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpBatchExecuteStatement) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*BatchExecuteStatementInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("RedshiftData.BatchExecuteStatement")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentBatchExecuteStatementInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpCancelStatement struct {
}

func (*awsAwsjson11_serializeOpCancelStatement) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpCancelStatement) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CancelStatementInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("RedshiftData.CancelStatement")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentCancelStatementInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDescribeStatement struct {
}

func (*awsAwsjson11_serializeOpDescribeStatement) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDescribeStatement) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeStatementInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("RedshiftData.DescribeStatement")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentDescribeStatementInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDescribeTable struct {
}

func (*awsAwsjson11_serializeOpDescribeTable) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDescribeTable) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeTableInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("RedshiftData.DescribeTable")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentDescribeTableInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpExecuteStatement struct {
}

func (*awsAwsjson11_serializeOpExecuteStatement) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpExecuteStatement) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ExecuteStatementInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("RedshiftData.ExecuteStatement")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentExecuteStatementInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpGetStatementResult struct {
}

func (*awsAwsjson11_serializeOpGetStatementResult) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpGetStatementResult) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetStatementResultInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("RedshiftData.GetStatementResult")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentGetStatementResultInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpListDatabases struct {
}

func (*awsAwsjson11_serializeOpListDatabases) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpListDatabases) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListDatabasesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("RedshiftData.ListDatabases")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentListDatabasesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpListSchemas struct {
}

func (*awsAwsjson11_serializeOpListSchemas) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpListSchemas) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListSchemasInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("RedshiftData.ListSchemas")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentListSchemasInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpListStatements struct {
}

func (*awsAwsjson11_serializeOpListStatements) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpListStatements) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListStatementsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("RedshiftData.ListStatements")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentListStatementsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpListTables struct {
}

func (*awsAwsjson11_serializeOpListTables) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpListTables) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListTablesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("RedshiftData.ListTables")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentListTablesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsAwsjson11_serializeDocumentSqlList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsjson11_serializeDocumentSqlParameter(v *types.SqlParameter, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Name != nil {
		ok := object.Key("name")
		ok.String(*v.Name)
	}

	if v.Value != nil {
		ok := object.Key("value")
		ok.String(*v.Value)
	}

	return nil
}

func awsAwsjson11_serializeDocumentSqlParametersList(v []types.SqlParameter, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson11_serializeDocumentSqlParameter(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeOpDocumentBatchExecuteStatementInput(v *BatchExecuteStatementInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ClientToken != nil {
		ok := object.Key("ClientToken")
		ok.String(*v.ClientToken)
	}

	if v.ClusterIdentifier != nil {
		ok := object.Key("ClusterIdentifier")
		ok.String(*v.ClusterIdentifier)
	}

	if v.Database != nil {
		ok := object.Key("Database")
		ok.String(*v.Database)
	}

	if v.DbUser != nil {
		ok := object.Key("DbUser")
		ok.String(*v.DbUser)
	}

	if v.SecretArn != nil {
		ok := object.Key("SecretArn")
		ok.String(*v.SecretArn)
	}

	if v.SessionId != nil {
		ok := object.Key("SessionId")
		ok.String(*v.SessionId)
	}

	if v.SessionKeepAliveSeconds != nil {
		ok := object.Key("SessionKeepAliveSeconds")
		ok.Integer(*v.SessionKeepAliveSeconds)
	}

	if v.Sqls != nil {
		ok := object.Key("Sqls")
		if err := awsAwsjson11_serializeDocumentSqlList(v.Sqls, ok); err != nil {
			return err
		}
	}

	if v.StatementName != nil {
		ok := object.Key("StatementName")
		ok.String(*v.StatementName)
	}

	if v.WithEvent != nil {
		ok := object.Key("WithEvent")
		ok.Boolean(*v.WithEvent)
	}

	if v.WorkgroupName != nil {
		ok := object.Key("WorkgroupName")
		ok.String(*v.WorkgroupName)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentCancelStatementInput(v *CancelStatementInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Id != nil {
		ok := object.Key("Id")
		ok.String(*v.Id)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentDescribeStatementInput(v *DescribeStatementInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Id != nil {
		ok := object.Key("Id")
		ok.String(*v.Id)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentDescribeTableInput(v *DescribeTableInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ClusterIdentifier != nil {
		ok := object.Key("ClusterIdentifier")
		ok.String(*v.ClusterIdentifier)
	}

	if v.ConnectedDatabase != nil {
		ok := object.Key("ConnectedDatabase")
		ok.String(*v.ConnectedDatabase)
	}

	if v.Database != nil {
		ok := object.Key("Database")
		ok.String(*v.Database)
	}

	if v.DbUser != nil {
		ok := object.Key("DbUser")
		ok.String(*v.DbUser)
	}

	if v.MaxResults != 0 {
		ok := object.Key("MaxResults")
		ok.Integer(v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	if v.Schema != nil {
		ok := object.Key("Schema")
		ok.String(*v.Schema)
	}

	if v.SecretArn != nil {
		ok := object.Key("SecretArn")
		ok.String(*v.SecretArn)
	}

	if v.Table != nil {
		ok := object.Key("Table")
		ok.String(*v.Table)
	}

	if v.WorkgroupName != nil {
		ok := object.Key("WorkgroupName")
		ok.String(*v.WorkgroupName)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentExecuteStatementInput(v *ExecuteStatementInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ClientToken != nil {
		ok := object.Key("ClientToken")
		ok.String(*v.ClientToken)
	}

	if v.ClusterIdentifier != nil {
		ok := object.Key("ClusterIdentifier")
		ok.String(*v.ClusterIdentifier)
	}

	if v.Database != nil {
		ok := object.Key("Database")
		ok.String(*v.Database)
	}

	if v.DbUser != nil {
		ok := object.Key("DbUser")
		ok.String(*v.DbUser)
	}

	if v.Parameters != nil {
		ok := object.Key("Parameters")
		if err := awsAwsjson11_serializeDocumentSqlParametersList(v.Parameters, ok); err != nil {
			return err
		}
	}

	if v.SecretArn != nil {
		ok := object.Key("SecretArn")
		ok.String(*v.SecretArn)
	}

	if v.SessionId != nil {
		ok := object.Key("SessionId")
		ok.String(*v.SessionId)
	}

	if v.SessionKeepAliveSeconds != nil {
		ok := object.Key("SessionKeepAliveSeconds")
		ok.Integer(*v.SessionKeepAliveSeconds)
	}

	if v.Sql != nil {
		ok := object.Key("Sql")
		ok.String(*v.Sql)
	}

	if v.StatementName != nil {
		ok := object.Key("StatementName")
		ok.String(*v.StatementName)
	}

	if v.WithEvent != nil {
		ok := object.Key("WithEvent")
		ok.Boolean(*v.WithEvent)
	}

	if v.WorkgroupName != nil {
		ok := object.Key("WorkgroupName")
		ok.String(*v.WorkgroupName)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentGetStatementResultInput(v *GetStatementResultInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Id != nil {
		ok := object.Key("Id")
		ok.String(*v.Id)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentListDatabasesInput(v *ListDatabasesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ClusterIdentifier != nil {
		ok := object.Key("ClusterIdentifier")
		ok.String(*v.ClusterIdentifier)
	}

	if v.Database != nil {
		ok := object.Key("Database")
		ok.String(*v.Database)
	}

	if v.DbUser != nil {
		ok := object.Key("DbUser")
		ok.String(*v.DbUser)
	}

	if v.MaxResults != 0 {
		ok := object.Key("MaxResults")
		ok.Integer(v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	if v.SecretArn != nil {
		ok := object.Key("SecretArn")
		ok.String(*v.SecretArn)
	}

	if v.WorkgroupName != nil {
		ok := object.Key("WorkgroupName")
		ok.String(*v.WorkgroupName)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentListSchemasInput(v *ListSchemasInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ClusterIdentifier != nil {
		ok := object.Key("ClusterIdentifier")
		ok.String(*v.ClusterIdentifier)
	}

	if v.ConnectedDatabase != nil {
		ok := object.Key("ConnectedDatabase")
		ok.String(*v.ConnectedDatabase)
	}

	if v.Database != nil {
		ok := object.Key("Database")
		ok.String(*v.Database)
	}

	if v.DbUser != nil {
		ok := object.Key("DbUser")
		ok.String(*v.DbUser)
	}

	if v.MaxResults != 0 {
		ok := object.Key("MaxResults")
		ok.Integer(v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	if v.SchemaPattern != nil {
		ok := object.Key("SchemaPattern")
		ok.String(*v.SchemaPattern)
	}

	if v.SecretArn != nil {
		ok := object.Key("SecretArn")
		ok.String(*v.SecretArn)
	}

	if v.WorkgroupName != nil {
		ok := object.Key("WorkgroupName")
		ok.String(*v.WorkgroupName)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentListStatementsInput(v *ListStatementsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxResults != 0 {
		ok := object.Key("MaxResults")
		ok.Integer(v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	if v.RoleLevel != nil {
		ok := object.Key("RoleLevel")
		ok.Boolean(*v.RoleLevel)
	}

	if v.StatementName != nil {
		ok := object.Key("StatementName")
		ok.String(*v.StatementName)
	}

	if len(v.Status) > 0 {
		ok := object.Key("Status")
		ok.String(string(v.Status))
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentListTablesInput(v *ListTablesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ClusterIdentifier != nil {
		ok := object.Key("ClusterIdentifier")
		ok.String(*v.ClusterIdentifier)
	}

	if v.ConnectedDatabase != nil {
		ok := object.Key("ConnectedDatabase")
		ok.String(*v.ConnectedDatabase)
	}

	if v.Database != nil {
		ok := object.Key("Database")
		ok.String(*v.Database)
	}

	if v.DbUser != nil {
		ok := object.Key("DbUser")
		ok.String(*v.DbUser)
	}

	if v.MaxResults != 0 {
		ok := object.Key("MaxResults")
		ok.Integer(v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	if v.SchemaPattern != nil {
		ok := object.Key("SchemaPattern")
		ok.String(*v.SchemaPattern)
	}

	if v.SecretArn != nil {
		ok := object.Key("SecretArn")
		ok.String(*v.SecretArn)
	}

	if v.TablePattern != nil {
		ok := object.Key("TablePattern")
		ok.String(*v.TablePattern)
	}

	if v.WorkgroupName != nil {
		ok := object.Key("WorkgroupName")
		ok.String(*v.WorkgroupName)
	}

	return nil
}
