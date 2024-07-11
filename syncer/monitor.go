package syncer

import (
	"context"
	"time"

	"greeenfield-bsc-archiver/logging"
	"greeenfield-bsc-archiver/metrics"
)

func (s *BlockIndexer) monitorQuota() {
	if s.spClient == nil {
		return
	}
	monitorTicket := time.NewTicker(MonitorQuotaInterval)
	for range monitorTicket.C {
		quota, err := s.spClient.GetBucketReadQuota(context.Background(), s.getBucketName())
		if err != nil {
			logging.Logger.Errorf("failed to get bucket info from SP, err=%s", err.Error())
			continue
		}
		remaining := quota.ReadQuotaSize + quota.MonthlyFreeQuota + quota.SPFreeReadQuotaSize - quota.ReadConsumedSize - quota.MonthlyFreeConsumedSize - quota.FreeConsumedSize
		metrics.BucketRemainingQuotaGauge.Set(float64(remaining))
		logging.Logger.Infof("remaining quota in bytes is %d", remaining)
	}
}
