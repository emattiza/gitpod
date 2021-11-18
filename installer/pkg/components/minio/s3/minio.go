// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package s3

import (
	"fmt"
	"github.com/gitpod-io/gitpod/installer/pkg/common"
	"github.com/gitpod-io/gitpod/installer/pkg/helm"
	"github.com/gitpod-io/gitpod/installer/third_party/charts"
	"helm.sh/helm/v3/pkg/cli/values"
	"strings"
)

var Helm = func(apiPort int32, consolePort int32) common.HelmFunc {
	return common.CompositeHelmFunc(
		helm.ImportTemplate(charts.Minio(), helm.TemplateConfig{}, func(cfg *common.RenderContext) (*common.HelmConfig, error) {
			return &common.HelmConfig{
				Enabled: true,
				Values: &values.Options{
					Values: []string{
						//helm.KeyValue("minio.auth.rootUser", cfg.Values.StorageAccessKey),
						//helm.KeyValue("minio.auth.rootPassword", cfg.Values.StorageSecretKey),
						helm.KeyValue("minio.auth.existingSecret", cfg.Config.ObjectStorage.S3.Credentials.Name),
						helm.KeyValue("minio.gateway.enabled", "true"),
						helm.KeyValue("minio.gateway.replicaCount", "1"),
						//helm.KeyValue("minio.gateway.replicaCount", "2"),
						helm.KeyValue("minio.gateway.auth.s3.accessKey", "not-used"), // Used to get past Helm validation - provided by the auth.existingSecret
						helm.KeyValue("minio.gateway.auth.s3.secretKey", "not-user"), // Ditto
						helm.KeyValue("minio.gateway.auth.s3.serviceEndpoint", strings.TrimSuffix(cfg.Config.ObjectStorage.S3.Endpoint, "/")),
						helm.KeyValue("minio.gateway.type", "s3"),
						helm.KeyValue("minio.persistence.enabled", "false"),
						helm.KeyValue("minio.service.ports.api", fmt.Sprintf("%d", apiPort)),
						helm.KeyValue("minio.service.ports.console", fmt.Sprintf("%d", consolePort)),
					},
				},
			}, nil
		}),
	)
}
