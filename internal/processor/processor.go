package processor

import "golang.org/x/net/html"

type ChildNodesProcessor interface {
	ProcessChildNodes(*html.Node) error
}
