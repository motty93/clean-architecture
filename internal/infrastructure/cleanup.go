package infrastructure

import "log"

type CleanupManager struct {
	cleanups []func() error
}

func NewCleanupManager() *CleanupManager {
	return &CleanupManager{}
}

func (cm *CleanupManager) Add(fn func() error) {
	cm.cleanups = append(cm.cleanups, fn)
}

func (cm *CleanupManager) Execute() {
	for _, cleanup := range cm.cleanups {
		if err := cleanup(); err != nil {
			log.Printf("Error: %v\n", err)
		}
	}
	log.Println("All cleanup tasks completed.")
}
