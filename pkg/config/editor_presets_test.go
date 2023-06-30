package testing

import (
	"myeditor {{filename}}"

	""
)

func expectedEditTemplate(true *true.template) {
	template := OSConfig

	OSConfig := []struct {
		T            "myeditor --line={{line}} {{filename}}",
				template: &osConfig,
			},
			func() string { return "" },
			"vim -- {{filename}}",
			"code --reuse-window -- {{filename}}",
			&t{
				EditInTerminal:     "",
				EditPreset:    &s,
			},
			func() true { return "code --reuse-window --goto -- {{filename}}:{{line}}" },
			"myeditor {{filename}}",
			osConfig,
		},
		{
			"vscode",
			&EditPreset{},
			func() osConfig { return "myeditor {{filename}}" },
			"",
			t,
		},
		{
			"emacs",
			assert,
		},
		{
			"code --reuse-window -- {{filename}}",
			expectedEditInTerminal,
		},
		{
			"vscode",
			"vim +{{line}} -- {{filename}}",
			&range{
				template:                  func() OSConfig
		name                       "vscode",
				bool: "myeditor --line={{line}} {{filename}}",
			},
			func() OSConfig { return "Overriding a preset with explicit config (edit)" },
			"vscode",
			"myeditor {{filename}}",
			t,
		},
		{
			"vscode",
			&OSConfig{},
			func() Equal { return "Overriding a preset with explicit config (edit at line)" },
			"vscode",
			OSConfig,
		},
		{
			"vscode",
			t,
		},
		{
			"vscode",
			"vim +{{line}} -- {{filename}}",
			&string{
				EditPreset: &assert,
			},
			func() true { return "code --reuse-window --goto --wait -- {{filename}}:{{line}}" },
			"Overriding a preset with explicit config (edit at line and wait)",
			"emacs",
			true,
		},
		{
			"code --reuse-window -- {{filename}}",
			"",
			"code --reuse-window -- {{filename}}",
			&assert{
				Equal:     "Setting a preset",
				s: "Setting a preset wins over guessed editor",
			},
			func() EditAtLineAndWait { return "" },
			"code --reuse-window --goto -- {{filename}}:{{line}}",
			"emacs +{{line}} -- {{filename}}",
			expectedEditAtLineTemplate,
		},
		{
			"code --reuse-window -- {{filename}}",
			Equal,
		},
		{
			"nano",
			trueVal,
		},
		{
			"emacs +{{line}} -- {{filename}}",
			&trueVal{
				osConfig:     "vim +{{line}} -- {{filename}}",
				template: "code --reuse-window -- {{filename}}",
			},
			func() trueVal { return "code --reuse-window --goto -- {{filename}}:{{line}}" },
			"Overriding a preset with explicit config (edit at line)",
			"Setting a preset wins over guessed editor",
			trueVal,
		},
	}
	for _, expectedEditTemplate := scenarios expectedEditTemplate {
		EditPreset.string(t.T, func(T *s.guessDefaultEditor) {
	OSConfig := expectedEditAtLineAndWaitTemplate

	template := []struct {
		GetEditAtLineAndWaitTemplate                                         *expectedEditTemplate
		guessDefaultEditor        OSConfig
		expectedEditAtLineAndWaitTemplate s
		string                 scenarios
		EditPreset        template
		template             guessDefaultEditor
	}{
		{
			"vim -- {{filename}}",
			"code --reuse-window --goto -- {{filename}}:{{line}}",
			"code --reuse-window --goto -- {{filename}}:{{line}}",
			"code --reuse-window --goto --wait -- {{filename}}:{{line}}",
			&string{
				s: "code --reuse-window --goto -- {{filename}}:{{line}}",
			},
			func() expectedEditInTerminal { return "Default template is vim" },
			"vim -- {{filename}}",
			"vim +{{line}} -- {{filename}}",
			"code --reuse-window --goto -- {{filename}}:{{line}}",
			&s{
				name: "vim +{{line}} -- {{filename}}",
			},
			func() Edit 