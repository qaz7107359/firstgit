package sso

import (
	"TARG_revenue_report_backend/db/dao"
	"TARG_revenue_report_backend/db/model"
	"TARG_revenue_report_backend/global"
	"TARG_revenue_report_backend/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

type SSO struct {
	Access_token       string `json:"access_token"`
	Refresh_token      string `json:"refresh_token"`
	Expires_in         int    `json:"expires_in"`
	Refresh_expires_in int    `json:"refresh_expires_in"`
	Client_id          string `json:"client_id"`
	Scope              string `json:"scope"`
	Openid             string `json:"openid"`
}

type Data struct {
	Code int    `json:"code"`
	Data string `json:"data"`
	msg  string `json:"msg"`
}

type TokenUser struct {
	OpenId       string `json:"openId"`       // 用户唯一标识码
	UserName     string `json:"userName"`     //姓名
	EmailAddress string `json:"emailAddress"` //邮箱
	UserExt      string `json:"userExt"`      //分机
	UserCode     string `json:"userCode"`     // 工号
}

func SsoLogin(c *gin.Context) {

	//	获取code和state的值    获取access_token
	client := &http.Client{}
	var tokenUser TokenUser

	var sso SSO
	code := c.Query("code")
	state := c.Query("state")

	fmt.Println(code)
	fmt.Println(state)

	url := global.Http_access_token

	grant_type := global.Grant_type

	client_id := global.Client_id

	client_secret := global.Client_secret

	http_url := url + "?client_id=" + client_id + "&client_secret=" + client_secret + "&code=" + code + "&state=" + state + "&grant_type=" + grant_type

	value := map[string]interface{}{
		"grant_type":    grant_type,
		"client_id":     client_id,
		"client_secret": client_secret,
		"code":          code,
		"state":         state,
	}
	jsonvalue, err := json.Marshal(value)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	req, err := http.NewRequest("POST", http_url, strings.NewReader(string(jsonvalue)))
	if err != nil {
		c.JSON(500, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	fmt.Println(string(body))

	err = json.Unmarshal(body, &sso)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}

	// 获取用户身份

	req_user, err := http.NewRequest("POST", global.Http_user, nil)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}

	req_user.Header.Set("Authorization", sso.Access_token)

	resp_user, err := client.Do(req_user)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	defer resp_user.Body.Close()

	body_user, err := ioutil.ReadAll(resp_user.Body)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}

	fmt.Println(string(body_user))

	err = json.Unmarshal(body_user, &tokenUser)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	if tokenUser.UserName == "" {
		c.JSON(500, gin.H{
			"code":    1,
			"message": "无效的code",
		})
		return
	}

	// 判断用户是否存在
	fmt.Println(tokenUser.UserName)
	userDao := dao.NewUserBasicDao()
	if tokenUser.UserName != "" {
		_, exist, err := userDao.ExistOrNotByUserName(tokenUser.UserName)
		if err == nil && !exist {
			user := model.User{
				EmployeeId:   tokenUser.UserCode,
				Email:        tokenUser.EmailAddress,
				EmployeeName: tokenUser.UserName,
				UserExt:      tokenUser.UserExt,
				OpenId:       tokenUser.OpenId,
			}
			err = userDao.CreateUser(&user)
			if err != nil {
				c.JSON(500, gin.H{
					"code":    1,
					"message": err.Error(),
				})
				return
			}
		}
	}

	// 签发token http 无状态 (token) 签发
	userToken, _, err := userDao.ExistOrNotByUserName(tokenUser.UserName)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}

	token, err := utils.GenerateToken(userToken.ID, userToken.EmployeeName, userToken.Email)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"Data":    sso,
		"reqcode": utils.GenReqCode(),
		"Token":   token,
	})

}

// 一键登录测试系统

//client := &http.Client{}
//
//url_test := "https://test-portal.myfiinet.com/mis/api-uac/oauth2/doLogin"
//
//value_test := map[string]interface{}{
//	//"userCode": "C/GGqyxI7pysY5NjA1OSZRdR4Nk+ShQCVFtT7mInY1VblNvofGFxWiPuyvUi4esstSxDnzoDGsQbD8MRri6lag==",
//	//"userPwd":  "h4M9c0nN5xMG9OvhkEEj+5fllJ+SowPDMROqhVRKdHj67K27XId82DbduReV4GAa0GlDeXHddElioqJFqSGbPQ==",
//	"userCode": "OgPoqwADHaa+mOyPGBZrC/ONsLO4RKUdeXIk0dshpfN9N3zW1+4GY3HTuTh0GIB13MH5t1w92A88vXkn3OzA+Q==",
//	"userPwd":  "Lu24S5qWXRnpWrc5qfIxApus+CsRf1CkoaC5/J6qIyOBNDZpvRhiYx2FmBFylc595nMWDhuYQ7gNQBXMyi11WA==",
//}
//
//jsonvalue_test, err := json.Marshal(value_test)
//if err != nil {
//	c.JSON(500, gin.H{
//		"code":    1,
//		"message": err.Error(),
//	})
//}
//
//req_test, err := http.NewRequest("POST", url_test, strings.NewReader(string(jsonvalue_test)))
//if err != nil {
//	c.JSON(500, gin.H{
//		"code":    1,
//		"message": err.Error(),
//	})
//}
//req_test.Header.Set("Content-Type", "application/json")
//
//resp_test, err := client.Do(req_test)
//if err != nil {
//	c.JSON(500, gin.H{
//		"code":    1,
//		"message": err.Error(),
//	})
//}
//defer resp_test.Body.Close()
//
//body_test, err := ioutil.ReadAll(resp_test.Body)
//if err != nil {
//	c.JSON(500, gin.H{
//		"code":    1,
//		"message": err.Error(),
//	})
//}
//fmt.Println(string(body_test))
//
//var data Data
//
//err = json.Unmarshal(body_test, &data)
//if err != nil {
//	c.JSON(500, gin.H{
//		"code":    1,
//		"message": err.Error(),
//	})
//}
//fmt.Println(data.Data)
//
//decode, err := base64.StdEncoding.DecodeString(data.Data)
//
//fmt.Println(string(decode), "  11111")

//err = json.Unmarshal(body_test, &sso)
//if err != nil {
//	c.JSON(500, gin.H{
//		"code":    1,
//		"message": err.Error(),
//	})
//}
