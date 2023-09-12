// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2023-09-12 13:44:27.9813739 +0200 CEST m=+0.011088001

package io


import (
	G "github.com/IBM/fp-go/io/generic"	
	T "github.com/IBM/fp-go/tuple"
)

// SequenceT1 converts 1 [IO[T]] into a [IO[T.Tuple1[T1]]]
func SequenceT1[T1 any](
  t1 IO[T1],
) IO[T.Tuple1[T1]] {
  return G.SequenceT1[
    IO[T.Tuple1[T1]],
    IO[T1],
  ](t1)
}

// SequenceTuple1 converts a [T.Tuple1[IO[T]]] into a [IO[T.Tuple1[T1]]]
func SequenceTuple1[T1 any](t T.Tuple1[IO[T1]]) IO[T.Tuple1[T1]] {
  return G.SequenceTuple1[
    IO[T.Tuple1[T1]],
    IO[T1],
  ](t)
}

// TraverseTuple1 converts a [T.Tuple1[IO[T]]] into a [IO[T.Tuple1[T1]]]
func TraverseTuple1[F1 ~func(A1) IO[T1], A1, T1 any](f1 F1) func(T.Tuple1[A1]) IO[T.Tuple1[T1]] {
  return G.TraverseTuple1[IO[T.Tuple1[T1]]](f1)
}

// SequenceT2 converts 2 [IO[T]] into a [IO[T.Tuple2[T1, T2]]]
func SequenceT2[T1, T2 any](
  t1 IO[T1],
  t2 IO[T2],
) IO[T.Tuple2[T1, T2]] {
  return G.SequenceT2[
    IO[T.Tuple2[T1, T2]],
    IO[T1],
    IO[T2],
  ](t1, t2)
}

// SequenceTuple2 converts a [T.Tuple2[IO[T]]] into a [IO[T.Tuple2[T1, T2]]]
func SequenceTuple2[T1, T2 any](t T.Tuple2[IO[T1], IO[T2]]) IO[T.Tuple2[T1, T2]] {
  return G.SequenceTuple2[
    IO[T.Tuple2[T1, T2]],
    IO[T1],
    IO[T2],
  ](t)
}

// TraverseTuple2 converts a [T.Tuple2[IO[T]]] into a [IO[T.Tuple2[T1, T2]]]
func TraverseTuple2[F1 ~func(A1) IO[T1], F2 ~func(A2) IO[T2], A1, A2, T1, T2 any](f1 F1, f2 F2) func(T.Tuple2[A1, A2]) IO[T.Tuple2[T1, T2]] {
  return G.TraverseTuple2[IO[T.Tuple2[T1, T2]]](f1, f2)
}

// SequenceT3 converts 3 [IO[T]] into a [IO[T.Tuple3[T1, T2, T3]]]
func SequenceT3[T1, T2, T3 any](
  t1 IO[T1],
  t2 IO[T2],
  t3 IO[T3],
) IO[T.Tuple3[T1, T2, T3]] {
  return G.SequenceT3[
    IO[T.Tuple3[T1, T2, T3]],
    IO[T1],
    IO[T2],
    IO[T3],
  ](t1, t2, t3)
}

// SequenceTuple3 converts a [T.Tuple3[IO[T]]] into a [IO[T.Tuple3[T1, T2, T3]]]
func SequenceTuple3[T1, T2, T3 any](t T.Tuple3[IO[T1], IO[T2], IO[T3]]) IO[T.Tuple3[T1, T2, T3]] {
  return G.SequenceTuple3[
    IO[T.Tuple3[T1, T2, T3]],
    IO[T1],
    IO[T2],
    IO[T3],
  ](t)
}

// TraverseTuple3 converts a [T.Tuple3[IO[T]]] into a [IO[T.Tuple3[T1, T2, T3]]]
func TraverseTuple3[F1 ~func(A1) IO[T1], F2 ~func(A2) IO[T2], F3 ~func(A3) IO[T3], A1, A2, A3, T1, T2, T3 any](f1 F1, f2 F2, f3 F3) func(T.Tuple3[A1, A2, A3]) IO[T.Tuple3[T1, T2, T3]] {
  return G.TraverseTuple3[IO[T.Tuple3[T1, T2, T3]]](f1, f2, f3)
}

// SequenceT4 converts 4 [IO[T]] into a [IO[T.Tuple4[T1, T2, T3, T4]]]
func SequenceT4[T1, T2, T3, T4 any](
  t1 IO[T1],
  t2 IO[T2],
  t3 IO[T3],
  t4 IO[T4],
) IO[T.Tuple4[T1, T2, T3, T4]] {
  return G.SequenceT4[
    IO[T.Tuple4[T1, T2, T3, T4]],
    IO[T1],
    IO[T2],
    IO[T3],
    IO[T4],
  ](t1, t2, t3, t4)
}

