package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var ConfigInstance Config

var UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/123.0.0.0"

type Config struct {
	PushPlusToken  string `yaml:"pushplus_token"`
	RefreshToken   string `yaml:"refresh_token"`
	BilibiliCookie string `yaml:"bilibili_cookie"`
	KKCookie       string `yaml:"ctoken=0Sd9U-Ym8OgMYsZRu_ZgBO8o; b-user-id=74f18365-8771-0ca1-a97a-5c91f706d24e; grey-id=f6f137aa-1966-033b-0e31-b18981404c09; grey-id.sig=q-PZkcFTieoNSy_YFLjDLYP4ayrbphk1vGjMPd51DXI; isQuark=false; isQuark.sig=DWPHMZYiiwQ-v58AbcP-rBdSIpzO8ZnrD67BdJuPatU; _UP_A4A_11_=wb9651462c9e43ce8b134dcf5c9a9335; _UP_D_=pc; __wpkreporterwid_=2ca11b68-81bd-4707-ad4c-fa73562887e3; _UP_F7E_8D_=aH%2Bh86gI1q4k6I85B%2FdJM0O93CyQD4M6GUOU%2BWEa1%2FCQ9pt%2Flxcgspx%2F7jdZxm88%2BDZR%2F3OCavVVlSrlA%2FHE35guhUiWeFXSFrpiz7iQup7LB%2BL83Dn0lIDh36hnRcflW%2FQJYV4NNTsN3RU4rAGaIRY%2BWfsgqVGQLbPeoOdhK0F%2B4CVxwd%2BA5kyclW53hcuTOm0qJXeruCuqKrFkR8sCHG%2FU0THtz7bCaLWY8apJjPRO89zY%2FAxDPIInCxl1g1C0s1YxLzJs52lWdoIk19nvWfytP81cJKxfv1GfwBTrR0bS%2FBGt91fGhZMwyZvOKbU%2F8ZkbgrGvmwxqzivluzXqQHrmHUwax9fVlczEGdRq2nJkSi56LgyDrXpsyYdGYA0F4Cj%2Fntr80vkUTca5E0LqLcBEt4UN6jDY; __pus=f1ea019e0c6fd98c1ca86a715e826010AARfirymQx0PCsBV9cOKfddgqscRbVcuA6tbHUggQwqbR9ZNUX1eZQPx5DA186MmODWZxxxvRGcHOUjpQb9l75Cm; __kp=b372cca0-0891-11ef-8d3e-9905b8fe724c; __kps=AARnlTM3B86cWNR4vRH5c1jB; __ktd=lbg8wm110ETlFcqP7Hd/5g==; __uid=AARnlTM3B86cWNR4vRH5c1jB; __itrace_wid=9825b44b-1eaa-4865-27f4-baa3a5be48bb; __puus=a7e2a68b7dde84e4a7c5e75ebdd99369AARLllthd64rtRbBvmBKi+fj9JAygJpK66K94m8Duwd1pI5eUM4G34tbJbTLXO0+qe7pQg3gwmgPqCQtJZik/YUrrQsTcx0HHgchH1OCjIhM/2xK0ERywfQ3Q+pw/KtRvyWWxZuW8RmmQqEZxBDwuLviE0Yciso4ve6ejKUvv+42Wd1ZFJ/CgJc7Wh8+wULFqhCX1SzD52TLlWLSg6IKMLpa
"`
	JdCookie       string `yaml:"__jdv=76161171|direct|-|none|-|1714660315240; areaId=12; ipLoc-djd=12-951-0-0; PCSYCityID=CN_320000_321000_0; shshshfpa=86c94375-8e82-860c-d8e0-ae820219f202-1714660317; shshshfpx=86c94375-8e82-860c-d8e0-ae820219f202-1714660317; pinId=M_OefvGsK10gQ7dS4IVahA; pin=15071221129_p; unick=evil%E4%B8%B6%E7%8B%BC; _tp=smEhSP3a2SIMLxQbgw8hrA%3D%3D; _pst=15071221129_p; source=PC; platform=pc; lpkLoginType=3; wlfstk_smdl=k0um16pxw8mo8i3xbuizqdwiakvswhlh; TrackID=1f6dMh0ARnotAmT7mQ6q3dNb-NZCBUvvf6R7xAfWSgvz7rvZ_FIjcIkrQK2do07usP6vhJlgqpPNr7Bc0nHC-E4JcGobGAGs4f1EPaaALm6UGxadnix4vJj4-nUdcNMLt; thor=5D05B833D101B8A273D8FD3F90EBC09C205AEF33CC5B0275CBAA1B84900318D40A0D6B6899FB5F383AAC0BA97EEB62DA588598EF8B16A86F50E47AA5965C3D8BFF4063D0001CF280D2432A0FDB939251EEB09F5625F21643FC79E5534EFB56970E244B7B241FE66C1FA5C69D35416782EE4AA1897A7CB880A16629D26DC95E2902872CC46C3B42AAD28A45CD4B51F912; flash=2_Lhl4GKceAo1osji2c0XePqtn8371btSTaoMDJ9CD-Ht9WF985NvpQvWAmSDdOoc4_FVr0r9JKaI-VFAcmQXrq4ag5QSVbko6PTFRh78B2XHTulBVH3260OJXVydFTf7iBTZYtFd19U5osIzjVLaUbfRBBtlbH6XPESGGNKx8ycq*; ceshi3.com=201; __jda=222648329.1714660315240578456522.1714660315.1714660315.1714660315.1; __jdc=222648329; shshshfpb=BApXcC_KyOupAP-1lfMVouGopx_W530PYBlYnBDho9xJ1MmGYLoC2; __jdb=222648329.10.1714660315240578456522|1.1714660315; cn=0
"`
}

func init() {
	LoadConfig()
}

func LoadConfig() {
	confFIle, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err.Error())
	}
	config := Config{}
	yaml.Unmarshal(confFIle, &config)
	ConfigInstance = config
}
