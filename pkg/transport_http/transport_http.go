package transport_http

import (
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	oldcontext "golang.org/x/net/context"
	"google.golang.org/grpc"
	"jaxf-github.fanatics.corp/apparel/partner-service/pkg/pb"
)

func MakeHTTPHandler(host string, dopts []grpc.DialOption, logger log.Logger) (http.Handler, error) {
	// mux for the reverse proxy
	gwmux := runtime.NewServeMux()

	// standard mux
	m := http.NewServeMux()

	ctx := oldcontext.Background()

	if err := pb.RegisterPartnerServiceHandlerFromEndpoint(ctx, gwmux, host, dopts); err != nil {
		return nil, errors.Wrap(err, "failed to register handler from endpoint")
	}

	// serve swagger at the /swagger endpoint
	m.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swagger"))))

	// otherwise redirect to reverse proxy
	m.Handle("/", gwmux)

	return m, nil
}
