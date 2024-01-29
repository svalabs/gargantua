package rbac

import (
	"context"
	"fmt"

	"github.com/golang/glog"
	hferrors "github.com/hobbyfarm/gargantua/v3/pkg/errors"
	"github.com/hobbyfarm/gargantua/v3/pkg/rbac"
	"github.com/hobbyfarm/gargantua/v3/pkg/util"
	rbacProto "github.com/hobbyfarm/gargantua/v3/protos/rbac"
	"google.golang.org/grpc/codes"
	empty "google.golang.org/protobuf/types/known/emptypb"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
)

func (rs *GrpcRbacServer) CreateRolebinding(c context.Context, cr *rbacProto.RoleBinding) (*empty.Empty, error) {
	rolebinding, err := rs.marshalRolebinding(c, cr)
	if err != nil {
		glog.Errorf("invalid rolebinding: %v", err)
		return &empty.Empty{}, hferrors.GrpcError(
			codes.InvalidArgument,
			"invalid rolebinding",
			cr,
		)
	}

	_, err = rs.kubeClientSet.RbacV1().RoleBindings(util.GetReleaseNamespace()).Create(c, rolebinding, metav1.CreateOptions{})
	if err != nil {
		glog.Errorf("error creating rolebinding in kubernetes: %v", err)
		return &empty.Empty{}, hferrors.GrpcError(
			codes.Internal,
			"error creating rolebinding",
			cr,
		)
	}

	return &empty.Empty{}, nil
}

func (rs *GrpcRbacServer) GetRolebinding(c context.Context, gr *rbacProto.ResourceId) (*rbacProto.RoleBinding, error) {
	rolebinding, err := rs.getRolebinding(c, gr)
	if err != nil {
		return nil, err
	}
	return unmarshalRolebinding(rolebinding), nil
}

func (rs *GrpcRbacServer) UpdateRolebinding(c context.Context, ur *rbacProto.RoleBinding) (*empty.Empty, error) {
	inputRolebinding, err := rs.marshalRolebinding(c, ur)
	if err != nil {
		glog.Errorf("invalid role: %v", err)
		return &empty.Empty{}, hferrors.GrpcError(
			codes.InvalidArgument,
			"invalid role",
			ur,
		)
	}

	k8sRolebinding, err := rs.getRolebinding(c, &rbacProto.ResourceId{Id: ur.GetName()})
	if err != nil {
		return &empty.Empty{}, err
	}

	k8sRolebinding.RoleRef = inputRolebinding.RoleRef
	k8sRolebinding.Subjects = inputRolebinding.Subjects

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		_, err := rs.kubeClientSet.RbacV1().RoleBindings(util.GetReleaseNamespace()).Update(c, k8sRolebinding, metav1.UpdateOptions{})
		if err != nil {
			glog.Errorf("error while updating rolebinding in kubernetes: %v", err)
			return hferrors.GrpcError(
				codes.Internal,
				"error while updating rolebinding",
				ur,
			)
		}
		return nil
	})

	if retryErr != nil {
		return &empty.Empty{}, retryErr
	}

	return &empty.Empty{}, nil
}

func (rs *GrpcRbacServer) DeleteRolebinding(c context.Context, dr *rbacProto.ResourceId) (*empty.Empty, error) {
	// we want to get the rolebinding first as this allows us to run it through the various checks before we attempt deletion
	// most important of which is checking that we have labeled it correctly
	// but it doesn't hurt to check if it exists before
	rolebinding, err := rs.getRolebinding(c, dr)
	if err != nil {
		return &empty.Empty{}, err
	}

	err = rs.kubeClientSet.RbacV1().RoleBindings(util.GetReleaseNamespace()).Delete(c, rolebinding.Name, metav1.DeleteOptions{})
	if err != nil {
		glog.Errorf("error deleting rolebinding in kubernetes: %v", err)
		return &empty.Empty{}, hferrors.GrpcError(
			codes.Internal,
			"error deleting rolebinding",
			dr,
		)
	}
	return &empty.Empty{}, nil
}

