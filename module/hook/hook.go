package hook

import (
	"sync"

	"github.com/insionng/prior"
)

var (
	QueuesMap  = new(sync.Map) //map[string]*pueue.PriorityQueue{}
	FiltersMap = new(sync.Map) //map[string][]byte{}

	// DefaultPriority 默认优先级为0值
	DefaultPriority int
	callback        func([]byte) []byte
)

// NewPriorityQueue New PriorityQueue
func NewPriorityQueue() *prior.PriorityQueue {
	return prior.NewPriorityQueue()
}

// SetPriorityQueueWith c.makross.QueuesMap[key] = NewPriorityQueue()
func SetPriorityQueueWith(key interface{}) *sync.Map {
	if QueuesMap == nil {
		QueuesMap = new(sync.Map)
	}
	QueuesMap.Store(key, NewPriorityQueue())
	return QueuesMap
}

// RemoveFilterHook  删除过滤钩子
func RemoveFilterHook(key string) {
	if HasFilterHook(key) {
		FiltersMap.Delete(key)
	}
}

// RemoveActionHook  删除动作钩子
func RemoveActionHook(key string) {
	if HasActionHook(key) {
		QueuesMap.Delete(key)
		FiltersMap.Delete(key)
	}
}

// RemoveActionsHook  删除所有动作钩子
func RemoveActionsHook() {
	QueuesMap = nil
}

// HasFilterHook  是否有过滤钩子
func HasFilterHook(key string) bool {
	if _, okay := FiltersMap.Load(key); okay {
		return true
	}
	return false
}

// HasQueuesMap  Has QueuesMap
func HasQueuesMap(key string) bool {
	if value, okay := QueuesMap.Load(key); okay {
		if pqueue, okay := value.(*prior.PriorityQueue); okay {
			if pqueue.Length() > 0 {
				return true
			}
		}
	}
	return false
}

// HasActionHook  是否有动作钩子
func HasActionHook(key string) bool {
	if value, okay := QueuesMap.Load(key); okay {
		if _, okay := value.(*prior.PriorityQueue); okay {
			return true
		}
	}
	return false
}

// AddActionHook  增加动作钩子
func AddActionHook(key string, function func(), priorities ...int) {
	AddFilterHook(key, func([]byte) []byte {
		function()
		return nil
	}, priorities...)
}

// AddFilterHook  增加过滤钩子
func AddFilterHook(key string, function func([]byte) []byte, priorities ...int) {

	if !HasQueuesMap(key) {
		SetPriorityQueueWith(key)
	}

	var priority int
	if len(priorities) > 0 {
		priority = priorities[0]
	} else {
		priority = DefaultPriority
	}
	pq := NewPriorityQueue()
	pq.Push(prior.NewNode(key, function, priority))
	QueuesMap.Store(key, pq)

}

// DoActionHook  动作钩子
func DoActionHook(key string) {
	DoFilterHook(key, nil)
}

// DoFilterHook  执行过滤钩子
func DoFilterHook(key string, function func() []byte) []byte {
	if !HasQueuesMap(key) {
		AddFilterHook(key, func(b []byte) []byte {
			if function == nil {
				return b
			}
			return function()
		})
		return DoFilterHook(key, function)
	}
	if !HasFilterHook(key) {
		if FiltersMap == nil {
			FiltersMap = new(sync.Map) //make(map[string][]byte)
		}
	}

	if HasActionHook(key) {
		if queue, okay := QueuesMap.Load(key); okay {
			pq, okay := queue.(*prior.PriorityQueue)
			if !okay {
				return nil
			}
			for pq.Length() > 0 {
				if node := pq.Pop(); node != nil {
					if value := node.GetValue(); value != nil {
						if callback, okay = value.(func([]byte) []byte); !okay {
							return nil
						}
					} else {
						continue
					}
				} else {
					continue
				}

				if function != nil { //for Global Filter
					if value, okay := FiltersMap.Load(key); okay {
						if b, okay := value.([]byte); okay {
							FiltersMap.Store(key, callback(b))
						}
					} else {
						FiltersMap.Store(key, callback(function()))
					}
				} else { //for Global Action
					FiltersMap.Store(key, callback(nil))
				}

			}
		}
	}

	if value, okay := FiltersMap.Load(key); okay {
		if b, okay := value.([]byte); okay {
			return b
		}
	}
	return nil
}

// RegisterActivationHook Set the activation hook for a plugin.
/* 启用插件
 * When a plugin is activated, the action 'activate_PLUGINNAME' hook is
 * called. In the name of this hook, PLUGINNAME is replaced with the name
 * of the plugin, including the optional subdirectory. For example, when the
 * plugin is located in wp-content/plugin/sampleplugin/sample.app, then
 * the name of this hook will become 'activate_sampleplugin/sample.app'.
 *
 * When the plugin consists of only one file and is (as by default) located at
 * wp-content/plugin/sample.app the name of this hook will be
 * 'activate_sample.app'.
 *
 * @param string   key    The filename of the plugin including the path.
 * @param callable function The function hooked to the 'activate_PLUGIN' action.
 */
