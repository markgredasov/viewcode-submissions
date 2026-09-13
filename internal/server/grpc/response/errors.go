package response

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetStatusCode(err error) codes.Code {
	statusCode := codes.OK
	if err != nil {
		if st, ok := status.FromError(err); ok {
			statusCode = st.Code()
		} else {
			statusCode = codes.Unknown
		}
	}
	return statusCode
}
