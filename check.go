package goi18n

import (
	"os"
	"strings"
)

// 默认的语言编码
var defaultCode Code = EN_US

// 语言MAP
var existMap = map[Code]struct{}{
	ZH_CN: {},
	ZH_TW: {},
	ZH_HK: {},
	EN_HK: {},
	EN_US: {},
	EN_GB: {},
	EN_WW: {},
	EN_CA: {},
	EN_AU: {},
	EN_IE: {},
	EN_FI: {},
	FI_FI: {},
	EN_DK: {},
	DA_DK: {},
	EN_IL: {},
	HE_IL: {},
	EN_ZA: {},
	EN_IN: {},
	EN_NO: {},
	EN_SG: {},
	EN_NZ: {},
	EN_ID: {},
	EN_PH: {},
	EN_TH: {},
	EN_MY: {},
	EN_XA: {},
	KO_KR: {},
	JA_JP: {},
	NL_NL: {},
	NL_BE: {},
	PT_PT: {},
	PT_BR: {},
	FR_FR: {},
	FR_LU: {},
	FR_CH: {},
	FR_BE: {},
	FR_CA: {},
	ES_LA: {},
	ES_ES: {},
	ES_AR: {},
	ES_US: {},
	ES_MX: {},
	ES_CO: {},
	ES_PR: {},
	DE_DE: {},
	DE_AT: {},
	DE_CH: {},
	RU_RU: {},
	IT_IT: {},
	EL_GR: {},
	NO_NO: {},
	HU_HU: {},
	TR_TR: {},
	CS_CZ: {},
	SL_SL: {},
	PL_PL: {},
	SV_SE: {},
}

// SetDefault 设置默认的语言编码
func SetDefault(code Code) {
	defaultCode = code
}

// GetEnv 获取环境变量当前配置的语言环境，未配置环境变量时将返回默认语言编码
func GetEnv() Code {
	// 检查语言是否存在
	code := Code(strings.ToLower(os.Getenv(Env)))
	if _, ok := existMap[code]; ok {
		return code
	}
	// 返回默认的语言
	return defaultCode
}
