package main

import (
	"flag"
	"os"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var (
	// TODO
	scheme     = runtime.NewScheme()
	managerLog = ctrl.Log.WithName("manager")
)

func init() {
	// TODO
	_ = clientgoscheme.AddToScheme(scheme)
}

func main() {
	var someSetting string

	flag.StringVar(&someSetting, "some-setting", "foobar", "Some setting to set.")

	flag.Parse()

	// TODO
	ctrl.SetLogger(zap.Logger(true))

	// Setup manager
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: scheme,
	})

	if err != nil {
		managerLog.Error(err, "failed to setup manager")
		os.Exit(1)
	}

	// Register controller
	err = ctrl.NewControllerManagedBy(mgr).
		For(&v1.ConfigMap{}).
		Complete(&configMapReconciler{
			Client: mgr.GetClient(),
			Log:    managerLog,
		})

	if err != nil {
		managerLog.Error(err, "failed to create controller")
	}

	// Start manager
	managerLog.Info("strating manager")
	// TODO
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		managerLog.Error(err, "problem running manager")
		os.Exit(1)
	}

}
