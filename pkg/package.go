package pkg

import "github.com/mhupman/go-zendesk/zendesk"

func Exported() int {
	o := zendesk.Organization{}
	return o.ExternalID
}
