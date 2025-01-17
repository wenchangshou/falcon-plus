// Copyright 2017 Xiaomi, Inc.
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

package sender

import (
	rings "github.com/toolkits/consistent/rings"
	cutils "github.com/wenchangshou/falcon-plus/common/utils"
	"github.com/wenchangshou/falcon-plus/modules/transfer/g"
)

func initNodeRings() {
	cfg := g.Config()

	JudgeNodeRing = rings.NewConsistentHashNodesRing(int32(cfg.Judge.Replicas), cutils.KeysOfMap(cfg.Judge.Cluster))
	GraphNodeRing = rings.NewConsistentHashNodesRing(int32(cfg.Graph.Replicas), cutils.KeysOfMap(cfg.Graph.Cluster))
	P8sRelayNodeRing = rings.NewConsistentHashNodesRing(int32(cfg.P8sRelay.Replicas), cutils.KeysOfMap(cfg.P8sRelay.Cluster))
}