// SequenceTuple4 converts a [T.Tuple4[IO[T]]] into a [IO[T.Tuple4[T1, T2, T3, T4]]]
func SequenceTuple4[T1, T2, T3, T4 any](t T.Tuple4[IO[T1], IO[T2], IO[T3], IO[T4]]) IO[T.Tuple4[T1, T2, T3, T4]] {
  return G.SequenceTuple4[
    IO[T.Tuple4[T1, T2, T3, T4]],
    IO[T1],
    IO[T2],
    IO[T3],
    IO[T4],
  ](t)
}

// TraverseTuple4 converts a [T.Tuple4[IO[T]]] into a [IO[T.Tuple4[T1, T2, T3, T4]]]
func TraverseTuple4[F1 ~func(A1) IO[T1], F2 ~func(A2) IO[T2], F3 ~func(A3) IO[T3], F4 ~func(A4) IO[T4], A1, A2, A3, A4, T1, T2, T3, T4 any](f1 F1, f2 F2, f3 F3, f4 F4) func(T.Tuple4[A1, A2, A3, A4]) IO[T.Tuple4[T1, T2, T3, T4]] {
  return G.TraverseTuple4[IO[T.Tuple4[T1, T2, T3, T4]]](f1, f2, f3, f4)
}

// SequenceT5 converts 5 [IO[T]] into a [IO[T.Tuple5[T1, T2, T3, T4, T5]]]
func SequenceT5[T1, T2, T3, T4, T5 any](
  t1 IO[T1],
  t2 IO[T2],
  t3 IO[T3],
  t4 IO[T4],
  t5 IO[T5],
) IO[T.Tuple5[T1, T2, T3, T4, T5]] {
  return G.SequenceT5[
    IO[T.Tuple5[T1, T2, T3, T4, T5]],
    IO[T1],
    IO[T2],
    IO[T3],
    IO[T4],
    IO[T5],
  ](t1, t2, t3, t4, t5)
}

// SequenceTuple5 converts a [T.Tuple5[IO[T]]] into a [IO[T.Tuple5[T1, T2, T3, T4, T5]]]
func SequenceTuple5[T1, T2, T3, T4, T5 any](t T.Tuple5[IO[T1], IO[T2], IO[T3], IO[T4], IO[T5]]) IO[T.Tuple5[T1, T2, T3, T4, T5]] {
  return G.SequenceTuple5[
    IO[T.Tuple5[T1, T2, T3, T4, T5]],
    IO[T1],
    IO[T2],
    IO[T3],
    IO[T4],
    IO[T5],
  ](t)
}

// TraverseTuple5 converts a [T.Tuple5[IO[T]]] into a [IO[T.Tuple5[T1, T2, T3, T4, T5]]]
func TraverseTuple5[F1 ~func(A1) IO[T1], F2 ~func(A2) IO[T2], F3 ~func(A3) IO[T3], F4 ~func(A4) IO[T4], F5 ~func(A5) IO[T5], A1, A2, A3, A4, A5, T1, T2, T3, T4, T5 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5) func(T.Tuple5[A1, A2, A3, A4, A5]) IO[T.Tuple5[T1, T2, T3, T4, T5]] {
  return G.TraverseTuple5[IO[T.Tuple5[T1, T2, T3, T4, T5]]](f1, f2, f3, f4, f5)
}

// SequenceT6 converts 6 [IO[T]] into a [IO[T.Tuple6[T1, T2, T3, T4, T5, T6]]]
func SequenceT6[T1, T2, T3, T4, T5, T6 any](
  t1 IO[T1],
  t2 IO[T2],
  t3 IO[T3],
  t4 IO[T4],
  t5 IO[T5],
  t6 IO[T6],
) IO[T.Tuple6[T1, T2, T3, T4, T5, T6]] {
  return G.SequenceT6[
    IO[T.Tuple6[T1, T2, T3, T4, T5, T6]],
    IO[T1],
    IO[T2],
    IO[T3],
    IO[T4],
    IO[T5],
    IO[T6],
  ](t1, t2, t3, t4, t5, t6)
}

// SequenceTuple6 converts a [T.Tuple6[IO[T]]] into a [IO[T.Tuple6[T1, T2, T3, T4, T5, T6]]]
func SequenceTuple6[T1, T2, T3, T4, T5, T6 any](t T.Tuple6[IO[T1], IO[T2], IO[T3], IO[T4], IO[T5], IO[T6]]) IO[T.Tuple6[T1, T2, T3, T4, T5, T6]] {
  return G.SequenceTuple6[
    IO[T.Tuple6[T1, T2, T3, T4, T5, T6]],
    IO[T1],
    IO[T2],
    IO[T3],
    IO[T4],
    IO[T5],
    IO[T6],
  ](t)
}

