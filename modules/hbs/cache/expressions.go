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

package cache

import (
	"github.com/wenchangshou/falcon-plus/common/model"
	"github.com/wenchangshou/falcon-plus/modules/hbs/db"
	"sync"
)

type SafeExpressionCache struct {
	sync.RWMutex
	L []*model.Expression
}

var ExpressionCache = &SafeExpressionCache{}

func (this *SafeExpressionCache) Get() []*model.Expression {
	this.RLock()
	defer this.RUnlock()
	return this.L
}

func (this *SafeExpressionCache) Init() {
	es, err := db.QueryExpressions()
	if err != nil {
		return
	}

	this.Lock()
	defer this.Unlock()
	this.L = es
}
