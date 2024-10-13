package timespec

import (
   "errors"
   "fmt"
   "regexp"
   "strconv"
   "time"
)

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

func In(t time.Time, spec string, timezone *time.Location) (bool, error) {
   fmt.Printf("spec %s\n", spec)
   matches := specRe.FindAllStringSubmatch(spec, -1)
   if matches == nil {
      return false, ErrMalformed
   }

   for _, match := range matches {
      startHour, err := strconv.Atoi(match[len(match)-2])
      if err != nil || startHour < 0 || startHour > 23 {
         return false, ErrMalformed
      }

      endHour, err := strconv.Atoi(match[len(match)-1])
      if err != nil || endHour < 0 || endHour > 23 {
         return false, ErrMalformed
      }

      startDayStr := "Mon"
      endDayStr   := "Sun"

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
         return false, ErrMalformed
      }

      endDay, ok := dayMap[endDayStr]
      if !ok {
         return false, ErrMalformed
      }

      fmt.Printf("startDay=%d endDay=%d startHour=%02d endHour=%02d\n", startDay, endDay, startHour, endHour)
   }

   return false, nil
}
