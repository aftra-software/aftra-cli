/*
Copyright © 2023 Syndis ehf. <syndis@syndis.is>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
	openapi "github.com/syndis-software/aftra-api/pkg/openapi"
)

var (
	submitCmd_filename string
	submitCmd_message  string

	submitCmd = &cobra.Command{
		Use:   "submit [scan-type] [scan-name] [scan-result]",
		Args:  cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
		Short: "Submit a json formatted scan result",
		Long: `Submit a json formatted scan result
	
Submit a scan result in the format for given scan-type. For example in
nessus format for syndis scans.
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			client := ctx.Value(clientKey).(*openapi.ClientWithResponses)
			scanType, scanName := args[0], args[1]

			switch {
			case ScanType(scanType) == syndis:

				var results openapi.BodySubmitScanResults
				// Submit a set of scan events
				if submitCmd_filename != "" {
					// We attempt to parse the contents to see if its a valid json file
					err := validate_json_file(submitCmd_filename)
					if err != nil {
						return err
					}

					company := ctx.Value(companyKey).(string)
					// Upload the file
					// 1 Get a signed upload url
					uploadInfo, err := openapi.DoGetUploadURL(ctx, client, company)
					if err != nil {
						return err
					}

					// 2 Upload the file
					err = upload_file(uploadInfo.Url, submitCmd_filename, uploadInfo.Fields)
					if err != nil {
						return err
					}

					// 3 Notify using submit scan results
					blobInfo := openapi.BlobUploadInfo{
						Bucket: uploadInfo.Bucket,
						Key:    uploadInfo.Key,
					}

					results.BlobUpload = &blobInfo
				} else {
					var scans []openapi.SyndisInternalScanEvent
					err := json.Unmarshal([]byte(submitCmd_message), &scans)
					if err != nil {
						return err
					}
					results.Events = &scans
				}

				resp, err := client.SubmitScanResults(ctx, scanName, results)

				if err != nil {
					return err
				}

				return openapi.CheckStatus(resp)
			default:
				return fmt.Errorf("unrecognised scan type %s", scanType)
			}

		},
	}
)

func validate_json_file(filename string) error {
	var scans []openapi.SyndisInternalScanEvent
	jsonFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &scans)
	return err
}
func createMultipartForm(filepath string, fields map[string]string) (bytes.Buffer, *multipart.Writer, error) {
	var b bytes.Buffer

	w := multipart.NewWriter(&b)
	var fw io.Writer
	file, err := os.Open(filepath)
	if err != nil {
		return b, nil, err
	}

	for key, val := range fields {
		if fw, err = w.CreateFormField(key); err != nil {
			return b, nil, err
		}
		if _, err = io.Copy(fw, strings.NewReader(val)); err != nil {
			return b, nil, err
		}
	}

	if fw, err = w.CreateFormFile("file", file.Name()); err != nil {
		return b, nil, err
	}
	if _, err = io.Copy(fw, file); err != nil {
		return b, nil, err
	}

	w.Close()
	return b, w, nil
}

func upload_file(url string, filepath string, fields map[string]string) error {
	byteBuffer, multiWriter, err := createMultipartForm(filepath, fields)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, &byteBuffer)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", multiWriter.FormDataContentType())

	// Submit the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return errors.New(string(b))
	}
	return nil
}

func init() {

	submitCmd.Flags().StringVarP(&submitCmd_filename, "filename", "f", "", "JSON file to submit")
	submitCmd.Flags().StringVarP(&submitCmd_message, "message", "m", "", "JSON string to submit")

	// Want this, but no way to clear flag values between tests at the moment
	// https://github.com/spf13/cobra/issues/1180
	// submitCmd.MarkFlagsMutuallyExclusive("filename", "message")

	rootCmd.AddCommand(submitCmd)

}
