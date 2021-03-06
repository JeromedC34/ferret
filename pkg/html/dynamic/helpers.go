package dynamic

import (
	"bytes"
	"context"
	"github.com/MontFerret/ferret/pkg/html/common"
	"github.com/MontFerret/ferret/pkg/html/dynamic/events"
	"github.com/MontFerret/ferret/pkg/runtime/values"
	"github.com/PuerkitoBio/goquery"
	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/protocol/dom"
	"github.com/mafredri/cdp/protocol/page"
	"golang.org/x/sync/errgroup"
	"strings"
)

type batchFunc = func() error

func runBatch(funcs ...batchFunc) error {
	eg := errgroup.Group{}

	for _, f := range funcs {
		eg.Go(f)
	}

	return eg.Wait()
}

func getRootElement(ctx context.Context, client *cdp.Client) (*dom.GetDocumentReply, error) {
	d, err := client.DOM.GetDocument(ctx, dom.NewGetDocumentArgs().SetDepth(1))

	if err != nil {
		return nil, err
	}

	return d, nil
}

func parseAttrs(attrs []string) *values.Object {
	var attr values.String

	res := values.NewObject()

	for _, el := range attrs {
		el = strings.TrimSpace(el)
		str := values.NewString(el)

		if common.IsAttribute(el) {
			attr = str
			res.Set(str, values.EmptyString)
		} else {
			current, ok := res.Get(attr)

			if ok {
				if current.String() != "" {
					res.Set(attr, current.(values.String).Concat(values.SpaceString).Concat(str))
				} else {
					res.Set(attr, str)
				}
			}
		}
	}

	return res
}

func loadInnerHTML(ctx context.Context, client *cdp.Client, id *HTMLElementIdentity) (values.String, error) {
	var args *dom.GetOuterHTMLArgs

	if id.objectID != "" {
		args = dom.NewGetOuterHTMLArgs().SetObjectID(id.objectID)
	} else if id.backendID > 0 {
		args = dom.NewGetOuterHTMLArgs().SetBackendNodeID(id.backendID)
	} else {
		args = dom.NewGetOuterHTMLArgs().SetNodeID(id.nodeID)
	}

	res, err := client.DOM.GetOuterHTML(ctx, args)

	if err != nil {
		return "", err
	}

	return values.NewString(res.OuterHTML), err
}

func parseInnerText(innerHTML string) (values.String, error) {
	buff := bytes.NewBuffer([]byte(innerHTML))

	parsed, err := goquery.NewDocumentFromReader(buff)

	if err != nil {
		return values.EmptyString, err
	}

	return values.NewString(parsed.Text()), nil
}

func createChildrenArray(nodes []dom.Node) []*HTMLElementIdentity {
	children := make([]*HTMLElementIdentity, len(nodes))

	for idx, child := range nodes {
		children[idx] = &HTMLElementIdentity{
			nodeID:    child.NodeID,
			backendID: child.BackendNodeID,
		}
	}

	return children
}

func contextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), DefaultTimeout)
}

func waitForLoadEvent(ctx context.Context, client *cdp.Client) error {
	loadEventFired, err := client.Page.LoadEventFired(ctx)

	if err != nil {
		return err
	}

	_, err = loadEventFired.Recv()

	if err != nil {
		return err
	}

	return loadEventFired.Close()
}

func createEventBroker(client *cdp.Client) (*events.EventBroker, error) {
	var err error
	var onLoad page.LoadEventFiredClient
	var onReload dom.DocumentUpdatedClient
	var onAttrModified dom.AttributeModifiedClient
	var onAttrRemoved dom.AttributeRemovedClient
	var onChildCountUpdated dom.ChildNodeCountUpdatedClient
	var onChildNodeInserted dom.ChildNodeInsertedClient
	var onChildNodeRemoved dom.ChildNodeRemovedClient
	ctx := context.Background()

	onLoad, err = client.Page.LoadEventFired(ctx)

	if err != nil {
		return nil, err
	}

	onReload, err = client.DOM.DocumentUpdated(ctx)

	if err != nil {
		onLoad.Close()
		return nil, err
	}

	onAttrModified, err = client.DOM.AttributeModified(ctx)

	if err != nil {
		onLoad.Close()
		onReload.Close()
		return nil, err
	}

	onAttrRemoved, err = client.DOM.AttributeRemoved(ctx)

	if err != nil {
		onLoad.Close()
		onReload.Close()
		onAttrModified.Close()
		return nil, err
	}

	onChildCountUpdated, err = client.DOM.ChildNodeCountUpdated(ctx)

	if err != nil {
		onLoad.Close()
		onReload.Close()
		onAttrModified.Close()
		onAttrRemoved.Close()
		return nil, err
	}

	onChildNodeInserted, err = client.DOM.ChildNodeInserted(ctx)

	if err != nil {
		onLoad.Close()
		onReload.Close()
		onAttrModified.Close()
		onAttrRemoved.Close()
		onChildCountUpdated.Close()
		return nil, err
	}

	onChildNodeRemoved, err = client.DOM.ChildNodeRemoved(ctx)

	if err != nil {
		onLoad.Close()
		onReload.Close()
		onAttrModified.Close()
		onAttrRemoved.Close()
		onChildCountUpdated.Close()
		onChildNodeInserted.Close()
		return nil, err
	}

	broker := events.NewEventBroker(
		onLoad,
		onReload,
		onAttrModified,
		onAttrRemoved,
		onChildCountUpdated,
		onChildNodeInserted,
		onChildNodeRemoved,
	)

	err = broker.Start()

	if err != nil {
		onLoad.Close()
		onReload.Close()
		onAttrModified.Close()
		onAttrRemoved.Close()
		onChildCountUpdated.Close()
		onChildNodeInserted.Close()
		onChildNodeRemoved.Close()
		return nil, err
	}

	return broker, nil
}
