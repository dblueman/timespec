package timespec

import (
   "testing"
)

func TestNew(t *testing.T) {
   tests := []struct{
      in, expected string
   }{
      {"08-20", "Sunday 08-20, Monday 08-20, Tuesday 08-20, Wednesday 08-20, Thursday 08-20, Friday 08-20, Saturday 08-20"},
      {"06-23", "Sunday 06-23, Monday 06-23, Tuesday 06-23, Wednesday 06-23, Thursday 06-23, Friday 06-23, Saturday 06-23"},
      {"Sat-Sun 10-00, Mon-Fri 08-00", "Sunday 10-00, Monday 08-00, Tuesday 08-00, Wednesday 08-00, Thursday 08-00, Friday 08-00, Saturday 10-00"},
      {"Sat 10-22, Sun 10-18, Mon-Tue 08-20, Wed-Fri 08-22", "Sunday 10-18, Monday 08-20, Tuesday 08-20, Wednesday 08-22, Thursday 08-22, Friday 08-22, Saturday 10-22"},
   }

   for _, test := range tests {
      hours, err := New(test.in)
      if err != nil {
         t.Error(err)
         continue
      }

      out := hours.String()

      if out != test.expected {
         t.Errorf("got %s but expected %s", out, test.expected)
      }
   }
}