func (rs *GrpcRbacServer) ListRolebinding(c context.Context, lr *empty.Empty) (*rbacProto.RoleBindings, error) {
	listOptions := metav1.ListOptions{
		LabelSelector: fmt.Sprintf("%s=%t", util.RBACManagedLabel, true),
	}

	rolebindings, err := rs.kubeClientSet.RbacV1().RoleBindings(util.GetReleaseNamespace()).List(c, listOptions)
	if err != nil {
		if errors.IsNotFound(err) {
			glog.Errorf("error: rolebindings not found")
			return &rbacProto.RoleBindings{}, hferrors.GrpcError(
				codes.NotFound,
				"rolebindings not found",
				lr,
			)
		}
		glog.Errorf("error in kubernetes while listing rolebindings %v", err)
		return &rbacProto.RoleBindings{}, hferrors.GrpcError(
			codes.Internal,
			"error listing rolebindings",
			lr,
		)
	}

	var preparedRolebindings = make([]*rbacProto.RoleBinding, 0)
	for _, r := range rolebindings.Items {
		pr := unmarshalRolebinding(&r)
		preparedRolebindings = append(preparedRolebindings, pr)
	}

	return &rbacProto.RoleBindings{Rolebindings: preparedRolebindings}, nil
}

func (rs *GrpcRbacServer) marshalRolebinding(ctx context.Context, preparedRoleBinding *rbacProto.RoleBinding) (*rbacv1.RoleBinding, error) {
	// first validation, the role it is referencing has to exist
	role, err := rs.kubeClientSet.RbacV1().Roles(util.GetReleaseNamespace()).Get(ctx, preparedRoleBinding.GetRole(), metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("invalid role ref")
	}

	rb := rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      preparedRoleBinding.GetName(),
			Namespace: util.GetReleaseNamespace(),
			Labels: map[string]string{
				util.RBACManagedLabel: "true",
			},
			OwnerReferences: []metav1.OwnerReference{
				{
					APIVersion: "rbac.authorization.k8s.io/v1",
					Kind:       "Role",
					Name:       role.Name,
					UID:        role.UID,
				},
			},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: rbac.RbacGroup,
			Name:     preparedRoleBinding.Role,
			Kind:     "Role",
		},
		Subjects: []rbacv1.Subject{},
	}

	for _, s := range preparedRoleBinding.GetSubjects() {
		if s.GetKind() != "Group" && s.GetKind() != "User" {
			return nil, fmt.Errorf("invalid subject kind")
		}

		rb.Subjects = append(rb.Subjects, rbacv1.Subject{
			Kind:     s.GetKind(),
			Name:     s.GetName(),
			APIGroup: rbac.RbacGroup,
		})
	}

	return &rb, nil
}

func unmarshalRolebinding(roleBinding *rbacv1.RoleBinding) *rbacProto.RoleBinding {
	prb := &rbacProto.RoleBinding{
		Name:     roleBinding.Name,
		Role:     roleBinding.RoleRef.Name,
		Subjects: []*rbacProto.Subject{},
	}

	for _, s := range roleBinding.Subjects {
		prb.Subjects = append(prb.GetSubjects(), &rbacProto.Subject{
			Kind: s.Kind,
			Name: s.Name,
		})
	}

	return prb
}

func (rs *GrpcRbacServer) getRolebinding(c context.Context, gr *rbacProto.ResourceId) (*rbacv1.RoleBinding, error) {
	if gr.GetId() == "" {
		glog.Errorf("invalid rolebinding id")
		return &rbacv1.RoleBinding{}, hferrors.GrpcError(
			codes.InvalidArgument,
			"invalid rolebinding id",
			gr,
		)
	}

	rolebinding, err := rs.kubeClientSet.RbacV1().RoleBindings(util.GetReleaseNamespace()).Get(c, gr.GetId(), metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			glog.Errorf("rolebinding not found")
			return &rbacv1.RoleBinding{}, hferrors.GrpcError(
				codes.NotFound,
				"rolebinding not found",
				gr,
			)
		}
		glog.Errorf("kubernetes error while getting rolebinding: %v", err)
		return &rbacv1.RoleBinding{}, hferrors.GrpcError(
			codes.Internal,
			"error while retrieving rolebinding",
			gr,
		)
	}

	if _, ok := rolebinding.Labels[util.RBACManagedLabel]; !ok {
		// this isn't a hobbyfarm rolebinding. we don't serve your kind here
		glog.Error("permission denied: rolebinding not managed by hobbyfarm")
		return &rbacv1.RoleBinding{}, hferrors.GrpcError(
			codes.PermissionDenied,
			"rolebinding not managed by hobbyfarm",
			gr,
		)
	}

	return rolebinding, nil
}
