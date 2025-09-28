package printer

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/TommyFiga/greq/internal/types"
)

func FormatResponse(resp *types.ResponseData , includeHeaders bool, prettyJSON bool) string {
	fmtResp := ""
	fmtResp += fmt.Sprintf("%s %s\n", resp.Protocol, resp.Status)

	if includeHeaders {
		for key, values := range resp.Headers {
			fmtResp += (key + ": ")

			for i, val := range values {
				fmtResp += val

				if i != len(values) - 1 {
					fmtResp += ", "
				}
			}

			fmtResp += "\n"
		}
	}

	bodyBytes := resp.Body

	if prettyJSON {
		var data any

		err := json.Unmarshal(bodyBytes, &data)
		if err != nil {		
			fmt.Fprintf(os.Stdout, "WARNING: failed to parse JSON for pretty-printing\n")
			return fmtResp + "\nData:\n" + string(bodyBytes)
		}

		indented, err := json.MarshalIndent(&data, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "WARNING: failed to format JSON: %v\n", err)
			return fmtResp + "\nData:\n" + string(bodyBytes)
		}

		bodyBytes = indented
	}

	return fmtResp + "\nData:\n" + string(bodyBytes)
}