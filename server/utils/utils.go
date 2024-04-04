package utils

import (
	"github.com/Xi-Yuer/cms/utils/modules/bcrypt"
	"github.com/Xi-Yuer/cms/utils/modules/captcha"
	"github.com/Xi-Yuer/cms/utils/modules/jwt"
	"github.com/Xi-Yuer/cms/utils/modules/logs"
	"github.com/Xi-Yuer/cms/utils/modules/response"
	"github.com/Xi-Yuer/cms/utils/modules/snowflake"
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
