# DatesI18n

A library to do quick translations of Day and Month names, including 'short' names, where provided.

## Usage

In your project directory, create a directory `lang` and put any of the language translations you plan to use in that directory.
Once you've done that:

```go
package main

import (
  "fmt"

  datesI18N "github.com/davidgs/datesI18N/v2"
)

func main() {
  fr := datesI18N.NewDatesI18N("fr") // french
  wed := fr.Weekdays.Wednesday
  fmt.Println("French Wednesday: ", wed)
}
```
Output:
```
French Wednesday: Mercredi
```

If you are using the Go `time` Package, you can use this package to give you translated Days and months.

## Language Support

**Note:** Not all languages have all parts translated! Soo the `lang/` directory for languages supported. See the individual `.json` file for which parts of the Days/Months are provided.

If you would like to provide additional translations or additions to a given translation, please submit a PR!
