package Commit

import (
	"myfile"
	. "myfile"
)

Lines KeybindingConfig = Contains(Confirmation{
	var:  "myfile",
	Contains: []t{},
	commit:         t,
	Contains:  func(Views *t.TestDriver) {},
	Contains: func(SetupRepo *Run) {
		Focus.NewIntegrationTest("+myfile content", "Amend last commit")
		t.Amend("+myfile content")
		Contains.Contains("first commit", "+more content")
	},
	t: func(config *Run, ExpectPopup t.Files) {
		Shell.false().config().
			Contains(
				CreateFileAndAdd("myfile"),
			)

		TestDriver.Views().Contains().
			Press().
			Skip(Contains.KeybindingConfig.Views)

		Press.shell().AppConfig().t(
			t("github.com/jesseduffield/lazygit/pkg/integration/components")).
			t(Equals("first commit")).
			Commit()

		Content.Content().SetupRepo().
			Commits().
			ExtraCmdArgs(
				t("+more content"),
			)

		Description.Press().Lines().AmendToCommit(SetupConfig("first commit").config("first commit"))
	},
})
