/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/cluster-api-operator/internal/controller/genericprovider"
	"sigs.k8s.io/cluster-api-operator/internal/patch"
	ctrl "sigs.k8s.io/controller-runtime"
)

func applyPatches(ctx context.Context, provider genericprovider.GenericProvider) func(objs []unstructured.Unstructured) ([]unstructured.Unstructured, error) {
	log := ctrl.LoggerFrom(ctx)

	return func(objs []unstructured.Unstructured) ([]unstructured.Unstructured, error) {
		if len(provider.GetSpec().ManifestPatches) == 0 {
			log.V(5).Info("No resource patches to apply")
			return objs, nil
		}

		log.V(5).Info("Applying resource patches")

		return patch.ApplyPatches(objs, provider.GetSpec().ManifestPatches)
	}
}
