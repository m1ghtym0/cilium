// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

//go:build !privileged_tests

package stream_test

import (
	"context"
	"errors"
	"io"
	"testing"

	. "github.com/cilium/cilium/pkg/stream"
)

func TestFirst(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 1. First on non-empty source
	fst, err := First(ctx, Range(42, 1000))
	assertNil(t, "First", err)

	if fst != 42 {
		t.Fatalf("expected 42, got %d", fst)
	}

	// 2. First on empty source
	_, err = First(ctx, Empty[int]())
	if !errors.Is(err, io.EOF) {
		t.Fatalf("expected EOF, got %s", err)
	}

	// 3. cancelled context
	cancel()
	_, err = First(ctx, Stuck[int]())
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected Canceled, got %s", err)
	}
}

func TestLast(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 1. Last on non-empty source
	fst, err := Last(ctx, Range(42, 100))
	assertNil(t, "Last", err)

	if fst != 99 {
		t.Fatalf("expected 99, got %d", fst)
	}

	// 2. First on empty source
	_, err = Last(ctx, Empty[int]())
	if !errors.Is(err, io.EOF) {
		t.Fatalf("expected EOF, got %s", err)
	}

	// 3. cancelled context
	cancel()
	_, err = Last(ctx, Stuck[int]())
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected Canceled, got %s", err)
	}
}

func TestToSlice(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	xs, err := ToSlice(ctx, Range(0, 5))
	assertNil(t, "ToSlice", err)
	assertSlice(t, "ToSlice", []int{0, 1, 2, 3, 4}, xs)
}
