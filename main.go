package main

import "fmt"

func Test_fDisConnect(lLoginID LLONG, pchDVRIP string, nDVRPort int, dwUser LLONG) {
	fmt.Println("Device disconnect callback.")
}

func main() {
	initParam := NETSDK_INIT_PARAM{}
	bRet := InitEx(Test_fDisConnect, &initParam)
	if false == bRet {
		fmt.Println("Init NetSDK failed")
		return
	}
	fmt.Println("Init NetSDK success")
	defer Cleanup()

	stNetParam := NET_PARAM{}
	stNetParam.ST_nWaittime = 8000
	SetNetworkParam(&stNetParam)

	var (
		ip             = ""
		port     int32 = 37777
		username       = ""
		passwd         = ""
	)

	fmt.Println("===>Please input device ip, port, username, password, e.g. 192.168.1.1 37777 abc abc123")
	fmt.Scanln(&ip, &port, &username, &passwd)

	stLoginIn := NET_IN_LOGIN_WITH_HIGHLEVEL_SECURITY{}
	copy(stLoginIn.ST_szIP[:], []byte(ip))
	stLoginIn.ST_nPort = port
	copy(stLoginIn.ST_szUserName[:], []byte(username))
	copy(stLoginIn.ST_szPassword[:], []byte(passwd))

	stLoginOut := NET_OUT_LOGIN_WITH_HIGHLEVEL_SECURITY{}
	lhandle := LoginWithHighLevelSecurity(&stLoginIn, &stLoginOut)
	if 0 == lhandle {
		fmt.Printf("LoginWithHighLevelSecurity failed, 0x%x\n", GetLastError())
		return
	}
	fmt.Println("LoginWithHighLevelSecurity success")
	defer Logout(lhandle)

	SerialNum := stLoginOut.ST_stDeviceInfo.ST_sSerialNumber[:]
	fmt.Println("Serial Num:", string(SerialNum))

}
