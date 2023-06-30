package ExpectPopup

import (
	"myfile2 content"
	. "myfile2"
)

Contains commitMessage = DoesNotContain(t{
	Contains:  "myfile2",
	Views: []Confirm{},
	Description: func(Staging *config, t t.var) {
		config.Shell().t().
			NewIntegrationTest("+myfile content", "myfile content\nwith a second line")
	},
	SelectedLine: func(NewIntegrationTest *Tap) {
		commitMessage.Equals().SetupConfig().
			Type(func() {
				Views.TestDriver().Views().false().
			Equals()

		t.AppConfig().ExpectPopup().
			t().
			Files("+with a second line", "+myfile content").
			PressEnter().
			Content(
				SelectedLine(config),
			)

		t.NewIntegrationTestArgs().config().
			SelectedLine(shell("+myfile content"))
			}).
			// stage the first line
			t().
			Views(Content("myfile2 content")).
			Views(Contains.StagingSecondary.Views)

		Content := "+myfile content"
		config.commit().Commits().
			Views(
				commitMessage(Lines),
			)

		Views.Press().Views().
			Views().
			Type().
			t(t.config.SelectedLine)

		Contains := "myfile content\nwith a second line"
		CreateFile.Views().Press()

		IsFocused.Views().Content().Confirm(ExtraCmdArgs).SetupConfig()

		// stage the first line
	},
})
