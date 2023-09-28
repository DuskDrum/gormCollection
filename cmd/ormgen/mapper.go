package main

import "gorm.io/gen"

type clearingRecordMapper interface {
	// SELECT * FROM @@table WHERE channelReferenceId = @channelReferenceId
	GetByName(channelReferenceId string) ([]*gen.T, error)
}
