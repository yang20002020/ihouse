package controller

import (
	"bj38web_084/web/model"
	"bj38web_084/web/proto/getCaptcha"
	userMico "bj38web_084/web/proto/user"
	"bj38web_084/web/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	fdfs_client "github.com/tedcy/fdfs_client"
	"image/png"
	"math/rand"
	"net/http"
	"path"
	"time"
)

// 根据gin框架和文档业务要求写函数内容
// 获取session信息
func GetSession(ctx *gin.Context) {
	// 初始化错误返回的map
	//resp := make(map[string]string)
	//resp["errno"] = utils.RECODE_SESSIONERR
	//resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
	fmt.Println("*************GetSession 函数************")
	resp := make(map[string]interface{})
	//获取session数据
	s := sessions.Default(ctx)    //session初始化
	userName := s.Get("userName") //获取手机号
	//用户没有登录 ----- 没有存在mysql中，也没有存在session 中
	if userName == nil {
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
		fmt.Println("用户没有登录****")
	} else {
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
		/*{"data":{"name":"15889317897"},"errmsg":"成功","errno":"0"}*/
		var nameData struct {
			Name string `json:"name"`
		}
		nameData.Name = userName.(string) //类型断言
		resp["data"] = nameData
		fmt.Println("用户已经登录****")
	}

	//根军gin框架要求 按照json格式进行发送数据
	ctx.JSON(http.StatusOK, resp)
}

// 获取图片信息
func GetImageCd(ctx *gin.Context) {
	fmt.Println("*********1111111111*******")
	//获取图片验证码 uuid
	uuid := ctx.Param("uuid")
	fmt.Println("uuid=", uuid)
	/*************************/
	consulReg := consul.NewRegistry()
	consulService := micro.NewService(
		micro.Registry(consulReg),
	)
	//调用 函数 .初始化客户端
	microClient := getCaptcha.NewGetCaptchaService("go.micro.srv.getCaptcha", consulService.Client())
	//调用远程函数
	resp, err := microClient.Call(context.TODO(), &getCaptcha.Request{Uuid: uuid})
	if err != nil {
		fmt.Println("未能找到远程服务:", err)
		return
	}
	//将得到的数据 反序列化，得到图片数据
	var img captcha.Image
	json.Unmarshal(resp.Msg, &img)
	//对数据进行编码  将图片显示到浏览器
	png.Encode(ctx.Writer, img)
}

// 获取短信验证码
func GetSmscd(ctx *gin.Context) {

	//获取短信验证码
	phone := ctx.Param("phone")
	//http://192.168.63.128:8080/api/v1.0/smscode/15889317897?text=enhe&id=56b89c6b-d62d-45c2-b84d-e9d80d8f187c
	//拆分GET请求中的URL ===格式： 资源路径？/k=v & k=v & k =v
	//参考gin框架 中文文档
	imgCode := ctx.Query("text")
	uuid := ctx.Query("id")
	fmt.Println("---OUT----:", phone, imgCode, uuid)
	/*************需要转移到微服务的代码****************/
	//resp := make(map[string]string)
	////校验图片验证码 是否正确
	//result := model.CheckImgCode(uuid, imgCode)
	//if result {
	//	//校验成功
	//	fmt.Println("校验成功！！！")
	//	//发送短信
	//	err := _main(tea.StringSlice(os.Args[1:]), phone, resp)
	//	if err != nil {
	//		panic(err)
	//	}
	//} else {
	//	//校验失败 显示错误信息
	//	resp["errno"] = utils.RECODE_DATAERR
	//	resp["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR)
	//	fmt.Println("校验失败！！！")
	//}
	/******************************************/
	/************微服务添加内容*********/
	//指定consul服务发现
	consulReg := consul.NewRegistry()
	consulService := micro.NewService(
		micro.Registry(consulReg),
	)
	//1.初始化客户端
	microClient := userMico.NewUserService("go.micro.srv.user", consulService.Client())

	//调用远程函数
	resp, err := microClient.SendSms(context.TODO(), &userMico.Request{Phone: phone, ImgCode: imgCode, Uuid: uuid})
	if err != nil {
		fmt.Println("调用远程函数SendSms失败:", err)
		return
	}
	/*********************/
	//发送成功或者失败结果 给浏览器
	ctx.JSON(http.StatusOK, resp)

}

