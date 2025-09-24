package printer

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
)

func FormatResponse(resp *http.Response, includeHeaders bool, prettyJSON bool) string {
	fmtResp := ""

	fmtResp += fmt.Sprintf("%s %s\n", resp.Proto, resp.Status)

	if includeHeaders {
		for key, values := range resp.Header {
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

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmtResp + "\n"
	}

	if prettyJSON && json.Valid(bodyBytes) {
		var data interface{}

		err := json.Unmarshal(bodyBytes, &data)
		if err != nil {
			// TODO			
			return "" 
		}

		bodyBytes, err = json.MarshalIndent(&data, "", "  ")
		if err != nil {
			// TODO			
			return "" 
		}
	}

	return fmtResp + string(bodyBytes)
}