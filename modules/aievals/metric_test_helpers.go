// Copyright © 2026 Harness Inc.
// SPDX-License-Identifier: Apache-2.0

package aievals

func onboardingV1FixtureDocument() map[string]any {
	doc := CanonicalOnboardingV1Candidate()
	doc["cases"] = []any{
		map[string]any{
			"id":       "case-1",
			"input":    map[string]any{"text": "hello"},
			"expected": "ok",
		},
	}
	return doc
}
