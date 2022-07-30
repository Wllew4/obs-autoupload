package w_cleanup

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/widget"
)

func add_to_playlist(label *widget.Label) {
	fmt.Println("Adding here :3")
	time.Sleep(5 * time.Second)
	label.SetText("Adding to Playlist: Complete!")
}

func cleanup_step(mode string, label *widget.Label) {

	label.SetText("Cleanup Step: " + mode + ": Complete!")
}
