package quandl

import (
	"fmt"
	"strings"
	"time"

	"github.com/jrlmx2/stockAnalysis/utils/influxdb"
)

const (
	hour             = time.Hour
	day              = hour * 24
	week             = day * 7
	fetch            = "select LAST(\"fetchTracker\"), frequency, time from markets.commodity./.*/ group by source, endpoint"
	formatFromInflux = "2006-01-02T15:04:05Z"
)

type code struct {
	endpoint    string
	lastUpdated time.Time
	frequency   string
}

func (c *code) SetFrequency(freq string) {
	c.frequency = freq
}

func (c *code) Frequency() string {
	return c.frequency
}

func (c *code) LastUpdated() time.Time {
	return c.lastUpdated
}

func (c *code) Update(t time.Time) {
	c.lastUpdated = t
}

func (c *code) Endpoint() string {
	return c.endpoint
}

func daysIn(m time.Month, year int) time.Duration {
	return time.Duration(time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day())
}

func (c *code) isNewQuarter() bool {
	now := time.Now().UTC()
	currentQuarter := now.Month() / 3 //quarters are 0 for Q1, 1 for Q2, 2 for Q3, 3 for Q4
	lastUpdatedQuarter := c.lastUpdated.Month() / 3
	fmt.Printf("%d==%d\n", now.Month()/3, c.lastUpdated.Month()/3)
	if lastUpdatedQuarter != currentQuarter {
		var firstOfNextQuarter time.Time
		if lastUpdatedQuarter == 3 { //Q4
			firstOfNextQuarter = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
		} else {
			firstOfNextQuarter = time.Date(c.lastUpdated.Year(), ((lastUpdatedQuarter+1)*3)+1, 1, 0, 0, 0, 0, time.UTC)
		}
		return now.After(firstOfNextQuarter)
	}
	return false
}

func (c *code) ShouldGetNew() bool {
	if c.frequency == "" || c.lastUpdated.IsZero() {
		fetchLatest()
	}
	if c.frequency == "" || c.lastUpdated.IsZero() {
		return true
	}

	now := time.Now().UTC()
	distanceFromLast := time.Since(c.lastUpdated)

	switch c.frequency {
	case "daily":
		return distanceFromLast >= day
	case "weekly":
		return distanceFromLast >= week
	case "monthly":
		return distanceFromLast >= day*daysIn(now.Month()-1, now.Year()) //day * days in last month
	case "quarterly":
		return c.isNewQuarter()
	case "annually":
		return c.lastUpdated.Year() < now.Year()
	case "yearly":
		return c.lastUpdated.Year() < now.Year()
	default:
		return true
	}
}

func fetchLatest() {
	resp, err := influxdb.QueryDB(fetch)
	if err != nil {
		panic(err)
	}

	for _, c := range codes {
		for _, point := range resp {
			for _, series := range point.Series {
				if series.Tags["source"]+"/"+series.Tags["endpoint"] == strings.ToLower(c.Endpoint()) {
					c.SetFrequency(series.Values[0][2].(string))
					d, _ := time.Parse(formatFromInflux, series.Values[0][0].(string))
					//fmt.Printf("updated %s with %v, %v \n", c.Endpoint(), series.Values[0][2], series.Values[0][0])
					c.Update(d)
					break
				}
			}
		}
	}

}

