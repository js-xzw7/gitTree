package helper

import (
	"time"
)

/**
*  TimeHelper
*  @Description:
 */
type TimeHelper interface {
}

type defaultTimeHelper struct {
}

func NewTimeHelper() defaultTimeHelper {
	return defaultTimeHelper{}
}

const (
	TimeEnumWorkHoursStartCommon = "09:30:00"
	TimeEnumWorkHoursEndCommon   = "19:00:00"

	TimeEnumWorkHoursStartUnusual = "10:00:00"
	TimeEnumWorkHoursEndUnusual   = "17:30:00"

	//下午1点
	TimeEnumWorkHourPmsStartCommon = "13:00:00"
	TimeEnumWorkHourAmEndCommon    = "12:00:00"

	//
	TimeEnumStartTime      = "00:00:00"
	TimeEnumEmptyStartTime = " 00:00:00"
	TimeEnumEndTime        = "23:59:59"
	TimeEnumEmptyEndTime   = " 23:59:59"

	//日期转换格式
	TimeEnumDateTimeLayout   = "2006-01-02 15:04:05"
	TimeEnumDateLayout       = "2006-01-02"
	TimeEnumDateLayoutAll    = "2006-01-02 00:00:00"
	TimeEnumDateLayoutNumber = "20060102"
	TimeEnumTimeLayout       = "15:04:05"
)

/**
 *  GetNowTimestamp
 *  @Description: 获取当前时间戳，精确到秒
 *  @receiver i
 *  @return int32
 */
func (i defaultTimeHelper) GetNowTimestamp() int64 {
	nowTimestamp := int64(time.Now().Unix())
	return nowTimestamp
}

/**
 *  GetNowDateTime
 *  @Description: 获取当前时间
 *  @receiver i
 *  @return string
 */
func (i defaultTimeHelper) GetNowDateTime() string {
	nowTimestamp := i.GetNowTimestamp()
	datetime := i.TimestampToDatetime(int64(nowTimestamp))
	return datetime
}

/**
 *  getDateTimeLayout
 *  @Description: 获取DateTime时间格式
 *  @receiver i
 *  @return string
 */
func (i defaultTimeHelper) GetDateTimeLayout() string {
	return TimeEnumDateTimeLayout
}

/**
 *  getDateTimeLayout
 *  @Description: 获取DateTime时间格式
 *  @receiver i
 *  @return string
 */
func (i defaultTimeHelper) GetTimeLayout() string {
	return TimeEnumTimeLayout
}

/**
 *  GetDateLayout
 *  @Description: 获取日期部分
 *  @receiver i
 *  @return string
 */
func (i defaultTimeHelper) GetDateLayout() string {
	return TimeEnumDateLayout
}

/**
 *  TimestampToDatetime
 *  @Description: 时间戳转 dateTime
 *  @receiver i
 *  @param timestamp
 *  @return string
 */
func (i defaultTimeHelper) TimestampToDatetime(timestamp int64) string {
	datetime := time.Unix(timestamp, 0).Format(i.GetDateTimeLayout())
	return datetime
}

/**
 *  DatetimeToTimestamp
 *  @Description: dateTime转时间戳
 *  @receiver i
 *  @param dataTime
 *  @return int64
 */
func (i defaultTimeHelper) GetTimestampByDatetime(dataTime string) int64 {
	Loc, _ := time.LoadLocation("Asia/Shanghai")
	t2, _ := time.ParseInLocation(i.GetDateTimeLayout(), dataTime, Loc)
	return t2.Unix()
}

/**
 *  DatetimeToDate
 *  @Description: 2006-01-02 00:00:00
 *  @receiver i
 *  @param dataTime
 */
func (i defaultTimeHelper) DatetimeToDate(dataTime string) string {
	Loc, _ := time.LoadLocation("Asia/Shanghai")
	t2, _ := time.ParseInLocation(i.GetDateTimeLayout(), dataTime, Loc)
	return t2.Format("2006-01-02 00:00:00")
}

/**
 *  GetYear
 *  @Description:
 *  @receiver i
 *  @param times
 *  @return int
 */
