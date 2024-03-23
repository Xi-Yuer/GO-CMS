package utils

import (
	"github.com/Xi-Yuer/cms/utils/modules/bcrypt"
	"github.com/Xi-Yuer/cms/utils/modules/logs"
	"github.com/Xi-Yuer/cms/utils/modules/response"
	"github.com/Xi-Yuer/cms/utils/modules/snowflake"
	"github.com/Xi-Yuer/cms/utils/modules/translator"
)

var Log = logsModules.Log
var Response = responseModules.Response
var GenID = snowflake.GenID
var Bcrypt = &bcrypt.Bcrypt{}
var Translator = translator.ValidatorTrans
var Trans = translator.Trans
