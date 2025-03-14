// Copyright 2021 Nitric Technologies Pty Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package deploytf

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	sql "github.com/nitrictech/nitric/cloud/aws/deploytf/generated/sql"
	"github.com/nitrictech/nitric/cloud/common/deploy/image"
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
	"github.com/samber/lo"
)

func (n *NitricAwsTerraformProvider) SqlDatabase(stack cdktf.TerraformStack, name string, config *deploymentspb.SqlDatabase) error {
	sqlConfig := sql.SqlConfig{
		DbName:                    jsii.String(name),
		RdsClusterEndpoint:        n.Rds.ClusterEndpointOutput(),
		RdsClusterUsername:        n.Rds.ClusterUsernameOutput(),
		RdsClusterPassword:        n.Rds.ClusterPasswordOutput(),
		SubnetIds:                 cdktf.Token_AsList(n.Vpc.PrivateSubnetIdsOutput(), &cdktf.EncodingOptions{}),
		SecurityGroupIds:          &[]*string{n.Rds.SecurityGroupIdOutput()},
		CreateDatabaseProjectName: n.Rds.CreateDatabaseProjectNameOutput(),
		VpcId:                     n.Vpc.VpcIdOutput(),
		CodebuildRoleArn:          n.Rds.CodebuildRoleArnOutput(),
		CodebuildRegion:           &n.Region,
		DependsOn:                 &[]cdktf.ITerraformDependable{n.Rds},
	}

	// Inspect the provided migration image and get its command and working directory
	if config.GetImageUri() != "" {
		inspect, err := image.CommandFromImageInspect(config.GetImageUri(), " ")
		if err != nil {
			return err
		}

		sqlConfig.Migrations = map[string]interface{}{
			"image_uri":       config.GetImageUri(),
			"migrate_command": inspect.Cmd,
			"work_dir":        lo.Ternary(inspect.WorkDir != "", inspect.WorkDir, "/"),
		}
	}

	sql.NewSql(stack, jsii.String(name), &sqlConfig)

	return nil
}
