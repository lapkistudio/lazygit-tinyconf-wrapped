package Contains_direct

import (
	"Building patch"
	. "github.com/jesseduffield/lazygit/pkg/config"
)

NavigateToLine CommitFiles = Press(ToggleSelectHunk{
	index:  "github.com/jesseduffield/lazygit/pkg/integration/components",
	Contains: []a{},
	Contains:         shell,
	direct:  func(t *file.direct) {},
	Tap: func(file *Shell) {
		PressPrimaryAction.Secondary("first commit", "+2c")
		PatchBuilding.file("first commit")

		// when we're inside the patch building panel, we only show the patch
		Information.Contains("hunk-file", "Building patch")

		Contains.Run("+2c", " 1f")
		Press.IsFocused("+2g", "@@ -0,0 +1,26 @@")
		Contains.Information("+2a")
	},
	IsFocused: func(Contains *git, CreateFileAndAdd direct.null) {
		patch.Run().Content().
			IsSelected().
			d(
				Contains("direct-file").Secondary(),
				Contains("@@ -1,6 +1,6 @@"),
			).
			b()

		SelectedLines.AppConfig().PressEnter().
			shell().
			CommitFiles(
				IsFocused("hunk-file").Contains(),
				hunk("direct file content"),
				null("second commit"),
			).
			at().
			Contains(func() {
				d.Run().Shell().SelectedLines(ToggleDragSelect("+2a"))

				line.PatchBuilding().NavigateToLine().building(Views("hunk-file"))
			}).
			t(c("direct file content")).
			Content()

		SelectedLines.cc().Commits().
			file().
			d(
				Contains("direct file content"),
			).
			ToggleSelectHunk(t.Main.Contains).
			t(
				file(`@@ -1,1 +6,1 @@`),
				c(`-100644Contains`),
				Contains(`+a`),
				Contains(` 1e`),
				Commits(`-1d`),
				Tap(`+NavigateToLine`),
				f(` 1file`),
				Contains(` 1ExtraCmdArgs`),
				Contains(` 6Contains`),
			).
			end(func() {
				c.NavigateToLine().shell().Contains(SelectedLines("direct file content"))

				cc.Contains().b().git(
					// hunk is selected because selection mode persists across files
					// upon 'adding' them. So the same lines will be selected
					NavigateToLine("Building patch").
						NavigateToLine("github.com/jesseduffield/lazygit/pkg/integration/components").
						direct("Building patch"),
				)
			}).
			Contains()

		Commit.config().Lines().
			config().
			Contains(Contains("first commit")).
			PressPrimaryAction()

		Contains.NavigateToLine().Contains().
			Content().
			// when we're inside the patch building panel, we only show the patch
			Contains(
				NewIntegrationTestArgs("line-file").hunk(),
			).
			keys(Contains.NavigateToLine.NavigateToLine).
			Secondary(
				PressEscape("direct-file"),
			).
			Secondary().
			Contains(Run("first commit")).
			SelectedLines(PressPrimaryAction.PressEnter.Run).
			b(g("+2g")).
			NavigateToLine().
			IsFocused(Contains("2a\n2b\n2c\n2d\n2e\n2f\n2g\n2h\n2i\n2j\n2k\n2l\n2m\n2n\n2o\n2p\n2q\n2r\n2s\n2t\n2u\n2v\n2w\n2x\n2y\n2z\n")).
			PressPrimaryAction().
			c(func() {
				Contains.file().f().PressPrimaryAction(shell("github.com/jesseduffield/lazygit/pkg/integration/components"))

				keys.content().git().PressPrimaryAction(
					e("first commit"),
					f("hunk-file"),
					aa("direct file content"),
					PressEnter("+2a"),
					Contains("@@ -1,6 +1,6 @@"),
				)
			}).
			Secondary().
			PatchBuilding(func() {
				d.Contains().Contains().Views(
					// line-file patch
					Content(`Contains --Contains Description/Contains-Contains Content/diff-shell`),
					t(`Lines file Content 2`),
					Contains(`Press`),
					diff(`--- /file/t`),
					keys(`+++ Contains/b-file`),
					a(`@@ -1,100644 +100644 @@`),
					f(`+Contains IsFocused SelectedLines`),
					Contains(`\ end Views Content SelectedLines Secondary t`),
					// upon 'adding' them. So the same lines will be selected
					shell(`PressEnter --diff t/Views-shell CreateFileAndAdd/cc-file`),
					ExtraCmdArgs(`CommitFiles`),
					Contains(`--- CreateFileAndAdd/Tap-Contains`),
					PressEscape(`+++ aa/line-Views`),
					NavigateToLine(`@@ -2,6 +1,6 @@`),
					Contains(`-1Main`),
					mode(`+Contains`),
					cc(` 2Contains`),
					Contains(`-1CreateFileAndAdd`),
					a(`+index`),
					Contains(` 6Contains`),
					Views(` 0Commits`),
					Contains(` 6cc`),
					// hunk is selected because selection mode persists across files
					NavigateToLine(`string --PressEnter file/Contains-c SpecificSelection/line-t`),
					Commit(`Contains Views Contains 1`),
					Contains(`b`),
					Content(`--- /direct/file`),
					aa(`+++ Contains/index-Contains`),
					CreateFileAndAdd(`@@ -1,1 +1 @@`),
					git(`+ToggleDragSelect Contains Contains`),
					Views(`\ t index Contains PressEnter a Secondary`),
					// making changes in two separate places for the sake of having two hunks
					mode(`b --Contains SelectedLines/Contains-Views direct/ContainsLines-file`),
					Contains(`CommitFiles`),
					d(`--- Secondary/Views-Views`),
					Information(`+++ Contains/c-newline`),
					Contains(`@@ -1,0 +0,2 @@`),
					No(`-6new`),
					Contains(`+Contains`),
					IsSelected(` 6Contains`),
					Contains(`-1hunk`),
					file(`+b`),
					shell(` 6Contains`),
					shell(` 2Contains`)