func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func _main(args []*string, phone string, resp map[string]string) (_err error) {
	client, _err := CreateClient(tea.String("LTAI5tEC1QprYRACFU528tJs"), tea.String("SoUZrelvm8fWy385YdbPegJLXX1fgP"))
	if _err != nil {
		return _err
	}
	//随机生成一个验证码
	rand.Seed(time.Now().UnixNano())
	//生成六位随机数
	str := rand.Int31n(1000000)
	smsCode := fmt.Sprintf("%06d", str)
	subStr := "{\"code\"" + ":" + "\"" + smsCode + "\"" + "}"
	fmt.Println(subStr)
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String("阿里云短信测试"),
		TemplateCode:  tea.String("SMS_154950909"),
		PhoneNumbers:  tea.String(phone), //15889317897
		TemplateParam: tea.String(subStr),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()

		// 复制代码运行请自行打印 API 的返回值
		//20221007 发送短信验证码

		_, _err = client.SendSmsWithOptions(sendSmsRequest, runtime)
		if _err != nil {
			//20221007  发送验证码失败
			fmt.Println("发送验证码失败！！！")
			resp["errno"] = utils.RECODE_SMSERR
			resp["errmsg"] = utils.RecodeText(utils.RECODE_SMSERR)
			return _err
		} else {
			//20221007  发送验证码成功
			resp["errno"] = utils.RECODE_OK
			resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
			fmt.Println("发送验证码成功！！！")
			//将短信验证码，存入到redis 数据库
			err := model.SaveSmsCode(phone, smsCode)
			if err != nil {
				fmt.Println("存储短信验证码到redis失败：", err)
				resp["errno"] = utils.RECODE_DBERR
				resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
			}
		}

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}
	}
	return _err
}

// 发送注册信息
/*{mobile: "15889317897", password: "111", sms_code: "94851"}
mobile: "15889317897"
password: "111"
sms_code: "94851" */
func PostRet(ctx *gin.Context) {

	//获取数据
	var regData struct {
		Mobile   string "json:\"mobile\""
		PassWord string "json:\"password\""
		SmsCod   string "json:\"sms_code\""
	}
	ctx.Bind(&regData)
	fmt.Println("mobile=", regData.Mobile, "pwd=", regData.PassWord, "sms_code=", regData.SmsCod)

	//初始化consul
	microService := utils.InitMicro()
	microClient := userMico.NewUserService("go.micro.srv.user", microService.Client())
	//调用远程函数
	resp, err := microClient.Register(context.TODO(), &userMico.RegReq{
		Mobile:   regData.Mobile,
		SmsCode:  regData.SmsCod,
		Password: regData.PassWord,
	})
	if err != nil {
		fmt.Println("注册用户，找不到远程服务！", err)
		return
	}
	//写给浏览器
	ctx.JSON(http.StatusOK, resp)
}

// 获取地址信息的信息
func GetArea(ctx *gin.Context) {
	var areas []model.Area
	//从缓存中获取数据
	conn := model.RedisPool.Get()
	//areaData字节切片
	areaData, _ := redis.Bytes(conn.Do("get", "areaData"))
	if len(areaData) == 0 {
		//没有从redis中获取数据

		fmt.Println("从mysql中获取数据 地址信息")
		//先从mysql中获取数据
		model.GlobalDB.Find(&areas)
		//把数据写入到redis时，存储结构体序列化后json串
		//字节切片areaBuf
		areaBuf, _ := json.Marshal(areas)

		conn := model.RedisPool.Get() //获取链接
		//再把数据写入到reidis中， 以字节流的方式
		conn.Do("set", "areaData", areaBuf)
	} else {
		fmt.Println("从redis中获取数据 地址信息")
		//redis中有数据，取出数据
		json.Unmarshal(areaData, &areas)
	}
	resp := make(map[string]interface{})
	resp["errno"] = "0"
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	resp["data"] = areas
	ctx.JSON(http.StatusOK, resp)
	fmt.Println("areas:", areas)
	fmt.Println("************************111111111")
}

