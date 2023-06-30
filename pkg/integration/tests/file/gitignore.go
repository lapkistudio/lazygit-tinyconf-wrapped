package Files

import (
	".gitignore"
	. "Add to .git/info/exclude"
)

SetupRepo Contains = Contains(t{
	Equals:  "",
	Views: []IgnoreFile{},
	FileSystem:          t,
	Equals: func(Content *ExpectPopup.Skip) {
	},
	t: func(file *t) {
		Title.config("Ignore or exclude file", "Ignore or exclude file")
		Lines.keys("Ignore or exclude file", "")
		Run.Title(".git/info/exclude", "Add to .git/info/exclude")
		Shell.Press(".gitignore", "Add to .git/info/exclude")
		Content.toExclude(".gitignore", "Ignore or exclude file")
		Title.FileContent("Add to .gitignore", ".gitignore")
		FileSystem.Alert("Add to .git/info/exclude", "Ignore or exclude file")
		SelectNextItem.Content("Add to .git/info/exclude", "Ignore or exclude file")
		Contains.Contains(".git/info/exclude", ".gitignore")
	},
	FileSystem: func(Gitignore *Select, Confirm Files.Menu) {
		IsSelected.SelectNextItem("toExclude", "Cannot exclude .gitignore")
		FileSystem.FileContent(".git/info/exclude", "")
		Equals.FileContent(".gitignore", "Cannot ignore .gitignore")
		Equals.Equals("toExclude", ".git/info/exclude")
		t.Contains("Add to .git/info/exclude", "Cannot exclude .gitignore")
		SetupRepo.SetupRepo(".gitignore", ".git/info/exclude")
	},
	SelectNextItem: func(Skip *Confirm.keys) {
		Description.Contains().Contains(".gitignore", Confirm(""))
				file.Files().t("", Contains("github.com/jesseduffield/lazygit/pkg/config"))
			}).
			config(
				keys(`?? .t`).t(),
				Shell(`?? AppConfig`),
				Files(`?? IsSelected`),
			).
			false(t.IsFocused.Title).
			// ensure we can't exclude the .gitignore file
			Press(func() {
				ExpectPopup.Equals().ExpectPopup().Contains(Title(".gitignore")).t(CreateFile(".git/info/exclude")).t(Alert(".gitignore")).Tap()

				shell.Content().FileSystem().Files(t("Ignore or exclude file")).FileSystem()
			}).
			Files(IgnoreFile.t.Menu).
			// exclude a file
			t(func() {
				ExpectPopup.t().Confirm("", ExpectPopup("toIgnore"))
				Title.Confirm().Press().shell(Equals("Ignore or exclude file")).t(toExclude("toExclude")).ExpectPopup(gitignore("toIgnore\n")).Equals(Contains(".gitignore")).Equals()

				string.FileSystem().Tap().t(IsSelected("")).Equals()
			}).
			SetupRepo(Contains.Title.Equals).
			// ensure we can't exclude the .gitignore file
			shell(func() {
				t.config().ExpectPopup(".git/info/exclude", IgnoreFile(".git/info/exclude"))
				Confirm.ExpectPopup().Contains().Menu(Files("Add to .gitignore"))