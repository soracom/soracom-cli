package cmd

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestEmbeddedSkillIsBundled(t *testing.T) {
	body, err := assets.ReadFile(embeddedSkillPath)
	if err != nil {
		t.Fatalf("embedded SKILL.md missing (did `make generate` copy it into assets/skill/?): %v", err)
	}
	s := string(body)
	if !strings.Contains(s, "name: soracom-cli") {
		t.Errorf("embedded SKILL.md does not contain the expected frontmatter name; got first 80 bytes: %q", s[:min(80, len(s))])
	}
	for _, marker := range []string{"soracom describe", "--dry-run", "--fields"} {
		if !strings.Contains(s, marker) {
			t.Errorf("embedded SKILL.md does not mention %q", marker)
		}
	}
}

func TestEmbeddedAgentsIsBundled(t *testing.T) {
	body, err := assets.ReadFile(embeddedAgentsPath)
	if err != nil {
		t.Fatalf("embedded AGENTS.md missing (did `make generate` copy it into assets/skill/?): %v", err)
	}
	if !strings.Contains(string(body), "AGENTS.md") {
		t.Errorf("embedded AGENTS.md does not look like the AGENTS guide")
	}
}

func TestSkillInstallTargetDir(t *testing.T) {
	cases := []struct {
		name    string
		scope   string
		dir     string
		home    string
		want    string
		wantErr bool
	}{
		{name: "default project", scope: "project", want: filepath.Join(".agents", "skills", "soracom-cli")},
		{name: "empty scope defaults to project", scope: "", want: filepath.Join(".agents", "skills", "soracom-cli")},
		{name: "user scope", scope: "user", home: "/home/me", want: filepath.Join("/home/me", ".claude", "skills", "soracom-cli")},
		{name: "explicit dir overrides scope", scope: "user", dir: "/tmp/skills", home: "/home/me", want: filepath.Join("/tmp/skills", "soracom-cli")},
		{name: "invalid scope", scope: "bogus", wantErr: true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := skillInstallTargetDir(c.scope, c.dir, c.home)
			if c.wantErr {
				if err == nil {
					t.Fatalf("expected error for scope %q, got dir %q", c.scope, got)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != c.want {
				t.Errorf("skillInstallTargetDir(%q,%q,%q) = %q, want %q", c.scope, c.dir, c.home, got, c.want)
			}
		})
	}
}

func TestSkillProvenanceFooterMentionsVersion(t *testing.T) {
	if !strings.Contains(string(skillProvenanceFooter()), "soracom-cli v") {
		t.Errorf("provenance footer should reference the CLI version")
	}
}
