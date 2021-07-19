package verify

import (
	utils "gozero-vue-admin/common/utils"
)

var (
	IdVerify               = utils.Rules{"ID": {utils.NotEmpty()}}
	ApiVerify              = utils.Rules{"Path": {utils.NotEmpty()}, "Description": {utils.NotEmpty()}, "ApiGroup": {utils.NotEmpty()}, "Method": {utils.NotEmpty()}}
	MenuVerify             = utils.Rules{"Path": {utils.NotEmpty()}, "ParentId": {utils.NotEmpty()}, "Name": {utils.NotEmpty()}, "Component": {utils.NotEmpty()}, "Sort": {utils.Ge("0")}}
	MenuMetaVerify         = utils.Rules{"Title": {utils.NotEmpty()}}
	LoginVerify            = utils.Rules{"CaptchaId": {utils.NotEmpty()}, "Captcha": {utils.NotEmpty()}, "Username": {utils.NotEmpty()}, "Password": {utils.NotEmpty()}}
	RegisterVerify         = utils.Rules{"Username": {utils.NotEmpty()}, "NickName": {utils.NotEmpty()}, "Password": {utils.NotEmpty()}, "AuthorityId": {utils.NotEmpty()}}
	PageInfoVerify         = utils.Rules{"Page": {utils.NotEmpty()}, "PageSize": {utils.NotEmpty()}}
	//CustomerVerify         = utils.Rules{"CustomerName": {utils.NotEmpty()}, "CustomerPhoneData": {utils.NotEmpty()}}
	//AutoCodeVerify         = utils.Rules{"Abbreviation": {utils.NotEmpty()}, "StructName": {utils.NotEmpty()}, "PackageName": {utils.NotEmpty()}, "Fields": {utils.NotEmpty()}}
	AuthorityVerify        = utils.Rules{"AuthorityId": {utils.NotEmpty()}, "AuthorityName": {utils.NotEmpty()}, "ParentId": {utils.NotEmpty()}}
	AuthorityIdVerify      = utils.Rules{"AuthorityId": {utils.NotEmpty()}}
	OldAuthorityVerify     = utils.Rules{"OldAuthorityId": {utils.NotEmpty()}}
	ChangePasswordVerify   = utils.Rules{"Username": {utils.NotEmpty()}, "Password": {utils.NotEmpty()}, "NewPassword": {utils.NotEmpty()}}
	SetUserAuthorityVerify = utils.Rules{"UUID": {utils.NotEmpty()}, "AuthorityId": {utils.NotEmpty()}}
)