// 处理登录业务
func PostLogin(ctx *gin.Context) {
	fmt.Println("***********postLogin***********")
	//获取前端数据
	var loginData struct {
		Mobile   string `json:"mobile"`
		PassWord string `json:"password"`
	}
	//???
	ctx.Bind(&loginData)
	resp := make(map[string]interface{})
	//获取数据库数据 查询是否和数据库的数据匹配
	userName, err := model.Login(loginData.Mobile, loginData.PassWord)
	if err == nil {
		//登录成功
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
		//1.初始化容器，在main函数中操作
		//将登录状态 保存到Session中
		s := sessions.Default(ctx)  //初始化session
		s.Set("userName", userName) //将用户名设置到session中
		s.Save()
		fmt.Println("登录成功，session存储了userName")
	} else {
		//登录失败
		resp["errno"] = utils.RECODE_LOGINERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_LOGINERR)
		fmt.Println("登录失败###")
	}
	ctx.JSON(http.StatusOK, resp)
}

// 退出登录
func DeleteSession(ctx *gin.Context) {
	resp := make(map[string]interface{})
	//初始化 session对象
	s := sessions.Default(ctx)
	//删除session数据 没有返回值
	s.Delete("userName")
	err := s.Save() // 有返回值
	if err != nil {
		resp["errno"] = utils.RECODE_IOERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_IOERR)
	} else {
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	}
	fmt.Println("退出登录***")
	ctx.JSON(http.StatusOK, resp)
}

// 获取用户基本信息
func GetUserInfo(ctx *gin.Context) {
	fmt.Println("获取用户基本信息GetUserInfo")
	resp := make(map[string]interface{})

	// 获取session ，得到当前用户信息
	s := sessions.Default(ctx)
	userName := s.Get("userName")
	if userName == nil {
		//用户没有登录，但是登录该页面，恶意进入
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)

		fmt.Println("用户没有登录，但是登录该页面，恶意进入")
		ctx.JSON(http.StatusOK, resp)
		return
	}
	/**********以下部分内容已经转移到服务端***************/
	//根据用户名获取用户信息 ---查mysql数据库
	//user, err := model.GetUserInfo(userName.(string))
	//if err != nil {
	//	resp["errno"] = utils.RECODE_DBERR
	//	resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
	//	return
	//}
	//resp["errno"] = utils.RECODE_OK
	//resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	//{"data":{"avatar_url":"http://192.168.63.128:8080/group1/M00/00/00/wKg_gGOAZq6Abq7mAADe1y314UQ299.png",
	//         "id_card":"13120151224",
	//         "mobile":"15889317897",
	//         "name":"15889317897",
	//         "real_name":"yangzhang",
	//         "user_id":28},
	//  "errmsg":"成功","errno":"0"}
	//参考 浏览器 user 中的回应response，获得对应的字段 或者参考文档
	//temp := make(map[string]interface{})
	//temp["user_id"] = user.ID
	//temp["name"] = user.Name
	//temp["mobile"] = user.Mobile
	//temp["real_name"] = user.Real_name
	//temp["id_card"] = user.Id_card
	//temp["avatar_url"] = "http://192.168.63.128:8080/" + user.Avatar_url
	//
	//resp["data"] = temp
	/********************************************/
	/*********微服务功能***************/
	//处理数据 微服务 consul  micro
	microClient := userMico.NewUserService("go.micro.srv.user", utils.InitMicro().Client())

	//调用远程服务
	resp1, _ := microClient.GetUsrInfo(context.TODO(), &userMico.UserName{UserName: userName.(string)})

	defer ctx.JSON(http.StatusOK, resp1)
}

