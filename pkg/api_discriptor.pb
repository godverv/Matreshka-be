
�
grpc/matreshka-be_api.protomatreshka_be_apigoogle/api/annotations.proto"7
AppInfo
name (	Rname
version (	Rversion";
ListRequest
limit (Rlimit
offset (Roffset"=

ApiVersion	
Request$
Response
version (	Rversion"]
	GetConfig,
Request!
service_name (	RserviceName"
Response
config (Rconfig"`

PostConfigF
Request
content (Rcontent!
service_name (	RserviceName

Response"�
PatchConfig�
Request!
service_name (	RserviceNameZ
path_to_value (26.matreshka_be_api.PatchConfig.Request.PathToValueEntryRpathToValue>
PathToValueEntry
key (	Rkey
value (	Rvalue:8

Response"�
ListConfigsr
Request@
list_request (2.matreshka_be_api.ListRequestRlistRequest%
search_pattern (	RsearchPatternA
Response5
services (2.matreshka_be_api.AppInfoRservices2�
MatreshkaBeAPIk

ApiVersion$.matreshka_be_api.ApiVersion.Request%.matreshka_be_api.ApiVersion.Response"���
/versionv
	GetConfig#.matreshka_be_api.GetConfig.Request$.matreshka_be_api.GetConfig.Response"���/config/{service_name}u
ListConfigs%.matreshka_be_api.ListConfigs.Request&.matreshka_be_api.ListConfigs.Response"���"/config/list:*|

PostConfig$.matreshka_be_api.PostConfig.Request%.matreshka_be_api.PostConfig.Response"!���"/config/{service_name}:*�
PatchConfig%.matreshka_be_api.PatchConfig.Request&.matreshka_be_api.PatchConfig.Response"'���!"/config/patch/{service_name}:*BZ/matreshka_apibproto3