// TraverseTuple6 converts a [T.Tuple6[IO[T]]] into a [IO[T.Tuple6[T1, T2, T3, T4, T5, T6]]]
func TraverseTuple6[F1 ~func(A1) IO[T1], F2 ~func(A2) IO[T2], F3 ~func(A3) IO[T3], F4 ~func(A4) IO[T4], F5 ~func(A5) IO[T5], F6 ~func(A6) IO[T6], A1, A2, A3, A4, A5, A6, T1, T2, T3, T4, T5, T6 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6) func(T.Tuple6[A1, A2, A3, A4, A5, A6]) IO[T.Tuple6[T1, T2, T3, T4, T5, T6]] {
  return G.TraverseTuple6[IO[T.Tuple6[T1, T2, T3, T4, T5, T6]]](f1, f2, f3, f4, f5, f6)
}

// SequenceT7 converts 7 [IO[T]] into a [IO[T.Tuple7[T1, T2, T3, T4, T5, T6, T7]]]
func SequenceT7[T1, T2, T3, T4, T5, T6, T7 any](
  t1 IO[T1],
  t2 IO[T2],
  t3 IO[T3],
  t4 IO[T4],
  t5 IO[T5],
  t6 IO[T6],
  t7 IO[T7],
) IO[T.Tuple7[T1, T2, T3, T4, T5, T6, T7]] {
  return G.SequenceT7[
    IO[T.Tuple7[T1, T2, T3, T4, T5, T6, T7]],
    IO[T1],
    IO[T2],
    IO[T3],
    IO[T4],
    IO[T5],
    IO[T6],
    IO[T7],
  ](t1, t2, t3, t4, t5, t6, t7)
}

// SequenceTuple7 converts a [T.Tuple7[IO[T]]] into a [IO[T.Tuple7[T1, T2, T3, T4, T5, T6, T7]]]
func SequenceTuple7[T1, T2, T3, T4, T5, T6, T7 any](t T.Tuple7[IO[T1], IO[T2], IO[T3], IO[T4], IO[T5], IO[T6], IO[T7]]) IO[T.Tuple7[T1, T2, T3, T4, T5, T6, T7]] {
  return G.SequenceTuple7[
    IO[T.Tuple7[T1, T2, T3, T4, T5, T6, T7]],
    IO[T1],
    IO[T2],
    IO[T3],
    IO[T4],
    IO[T5],
    IO[T6],
    IO[T7],
  ](t)
}

// TraverseTuple7 converts a [T.Tuple7[IO[T]]] into a [IO[T.Tuple7[T1, T2, T3, T4, T5, T6, T7]]]
func TraverseTuple7[F1 ~func(A1) IO[T1], F2 ~func(A2) IO[T2], F3 ~func(A3) IO[T3], F4 ~func(A4) IO[T4], F5 ~func(A5) IO[T5], F6 ~func(A6) IO[T6], F7 ~func(A7) IO[T7], A1, A2, A3, A4, A5, A6, A7, T1, T2, T3, T4, T5, T6, T7 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6, f7 F7) func(T.Tuple7[A1, A2, A3, A4, A5, A6, A7]) IO[T.Tuple7[T1, T2, T3, T4, T5, T6, T7]] {
  return G.TraverseTuple7[IO[T.Tuple7[T1, T2, T3, T4, T5, T6, T7]]](f1, f2, f3, f4, f5, f6, f7)
}

// SequenceT8 converts 8 [IO[T]] into a [IO[T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]]]
func SequenceT8[T1, T2, T3, T4, T5, T6, T7, T8 any](
  t1 IO[T1],
  t2 IO[T2],
  t3 IO[T3],
  t4 IO[T4],
  t5 IO[T5],
  t6 IO[T6],
  t7 IO[T7],
  t8 IO[T8],
) IO[T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]] {
  return G.SequenceT8[
    IO[T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]],
    IO[T1],
    IO[T2],
    IO[T3],
    IO[T4],
    IO[T5],
    IO[T6],
    IO[T7],
    IO[T8],
  ](t1, t2, t3, t4, t5, t6, t7, t8)
}

// SequenceTuple8 converts a [T.Tuple8[IO[T]]] into a [IO[T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]]]
func SequenceTuple8[T1, T2, T3, T4, T5, T6, T7, T8 any](t T.Tuple8[IO[T1], IO[T2], IO[T3], IO[T4], IO[T5], IO[T6], IO[T7], IO[T8]]) IO[T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]] {
  return G.SequenceTuple8[
    IO[T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]],
    IO[T1],
    IO[T2],
    IO[T3],
    IO[T4],
    IO[T5],
    IO[T6],
    IO[T7],
    IO[T8],
  ](t)
}

