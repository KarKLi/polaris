/*
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package config

import (
	"context"

	apiconfig "github.com/polarismesh/specification/source/go/api/v1/config_manage"

	"github.com/polarismesh/polaris/apiserver/grpcserver"
)

// GetConfigFile 拉取配置
func (g *ConfigGRPCServer) GetConfigFile(ctx context.Context,
	configFile *apiconfig.ClientConfigFileInfo) (*apiconfig.ConfigClientResponse, error) {
	ctx = grpcserver.ConvertContext(ctx)
	response := g.configServer.GetConfigFileForClient(ctx, configFile)
	return response, nil
}

// WatchConfigFiles 订阅配置变更
func (g *ConfigGRPCServer) WatchConfigFiles(ctx context.Context,
	request *apiconfig.ClientWatchConfigFileRequest) (*apiconfig.ConfigClientResponse, error) {
	ctx = grpcserver.ConvertContext(ctx)

	// 阻塞等待响应
	callback, err := g.configServer.WatchConfigFiles(ctx, request)
	if err != nil {
		return nil, err
	}

	return callback(), nil
}
