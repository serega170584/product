package server

import (
	"github.com/stretchr/testify/require"
	"product/internal/proto"
	servererr "product/internal/server/error"
	"testing"
)

const (
	ErrIncorrectToValue string = "Value '%s' of field 'to' is not correct\n %v"
)

func TestValidate(t *testing.T) {
	req := &proto.EmailRequest{To: []string{"asdadasd"}, BodyType: proto.EmailRequest_TEXT}
	err := validate(req)
	require.NoError(t, err)
}

func TestValidateError(t *testing.T) {
	cases := []struct {
		name string
		in   *proto.EmailRequest
		err  error
	}{
		{
			name: "bad_to",
			in:   &proto.EmailRequest{To: nil, BodyType: proto.EmailRequest_HTML},
			err:  servererr.NewToError(nil, nil),
		},
		{
			name: "bad_body_type",
			in:   &proto.EmailRequest{To: []string{"asdasdadad"}, BodyType: proto.EmailRequest__UNSPECIFIED},
			err:  servererr.NewBodyTypeError(proto.EmailRequest__UNSPECIFIED, nil),
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {})
		err := validate(tCase.in)
		require.Error(t, err)
		require.EqualError(t, err, tCase.err.Error())
	}
}
