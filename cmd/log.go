/*
Copyright © 2023 Syndis ehf. <syndis@syndis.is>

*/
package cmd

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"time"

	"github.com/spf13/cobra"
	openapi "github.com/syndis-software/iris-api/pkg/openapi"
)

// logCmd represents the log command
var (
	logCmd = &cobra.Command{
		Use:   "log",
		Short: "Submit a log message for the current token",
		Long: `Submit a log message for the current token
	
Log messages can be viewed against the token in the Iris UI. This can provide a
useful feedback loop if you are configuring your token-using-application via
the token config in iris. Simply pass in any string and it will appear there.
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			if len(args) > 0 {
				// Log a single message
				logs := make([]openapi.SubmitLogEvent, 0, 1)
				logs = append(logs, openapi.SubmitLogEvent{
					Message:   args[0],
					Timestamp: time.Now().UnixMilli(),
				})
				err := upload(ctx, logs)

				if err != nil {
					return err
				}

			} else {
				// Assume we're in stdin mode.
				// 4 channels
				//  - message handler
				//	- stop (used to indicate exit)
				//  - ticker (controls how frequently we save)
				//  - done (used to force main thread to wait for goroutine completion)

				// 1 worker
				//  - api_worker: Add scanned text to a list. If it gets told by ticker to upload
				//                do that.

				stdIn := ctx.Value(stdInKey).(io.Reader)
				scanner := bufio.NewScanner(stdIn)

				messages := make(chan string)
				done := make(chan bool)
				stop := make(chan bool)

				var uploadFrequency time.Duration = 3 * time.Second
				ticker := time.NewTicker(uploadFrequency)

				go api_worker(messages, ticker.C, stop, done, ctx, cmd.ErrOrStderr())

				for scanner.Scan() {
					messages <- scanner.Text()
				}

				stop <- true

				<-done

				close(messages)
				close(stop)
			}
			return nil

		},
	}
)

func upload(ctx context.Context, logs []openapi.SubmitLogEvent) error {
	if len(logs) == 0 {
		return nil
	}
	client := ctx.Value(clientKey).(*openapi.ClientWithResponses)

	foo := openapi.SubmitLogsForTokenJSONRequestBody(logs)

	resp, err := client.SubmitLogsForToken(ctx, foo)

	if err != nil {
		return err
	}

	return openapi.CheckStatus(resp)
}

func api_worker(messages <-chan string, control <-chan time.Time, stop <-chan bool, done chan<- bool, ctx context.Context, errOut io.Writer) {
	logs := make([]openapi.SubmitLogEvent, 0, 1)
	count := 0
	var err error
	for {
		select {
		case message := <-messages:
			logs = append(logs, openapi.SubmitLogEvent{
				Message:   message,
				Timestamp: time.Now().UnixMilli(),
			})
			count += 1
		case <-stop:
			err = upload(ctx, logs)
			if err != nil {
				fmt.Fprintf(errOut, "%s Error: %s\n", time.Now().Format(time.RFC3339), err.Error())
			}
			done <- true
			return
		case <-control:
			err = upload(ctx, logs)
			if err != nil {
				fmt.Fprintf(errOut, "%s Error: %s\n", time.Now().Format(time.RFC3339), err.Error())
			}
			logs = nil
		}
	}

}

func init() {
	rootCmd.AddCommand(logCmd)

}
