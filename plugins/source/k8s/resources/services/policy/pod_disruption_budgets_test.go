package policy

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	v1 "k8s.io/api/core/v1"
	policy "k8s.io/api/policy/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/policy/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func createPodDisruptionBudgets(t *testing.T, ctrl *gomock.Controller) client.Services {
	r := policy.PodDisruptionBudget{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}
	r.Spec.MinAvailable = &intstr.IntOrString{IntVal: 1}
	r.Spec.MaxUnavailable = &intstr.IntOrString{IntVal: 2}

	resourceClient := resourcemock.NewMockPodDisruptionBudgetInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&policy.PodDisruptionBudgetList{Items: []policy.PodDisruptionBudget{r}}, nil,
	)

	serviceClient := resourcemock.NewMockPolicyV1Interface(ctrl)

	serviceClient.EXPECT().PodDisruptionBudgets(metav1.NamespaceAll).Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().PolicyV1().Return(serviceClient)

	return client.Services{CoreAPI: cl}
}

func TestPodDisruptionBudgets(t *testing.T) {
	client.MockTestHelper(t, PodDisruptionBudgets(), createPodDisruptionBudgets, client.WithTestNamespaces(v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "default"}}))
}
