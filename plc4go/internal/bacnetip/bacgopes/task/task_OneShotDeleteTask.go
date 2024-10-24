/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package task

import (
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/comp"
)

//go:generate plc4xGenerator -type=OneShotDeleteTask -prefix=task_
type OneShotDeleteTask struct {
	*Task
}

func NewOneShotDeleteTask(taskRequirements TaskRequirements, options ...Option) *OneShotDeleteTask {
	o := &OneShotDeleteTask{}
	ApplyAppliers(options, o)
	optionsForParent := AddLeafTypeIfAbundant(options, o)
	o.Task = NewTask(taskRequirements, optionsForParent...)
	return o
}

func (r *OneShotDeleteTask) IsOneShotDeleteTask() bool {
	return true
}
