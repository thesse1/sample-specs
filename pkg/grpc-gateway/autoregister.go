package gateway

import (
    "context"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	//sample.Samples
	samplepb "github.com/thesse1/sample-specs/dist/pb/sample"
    
	"google.golang.org/grpc"
	"net/http"
    _ "github.com/thesse1/sample-specs/dist/pb/sample"
    
)


// newGateway returns a new gateway server which translates HTTP into gRPC.
func NewGateway(ctx context.Context, conn *grpc.ClientConn, opts []gwruntime.ServeMuxOption) (http.Handler, error) {

	mux := gwruntime.NewServeMux(opts...)

	for _, f := range []func(context.Context, *gwruntime.ServeMux, *grpc.ClientConn) error{
	// sample.Samples
	samplepb.RegisterSamplesHandler,


//installed services
	} {
		if err := f(ctx, mux, conn); err != nil {
			return nil, err
		}
	}
	return mux, nil
}
