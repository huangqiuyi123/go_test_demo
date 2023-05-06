package main

import (
	"log"

	"testDemo/common"
)

func main() {

	var request = common.ReadJson("test")

	header := map[string]string{
		"cookie": "_BEAMER_USER_ID_whBStDUm30458=b99b65e5-1186-4c59-933d-f8e9e5f6f984; _BEAMER_FIRST_VISIT_whBStDUm30458=2023-02-01T06:40:35.843Z; _BEAMER_FIRST_VISIT_undefined=2023-02-06T06:09:10.546Z; n_u=24c37fcbcf4d401d2541fbdc24873b94; r_b_ined=1; mafrc=8XYE699L9E17%2Csales_plugin_af; a_ste=m7tH55QQBALTFezrpyp0YE0xb+hZjcgYSHtmcFcx5Iq01UjUVhSQaMCYXlAAsFoNb8aFMObih5h5l/EAByfhmg==; store_id=1665628907335; merchant_id=4211387386; lang.sig=HPZEXM6qRQA3fl9QF0Gl5KM_KZ7FwUtDpVV9UEUrrek; addressLang.sig=fZhLaUxh_564Gt_Ygb8agf56cVb1lYYp6NMpk7wfgaM; userSelectLocale.sig=xaWhkiDLccJKOWtBx98z0KVVx7o_iP0WoEYPBrEqJCw; s_id=432C37B3A7ECD1BD874A64811EFEFE63; s_id.sig=3afec40af11c2c5d2917020221fa9f02; osudb_lang=en; osudb_hdid=2217590e27c9544604c3575811dae79c; f_ds_info.sig=3CKJGA_m714rZP1Sd6rVfiPZrE-e85dCA1u-vNls85c; store_id.sig=jkiGcSihCNQHcIwuJeDT90gbTVd1S9PYIlgxgvQOJtM; merchant_id.sig=JdG9OdjYaypCshQ26ZRHpUNyoOuPv3UW7kt9Ba4A7VU; tsao=c79f003d-f1bd-4e08-b0a1-fe5c1198b515#ad50675722ed913e; tsui=03eb85fa-4176-4e1e-be02-463406306734#ca17543de9e9389f; f_ds_info=vIkVV50Dkzv2zmQ98TKTgtk3b1cl9Ma+OJEOEXxG5CSnNgeHzfIbqulQrZ39Pa3cYEZKj5kOO+OxU5eNBLt8HQ==; d_cstgc=1163336839_2_4201796168_1683167613661QBP0JE; osudb_param=; d_osudb_appid=1163336839; osudb_ustate=1; d_osudb_subappid=2; d_osudb_uid=4201796168; d_osudb_oar=#01#SID0000073BJT6aUjH60NhP34deK+FN2j9slMLBzhCU3puj89p4h6VoyFF+JHbxyJPtsSLq0/+UTogRc+vaMAO8gHTp/OMeA7FA60s2G9OFP0sefEedrcKKfcTy5JNXxiB2ty4W/6iBiYrbQb+pVe+k7ron69UUA==; d_osudb_third=#T01#SID0000073BEodgmLVBAigi0VmlmjjrcuxaTwBIOQte6VsnZThAjj5bHqO2xxD37E1yniKkHGxgeWveLDUTqCylrk17UZWe9IOBnpodVl5uAmJTRE/n/yWuKrQEjniiFhn2FMCqj+QJFNUTRmgiLN59GoQXaJnyP63gNlyMp7iCQ==; lp_url={%22landingPageHtml%22:%22https://niki.myshoplinestg.com/%22%2C%22occurredAt%22:1683168211423}; currency_code=USD; currency_code.sig=nEGddW1-E-8oJfI_Pm_5XNzC2sMi1n3aVzZ3v01csyY; localization=US; currency_code_userSetting=USD; currency_code_userSetting.sig=wreMdGqvcOcZfYXi-Fd1QDxl5OWoQm3s2QLyXkCpvxE; osudb_appid=1163336839; osudb_uid=4213749186; osudb_oar=#01#SID0000073BGE8lnusLx5LGLZaBwIgEnu5o6e1g/zVjyrSNfbNRmxNOy8TkrvS2IsFrRGGIQtRHih59wMlvR1TwwAU/eq0/dPe5j/NIclGmeEbeiRY4l6ADW5IYiRpszzuVf+wjuHptFyMm+CbcV+RRHg5feiVPzShHiTRUfXTEjwNHg; osudb_subappid=1665628907335; lang=en; addressLang=en; userSelectLocale=en; JSESSIONID=E3218C08629BFEE350DD3068E74048FA; n_sess={\"session_id\":\"23583ba1-6481-4c04-b6ab-ad4ad4a714ec\",\"created_at\":1683366560337,\"last_session_id\":\"9f307569-b794-4c48-9c31-1b0511d2a0d4\",\"session_create_type\":102}; history_browse_products=16058402017068989883880637,16056560679292150805570637,16058400049415287022080637,16058402029173583456560637,16057240158879062455150637,16058402539549580405810637,16057240159038278226880637,16057240160812636592860637,16057240158524224331930637,16059098471082035866340637; history_browse_products.sig=PNF0hGNKWd-zeLXdNdj8oJ3B5Im4xDs1gJBj0dYSFDQ; client_abandoned_order_seq=2405923015635437618374%233c14101d538785a5; t_cart=7adb064ae660432383cadf8175b210b9; t_cart.sig=57e1a8ef26a9f39790d79616b463bfe5; a_osudb_appid=1163336839; a_cstgc=1163336839_1_4211387386_1683369460462nhPWIb; a_osudb_uid=4211387386; a_osudb_oar=#01#SID0000073BL3dkxsKqbvYOkdIBisV6ItHqMY5Og9tZ+FycZVElqCmW0KG2e/VdtNPNIZdA3EiWih+S6QbyvsZXCk+1PNYm5cRfRUbbyHpRjiEiZltVwMu3Jv/W+0/wkTr9d4o74xRmnA/cRDenijf/dG5U3AUyA==; a_osudb_subappid=1; a_dhch=niki; a_lang=zh-hant-tw; _BEAMER_FILTER_BY_URL_whBStDUm30458=false",
	}

	// 循环发起请求json文件的数据
	for i := 0; i < len(request); i++ {
		url := request[i].Url
		resBody := request[i].ResBody
		name := request[i].Name
		log.Println(name)
		common.Post(url, resBody, nil, header)
	}
}
