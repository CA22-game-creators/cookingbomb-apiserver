package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// InvalidArgument 400:BadRequest
func InvalidArgument(v string) error {
	return status.Error(codes.InvalidArgument, v)
}

// Unauthenticated 401:Unauthorized
func Unauthenticated(v string) error {
	return status.Error(codes.Unauthenticated, v)
}

// Internal 500:InternalError
func Internal(v string) error {
	return status.Error(codes.Internal, v)
}