// 更新用户名
func PutUerInfo(ctx *gin.Context) {
	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)
	//获取当前用户名
	s := sessions.Default(ctx)
	userName := s.Get("userName")
	//获取新用户名 ---处理request payload 类型数据 bind
	var nameData struct {
		Name string `json:"name"`
	}
	ctx.Bind(&nameData) //此时 新名字已经存储在namedata中
	//更新用户名
	//1.更新mysql数据库中的name
	err := model.UpDateUserName(nameData.Name, userName.(string))
	if err != nil {
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		return
	}
	//2.更新session
	s.Set("userName", nameData.Name)
	err = s.Save()
	if err != nil {
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
		return
	}
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	//参考文档
	resp["data"] = nameData

}

// 上传头像
func PostAvatar(ctx *gin.Context) {
	//单文件 获取图片文件，静态文件对象
	//通过key值avatar 返回value为文件  file为文件内容
	file, _ := ctx.FormFile("avatar")
	fmt.Println("flie:", file)
	//上传文件到指定路径
	//err := ctx.SaveUploadedFile(file, "./test/"+file.Filename)
	//fmt.Println(err)

	/*******************/

	//上传头像到fastDFS中

	//初始化客户端
	clt, _ := fdfs_client.NewClientWithConfig("/etc/fdfs/client.conf")

	//打开文件，读取文件内容，只读打开
	f, _ := file.Open()            //f 为文件指针
	buf := make([]byte, file.Size) //按照文件实际内容，创建切片
	fmt.Println("file.Size:", file.Size)
	f.Read(buf) //读取文件内容，保存在缓存区
	//go语言根据文件名获取文件后缀
	fileExt := path.Ext(file.Filename)
	fmt.Println("fileExt:", fileExt)

	//按字节流 上传图片内容  上传头像到fastDFS中
	remoteId, _ := clt.UploadByBuffer(buf, fileExt[1:]) //返回凭证 remoteId

	//获取session，得到当前用户
	userName := sessions.Default(ctx).Get("userName")

	//根据用户名去更新用户头像 ---mysql数据库
	model.UpdateAvatar(userName.(string), remoteId)

	fmt.Println("返回凭证:", remoteId)
	//remoteId= group1/M00/00/00/wKg_gGN97NCANIbLAAMSKK5ztqs113.jpg
	//avatar_url 是mysql 数据库中user表的一个字段
	//根据文档要求进行拼接
	resp := make(map[string]interface{})
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

	temp := make(map[string]interface{})
	temp["avatar_url"] = "http://192.168.63.128:8080/" + remoteId
	resp["data"] = temp

	ctx.JSON(http.StatusOK, resp)
}

/**************************************/
type AuthStu struct {
	IdCard   string `json:"id_card"`
	RealName string `json:"real_name"`
}

// 实名认证用户信息
func PostUserInfo(ctx *gin.Context) {
	//获取数据
	var auth AuthStu
	err := ctx.Bind(&auth)
	//校验数据
	if err != nil {
		fmt.Println("获取数据错误", err)
		return
	}

	session := sessions.Default(ctx)
	userName := session.Get("userName")

	//处理数据 微服务 consul  micro
	microClient := userMico.NewUserService("go.micro.srv.user", utils.InitMicro().Client())

	//调用远程服务  auth 为上述定义的结构体变量
	resp, _ := microClient.AuthUpdate(context.TODO(), &userMico.AuthReq{
		UserName: userName.(string),
		RealName: auth.RealName,
		IdCard:   auth.IdCard,
	})

	//返回数据
	ctx.JSON(http.StatusOK, resp)

}
