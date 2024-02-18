// Copyright (c) 2024 IBM Corp.
// All rights reserved.
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

package statereaderioeither

import (
	G "github.com/IBM/fp-go/statereaderioeither/generic"
)

func Right[S, R, E, A any](a A) StateReaderIOEither[S, R, E, A] {
	return G.Right[StateReaderIOEither[S, R, E, A]](a)
}

func Of[S, R, E, A any](a A) StateReaderIOEither[S, R, E, A] {
	return G.Of[StateReaderIOEither[S, R, E, A]](a)
}

func MonadMap[S, R, E, A, B any](fa StateReaderIOEither[S, R, E, A], f func(A) B) StateReaderIOEither[S, R, E, B] {
	return G.MonadMap[StateReaderIOEither[S, R, E, A], StateReaderIOEither[S, R, E, B]](fa, f)
}

func Map[S, R, E, A, B any](f func(A) B) func(StateReaderIOEither[S, R, E, A]) StateReaderIOEither[S, R, E, B] {
	return G.Map[StateReaderIOEither[S, R, E, A], StateReaderIOEither[S, R, E, B]](f)
}

func MonadChain[S, R, E, A, B any](fa StateReaderIOEither[S, R, E, A], f func(A) StateReaderIOEither[S, R, E, B]) StateReaderIOEither[S, R, E, B] {
	return G.MonadChain[StateReaderIOEither[S, R, E, A], StateReaderIOEither[S, R, E, B]](fa, f)
}

func Chain[S, R, E, A, B any](f func(A) StateReaderIOEither[S, R, E, B]) func(StateReaderIOEither[S, R, E, A]) StateReaderIOEither[S, R, E, B] {
	return G.Chain[StateReaderIOEither[S, R, E, A], StateReaderIOEither[S, R, E, B]](f)
}

func MonadAp[S, R, E, A, B any](fab StateReaderIOEither[S, R, E, func(A) B], fa StateReaderIOEither[S, R, E, A]) StateReaderIOEither[S, R, E, B] {
	return G.MonadAp[StateReaderIOEither[S, R, E, A], StateReaderIOEither[S, R, E, B], StateReaderIOEither[S, R, E, func(A) B]](fab, fa)
}

func Ap[S, R, E, A, B any](fa StateReaderIOEither[S, R, E, A]) func(StateReaderIOEither[S, R, E, func(A) B]) StateReaderIOEither[S, R, E, B] {
	return G.Ap[StateReaderIOEither[S, R, E, A], StateReaderIOEither[S, R, E, B], StateReaderIOEither[S, R, E, func(A) B]](fa)
}