package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// The agent skill is embedded into the binary at build time. `make generate`
// copies the canonical SKILL.md and AGENTS.md into assets/skill/ so the
// existing `//go:embed assets/*` directive (see lang_utils.go) bundles them.
const (
	embeddedSkillPath  = "assets/skill/SKILL.md"
	embeddedAgentsPath = "assets/skill/AGENTS.md"

	skillName = "soracom-cli"
)

var skillShowAgents bool
var skillInstallScope string
var skillInstallDir string

func init() {
	SkillShowCmd.Flags().BoolVar(&skillShowAgents, "agents", false, TRCLI("cli.skill.show.agents"))
	SkillInstallCmd.Flags().StringVar(&skillInstallScope, "scope", "project", TRCLI("cli.skill.install.scope"))
	SkillInstallCmd.Flags().StringVar(&skillInstallDir, "dir", "", TRCLI("cli.skill.install.dir"))

	SkillCmd.AddCommand(SkillShowCmd)
	SkillCmd.AddCommand(SkillInstallCmd)
	RootCmd.AddCommand(SkillCmd)
}

// SkillCmd defines 'skill' command
var SkillCmd = &cobra.Command{
	Use:   "skill",
	Short: TRCLI("cli.skill.summary"),
	Long:  TRCLI("cli.skill.description"),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return fmt.Errorf("unknown command %q for %q", args[0], cmd.CommandPath())
		}
		return cmd.Help()
	},
}

// SkillShowCmd defines 'skill show' subcommand
var SkillShowCmd = &cobra.Command{
	Use:   "show",
	Short: TRCLI("cli.skill.show.summary"),
	Long:  TRCLI("cli.skill.show.description"),
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		path := embeddedSkillPath
		if skillShowAgents {
			path = embeddedAgentsPath
		}
		body, err := assets.ReadFile(path)
		if err != nil {
			return fmt.Errorf("embedded agent skill not found: %w", err)
		}
		_, err = os.Stdout.Write(body)
		return err
	},
}

// SkillInstallCmd defines 'skill install' subcommand
var SkillInstallCmd = &cobra.Command{
	Use:   "install",
	Short: TRCLI("cli.skill.install.summary"),
	Long:  TRCLI("cli.skill.install.description"),
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		body, err := assets.ReadFile(embeddedSkillPath)
		if err != nil {
			return fmt.Errorf("embedded agent skill not found: %w", err)
		}

		home := ""
		if skillInstallDir == "" && skillInstallScope == "user" {
			h, err := os.UserHomeDir()
			if err != nil {
				return fmt.Errorf("cannot determine home directory for --scope user: %w", err)
			}
			home = h
		}
		dir, err := skillInstallTargetDir(skillInstallScope, skillInstallDir, home)
		if err != nil {
			return err
		}

		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
		dest := filepath.Join(dir, "SKILL.md")
		body = append(body, skillProvenanceFooter()...)
		if err := os.WriteFile(dest, body, 0o644); err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "Installed %s agent skill to %s\n", skillName, dest)
		return nil
	},
}

// skillInstallTargetDir resolves the skill directory (the directory that will
// directly contain SKILL.md) from the requested scope/dir. It is pure so it can
// be unit-tested without touching the filesystem.
func skillInstallTargetDir(scope, dir, home string) (string, error) {
	if dir != "" {
		return filepath.Join(dir, skillName), nil
	}
	switch scope {
	case "", "project":
		return filepath.Join(".agents", "skills", skillName), nil
	case "user":
		return filepath.Join(home, ".claude", "skills", skillName), nil
	default:
		return "", fmt.Errorf("invalid --scope %q (expected \"project\" or \"user\")", scope)
	}
}

// skillProvenanceFooter records which CLI version produced the installed copy,
// mirroring how skill installers stamp provenance, so a stale copy is visible.
func skillProvenanceFooter() []byte {
	v := version
	if v == "" {
		v = "dev"
	}
	return []byte(fmt.Sprintf("\n<!-- Installed by `soracom skill install` from soracom-cli v%s. Re-run after upgrading the CLI to refresh. -->\n", v))
}
