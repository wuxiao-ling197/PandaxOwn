package devicerpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"math/rand"
)

const (
	jsonEndpoint = "http://192.168.0.30:8068/jsonrpc"
	dbname       = "odoo18"                                   //"o18"
	login        = "admin"                                    //"odoo18"
	pwd          = "e3e5e4f27d4482e2982a02d62c04a9561bb77f20" //"7bd8edc57978d705950009782e1af8259a501914"
)

// RPCPayload 结构体定义
type RPCPayloads struct {
	JsonRPC string    `json:"jsonrpc"`
	Method  string    `json:"method"`
	Params  RPCParams `json:"params"`
	ID      int       `json:"id"`
}

// RPCParams 结构体定义
type RPCParams struct {
	Service string        `json:"service"`
	Method  string        `json:"method"`
	Args    []interface{} `json:"args"`
}

// ResultData 结构体用于返回结果
type ResultData struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// getPayload 构造RPC负载
func getPayload(service, method string, args []interface{}) RPCPayloads {
	return RPCPayloads{
		JsonRPC: "2.0",
		Method:  "call",
		Params:  RPCParams{Service: service, Method: method, Args: args},
		ID:      random(0, 100000000),
	}
}

// odoo用户认证 如果报错odoo server error 200 可以尝试更新apikey
func Authenticate() (ResultData, error) {
	payload := getPayload("common", "login", []interface{}{dbname, login, pwd})
	// jsonPayload, err := json.Marshal(payload)
	// if err != nil {
	// 	return 0, fmt.Errorf("failed to marshal payload: %w", err)
	// }
	// req, err := http.NewRequest("POST", jsonEndpoint, bytes.NewBuffer(jsonPayload))
	// if err != nil {
	// 	return 0, fmt.Errorf("failed to create request: %w", err)
	// }
	// req.Header.Set("Content-Type", "application/json")
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	return 0, fmt.Errorf("request failed: %w", err)
	// }
	// defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return 0, fmt.Errorf("failed to read response body: %w", err)
	// }
	// if resp.StatusCode != http.StatusOK {
	// 	return 0, fmt.Errorf("request failed with status %s: %s", resp.Status, body)
	// }
	// var response struct {
	// 	Result int `json:"result"`
	// 	Error  *struct {
	// 		Code    int    `json:"code"`
	// 		Message string `json:"message"`
	// 	} `json:"error"`
	// }
	// if err := json.Unmarshal(body, &response); err != nil {
	// 	return 0, fmt.Errorf("failed to decode response: %w", err)
	// }
	// if response.Error != nil {
	// 	return 0, fmt.Errorf("Odoo authentication error: %s", response.Error.Message)
	// }
	// return response.Result, nil
	result, err := sendRequest(payload, "GET")
	if err != nil {
		return ResultData{Status: "fail", Code: 500, Message: "参数解析失败"}, err
	}
	// fmt.Printf("%+v\n", result)
	return result, nil
}

func AllUser() (ResultData, error) {
	domain := [][]interface{}{{"active", "=", "true"}}
	userId, err := Authenticate()
	if err != nil {
		return ResultData{Status: "fail", Code: 500, Message: "用户验证失败"}, err
	}
	payload := getPayload("object", "execute_kw", []interface{}{dbname, userId, pwd, "res_users", "read", []interface{}{domain}})
	result, err := sendRequest(payload, "GET")
	if err != nil {
		return ResultData{Status: "fail", Code: 500, Message: "参数解析失败"}, err
	}
	return result, nil
}

func Search(domain [][]interface{}, model string) (ResultData, error) {
	auth, err := Authenticate()
	if err != nil {
		return ResultData{Status: "fail", Code: 500, Message: "用户验证失败"}, err
	}
	userId := auth.Data
	payload := getPayload("object", "execute_kw", []interface{}{dbname, userId, pwd, model, "search", []interface{}{domain}})
	fmt.Printf("\npayload= %v\n", payload)
	return sendRequest(payload, "GET")
}

func FindOne(id int) (ResultData, error) {
	userId, err := Authenticate()
	if err != nil {
		return ResultData{Status: "fail", Code: 500, Message: "用户验证失败"}, err
	}

	payload := getPayload("object", "execute_kw", []interface{}{dbname, userId.Data, pwd, "res.users", "read", []interface{}{id}})
	return sendRequest(payload, "POST")
}

/*
表单新建按钮 web_save
*创建user时默认为内容用户，同时创建员工和联系人，参数可选：参照 addons\auth_signup\models\res_users.py create函数
[{'employee_ids': [], 'karma': 0, 'is_published': False, 'image_1920': False, 'name': 'create',

	'email': False, 'login': 'create', 'company_ids': [[4, 1]], 'company_id': 1, 'sel_groups_1_10_11': 1, ...权限组内容... ',in_group_303': True,
	lang': 'zh_CN', 'tz': 'Egypt', 'tour_enabled': True, 'action_id': False, 'notification_type': 'email',
	'odoobot_state': False, 'signature': '<p data-o-mail-quote="1">--<br data-o-mail-quote="1">create</p>',
	'calendar_default_privacy': 'public', 'livechat_username': False, 'livechat_lang_ids': [], 'oauth_provider_id': False,
	'oauth_uid': False}]

但是创建employee时不会创建用户只会创建联系人
odoo页面上新建时会自动生成res.partner 对象，直接调用create函数时请传入员工 name参数
*/
func Create(createData map[string]interface{}, model string) (ResultData, error) {
	userId, err := Authenticate()
	if err != nil {
		return ResultData{Status: "fail", Code: 500, Message: "参数解析失败"}, err
	}
	payload := getPayload("object", "execute_kw", []interface{}{dbname, userId.Data, pwd, model, "create", []interface{}{createData}})
	fmt.Printf("createPayload= %+v\n", payload)
	return sendRequest(payload, "POST")
}