func (i defaultTimeHelper) GetYear(times string) int {
	s, _ := time.Parse(i.GetDateTimeLayout(), times)
	return s.Year()
}

/**
 *  GetMonth
 *  @Description:
 *  @receiver i
 *  @param times
 *  @return int
 */
func (i defaultTimeHelper) GetMonth(times string) int {
	s, _ := time.Parse(i.GetDateTimeLayout(), times)
	return int(s.Month())
}

/**
 *  GetDay
 *  @Description:
 *  @receiver i
 *  @param times
 *  @return int
 */
func (i defaultTimeHelper) GetDay(times string) int {
	s, _ := time.Parse(i.GetDateTimeLayout(), times)
	return s.Day()
}

/**
 *  GetDate
 *  @Description:
 *  @receiver i
 *  @param times
 *  @return int
 *  @return time.Month
 *  @return int
 */
func (i defaultTimeHelper) GetDate(times string) (int, time.Month, int) {
	s, _ := time.Parse(i.GetDateTimeLayout(), times)
	return s.Date()
}

/**
 *  GetYearDay
 *  @Description:
 *  @receiver i
 *  @param times
 *  @return int
 */
func (i defaultTimeHelper) GetYearDay(times string) int {
	s, _ := time.Parse(i.GetDateTimeLayout(), times)
	return s.YearDay()
}

/**
 *  GetMonthStartEnd
 *  @Description: 获取本月的开始时间结束时间
 *  @receiver i
 *  @param t
 *  @return time.Time
 *  @return time.Time
 */
func (i defaultTimeHelper) GetMonthStartEndCommon(t time.Time) (time.Time, time.Time) {
	monthStartDay := t.AddDate(0, 0, -t.Day()+1)
	monthStartTime := time.Date(monthStartDay.Year(), monthStartDay.Month(), monthStartDay.Day(), 0, 0, 0, 0, t.Location())
	monthEndDay := monthStartTime.AddDate(0, 1, -1)
	monthEndTime := time.Date(monthEndDay.Year(), monthEndDay.Month(), monthEndDay.Day(), 23, 59, 59, 0, t.Location())
	return monthStartTime, monthEndTime

	//currentYear,currentMonth,_ := now.Date()
	//currentLocation := now.Location()
	//firstOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, currentLocation)//月初开始
	//EndOfMonth := time.Date(now.Year(), now.Month(), 1, 23, 59, 59, 999999999, currentLocation)//月初开始
	//lastOfMonth := EndOfMonth.AddDate(0, 1, -1)//月底结束
	//fmt.Println("月初",firstOfMonth)
	//fmt.Println("月末",lastOfMonth)
}

/**
 *  GetWeekStartDateTime
 *  @Description: 获取周开始和结束时间
 *  @receiver i
 *  @param nowTime
 *  @return time.Time
 *  @return time.Time
 */
func (i defaultTimeHelper) GetWeekStartAndEnd(nowTime time.Time) (time.Time, time.Time) {
	week := nowTime.Weekday()
	offset := int(time.Monday - week)
	if offset > 0 {
		offset = -6
	}
	weekStart := nowTime.AddDate(0, 0, offset).Format(TimeEnumDateLayout) + TimeEnumEmptyStartTime
	weekEnd := nowTime.AddDate(0, 0, offset).AddDate(0, 0, 6).Format(TimeEnumDateLayout) + TimeEnumEmptyEndTime
	Loc, _ := time.LoadLocation("Asia/Shanghai")
	start, _ := time.ParseInLocation(i.GetDateTimeLayout(), weekStart, Loc)
	end, _ := time.ParseInLocation(i.GetDateTimeLayout(), weekEnd, Loc)
	return start, end
}

/**
 *  GetTimeToString
 *  @Description: time格式转换成string格式
 *  @receiver i
 *  @param time.timne
 *  @return string
 */
func (i defaultTimeHelper) GetTimeToString(t time.Time) string {
	datetime := t.Format(i.GetDateTimeLayout())
	return datetime
}
