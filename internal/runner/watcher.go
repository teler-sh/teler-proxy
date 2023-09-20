package runner

import "github.com/fsnotify/fsnotify"

type watcher struct {
	config, datasets *fsnotify.Watcher
}
