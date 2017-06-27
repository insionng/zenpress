package hook

import (
	"github.com/insionng/zenpress/module/hook"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/insionng/zenpress/module/hook",

	"DefaultPriority": hook.DefaultPriority,
	"FiltersMap":      hook.FiltersMap,
	"QueuesMap":       hook.QueuesMap,

	"AddActionHook":            hook.AddActionHook,
	"AddFilterHook":            hook.AddFilterHook,
	"DoActionHook":             hook.DoActionHook,
	"DoFilterHook":             hook.DoFilterHook,
	"HasActionHook":            hook.HasActionHook,
	"HasFilterHook":            hook.HasFilterHook,
	"HasQueuesMap":             hook.HasQueuesMap,
	"NewPriorityQueue":         hook.NewPriorityQueue,
	"RegisterActivationHook":   hook.RegisterActivationHook,
	"RegisterDeactivationHook": hook.RegisterDeactivationHook,
	"RegisterUninstallHook":    hook.RegisterUninstallHook,
	"RemoveActionHook":         hook.RemoveActionHook,
	"RemoveActionsHook":        hook.RemoveActionsHook,
	"RemoveFilterHook":         hook.RemoveFilterHook,
	"SetPriorityQueueWith":     hook.SetPriorityQueueWith,
}
