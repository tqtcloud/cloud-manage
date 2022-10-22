package utils_test

import (
	"testing"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/tqtcloud/cloud-manage/utils"
)

func TestParseTime(t *testing.T) {
	utils.ParseDefaultSecondTime("")
}

func init() {
	zap.DevelopmentSetup()
}
