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
)

func In(time time.Time, spec string, timezone *time.Location) (bool, error) {
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

      startDay := "Mon"
      endDay   := "Sun"

      switch len(match) {
      case 3:
         startDay = match[1]
         endDay   = startDay
      case 4:
         startDay = match[1]
         endDay   = match[2]
      }

      fmt.Printf("startDay=%s endDay=%s startHour=%02d endHour=%02d\n", startDay, endDay, startHour, endHour)
   }

   return false, nil
}
