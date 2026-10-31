# systemone

A small Go client for SystemOne-style classification models, such as Cloudflare Workers AI's `clef-flash`.

You give the model some text (the *state*) and a set of named questions. It answers each one with a probability, a choice, or a score.

## Install

```sh
go get github.com/MarcelArt/systemone
```

## Configuration

`systemone.New()` reads these environment variables (see [example.env](example.env)):

| Variable             | Description                                  |
| -------------------- | -------------------------------------------- |
| `SYSTEMONE_MODEL`    | Model name sent in the request body          |
| `SYSTEMONE_BASE_URL` | Full URL of the run endpoint                 |
| `SYSTEMONE_API_KEY`  | Bearer token                                 |

For Cloudflare Workers AI the URL looks like:

```
https://api.cloudflare.com/client/v4/accounts/<ACCOUNT_ID>/ai/run/@cf/cloudflare/clef-flash
```

You can also configure the client in code. Options override the environment:

```go
s1 := systemone.New(
    systemone.WithModel("clef-flash"),
    systemone.WithBaseURL("https://..."),
    systemone.WithAPIKey("..."),
)
```

## Usage

```go
package main

import (
    "context"
    "log"

    "github.com/MarcelArt/systemone"
)

func main() {
    ctx := context.Background()

    s1 := systemone.New()
    defer s1.Close()

    res, err := s1.Post(ctx, systemone.Body{
        Model: s1.Model,
        State: "Checkout has been failing for every customer for the last hour.",
        Questions: systemone.Questions{
            "urgent": systemone.Noul("Is this support request urgent?", nil),
            "team": systemone.Choice(
                "Which team should handle this request?",
                systemone.ChoiceCriteria{
                    "billing":   "Payments, invoices, and refunds",
                    "sales":     "Plans and upgrades",
                    "technical": "Outages, errors, and configuration",
                },
            ),
            "severity": systemone.Score(
                "How severe is the customer impact?",
                systemone.ScoreCriteria{"No impact", "Minor", "Major", "Critical"},
            ),
        },
    })
    if err != nil {
        log.Fatalln(err)
    }

    log.Println(res)
}
```

A runnable version is in [examples/clef](examples/clef). Copy `example.env` to `.env`, fill it in, then run `make clef`.

### Question types

| Constructor                      | Answer                                                   |
| -------------------------------- | -------------------------------------------------------- |
| `Noul(instructions, criteria)`   | Probability (0 to 1) that the statement is true. `criteria` may be `nil` |
| `Choice(instructions, criteria)` | One key of `criteria`, plus a probability for each key   |
| `Score(instructions, criteria)`  | A score from 0 to `len(criteria)-1`, plus a probability per level |

### Reading the response

`Post` returns the decoded JSON body as a `map[string]any`, exactly as the endpoint sent it. With Cloudflare the answers are under `result.answers`, keyed by the names you used in `Questions`:

```json
{
  "result": {
    "answers": {
      "team":   { "type": "choice", "choice": "technical", "probabilities": { "billing": 0.05, "sales": 0.01, "technical": 0.94 }, "confidence": 0.83 },
      "urgent": { "type": "noul", "noul": 0.92 }
    },
    "usage": { "input_tokens": 346, "output_tokens": 0 }
  },
  "success": true
}
```

`PostTyped` is an alternative that decodes the Cloudflare envelope into a `*Response` (`Model`, `Answers`, `Usage`). It only works with endpoints that use that envelope.

## Errors

`Post` returns a non-nil `error` on transport failures. A non-2xx status is returned as a `*ResponseError`, which you can extract with `errors.As`.

## Documentation

All exported identifiers have doc comments:

```sh
go doc -all github.com/MarcelArt/systemone
```
