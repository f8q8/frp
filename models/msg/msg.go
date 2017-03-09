// Copyright 2016 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package msg

import "reflect"

const (
	TypeLogin         = 'o'
	TypeLoginResp     = '1'
	TypeNewProxy      = 'p'
	TypeNewProxyResp  = '2'
	TypeNewWorkConn   = 'w'
	TypeReqWorkConn   = 'r'
	TypeStartWorkConn = 's'
	TypePing          = 'h'
	TypePong          = '4'
)

var (
	TypeMap       map[byte]reflect.Type
	TypeStringMap map[reflect.Type]byte
)

func init() {
	TypeMap = make(map[byte]reflect.Type)
	TypeStringMap = make(map[reflect.Type]byte)

	TypeMap[TypeLogin] = getTypeFn((*Login)(nil))
	TypeMap[TypeLoginResp] = getTypeFn((*LoginResp)(nil))
	TypeMap[TypeNewProxy] = getTypeFn((*NewProxy)(nil))
	TypeMap[TypeNewProxyResp] = getTypeFn((*NewProxyResp)(nil))
	TypeMap[TypeNewWorkConn] = getTypeFn((*NewWorkConn)(nil))
	TypeMap[TypeReqWorkConn] = getTypeFn((*ReqWorkConn)(nil))
	TypeMap[TypeStartWorkConn] = getTypeFn((*StartWorkConn)(nil))
	TypeMap[TypePing] = getTypeFn((*Ping)(nil))
	TypeMap[TypePong] = getTypeFn((*Pong)(nil))

	for k, v := range TypeMap {
		TypeStringMap[v] = k
	}
}

func getTypeFn(obj interface{}) reflect.Type {
	return reflect.TypeOf(obj).Elem()
}

// Message wraps socket packages for communicating between frpc and frps.
type Message interface{}

// When frpc start, client send this message to login to server.
type Login struct {
	Version      string `json:"version"`
	Hostname     string `json:"hostname"`
	Os           string `json:"os"`
	Arch         string `json:"arch"`
	User         string `json:"user"`
	PrivilegeKey string `json:"privilege_key"`
	Timestamp    int64  `json:"timestamp"`
	RunId        string `json:"run_id"`

	// Some global configures.
	PoolCount int `json:"pool_count"`
}

type LoginResp struct {
	Version string `json:"version"`
	RunId   string `json:"run_id"`
	Error   string `json:"error"`
}

// When frpc login success, send this message to frps for running a new proxy.
type NewProxy struct {
	ProxyName      string `json:"proxy_name"`
	ProxyType      string `json:"proxy_type"`
	UseEncryption  bool   `json:"use_encryption"`
	UseCompression bool   `json:"use_compression"`

	// tcp and udp only
	RemotePort int64 `json:"remote_port"`

	// http and https only
	CustomDomains     []string `json:"custom_domains"`
	SubDomain         string   `json:"subdomain"`
	Locations         []string `json:"locations"`
	HostHeaderRewrite string   `json:"host_header_rewrite"`
	HttpUser          string   `json:"http_user"`
	HttpPwd           string   `json:"http_pwd"`
}

type NewProxyResp struct {
	ProxyName string `json:"proxy_name"`
	Error     string `json:"error"`
}

type NewWorkConn struct {
	RunId string `json:"run_id"`
}

type ReqWorkConn struct {
}

type StartWorkConn struct {
	ProxyName string `json:"proxy_name"`
}

type Ping struct {
}

type Pong struct {
}