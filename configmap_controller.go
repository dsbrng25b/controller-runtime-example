package main

import (
	"context"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type configMapReconciler struct {
	Client client.Client
	Log    logr.Logger
}

func (r *configMapReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()

	log := r.Log.WithValues("task", req.NamespacedName)

	log.Info("reconcile now", "name", req.NamespacedName)

	return ctrl.Result{}, nil
}
