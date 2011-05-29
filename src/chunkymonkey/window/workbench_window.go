package window

import (
	"chunkymonkey/inventory"
	"chunkymonkey/slot"
	. "chunkymonkey/types"
)

const (
	workbenchInvCraftNum   = 1 + inventory.WorkbenchInvCraftWidth*inventory.WorkbenchInvCraftHeight
	workbenchInvCraftStart = 0
	workbenchInvCraftEnd   = workbenchInvCraftStart + workbenchInvCraftNum

	workbenchInvMainStart = workbenchInvCraftEnd
	workbenchInvMainEnd   = workbenchInvMainStart + playerInvMainNum

	workbenchInvHoldingStart = workbenchInvMainEnd
	workbenchInvHoldingEnd   = workbenchInvHoldingStart + playerInvHoldingNum
)

type WorkbenchWindow struct {
	Window
	crafting *inventory.WorkbenchInventory
	main     *inventory.Inventory
	holding  *inventory.Inventory
}

func NewWorkbenchWindow(entityId EntityId, viewer IWindowViewer, windowId WindowId, crafting *inventory.WorkbenchInventory, main, holding *inventory.Inventory) (w *WorkbenchWindow) {
	w = &WorkbenchWindow{
		crafting: crafting,
		main:     main,
		holding:  holding,
	}
	w.Window.Init(
		windowId,
		InvTypeIdWorkbench,
		viewer,
		"Crafting",
		&w.crafting.Inventory,
		main,
		holding)

	return
}

func (w *WorkbenchWindow) Click(slotId SlotId, cursor *slot.Slot, rightClick bool, shiftClick bool) (accepted bool) {
	switch {
	case slotId < 0:
		return false
	case slotId < workbenchInvCraftEnd:
		accepted = w.crafting.Click(
			slotId-workbenchInvCraftStart,
			cursor, rightClick, shiftClick)
	case slotId < workbenchInvMainEnd:
		accepted = w.main.StandardClick(
			slotId-workbenchInvMainStart,
			cursor, rightClick, shiftClick)
	case slotId < workbenchInvHoldingEnd:
		accepted = w.holding.StandardClick(
			slotId-workbenchInvHoldingStart,
			cursor, rightClick, shiftClick)
	}
	return
}