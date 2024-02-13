package occurrences_test

import (
	"context"
	"testing"

	"buf.build/gen/go/matheusslima/go-poc/grpc/go/occurrences/occurrencesgrpc"
	occurrenceGrpc "buf.build/gen/go/matheusslima/go-poc/protocolbuffers/go/occurrences"
	grpcServer "github.com/mateusfdl/go-poc/internal/grpc"
	"github.com/mateusfdl/go-poc/internal/logger"
	"github.com/mateusfdl/go-poc/internal/occurrences"
	grpcInternal "github.com/mateusfdl/go-poc/internal/occurrences/grpc"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

var (
	TestApp = fx.New(
		occurrences.Module,
		grpcServer.Module,
		logger.Module,
	)

	clientConn       *grpc.ClientConn
	OccurrenceClient occurrencesgrpc.OccurrenceServiceClient
)

func TestOrderGRPCHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Order GRPC handler suite")
}

var _ = BeforeSuite(func() {
	Expect(TestApp.Start(context.Background())).To(Succeed())

	//var dialErr error
	//clientConn, dialErr = grpc.Dial(
	//	"",
	//	grpc.WithTransportCredentials(insecure.NewCredentials()),
	//)

	//Expect(dialErr).NotTo(HaveOccurred())

	//OccurrenceClient = occurrencesgrpc.NewOccurrenceServiceClient(clientConn)
})

var _ = AfterSuite(func() {
	//Expect(clientConn.Close()).To(Succeed())
	Expect(TestApp.Stop(context.Background())).To(Succeed())
})

var _ = Describe("grpc occurrence handler", func() {
	Describe("CreateOccurrence", func() {
		It("sends a create event to the creation producer channel", func() {
			resp, err := OccurrenceClient.CreateOccurrence(
				context.Background(),
				&occurrenceGrpc.NewOccurrence{},
			)
			Expect(err).To(Equal(nil))

			Expect(resp.OccurrenceId).To(Equal("1"))
		})
	})
})
