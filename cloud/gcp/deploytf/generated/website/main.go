// website
package website

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"website.Website",
		reflect.TypeOf((*Website)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addProvider", GoMethod: "AddProvider"},
			_jsii_.MemberProperty{JsiiProperty: "basePath", GoGetter: "BasePath"},
			_jsii_.MemberProperty{JsiiProperty: "bucketNameOutput", GoGetter: "BucketNameOutput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "errorDocument", GoGetter: "ErrorDocument"},
			_jsii_.MemberProperty{JsiiProperty: "errorDocumentOutput", GoGetter: "ErrorDocumentOutput"},
			_jsii_.MemberProperty{JsiiProperty: "fileMd5SOutput", GoGetter: "FileMd5SOutput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberProperty{JsiiProperty: "indexDocument", GoGetter: "IndexDocument"},
			_jsii_.MemberProperty{JsiiProperty: "indexDocumentOutput", GoGetter: "IndexDocumentOutput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForOutput", GoMethod: "InterpolationForOutput"},
			_jsii_.MemberProperty{JsiiProperty: "localDirectory", GoGetter: "LocalDirectory"},
			_jsii_.MemberProperty{JsiiProperty: "localDirectoryOutput", GoGetter: "LocalDirectoryOutput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "providers", GoGetter: "Providers"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "skipAssetCreationFromLocalModules", GoGetter: "SkipAssetCreationFromLocalModules"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "stackId", GoGetter: "StackId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
			_jsii_.MemberProperty{JsiiProperty: "websiteName", GoGetter: "WebsiteName"},
		},
		func() interface{} {
			j := jsiiProxy_Website{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformModule)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"website.WebsiteConfig",
		reflect.TypeOf((*WebsiteConfig)(nil)).Elem(),
	)
}
