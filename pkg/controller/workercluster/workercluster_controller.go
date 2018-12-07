package workercluster

import (
	"context"

	travisciv1alpha1 "github.com/travis-ci/worker-operator/pkg/apis/travisci/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	//	deploymentutil "k8s.io/kubernetes/pkg/controller/deployment/util"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_workercluster")

// Add creates a new WorkerCluster Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileWorkerCluster{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("workercluster-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource WorkerCluster
	err = c.Watch(&source.Kind{Type: &travisciv1alpha1.WorkerCluster{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource Deployments and requeue the owner WorkerCluster
	err = c.Watch(&source.Kind{Type: &appsv1.Deployment{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &travisciv1alpha1.WorkerCluster{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileWorkerCluster{}

// ReconcileWorkerCluster reconciles a WorkerCluster object
type ReconcileWorkerCluster struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a WorkerCluster object and makes changes based on the state read
// and what is in the WorkerCluster.Spec
func (r *ReconcileWorkerCluster) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling WorkerCluster")

	// Fetch the WorkerCluster instance
	instance := &travisciv1alpha1.WorkerCluster{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	deployment := newDeploymentForCluster(instance)

	// Set WorkerCluster instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, deployment, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this Deployment already exists
	found := &appsv1.Deployment{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: deployment.Name, Namespace: deployment.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Deployment", "Deployment.Namespace", deployment.Namespace, "Deployment.Name", deployment.Name)
		err = r.client.Create(context.TODO(), deployment)
		if err != nil {
			return reconcile.Result{}, err
		}

		found = deployment
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// TODO figure out a way to update the correct properties of the deployment without causing a cycle

	//	if deploymentutil.EqualIgnoreHash(&deployment.Spec.Template, &found.Spec.Template) {
	//		// Deployment already exists - don't requeue
	//		reqLogger.Info("Skip reconcile: Deployment already exists", "Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name)
	//		return reconcile.Result{}, nil
	//	}
	//
	//	reqLogger.Info("Updating the existing Deployment", "Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name)
	//	err = r.client.Update(context.TODO(), deployment)
	//	if err != nil {
	//		return reconcile.Result{}, err
	//	}

	// List the pods for the deployment, and determine their pool sizes
	podList := &corev1.PodList{}
	labelSelector := labels.SelectorFromSet(found.Spec.Selector.MatchLabels)
	listOps := &client.ListOptions{Namespace: instance.Namespace, LabelSelector: labelSelector}
	if err = r.client.List(context.TODO(), listOps, podList); err != nil {
		return reconcile.Result{}, err
	}

	var statuses []travisciv1alpha1.WorkerStatus
	for _, pod := range podList.Items {
		statuses = append(statuses, travisciv1alpha1.WorkerStatus{
			Name: pod.Name,
			// TODO actually query the workers for this information
			CurrentPoolSize:   1,
			ExpectedPoolSize:  1,
			RequestedPoolSize: 1,
		})
	}

	instance.Status = travisciv1alpha1.WorkerClusterStatus{
		WorkerStatuses: statuses,
	}
	if err = r.client.Status().Update(context.TODO(), instance); err != nil {
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}

func newDeploymentForCluster(cluster *travisciv1alpha1.WorkerCluster) *appsv1.Deployment {
	maxUnavailable := intstr.FromInt(0)
	maxSurge := intstr.FromInt(1)

	template := cluster.Spec.Template.DeepCopy()
	configureContainer(&template.Spec.Containers[0])

	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cluster.Name,
			Namespace: cluster.Namespace,
			Labels:    cluster.Labels,
		},
		Spec: appsv1.DeploymentSpec{
			Selector: cluster.Spec.Selector,
			Template: *template,
			Strategy: appsv1.DeploymentStrategy{
				Type: appsv1.RollingUpdateDeploymentStrategyType,
				RollingUpdate: &appsv1.RollingUpdateDeployment{
					MaxUnavailable: &maxUnavailable,
					MaxSurge:       &maxSurge,
				},
			},
		},
	}
}

func configureContainer(c *corev1.Container) {
	newEnvVars := []corev1.EnvVar{
		{
			// The remote controller API is used to adjust pool sizes on the fly
			Name:  "TRAVIS_WORKER_REMOTE_CONTROLLER_ADDR",
			Value: "0.0.0.0:8080",
		},
		{
			Name: "TRAVIS_WORKER_REMOTE_CONTROLLER_AUTH",
			// TODO make this randomly assigned.
			// The operator needs to know what this is to talk to the worker, but it will have the pod definition,
			// so it could just read it from there when it needs to query the worker API
			Value: "worker:worker",
		},
		{
			// Don't start any processors when the worker starts.
			// Instead, let this operator use the API to assign a pool size.
			Name:  "TRAVIS_WORKER_POOL_SIZE",
			Value: "0",
		},
	}

	c.Env = append(c.Env, newEnvVars...)
}
