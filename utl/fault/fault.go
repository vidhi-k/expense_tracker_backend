package fault

import "fmt"

type AppError struct {
	Err     error
	Message string
	Status  int
}

func (appErr *AppError) SetError(err error) error {
	appErr.Err = err
	return appErr
}

type HTTPErrorResp struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Err     string `json:"error"`
}

func (appErr *AppError) Error() string {
	errMsg := ""
	errMsg = appErr.Message
	err := appErr

	for {
		if ae, ok := err.Err.(*AppError); ok {
			err = ae

			errMsg += "->" + ae.Message
		} else {
			if err.Err != nil {
				errMsg += fmt.Sprintf("-> internal error: %s", err.Err.Error())
			}

			break
		}
	}

	return errMsg
}
