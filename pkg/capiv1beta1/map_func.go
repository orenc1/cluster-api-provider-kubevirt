package capiv1beta1

import (
	context "context"

	infrav1 "sigs.k8s.io/cluster-api-provider-kubevirt/api/v1alpha1"
	clusterv1 "sigs.k8s.io/cluster-api/api/core/v1beta2"
	"sigs.k8s.io/cluster-api/util"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func MapV1beta1MachineToKVMachine(ctx context.Context, object client.Object) []reconcile.Request {
	machine, ok := object.(*clusterv1.Machine)
	if !ok {
		return nil
	}
	mapFunc := util.MachineToInfrastructureMapFunc(infrav1.GroupVersion.WithKind("KubevirtMachine"))
	return mapFunc(ctx, machine)
}

func MapV1beta1ClusterToKVKind(cl client.Client, kind string) handler.MapFunc {
	return func(ctx context.Context, object client.Object) []reconcile.Request {
		cluster, ok := object.(*clusterv1.Cluster)
		if !ok {
			return nil
		}
		mapFunc := util.ClusterToInfrastructureMapFunc(ctx,
			infrav1.GroupVersion.WithKind(kind),
			cl,
			&infrav1.KubevirtCluster{},
		)
		return mapFunc(ctx, cluster)
	}
}