func RegisterActivationHook(function func(), key ...string) {
	var iKey = "activate"
	if len(key) > 0 {
		iKey = iKey + "_" + key[0]
	}
	AddActionHook(iKey, function)
}

// RegisterDeactivationHook Set the deactivation hook for a plugin.
/* 禁用插件
 * When a plugin is deactivated, the action 'deactivate_PLUGINNAME' hook is
 * called. In the name of this hook, PLUGINNAME is replaced with the name
 * of the plugin, including the optional subdirectory. For example, when the
 * plugin is located in content/plugin/sampleplugin/sample.app, then
 * the name of this hook will become 'deactivate_sampleplugin/sample.app'.
 *
 * When the plugin consists of only one file and is (as by default) located at
 * content/plugin/sample.app the name of this hook will be
 * 'deactivate_sample.app'.
 *
 *
 * @param string   key     The filename of the plugin including the path.
 * @param callable function The function hooked to the 'deactivate_PLUGIN' action.
 */
func RegisterDeactivationHook(function func(), key ...string) {
	var iKey = "deactivate"
	if len(key) > 0 {
		iKey = iKey + "_" + key[0]
	}
	AddActionHook(iKey, function)
}

// RegisterUninstallHook 卸载插件
func RegisterUninstallHook(function func(), key ...string) {
	var iKey = "uninstall"
	if len(key) > 0 {
		iKey = iKey + "_" + key[0]
	}
	AddActionHook(iKey, function)
}

//***********************************************************************//

/**
 * Add a submenu page.
 *
 * This function takes a capability which will be used to determine whether
 * or not a page is included in the menu.
 *
 * The function which is hooked in to handle the output of the page must check
 * that the user has the required capability as well.
 *
 * @global arraysubmenu
 * @global arraymenu
 * @global array_wp_real_parent_file
 * @global bool _wp_submenu_nopriv
 * @global array_registered_pages
 * @global array_parent_pages
 *
 * @param string  parent_slug The slug name for the parent menu (or the file name of a standard WordPress admin page).
 * @param string  page_title  The text to be displayed in the title tags of the page when the menu is selected.
 * @param string  menu_title  The text to be used for the menu.
 * @param string  capability  The capability required for this menu to be displayed to the user.
 * @param string  menu_slug   The slug name to refer to this menu by (should be unique for this menu).
 * @param callable function    The function to be called to output the content for this page.
 * @return false|string The resulting page's hook_suffix, or false if the user does not have the capability required.
 */
/*
func  AddSubmenuPage(parent_slug,page_title,menu_title,capability,menu_slug string,function ...func() ) {

	if ( !current_user_can(capability ) ) {
		_wp_submenu_nopriv[parent_slug][menu_slug] = true;
		return false;
	}


	 // If the parent doesn't already have a submenu, add a link to the parent
	 // as the first item in the submenu. If the submenu file is the same as the
	 // parent file someone is trying to link back to the parent manually. In
	 // this case, don't automatically add a link back to avoid duplication.

	if (!isset(submenu[parent_slug] ) &&menu_slug !=parent_slug ) {
		foreach ( (array)menu asparent_menu ) {
			if (parent_menu[2] ==parent_slug && current_user_can(parent_menu[1] ) )
				submenu[parent_slug][] = array_slice(parent_menu, 0, 4 );
		}
	}

	submenu[parent_slug][] = array (menu_title,capability,menu_slug,page_title );

	hookname :=  menu_slug+parent_slug;

if (len(function)>0)&&(len(hookname)>0) {
	AddAction(hookname,function[0])
}
	_registered_pages[hookname] = true;

	// No parent as top level.
	_parent_pages[menu_slug] =parent_slug;

	return hookname;
}
*/

/**
 * Add submenu page to the Settings main menu.
 *
 * This function takes a capability which will be used to determine whether
 * or not a page is included in the menu.
 *
 * The function which is hooked in to handle the output of the page must check
 * that the user has the required capability as well.
 *
 * @param string   page_title The text to be displayed in the title tags of the page when the menu is selected.
 * @param string   menu_title The text to be used for the menu.
 * @param string   capability The capability required for this menu to be displayed to the user.
 * @param string   menu_slug  The slug name to refer to this menu by (should be unique for this menu).
 * @param callable function   The function to be called to output the content for this page.
 * @return false|string The resulting page's hook_suffix, or false if the user does not have the capability required.
 */
/*
func AddOptionsPage(page_title, menu_title, capability, menu_slug string, function ...func()) {
	return AddSubmenuPage("options-general", page_title, menu_title, capability, menu_slug, function...)
}
*/
