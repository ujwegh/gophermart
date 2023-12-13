// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package clients

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson72d98f8bDecodeGithubComUjweghGophermartInternalAppServiceClients(in *jlexer.Lexer, out *AccrualResponseDto) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "order":
			out.OrderID = string(in.String())
		case "status":
			out.AccrualStatus = AccrualStatus(in.String())
		case "accrual":
			out.Accrual = float64(in.Float64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson72d98f8bEncodeGithubComUjweghGophermartInternalAppServiceClients(out *jwriter.Writer, in AccrualResponseDto) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"order\":"
		out.RawString(prefix[1:])
		out.String(string(in.OrderID))
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.AccrualStatus))
	}
	{
		const prefix string = ",\"accrual\":"
		out.RawString(prefix)
		out.Float64(float64(in.Accrual))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AccrualResponseDto) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson72d98f8bEncodeGithubComUjweghGophermartInternalAppServiceClients(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AccrualResponseDto) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson72d98f8bEncodeGithubComUjweghGophermartInternalAppServiceClients(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AccrualResponseDto) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson72d98f8bDecodeGithubComUjweghGophermartInternalAppServiceClients(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AccrualResponseDto) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson72d98f8bDecodeGithubComUjweghGophermartInternalAppServiceClients(l, v)
}
