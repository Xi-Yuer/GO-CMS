package utils

import (
	"github.com/Xi-Yuer/cms/utils/modules/bcrypt"
	"github.com/Xi-Yuer/cms/utils/modules/buildTree"
	"github.com/Xi-Yuer/cms/utils/modules/buildZip"
	"github.com/Xi-Yuer/cms/utils/modules/captcha"
	"github.com/Xi-Yuer/cms/utils/modules/contain"
	"github.com/Xi-Yuer/cms/utils/modules/exportExcel"
	"github.com/Xi-Yuer/cms/utils/modules/file"
	"github.com/Xi-Yuer/cms/utils/modules/jwt"
	"github.com/Xi-Yuer/cms/utils/modules/logs"
	"github.com/Xi-Yuer/cms/utils/modules/response"
	"github.com/Xi-Yuer/cms/utils/modules/snowflake"
	"github.com/Xi-Yuer/cms/utils/modules/timeTask"
	"github.com/Xi-Yuer/cms/utils/modules/translator"
	"github.com/Xi-Yuer/cms/utils/modules/unique"
)

var Log = logsModules.Log
var Response = responseModules.Response
var GenID = snowflake.GenID
var Bcrypt = &bcrypt.Bcrypt{}
var Translator = translator.ValidatorTrans
var Trans = translator.Trans
var Captcha = &captcha.Captcha{}
var Jsonwebtoken = &jwt.Jsonwebtoken{}
var Unique = unique.RemoveDuplicatesAndEmpty
var BuildPages = buildTree.BuildMenu
var BuildDepartment = buildTree.BuildDepartment
var Contain = contain.StringInSlice
var ExportExcel = exportExcel.ExportExcel
var TimeTask = timeTask.TimeTask
var File = file.File
var FormatCommits = buildTree.FormatCommits
var GroupCommitsByDate = buildTree.GroupCommitsByDate
var CreateFilesAndZip = buildZip.CreateFilesAndZip
