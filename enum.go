package goi18n

// Code 语言编码
type Code string

// 配置语言的环境变量
const (
	Env1 string = "GO_I18N_CODE"
	Env2 string = "GO_LANGUAGE_CODE"
)

const (
	ZH_CN Code = "zh-cn" // 简体中文(中国)
	ZH_TW Code = "zh-tw" // 繁体中文(台湾地区)
	ZH_HK Code = "zh-hk" // 繁体中文(香港)
	EN_HK Code = "en-hk" // 英语(香港)
	EN_US Code = "en-us" // 英语(美国)
	EN_GB Code = "en-gb" // 英语(英国)
	EN_WW Code = "en-ww" // 英语(全球)
	EN_CA Code = "en-ca" // 英语(加拿大)
	EN_AU Code = "en-au" // 英语(澳大利亚)
	EN_IE Code = "en-ie" // 英语(爱尔兰)
	EN_FI Code = "en-fi" // 英语(芬兰)
	FI_FI Code = "fi-fi" // 芬兰语(芬兰)
	EN_DK Code = "en-dk" // 英语(丹麦)
	DA_DK Code = "da-dk" // 丹麦语(丹麦)
	EN_IL Code = "en-il" // 英语(以色列)
	HE_IL Code = "he-il" // 希伯来语(以色列)
	EN_ZA Code = "en-za" // 英语(南非)
	EN_IN Code = "en-in" // 英语(印度)
	EN_NO Code = "en-no" // 英语(挪威)
	EN_SG Code = "en-sg" // 英语(新加坡)
	EN_NZ Code = "en-nz" // 英语(新西兰)
	EN_ID Code = "en-id" // 英语(印度尼西亚)
	EN_PH Code = "en-ph" // 英语(菲律宾)
	EN_TH Code = "en-th" // 英语(泰国)
	EN_MY Code = "en-my" // 英语(马来西亚)
	EN_XA Code = "en-xa" // 英语(阿拉伯)
	KO_KR Code = "ko-kr" // 韩文(韩国)
	JA_JP Code = "ja-jp" // 日语(日本)
	NL_NL Code = "nl-nl" // 荷兰语(荷兰)
	NL_BE Code = "nl-be" // 荷兰语(比利时)
	PT_PT Code = "pt-pt" // 葡萄牙语(葡萄牙)
	PT_BR Code = "pt-br" // 葡萄牙语(巴西)
	FR_FR Code = "fr-fr" // 法语(法国)
	FR_LU Code = "fr-lu" // 法语(卢森堡)
	FR_CH Code = "fr-ch" // 法语(瑞士)
	FR_BE Code = "fr-be" // 法语(比利时)
	FR_CA Code = "fr-ca" // 法语(加拿大)
	ES_LA Code = "es-la" // 西班牙语(拉丁美洲)
	ES_ES Code = "es-es" // 西班牙语(西班牙)
	ES_AR Code = "es-ar" // 西班牙语(阿根廷)
	ES_US Code = "es-us" // 西班牙语(美国)
	ES_MX Code = "es-mx" // 西班牙语(墨西哥)
	ES_CO Code = "es-co" // 西班牙语(哥伦比亚)
	ES_PR Code = "es-pr" // 西班牙语(波多黎各)
	DE_DE Code = "de-de" // 德语(德国)
	DE_AT Code = "de-at" // 德语(奥地利)
	DE_CH Code = "de-ch" // 德语(瑞士)
	RU_RU Code = "ru-ru" // 俄语(俄罗斯)
	IT_IT Code = "it-it" // 意大利语(意大利)
	EL_GR Code = "el-gr" // 希腊语(希腊)
	NO_NO Code = "no-no" // 挪威语(挪威)
	HU_HU Code = "hu-hu" // 匈牙利语(匈牙利)
	TR_TR Code = "tr-tr" // 土耳其语(土耳其)
	CS_CZ Code = "cs-cz" // 捷克语(捷克共和国)
	SL_SL Code = "sl-sl" // 斯洛文尼亚语
	PL_PL Code = "pl-pl" // 波兰语(波兰)
	SV_SE Code = "sv-se" // 瑞典语(瑞典)
)

func (c Code) String() string {
	return string(c)
}
