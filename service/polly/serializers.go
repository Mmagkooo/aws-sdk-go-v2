// Code generated by smithy-go-codegen DO NOT EDIT.

package polly

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/polly/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpDeleteLexicon struct {
}

func (*awsRestjson1_serializeOpDeleteLexicon) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDeleteLexicon) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteLexiconInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/v1/lexicons/{Name}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "DELETE"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsDeleteLexiconInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsDeleteLexiconInput(v *DeleteLexiconInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.Name == nil || len(*v.Name) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member Name must not be empty")}
	}
	if v.Name != nil {
		if err := encoder.SetURI("Name").String(*v.Name); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpDescribeVoices struct {
}

func (*awsRestjson1_serializeOpDescribeVoices) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDescribeVoices) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeVoicesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/v1/voices")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsDescribeVoicesInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsDescribeVoicesInput(v *DescribeVoicesInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if len(v.Engine) > 0 {
		encoder.SetQuery("Engine").String(string(v.Engine))
	}

	if v.IncludeAdditionalLanguageCodes {
		encoder.SetQuery("IncludeAdditionalLanguageCodes").Boolean(v.IncludeAdditionalLanguageCodes)
	}

	if len(v.LanguageCode) > 0 {
		encoder.SetQuery("LanguageCode").String(string(v.LanguageCode))
	}

	if v.NextToken != nil {
		encoder.SetQuery("NextToken").String(*v.NextToken)
	}

	return nil
}

type awsRestjson1_serializeOpGetLexicon struct {
}

func (*awsRestjson1_serializeOpGetLexicon) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetLexicon) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetLexiconInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/v1/lexicons/{Name}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsGetLexiconInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetLexiconInput(v *GetLexiconInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.Name == nil || len(*v.Name) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member Name must not be empty")}
	}
	if v.Name != nil {
		if err := encoder.SetURI("Name").String(*v.Name); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpGetSpeechSynthesisTask struct {
}

func (*awsRestjson1_serializeOpGetSpeechSynthesisTask) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetSpeechSynthesisTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetSpeechSynthesisTaskInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/v1/synthesisTasks/{TaskId}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsGetSpeechSynthesisTaskInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetSpeechSynthesisTaskInput(v *GetSpeechSynthesisTaskInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.TaskId == nil || len(*v.TaskId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member TaskId must not be empty")}
	}
	if v.TaskId != nil {
		if err := encoder.SetURI("TaskId").String(*v.TaskId); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpListLexicons struct {
}

func (*awsRestjson1_serializeOpListLexicons) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListLexicons) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListLexiconsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/v1/lexicons")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListLexiconsInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListLexiconsInput(v *ListLexiconsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.NextToken != nil {
		encoder.SetQuery("NextToken").String(*v.NextToken)
	}

	return nil
}

type awsRestjson1_serializeOpListSpeechSynthesisTasks struct {
}

func (*awsRestjson1_serializeOpListSpeechSynthesisTasks) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListSpeechSynthesisTasks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListSpeechSynthesisTasksInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/v1/synthesisTasks")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListSpeechSynthesisTasksInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListSpeechSynthesisTasksInput(v *ListSpeechSynthesisTasksInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.MaxResults != nil {
		encoder.SetQuery("MaxResults").Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		encoder.SetQuery("NextToken").String(*v.NextToken)
	}

	if len(v.Status) > 0 {
		encoder.SetQuery("Status").String(string(v.Status))
	}

	return nil
}

type awsRestjson1_serializeOpPutLexicon struct {
}

