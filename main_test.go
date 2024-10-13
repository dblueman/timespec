package timespec

import (
   "testing"
)

func TestNew(t *testing.T) {
   tests := []string{
      "08-20", "06-23", "Sat-Sun 10-00, Mon-Fri 08-00", "Sat 10-22, Sun 10-18, Mon-Tue 08-20, Wed-Fri 08-22",
   }

   for _, test := range tests {
      _, err := New(test)
      if err != nil {
         t.Error(err)
      }
   }
}
