package rpccontext

import (
	"github.com/Metchain/Metblock/appmessage"
	"github.com/Metchain/Metblock/mconfig/dagconfig"
	"github.com/Metchain/Metblock/protoserver/routerpkg"
	"sync"
)

// NotificationManager manages notifications for the RPC
type NotificationManager struct {
	sync.RWMutex
	listeners map[*routerpkg.Router]*NotificationListener
	params    *dagconfig.Params
}

// NotificationListener represents a registered RPC notification listener
type NotificationListener struct {
	params *dagconfig.Params

	propagateBlockAddedNotifications                            bool
	propagateVirtualSelectedParentChainChangedNotifications     bool
	propagateFinalityConflictNotifications                      bool
	propagateFinalityConflictResolvedNotifications              bool
	propagateUTXOsChangedNotifications                          bool
	propagateVirtualSelectedParentBlueScoreChangedNotifications bool
	propagateVirtualDaaScoreChangedNotifications                bool
	propagatePruningPointUTXOSetOverrideNotifications           bool
	propagateNewBlockTemplateNotifications                      bool

	includeAcceptedTransactionIDsInVirtualSelectedParentChainChangedNotifications bool
}

// NotifyNewBlockTemplate notifies the notification manager that a new
// block template is available for miners
func (nm *NotificationManager) NotifyNewBlockTemplate(
	notification *appmessage.NewBlockTemplateNotificationMessage) error {

	nm.RLock()
	defer nm.RUnlock()

	for router, listener := range nm.listeners {
		if listener.propagateNewBlockTemplateNotifications {
			err := router.OutgoingRoute().Enqueue(notification)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
