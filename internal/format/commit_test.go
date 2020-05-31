package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommitMessage(t *testing.T) {
	tcs := map[string]struct {
		o        *CommitMessageOption
		expected string
	}{
		"Simple subject feature": {
			o:        &CommitMessageOption{Ctype: FeatureCommit, Description: "commit description"},
			expected: "feat: commit description",
		},
		"Subject + scope perf": {
			o:        &CommitMessageOption{Ctype: PerfCommit, Description: "message", Scope: "scope"},
			expected: "perf(scope): message",
		},
		"Breaking change refactor": {
			o:        &CommitMessageOption{Ctype: RefactorCommit, Description: "message", BreakingChanges: true},
			expected: "refactor!: message",
		},
		"Full stuff": {
			o:        &CommitMessageOption{Ctype: FeatureCommit, Scope: "scope", Description: "message", BreakingChanges: true, Body: "The message body", Footers: []string{"First foot", "Second foot"}},
			expected: "feat(scope)!: message\n\nThe message body\n\nFirst foot\nSecond foot",
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, CommitMessage(tc.o))
		})
	}
}

func TestFindCommitType(t *testing.T) {
	tcs := map[string]struct {
		str      string
		expected CommitType
	}{
		"Nil": {"fail", NilCommit},

		"B":      {"b", BuildCommit},
		"Build":  {"bUild", BuildCommit},
		"Builds": {"builds", BuildCommit},

		"Ci": {"ci", CiCommit},

		"Ch":     {"ch", ChoreCommit},
		"Chore":  {"chore", ChoreCommit},
		"Chores": {"chOreS", ChoreCommit},

		"D":    {"d", DocCommit},
		"Doc":  {"Doc", DocCommit},
		"Docs": {"docs", DocCommit},

		"Fe":       {"fe", FeatureCommit},
		"Feat":     {"feAt", FeatureCommit},
		"Feats":    {"feats", FeatureCommit},
		"Feature":  {"feature", FeatureCommit},
		"Features": {"features", FeatureCommit},

		"Fi":    {"fi", FixCommit},
		"Fix":   {"Fix", FixCommit},
		"Fixes": {"fixEs", FixCommit},

		"P":            {"p", PerfCommit},
		"Perf":         {"perf", PerfCommit},
		"Perfs":        {"pErFs", PerfCommit},
		"Performance":  {"performance", PerfCommit},
		"Performances": {"performances", PerfCommit},

		"R":         {"r", RefactorCommit},
		"Refactor":  {"reFactor", RefactorCommit},
		"Refactors": {"reFactors", RefactorCommit},

		"S":      {"s", StyleCommit},
		"Style":  {"style", StyleCommit},
		"Styles": {"stYles", StyleCommit},

		"T":     {"t", TestCommit},
		"Test":  {"Test", TestCommit},
		"Tests": {"tests", TestCommit},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, FindCommitType(tc.str))
		})
	}
}

func TestParseCommitMsg(t *testing.T) {
	tcs := map[string]struct {
		str      string
		expected *CommitMessageOption
	}{
		"Bad": {"i'm bad", nil},
		"Simple": {"feat: message description",
			&CommitMessageOption{Ctype: FeatureCommit, Description: "message description"}},
		"Scoped": {"fix(scope): message description",
			&CommitMessageOption{Ctype: FixCommit, Description: "message description", Scope: "scope"}},
		"Breaking change": {"feat!: message description",
			&CommitMessageOption{Ctype: FeatureCommit, Description: "message description", BreakingChanges: true}},
		"With body": {"feat: message description\n\nCommit body\n",
			&CommitMessageOption{Ctype: FeatureCommit, Description: "message description", Body: "Commit body"}},
		"With footers": {"feat: message description\n\nCommit body\n\nFooter: 1\nFooter #2",
			&CommitMessageOption{Ctype: FeatureCommit, Description: "message description", Body: "Commit body", Footers: []string{"Footer: 1", "Footer #2"}}},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, ParseCommitMsg(tc.str))
		})
	}
}
