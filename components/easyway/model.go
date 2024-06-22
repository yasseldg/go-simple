package easyway

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/types/sTime"

	"github.com/yasseldg/mgm/v4"
)

type InterEasyWay interface {
	String(ew_type string) string
	Log(ew_type string)

	Ts() int64
	Ew(ew_type string) InterEwType
}

type EwMap map[string]*EwType

type EasyWay struct {
	mgm.DefaultModel

	M_ts  int64 `bson:"ts" json:"ts"`
	M_map EwMap `bson:"ew" json:"ew"`
}
type EasyWays []*EasyWay

func (ew *EasyWay) String(ew_type string) string {

	msg := fmt.Sprintf("EW %s: ts: %s", ew_type, sTime.ForLog(ew.M_ts, 0))

	_ew := ew.Ew(ew_type)
	if _ew == nil {
		return fmt.Sprintf("%s  ..  failed to get ew.", msg)

	}

	return fmt.Sprintf("%s  ..  %s", msg, _ew.String())
}

func (ew *EasyWay) Log(ew_type string) {

	sLog.Info(ew.String(ew_type))
}

func (ew *EasyWay) Ts() int64 {
	return ew.M_ts
}

func (ew *EasyWay) Ew(ew_type string) InterEwType {
	if ewType, ok := ew.M_map[ew_type]; ok {
		return ewType
	}
	return nil
}

// {
//   "_id": {
//     "$oid": "665df6a5a1f043c3ccb07283"
//   },
//   "ts": 1717433700,
//   "ew": {
//     "t5000": {
//       "b": {
//         "h": 68298.5,
//         "t": 67957.974
//       },
//       "s": {
//         "h": 70377.5,
//         "t": 70871.6008
//       }
//     }
//   }
// }
