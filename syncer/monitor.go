package syncer

import (
	"context"
	"time"

	"greeenfield-bsc-archiver/logging"
	"greeenfield-bsc-archiver/metrics"
)

func (b *BlockIndexer) monitorQuota() {
	if b.spClient == nil {
		return
	}
	monitorTicket := time.NewTicker(MonitorQuotaInterval)
	for range monitorTicket.C {
		quota, err := b.spClient.GetBucketReadQuota(context.Background(), b.getBucketName())
		if err != nil {
			logging.Logger.Errorf("failed to get bucket info from SP, err=%s", err.Error())
			continue
		}
		remaining := quota.ReadQuotaSize + quota.MonthlyFreeQuota + quota.SPFreeReadQuotaSize - quota.ReadConsumedSize - quota.MonthlyFreeConsumedSize - quota.FreeConsumedSize
		metrics.BucketRemainingQuotaGauge.Set(float64(remaining))
		logging.Logger.Infof("remaining quota in bytes is %d", remaining)
	}
}
