package manager

import (
	"fmt"
	"github.com/onepanelio/core/util"
	"github.com/onepanelio/core/util/logging"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"

	"github.com/onepanelio/core/model"
)

var onepanelEnabledLabelKey = labelKeyPrefix + "enabled"

func (r *ResourceManager) ListNamespaces() (namespaces []*model.Namespace, err error) {
	namespaces, err = r.kubeClient.ListNamespaces(model.ListOptions{
		LabelSelector: fmt.Sprintf("%s=%s", onepanelEnabledLabelKey, "true"),
	})
	if err != nil {
		logging.Logger.Log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Error("ListNamespaces failed.")
		err = util.NewUserError(codes.Unknown, "List namespaces failed.")
	}
	return
}
