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
	KKCookie       string `yaml:"kk_cookie"`
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