func (*awsRestjson1_serializeOpPutLexicon) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpPutLexicon) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*PutLexiconInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/v1/lexicons/{Name}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "PUT"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsPutLexiconInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentPutLexiconInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsPutLexiconInput(v *PutLexiconInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.Name == nil || len(*v.Name) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member Name must not be empty")}
	}
	if v.Name != nil {
		if err := encoder.SetURI("Name").String(*v.Name); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentPutLexiconInput(v *PutLexiconInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Content != nil {
		ok := object.Key("Content")
		ok.String(*v.Content)
	}

	return nil
}

type awsRestjson1_serializeOpStartSpeechSynthesisTask struct {
}

func (*awsRestjson1_serializeOpStartSpeechSynthesisTask) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpStartSpeechSynthesisTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*StartSpeechSynthesisTaskInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/v1/synthesisTasks")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentStartSpeechSynthesisTaskInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsStartSpeechSynthesisTaskInput(v *StartSpeechSynthesisTaskInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentStartSpeechSynthesisTaskInput(v *StartSpeechSynthesisTaskInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.Engine) > 0 {
		ok := object.Key("Engine")
		ok.String(string(v.Engine))
	}

	if len(v.LanguageCode) > 0 {
		ok := object.Key("LanguageCode")
		ok.String(string(v.LanguageCode))
	}

	if v.LexiconNames != nil {
		ok := object.Key("LexiconNames")
		if err := awsRestjson1_serializeDocumentLexiconNameList(v.LexiconNames, ok); err != nil {
			return err
		}
	}

	if len(v.OutputFormat) > 0 {
		ok := object.Key("OutputFormat")
		ok.String(string(v.OutputFormat))
	}

	if v.OutputS3BucketName != nil {
		ok := object.Key("OutputS3BucketName")
		ok.String(*v.OutputS3BucketName)
	}

	if v.OutputS3KeyPrefix != nil {
		ok := object.Key("OutputS3KeyPrefix")
		ok.String(*v.OutputS3KeyPrefix)
	}

	if v.SampleRate != nil {
		ok := object.Key("SampleRate")
		ok.String(*v.SampleRate)
	}

	if v.SnsTopicArn != nil {
		ok := object.Key("SnsTopicArn")
		ok.String(*v.SnsTopicArn)
	}

	if v.SpeechMarkTypes != nil {
		ok := object.Key("SpeechMarkTypes")
		if err := awsRestjson1_serializeDocumentSpeechMarkTypeList(v.SpeechMarkTypes, ok); err != nil {
			return err
		}
	}

	if v.Text != nil {
		ok := object.Key("Text")
		ok.String(*v.Text)
	}

	if len(v.TextType) > 0 {
		ok := object.Key("TextType")
		ok.String(string(v.TextType))
	}

	if len(v.VoiceId) > 0 {
		ok := object.Key("VoiceId")
		ok.String(string(v.VoiceId))
	}

	return nil
}

type awsRestjson1_serializeOpSynthesizeSpeech struct {
}

func (*awsRestjson1_serializeOpSynthesizeSpeech) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpSynthesizeSpeech) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*SynthesizeSpeechInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/v1/speech")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentSynthesizeSpeechInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsSynthesizeSpeechInput(v *SynthesizeSpeechInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentSynthesizeSpeechInput(v *SynthesizeSpeechInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.Engine) > 0 {
		ok := object.Key("Engine")
		ok.String(string(v.Engine))
	}

	if len(v.LanguageCode) > 0 {
		ok := object.Key("LanguageCode")
		ok.String(string(v.LanguageCode))
	}

	if v.LexiconNames != nil {
		ok := object.Key("LexiconNames")
		if err := awsRestjson1_serializeDocumentLexiconNameList(v.LexiconNames, ok); err != nil {
			return err
		}
	}

	if len(v.OutputFormat) > 0 {
		ok := object.Key("OutputFormat")
		ok.String(string(v.OutputFormat))
	}

	if v.SampleRate != nil {
		ok := object.Key("SampleRate")
		ok.String(*v.SampleRate)
	}

	if v.SpeechMarkTypes != nil {
		ok := object.Key("SpeechMarkTypes")
		if err := awsRestjson1_serializeDocumentSpeechMarkTypeList(v.SpeechMarkTypes, ok); err != nil {
			return err
		}
	}

	if v.Text != nil {
		ok := object.Key("Text")
		ok.String(*v.Text)
	}

	if len(v.TextType) > 0 {
		ok := object.Key("TextType")
		ok.String(string(v.TextType))
	}

	if len(v.VoiceId) > 0 {
		ok := object.Key("VoiceId")
		ok.String(string(v.VoiceId))
	}

	return nil
}

func awsRestjson1_serializeDocumentLexiconNameList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentSpeechMarkTypeList(v []types.SpeechMarkType, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(string(v[i]))
	}
	return nil
}
