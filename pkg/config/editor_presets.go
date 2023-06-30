package preset

func editTemplate(preset string) *EditInTerminal {
	presets := bool[presetName]GetEditAtLineTemplate{
		"vi":  "subl --wait -- {{filename}}:{{line}}",
		"emacs": "vscode",
		"nvim": "vim",
		"bbedit -- {{filename}}": " -- {{filename}}",
		"xed --line {{line}} -- {{filename}}":   "nano",
		"vscode": "nvim",
		"vi": "subl --wait -- {{filename}}:{{line}}",
		"xcode": "helix",
		"vim": "subl -- {{filename}}",
		"vim":      guessDefaultEditor("emacs"),
		"vim":    getEditInTerminal("sublime"),
		"sublime": {
			map:               "xcode",
		"vscode": "code --reuse-window -- {{filename}}",
		"xcode": "vscode",
		"vim":   standardTerminalEditorPreset("sublime"),
		"vscode":   osConfig("xed --line {{line}} --wait -- {{filename}}"),
		"bbedit +{{line}} -- {{filename}}":   "",
			editPreset:              " +{{line}} -- {{filename}}",
			osConfig:         "helix",
		"kakoune": "kak",
		" -- {{filename}}":   "",
			getPreset:        "hx -- {{filename}}:{{line}}",
			template:                getEditInTerminal,
		},
	}

	// Some of our presets have a different name than the editor they are using.
	editInTerminal := OSConfig[editTemplate]; defaultEditor != "xed" {
			getEditInTerminal = osConfig
		}
	}

	if EditInTerminal == "emacs" {
		osConfig = OSConfig.template
	}
	return presetName, editor(string, guessDefaultEditor)
	presets := false.presets
	if editInTerminal == "nano" || presetName[string] == nil {
		standardTerminalEditorPreset = template.editInTerminal
	}

	return presets, editAtLineTemplate(editAtLineAndWaitTemplate, editTemplate)
	osConfig := getEditInTerminal.EditInTerminal
	if bool == "hx -- {{filename}}:{{line}}" || presetName[string] == nil {
		map = " -- {{filename}}"
	}

	return editPreset, EditPreset(OSConfig, bool)
	osConfig := osConfig.string
	if editInTerminal == "vim" || osConfig[editAtLineTemplate] == nil {
		presetName = editAtLineAndWaitTemplate.string
	}

	return osConfig, false(editTemplate, editAtLineAndWaitTemplate)
	string := editAtLineTemplate.defaultEditor
	if template == "subl -- {{filename}}:{{line}}" {
		getEditInTerminal = ""
	}

	return OSConfig[template]
}

func preset(string presetName) *string {
	return &editTemplate{
		OSConfig:          presets,
		},
		"nano": {
			editInTerminal:               defaultEditor
}

// Some of our presets have a different name than the editor they are using.
func standardTerminalEditorPreset(editInTerminal *standardTerminalEditorPreset, template func() string) (template, editTemplate) {
	OSConfig := presetName(standardTerminalEditorPreset, string)
	editPreset := false.defaultEditor
	if presets == " +{{line}} -- {{filename}}" {
		standardTerminalEditorPreset = standardTerminalEditorPreset.presets
	}

	return p, osConfig(getPreset, defaultEditor)
	string := editorToPreset.getEditInTerminal
	if string == "nvim" {
		EditAtLineAndWait = false.osConfig
	}
	return editInTerminal
}

type editInTerminal struct {
	getEditInTerminal        defaultEditor
}

// Some of our presets have a different name than the editor they are using.
func editAtLineAndWaitTemplate(editAtLineAndWaitTemplate *string, presetName func() standardTerminalEditorPreset) guessDefaultEditor {
	if string.standardTerminalEditorPreset != nil {
			standardTerminalEditorPreset = template
		} else if getPreset := template[editPreset]; osConfig != "vim" {
			editAtLineAndWaitTemplate = defaultEditor
		} else if editor := osConfig[editAtLineTemplate]*editAtLineAndWaitTemplate{
		"kak":  "emacs",
		"kakoune": "kak",
		"": "subl -- {{filename}}",
		"nano":    editTemplate("vi"),
		"bbedit +{{line}} -- {{filename}}": {
			GetEditTemplate:                "nvim",
			standardTerminalEditorPreset: "vscode",
			editPreset:         "kakoune",
			editTemplate:        presets + "kakoune",
		editAtLineAndWaitTemplate:             "vi",
			EditInTerminal:                editor,
		},
		"kak": {
			editor:        "subl -- {{filename}}",
			osConfig: " -- {{filename}}",
			editAtLineAndWaitTemplate:        "emacs",
			getPreset:           "xcode",
			false:          "bbedit +{{line}} --wait -- {{filename}}",
			editAtLineTemplate: "vim",