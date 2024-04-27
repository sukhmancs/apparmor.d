// apparmor.d - Full set of apparmor profiles
// Copyright (C) 2021-2024 Alexandre Pujol <alexandre@pujol.io>
// SPDX-License-Identifier: GPL-2.0-only

package builder

import (
	"slices"
	"strings"

	"github.com/roddhjav/apparmor.d/pkg/prebuild/cfg"
)

type Enforce struct {
	cfg.Base
}

func init() {
	RegisterBuilder(&Enforce{
		Base: cfg.Base{
			Keyword: "enforce",
			Msg:     "All profiles have been enforced",
		},
	})
}

func (b Enforce) Apply(profile string) string {
	matches := regFlags.FindStringSubmatch(profile)
	if len(matches) == 0 {
		return profile
	}

	flags := strings.Split(matches[1], ",")
	idx := slices.Index(flags, "complain")
	if idx == -1 {
		return profile
	}
	flags = slices.Delete(flags, idx, idx+1)
	strFlags := "{"
	if len(flags) >= 1 {
		strFlags = " flags=(" + strings.Join(flags, ",") + ") {"
	}

	// Remove all flags definition, then set new flags
	profile = regFlags.ReplaceAllLiteralString(profile, "")
	return regProfileHeader.ReplaceAllLiteralString(profile, strFlags)
}
