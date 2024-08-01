package networking

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resource "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/networking/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func createIngressClasses(t *testing.T, ctrl *gomock.Controller) client.Services {
	r := resource.IngressClass{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	resourceClient := resourcemock.NewMockIngressClassInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.IngressClassList{Items: []resource.IngressClass{r}}, nil,
	)

	serviceClient := resourcemock.NewMockNetworkingV1Interface(ctrl)

	serviceClient.EXPECT().IngressClasses().Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().NetworkingV1().Return(serviceClient)

	return client.Services{CoreAPI: cl}
}

func TestIngressClasses(t *testing.T) {
	client.MockTestHelper(t, IngressClasses(), createIngressClasses)
}
