// Copyright (c) 2018 The Jaeger Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package calculationstrategy

// ProbabilityCalculator calculates the new probability given the current and target QPS
type ProbabilityCalculator interface {
	Calculate(targetQPS, curQPS, prevProbability float64) (newProbability float64)
}

// Func wraps a function of appropriate signature and makes a ProbabilityCalculator from it.
type Func func(targetQPS, curQPS, prevProbability float64) (newProbability float64)

// Calculate implements Calculator interface.
func (f Func) Calculate(targetQPS, curQPS, prevProbability float64) float64 {
	return f(targetQPS, curQPS, prevProbability)
}