// TraverseTuple8 converts a [T.Tuple8[IO[T]]] into a [IO[T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]]]
func TraverseTuple8[F1 ~func(A1) IO[T1], F2 ~func(A2) IO[T2], F3 ~func(A3) IO[T3], F4 ~func(A4) IO[T4], F5 ~func(A5) IO[T5], F6 ~func(A6) IO[T6], F7 ~func(A7) IO[T7], F8 ~func(A8) IO[T8], A1, A2, A3, A4, A5, A6, A7, A8, T1, T2, T3, T4, T5, T6, T7, T8 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6, f7 F7, f8 F8) func(T.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]) IO[T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]] {
  return G.TraverseTuple8[IO[T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]]](f1, f2, f3, f4, f5, f6, f7, f8)
}

// SequenceT9 converts 9 [IO[T]] into a [IO[T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]]]
func SequenceT9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](
  t1 IO[T1],
  t2 IO[T2],
  t3 IO[T3],
  t4 IO[T4],
  t5 IO[T5],
  t6 IO[T6],
  t7 IO[T7],
  t8 IO[T8],
  t9 IO[T9],
) IO[T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]] {
  return G.SequenceT9[
    IO[T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]],
    IO[T1],
    IO[T2],
    IO[T3],
    IO[T4],
    IO[T5],
    IO[T6],
    IO[T7],
    IO[T8],
    IO[T9],
  ](t1, t2, t3, t4, t5, t6, t7, t8, t9)
}

// SequenceTuple9 converts a [T.Tuple9[IO[T]]] into a [IO[T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]]]
func SequenceTuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](t T.Tuple9[IO[T1], IO[T2], IO[T3], IO[T4], IO[T5], IO[T6], IO[T7], IO[T8], IO[T9]]) IO[T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]] {
  return G.SequenceTuple9[
    IO[T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]],
    IO[T1],
    IO[T2],
    IO[T3],
    IO[T4],
    IO[T5],
    IO[T6],
    IO[T7],
    IO[T8],
    IO[T9],
  ](t)
}

// TraverseTuple9 converts a [T.Tuple9[IO[T]]] into a [IO[T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]]]
func TraverseTuple9[F1 ~func(A1) IO[T1], F2 ~func(A2) IO[T2], F3 ~func(A3) IO[T3], F4 ~func(A4) IO[T4], F5 ~func(A5) IO[T5], F6 ~func(A6) IO[T6], F7 ~func(A7) IO[T7], F8 ~func(A8) IO[T8], F9 ~func(A9) IO[T9], A1, A2, A3, A4, A5, A6, A7, A8, A9, T1, T2, T3, T4, T5, T6, T7, T8, T9 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6, f7 F7, f8 F8, f9 F9) func(T.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) IO[T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]] {
  return G.TraverseTuple9[IO[T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]]](f1, f2, f3, f4, f5, f6, f7, f8, f9)
}

// SequenceT10 converts 10 [IO[T]] into a [IO[T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]]]
func SequenceT10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](
  t1 IO[T1],
  t2 IO[T2],
  t3 IO[T3],
  t4 IO[T4],
  t5 IO[T5],
  t6 IO[T6],
  t7 IO[T7],
  t8 IO[T8],
  t9 IO[T9],
  t10 IO[T10],
) IO[T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]] {
  return G.SequenceT10[
    IO[T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]],
    IO[T1],
    IO[T2],
    IO[T3],
    IO[T4],
    IO[T5],
    IO[T6],
    IO[T7],
    IO[T8],
    IO[T9],
    IO[T10],
  ](t1, t2, t3, t4, t5, t6, t7, t8, t9, t10)
}

// SequenceTuple10 converts a [T.Tuple10[IO[T]]] into a [IO[T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]]]
func SequenceTuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](t T.Tuple10[IO[T1], IO[T2], IO[T3], IO[T4], IO[T5], IO[T6], IO[T7], IO[T8], IO[T9], IO[T10]]) IO[T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]] {
  return G.SequenceTuple10[
    IO[T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]],
    IO[T1],
    IO[T2],
    IO[T3],
    IO[T4],
    IO[T5],
    IO[T6],
    IO[T7],
    IO[T8],
    IO[T9],
    IO[T10],
  ](t)
}

// TraverseTuple10 converts a [T.Tuple10[IO[T]]] into a [IO[T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]]]
func TraverseTuple10[F1 ~func(A1) IO[T1], F2 ~func(A2) IO[T2], F3 ~func(A3) IO[T3], F4 ~func(A4) IO[T4], F5 ~func(A5) IO[T5], F6 ~func(A6) IO[T6], F7 ~func(A7) IO[T7], F8 ~func(A8) IO[T8], F9 ~func(A9) IO[T9], F10 ~func(A10) IO[T10], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6, f7 F7, f8 F8, f9 F9, f10 F10) func(T.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) IO[T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]] {
  return G.TraverseTuple10[IO[T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]]](f1, f2, f3, f4, f5, f6, f7, f8, f9, f10)
}