// // 失败 不是规范调用？？
// func OnChange(createData map[string]interface{}, model string) (ResultData, error) {
// 	userId, err := Authenticate()
// 	if err != nil {
// 		return ResultData{Status: "fail", Code: 500, Message: "参数解析失败"}, err
// 	}

// 	payload := getPayload("object", "execute_kw", []interface{}{dbname, userId.Data, pwd, model, "onchange", []interface{}{createData}})
// 	return sendRequest(payload, "POST")
// }

func Write(writeData map[string]interface{}, model string, id int64) (ResultData, error) {
	userId, err := Authenticate()
	if err != nil {
		return ResultData{Status: "fail", Code: 500, Message: "参数解析失败"}, err
	}

	payload := getPayload("object", "execute_kw", []interface{}{dbname, userId.Data, pwd, model, "write", []interface{}{id, writeData}})
	fmt.Printf("writePayload= %+v\n", payload)
	return sendRequest(payload, "POST")
}

// only send their signup url by email, not change directly
func ResetPassword(user int64) (ResultData, error) {
	userId, err := Authenticate()
	if err != nil {
		return ResultData{Status: "fail", Code: 500, Message: "参数解析失败"}, err
	}

	payload := getPayload("object", "execute_kw", []interface{}{dbname, userId.Data, pwd, "res.users", "action_reset_password", []interface{}{user}})
	fmt.Printf("重置密码Payload= %+v\n", payload)
	return sendRequest(payload, "POST")
}

// 删除记录
func Delete(id int, model string) (ResultData, error) {
	userId, err := Authenticate()
	if err != nil {
		return ResultData{Status: "fail", Code: 500, Message: "参数解析失败"}, err
	}

	payload := getPayload("object", "execute_kw", []interface{}{dbname, userId.Data, pwd, model, "unlink", []interface{}{id}})
	fmt.Printf("deletePayload= %+v\n", payload)
	return sendRequest(payload, "POST")
}

// 创建关联用户
func CreateRalativeUser(login string) (ResultData, error) {
	userId, err := Authenticate()
	if err != nil {
		return ResultData{Status: "fail", Code: 500, Message: "参数解析失败"}, err
	}

	payload := getPayload("object", "execute_kw", []interface{}{dbname, userId.Data, pwd, "hr.employee", "action_create_user", []interface{}{login}})
	fmt.Printf("创建关联用户load= %+v\n", payload)
	return sendRequest(payload, "POST")
}

// 启用用户totp 有问题，没梳理完
func EnableTotp(data map[string]interface{}) (ResultData, error) {
	userId, err := Authenticate()
	if err != nil {
		return ResultData{Status: "fail", Code: 500, Message: "参数解析失败"}, err
	}

	payload := getPayload("object", "execute_kw", []interface{}{dbname, userId.Data, pwd, "res.users", "action_totp_enable_wizard", []interface{}{data}})
	fmt.Printf("模型内自定义函数payload= %+v\n", payload)
	return sendRequest(payload, "POST")
}

// 传参是个问题，需要对action
func CustomAction(model string, action string, data map[string]interface{}) (ResultData, error) {
	userId, err := Authenticate()
	if err != nil {
		return ResultData{Status: "fail", Code: 500, Message: "参数解析失败"}, err
	}

	payload := getPayload("object", "execute_kw", []interface{}{dbname, userId.Data, pwd, model, action, []interface{}{data}})
	fmt.Printf("模型内自定义函数payload= %+v\n", payload)
	return sendRequest(payload, "POST")
}

func sendRequest(payload RPCPayloads, method string) (ResultData, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return ResultData{Status: "failed"}, fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequest(method, jsonEndpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return ResultData{Status: "failed"}, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ResultData{Status: "failed", Code: 500}, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResultData{Status: "failed", Code: 500}, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return ResultData{Status: "failed", Code: 500}, fmt.Errorf("request failed with status %s: %s", resp.Status, body)
	}

	var response struct {
		Result interface{} `json:"result"`
		Error  *struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return ResultData{Status: "failed", Code: 500}, fmt.Errorf("failed to decode response: %w", err)
	}
	// fmt.Printf("req=%v\n", req)

	if err := json.Unmarshal(body, &response); err != nil {
		return ResultData{Status: "failed"}, fmt.Errorf("failed to decode response: %w", err)
	}

	if response.Error != nil {
		return ResultData{
			Status:  "failed",
			Code:    response.Error.Code,
			Message: response.Error.Message,
		}, fmt.Errorf("RPC error: %s", response.Error.Message)
	}

	// 将结果存储到 ResultData 中
	return ResultData{
		Status: "success",
		Code:   200, // 或其他合适的状态码
		Data:   response.Result,
	}, nil
}

// random 生成随机数
func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
