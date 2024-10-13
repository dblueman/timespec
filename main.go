package timespec

import (
   "errors"
   "fmt"
   "regexp"
   "strconv"
   "time"
)

type Timespec [7][2]int

var (
   ErrMalformed = errors.New("In: malformed timespec")
   specRe = regexp.MustCompile(`(?:(\w{3})(?:-(\w{3}))? )?(\d{2})-(\d{2})`)
   dayMap = map[string]time.Weekday{
      "Sun": time.Sunday,
      "Mon": time.Monday,
      "Tue": time.Tuesday,
      "Wed": time.Wednesday,
      "Thu": time.Thursday,
      "Fri": time.Friday,
      "Sat": time.Saturday,
   }
)

func New(spec string) (*Timespec, error) {
   fmt.Printf("spec %s\n", spec)

   matches := specRe.FindAllStringSubmatch(spec, -1)
   if matches == nil {
      return nil, ErrMalformed
   }

   var hours Timespec

   for _, match := range matches {
      startHour, err := strconv.Atoi(match[len(match)-2])
      if err != nil || startHour < 0 || startHour > 23 {
         return nil, ErrMalformed
      }

      endHour, err := strconv.Atoi(match[len(match)-1])
      if err != nil || endHour < 0 || endHour > 23 {
         return nil, ErrMalformed
      }

      startDayStr := "Sun"
      endDayStr   := "Sat"

      switch len(match) {
      case 3:
         startDayStr = match[1]
         endDayStr   = startDayStr
      case 4:
         startDayStr = match[1]
         endDayStr   = match[2]
      }

      startDay, ok := dayMap[startDayStr]
      if !ok {
         return nil, ErrMalformed
      }

      endDay, ok := dayMap[endDayStr]
      if !ok {
         return nil, ErrMalformed
      }

      for day := startDay; day <= endDay; day = (day + 1) % 8 {
fmt.Printf("startDay=%d endDay=%d\n", startDay, endDay)
         hours[day][0] = startHour
         hours[day][1] = endHour
      }

      fmt.Printf("hours=%+v\n", hours)
   }

   return &hours, nil
}

func (hours *Timespec) In(t *time.Time, timezone *time.Location) bool {
   return false
}