var codes = [...]*code{
	&code{endpoint: "WORLDBANK/WLD_IFERTILIZERS"},
	&code{endpoint: "WORLDBANK/WLD_IFOOD"},
	&code{endpoint: "WORLDBANK/WLD_ITIMBER"},
	&code{endpoint: "WORLDBANK/WLD_IMETMIN"},
	&code{endpoint: "WORLDBANK/WLD_IBEVERAGES"},
	&code{endpoint: "WORLDBANK/WLD_IENERGY"},
	&code{endpoint: "WORLDBANK/WLD_INONFUEL"},
	&code{endpoint: "WORLDBANK/WLD_IFATS_OILS"},
	&code{endpoint: "WORLDBANK/WLD_IAGRICULTURE"},
	&code{endpoint: "WORLDBANK/WLD_IGRAINS"},
	&code{endpoint: "WORLDBANK/WLD_IRAW_MATERIAL"},
	&code{endpoint: "WORLDBANK/WLD_IOTHERRAWMAT"},
	&code{endpoint: "WORLDBANK/WLD_IOTHERFOOD"},
	&code{endpoint: "LBMA/GOLD"},
	&code{endpoint: "BMA/GOLD"},
	&code{endpoint: "WGC/GOLD_DAILY_USD"},
	&code{endpoint: "BUNDESBANK/BBK01_WT5511"},
	&code{endpoint: "CHRIS/CME_GC1"},
	&code{endpoint: "SHFE/AUV2013"},
	&code{endpoint: "WORLDBANK/WLD_GOLD"},
	&code{endpoint: "LBMA/SILVER"},
	&code{endpoint: "SHFE/AGV2013"},
	&code{endpoint: "CHRIS/CME_SI1"},
	&code{endpoint: "WORLDBANK/WLD_SILVER"},
	&code{endpoint: "JOHNMATT/PLAT"},
	&code{endpoint: "CHRIS/CME_PL1"},
	&code{endpoint: "JOHNMATT/PALL"},
	&code{endpoint: "CHRIS/CME_PA1"},
	&code{endpoint: "JOHNMATT/IRID"},
	&code{endpoint: "JOHNMATT/RHOD"},
	&code{endpoint: "JOHNMATT/RUTH"},
	&code{endpoint: "LME/PR_AL"},
	&code{endpoint: "ODA/PALUM_USD"},
	&code{endpoint: "SHFE/ALV2013"},
	&code{endpoint: "LME/PR_CO"},
	&code{endpoint: "LME/PR_CU"},
	&code{endpoint: "ODA/PCOPP_USD"},
	&code{endpoint: "CHRIS/CME_HG1"},
	&code{endpoint: "SHFE/CUV2013"},
	&code{endpoint: "ODA/PIORECR_USD"},
	&code{endpoint: "LME/PR_PB"},
	&code{endpoint: "ODA/PLEAD_USD"},
	&code{endpoint: "SHFE/PBV2013"},
	&code{endpoint: "LME/PR_MO"},
	&code{endpoint: "LME/PR_NI"},
	&code{endpoint: "ODA/PNICK_USD"},
	&code{endpoint: "LME/PR_FM"},
	&code{endpoint: "WORLDBANK/WLD_STL_JP_WIROD"},
	&code{endpoint: "SHFE/RBV2013"},
	&code{endpoint: "SHFE/WRV2013"},
	&code{endpoint: "LME/PR_TN"},
	&code{endpoint: "ODA/PTIN_USD"},
	&code{endpoint: "LME/PR_ZI"},
	&code{endpoint: "ODA/PZINC_USD"},
	&code{endpoint: "SHFE/ZNV2013"},
	&code{endpoint: "ODA/PBARL_USD"},
	&code{endpoint: "WORLDBANK/WLD_BARLEY"},
	&code{endpoint: "TFGRAIN/CORN"},
	&code{endpoint: "CHRIS/CME_C1"},
	&code{endpoint: "WORLDBANK/WLD_MAIZE"},
	&code{endpoint: "CHRIS/CME_O1"},
	&code{endpoint: "ODA/PRICENPQ_USD"},
	&code{endpoint: "CHRIS/CME_RR1"},
	&code{endpoint: "TFGRAIN/SOYBEANS"},
	&code{endpoint: "CHRIS/CME_S1"},
	&code{endpoint: "CHRIS/CME_SM1"},
	&code{endpoint: "WORLDBANK/WLD_SOYBEAN_OIL"},
	&code{endpoint: "CHRIS/CME_BO1"},
	&code{endpoint: "ODA/PWHEAMT_USD"},
	&code{endpoint: "CHRIS/CME_W1"},
	&code{endpoint: "CHRIS/CME_DA1"},
	&code{endpoint: "ODA/PBEEF_USD"},
	&code{endpoint: "CHRIS/CME_LC1"},
	&code{endpoint: "CHRIS/CME_FC1"},
	&code{endpoint: "ODA/PHIDE_USD"},
	&code{endpoint: "ODA/PPOULT_USD"},
	&code{endpoint: "ODA/PPORK_USD"},
	&code{endpoint: "CHRIS/CME_LN1"},
	&code{endpoint: "ODA/PLAMB_USD"},
	&code{endpoint: "ODA/PWOOLC_USD"},
	&code{endpoint: "ODA/PSALM_USD"},
	&code{endpoint: "ODA/PSHRI_USD"},
	&code{endpoint: "ODA/PFISH_USD"},
	&code{endpoint: "CHRIS/CME_RB1"},
	&code{endpoint: "CHRIS/CME_CL1"},
	&code{endpoint: "CHRIS/ICE_B1"},
	&code{endpoint: "DOE/I19263000008"},
	&code{endpoint: "DOE/EER_EPMRR_PF4_Y05LA_DPG"},
	&code{endpoint: "DOE/RWTC"},
	&code{endpoint: "DOE/RBRTE"},
	&code{endpoint: "DOE/I070000004"},
	&code{endpoint: "DOE/I060000004"},
	&code{endpoint: "BRP/CRUDE_OIL_PRICES"},
	&code{endpoint: "OPEC/ORB"},
	&code{endpoint: "ODA/POILDUB_USD"},
	&code{endpoint: "ODA/POILAPSP_USD"},
	&code{endpoint: "ODA/POILWTI_USD"},
	&code{endpoint: "ODA/POILBRE_USD"},
	&code{endpoint: "WORLDBANK/ECS_EP_PMP_SGAS_CD"},
	&code{endpoint: "FRED/GASMIDCOVW"},
	&code{endpoint: "FRED/DGASNYH"},
	&code{endpoint: "FRED/GASREGCOVW"},
	&code{endpoint: "FRED/DGASUSGULF"},
	&code{endpoint: "FRED/CUUR0000SETB01"},
	&code{endpoint: "FRED/GASALLCOVW"},
	&code{endpoint: "FRED/GASPRMCOVW"},
	&code{endpoint: "CHRIS/CME_NG1"},
	&code{endpoint: "CHRIS/ICE_M1"},
	&code{endpoint: "BRP/GAS_PRICES"},
	&code{endpoint: "ODA/PNGASEU_USD"},
	&code{endpoint: "ODA/PNGASUS_USD"},
	&code{endpoint: "ODA/PNGASJP_USD"},
	&code{endpoint: "FRED/GASPRICE"},
	&code{endpoint: "FRED/IR10110"},
	//&code{endpoint:"EPI/152"},
	&code{endpoint: "DOE/COAL"},
	&code{endpoint: "BRP/COAL_PRICES"},
	&code{endpoint: "INDEXMUNDI/COMMODITY_COALSOUTHAFRICANEXPORTPRICE"},
	&code{endpoint: "ODA/PCOALAU_USD"},
	&code{endpoint: "WORLDBANK/WLD_COAL_AUS"},
	&code{endpoint: "FRED/M04111GBM318NNBR"},
	&code{endpoint: "FRED/M04I1ADE00ESSM372NNBR"},
	&code{endpoint: "ODA/PCOFFOTM_USD"},
	&code{endpoint: "ODA/PCOFFROB_USD"},
	&code{endpoint: "CHRIS/ICE_KC1"},
	&code{endpoint: "ODA/PCOCO_USD"},
	&code{endpoint: "CHRIS/ICE_CC1"},
	&code{endpoint: "WORLDBANK/WLD_TOBAC_US"},
	&code{endpoint: "ODA/PTEA_USD"},
	&code{endpoint: "WORLDBANK/WLD_TEA_MOMBASA"},
	&code{endpoint: "ODA/PSUGAUSA_USD"},
	&code{endpoint: "CHRIS/ICE_SB1"},
	&code{endpoint: "ODA/PGNUTS_USD"},
	&code{endpoint: "ODA/PORANG_USD"},
	&code{endpoint: "CHRIS/ICE_OJ1"},
	&code{endpoint: "ODA/PBANSOP_USD"},
	&code{endpoint: "WORLDBANK/WLD_BANANA_US"},
	&code{endpoint: "ODA/PROIL_USD"},
	&code{endpoint: "ODA/PSUNO_USD"},
	&code{endpoint: "ODA/POLVOIL_USD"},
	&code{endpoint: "ODA/PPOIL_USD"},
	&code{endpoint: "WORLDBANK/WLD_PLMKRNL_OIL"},
	&code{endpoint: "WORLDBANK/WLD_GRNUT_OIL"},
	&code{endpoint: "ODA/PLOGSK_USD"},
	&code{endpoint: "ODA/PLOGORE_USD"},
	&code{endpoint: "WORLDBANK/WLD_COPRA"},
	&code{endpoint: "ODA/PSAWORE_USD"},
	&code{endpoint: "ODA/PSAWMAL_USD"},
	&code{endpoint: "WORLDBANK/WLD_WOODPULP"},
	&code{endpoint: "WORLDBANK/WLD_PLYWOOD"},
	&code{endpoint: "CHRIS/CME_LB1"},
	&code{endpoint: "ODA/PRUBB_USD"},
	&code{endpoint: "ODA/PCOTTIND_USD"},
	&code{endpoint: "CHRIS/ICE_CT1"},
	&code{endpoint: "WORLDBANK/WLD_UREA_EE_BULK"},
	&code{endpoint: "WORLDBANK/WLD_PHOSROCK"},
	&code{endpoint: "WORLDBANK/WLD_POTASH"},
	&code{endpoint: "WORLDBANK/WLD_TSP"},
	&code{endpoint: "WORLDBANK/WLD_DAP"},
}
