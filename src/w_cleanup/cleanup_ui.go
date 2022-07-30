package w_cleanup

import (
	"auto_upload/src/secrets"
	"auto_upload/src/util"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type (
	cleanupui struct {
		post_upload_dropdown *widget.Select
		archive_dir_label    *widget.Label
		archive_dir_entry    *widget.Entry

		playlist_enabled       bool
		playlist_enabled_check *widget.Check
		playlist_label         *widget.Label
		playlist_entry         *widget.Entry
	}
)

func NewCleanupUI() *cleanupui {

	post_upload_options := []string{
		"move", "delete", "noting",
	}

	cleanup_widgets := cleanupui{
		archive_dir_label: widget.NewLabel("VOD Archive Directory:\t"),
		archive_dir_entry: util.PopulatedEntry(&secrets.Config().Files.VOD_ARCHIVE_DIR),

		playlist_enabled: secrets.Config().Upload.PLAYLIST_ID != "",
		playlist_label:   widget.NewLabel("Playlist ID:          \t"),
		playlist_entry:   util.PopulatedEntry(&secrets.Config().Upload.PLAYLIST_ID),
	}
	cleanup_widgets.post_upload_dropdown = widget.NewSelect(post_upload_options, cleanup_widgets.updateVODArchiveOptions)
	cleanup_widgets.post_upload_dropdown.SetSelected(secrets.Config().Files.POST_UPLOAD_STEP)
	cleanup_widgets.updateVODArchiveOptions(secrets.Config().Files.POST_UPLOAD_STEP)

	cleanup_widgets.playlist_enabled_check = widget.NewCheck("", cleanup_widgets.updatePlaylistCheck)
	cleanup_widgets.playlist_enabled_check.SetChecked(cleanup_widgets.playlist_enabled)
	cleanup_widgets.updatePlaylistCheck(cleanup_widgets.playlist_enabled)

	return &cleanup_widgets
}

func (widgets cleanupui) updateVODArchiveOptions(step string) {
	if step != "move" {
		widgets.archive_dir_label.Hide()
		widgets.archive_dir_entry.Hide()
	} else {
		widgets.archive_dir_label.Show()
		widgets.archive_dir_entry.Show()
	}
}

func (widgets cleanupui) updatePlaylistCheck(enabled bool) {
	if !enabled {
		widgets.playlist_label.Hide()
		widgets.playlist_entry.Hide()
	} else {
		widgets.playlist_label.Show()
		widgets.playlist_entry.Show()
	}
}

func start_cleanup(ui_context util.UIContext, cleanup_mode string, playlist_enabled bool) {
	playlist_label := widget.NewLabel("Adding to Playlist: in progress...")
	playlist_label.Hide()
	cleanup_label := widget.NewLabel("Cleanup Step: " + cleanup_mode + ": in progress...")
	ui_context.SetContent(
		nil,
		container.NewVBox(
			playlist_label,
			cleanup_label,
		),
	)
	if playlist_enabled {
		playlist_label.Show()
		go add_to_playlist(playlist_label)
	}
	go cleanup_step(cleanup_mode, cleanup_label)
}

func UI_cleanup(ui_context util.UIContext) {
	cleanup_widgets := NewCleanupUI()

	ui_context.SetContent(
		func() {
			start_cleanup(
				ui_context,
				cleanup_widgets.post_upload_dropdown.Selected,
				cleanup_widgets.playlist_enabled_check.Checked,
			)
		},
		widget.NewLabel("Confirm cleanup steps"),
		container.New(layout.NewFormLayout(),
			container.NewVBox(
				widget.NewLabel("Post-Upload Step:     \t"),
				cleanup_widgets.archive_dir_label,
				widget.NewLabel("Add to Playlist:      \t"),
				cleanup_widgets.playlist_label,
			),
			container.NewVBox(
				cleanup_widgets.post_upload_dropdown,
				cleanup_widgets.archive_dir_entry,
				cleanup_widgets.playlist_enabled_check,
				cleanup_widgets.playlist_entry,
			),
		),
	)
}
