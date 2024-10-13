package timespec

import (
   "errors"
   "fmt"
   "regexp"
   "strconv"
   "strings"
   "time"
)

type Timespec [7][2]int

var (
   ErrMalformed = errors.New("Malformed timespec")
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
   matches := specRe.FindAllStringSubmatch(spec, -1)
   if matches == nil {
      return nil, ErrMalformed
   }

   var hours Timespec

   for _, match := range matches {
      if len(match) != 5 {
         return nil, ErrMalformed
      }

      startHour, err := strconv.Atoi(match[len(match)-2])
      if err != nil || startHour < 0 || startHour > 23 {
         return nil, ErrMalformed
      }

      endHour, err := strconv.Atoi(match[len(match)-1])
      if err != nil || endHour < 0 || endHour > 23 {
         return nil, ErrMalformed
      }

      startDayStr := match[1]
      endDayStr   := match[2]

      if startDayStr == "" {
         if endDayStr == "" {
            startDayStr = "Sun"
            endDayStr = "Sat"
         }
      } else {
         if endDayStr == startDayStr {
            return nil, ErrMalformed
         }

         if endDayStr == "" {
            endDayStr = startDayStr
         }
      }

      startDay, ok := dayMap[startDayStr]
      if !ok {
         return nil, ErrMalformed
      }

      endDay, ok := dayMap[endDayStr]
      if !ok {
         return nil, ErrMalformed
      }
      day := startDay
      for {
         hours[day][0] = startHour
         hours[day][1] = endHour

         if day == endDay {
            break
         }

         day = (day + 1) % 7
      }
   }

   return &hours, nil
}

func (hours *Timespec) String() (string) {
   var elems []string

   for i, day := range hours {
      elems = append(elems, fmt.Sprintf("%s %02d-%02d", time.Weekday(i), day[0], day[1]))
   }

   return strings.Join(elems, ", ")
}

func (hours *Timespec) In(t *time.Time, timezone *time.Location) bool {
   return false
}
